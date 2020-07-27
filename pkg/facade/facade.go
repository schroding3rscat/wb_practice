package facade

type db interface {
	GetTenantID(subscriptionID int) (tenantID int, err error)
	CreateSubscription(tenantID int, plan string, billingPeriodDay int) (subscriptionID int, err error)
	DeleteSubscription(subscriptionID int) error
	UpdateSubscription(subscriptionID int, plan string) error
}

type cloud interface {
	CreateTenant(name string) (tenantID int, err error)
	DeleteTenant(tenantID int) error
	SetTenantQuota(tenantID int, storageBytes int64) error
}

type plan interface {
	GetQuota(plan string) (storageBytes int64, err error)
}

// Provisioner is an interface describing all necessary methods for subscription handling.
type Provisioner interface {
	CreateSubscription(name string, plan string, billingPeriodDay int) (subscriptionID int, err error)
	ChangeSubscription(subscriptionID int, plan string) error
	TerminateSubscription(subscriptionID int) error
}

type provisioner struct {
	db db
	cl cloud
	pl plan
}

// CreateCreateSubscription provisions a new tenant and sets its quota according the plan.
func (p *provisioner) CreateSubscription(name string, plan string, billingPeriodDay int) (subscriptionID int, err error) {
	storageBytes, err := p.pl.GetQuota(plan)
	if err != nil {
		return
	}

	tenantID, err := p.cl.CreateTenant(name)
	if err != nil {
		return
	}

	err = p.cl.SetTenantQuota(tenantID, storageBytes)
	if err != nil {
		return
	}

	subscriptionID, err = p.db.CreateSubscription(tenantID, plan, billingPeriodDay)
	if err != nil {
		return
	}

	return
}

func (p *provisioner) ChangeSubscription(subscriptionID int, plan string) (err error) {
	tenantID, err := p.db.GetTenantID(subscriptionID)
	if err != nil {
		return
	}

	storageBytes, err := p.pl.GetQuota(plan)
	if err != nil {
		return
	}

	err = p.cl.SetTenantQuota(tenantID, storageBytes)
	if err != nil {
		return
	}

	err = p.db.UpdateSubscription(subscriptionID, plan)
	if err != nil {
		return
	}

	return
}

func (p *provisioner) TerminateSubscription(subscriptionID int) (err error) {
	tenantID, err := p.db.GetTenantID(subscriptionID)
	if err != nil {
		return
	}

	err = p.cl.DeleteTenant(tenantID)
	if err != nil {
		return
	}

	err = p.db.DeleteSubscription(subscriptionID)
	if err != nil {
		return
	}

	return
}

// NewProvisioner returns implementation of Provisioner interface.
func NewProvisioner(db db, cl cloud, pl plan) Provisioner {
	return &provisioner{
		db: db,
		cl: cl,
		pl: pl,
	}
}
