package cloud

import (
	"fmt"
	"math/rand"
)

// Cloud contains methods for tenant handling.
type Cloud interface {
	CreateTenant(name string) (tenantID int, err error)
	DeleteTenant(tenantID int) (err error)
	SetTenantQuota(tenantID int, storageBytes int64) (err error)
}

type cloudClient struct{}

// CreateTenant adds a new tenant to the cloud.
func (c *cloudClient) CreateTenant(name string) (tenantID int, err error) {
	tenantID = rand.Int()
	fmt.Printf("created tenant %v\n", tenantID)
	return
}

// DeleteTenant removes the tenant from the cloud.
func (c *cloudClient) DeleteTenant(tenantID int) (err error) {
	fmt.Printf("deleted tenant %v\n", tenantID)
	return
}

// SetTenantQuota updated storage quota for the tenant.
func (c *cloudClient) SetTenantQuota(tenantID int, storageBytes int64) (err error) {
	fmt.Printf("set quota %v for tenant %v\n", storageBytes, tenantID)
	return
}

// NewClient returns the interface for tenant handling.
func NewClient(endpoint string) Cloud {
	return &cloudClient{}
}
