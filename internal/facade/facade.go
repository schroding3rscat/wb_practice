package facade

import "github.com/schroding3rscat/wb_practice/internal/models"

type db interface {
	GetTenant(subscriptionID int) (tenant models.Tenant, err error)
	CreateSubscription(subscription models.Subscription) (subscriptionID int, err error)
	DeleteSubscription(subscriptionID int) (err error)
	UpdateSubscription(subscriptionID int, plan string) (err error)
}

type cloud interface {
	CreateTenant(tenant models.Tenant) (tenantID int, err error)
	DeleteTenant(tenantID int) (err error)
	SetTenantQuota(tenantID int, storageBytes int64) (err error)
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
	db    db
	cloud cloud
	plan  plan
}

// CreateCreateSubscription provisions a new tenant and sets its quota according the plan.
func (p *provisioner) CreateSubscription(name string, plan string, billingPeriodDay int) (subscriptionID int, err error) {
	storageBytes, err := p.plan.GetQuota(plan)
	if err != nil {
		return
	}

	tenantID, err := p.cloud.CreateTenant(models.Tenant{Name: name})
	if err != nil {
		return
	}

	err = p.cloud.SetTenantQuota(tenantID, storageBytes)
	if err != nil {
		return
	}

	subscriptionID, err = p.db.CreateSubscription(models.Subscription{
		TenantID:         tenantID,
		Plan:             plan,
		BillingPeriodDay: billingPeriodDay,
	})
	if err != nil {
		return
	}

	return
}

func (p *provisioner) ChangeSubscription(subscriptionID int, plan string) (err error) {
	tenant, err := p.db.GetTenant(subscriptionID)
	if err != nil {
		return
	}

	storageBytes, err := p.plan.GetQuota(plan)
	if err != nil {
		return
	}

	err = p.cloud.SetTenantQuota(tenant.ID, storageBytes)
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
	tenant, err := p.db.GetTenant(subscriptionID)
	if err != nil {
		return
	}

	err = p.cloud.DeleteTenant(tenant.ID)
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
		db:    db,
		cloud: cl,
		plan:  pl,
	}
}
