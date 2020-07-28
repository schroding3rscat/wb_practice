package models

// Subscription represents resources provisioned to a customer.
type Subscription struct {
	TenantID         int
	Plan             string
	BillingPeriodDay int
}
