package cloud

import (
	"fmt"
	"math/rand"
)

// Cloud contains methods for tenant handling.
type Cloud interface {
	CreateTenant(name string) (tenantID int, err error)
	DeleteTenant(tenantID int) error
	SetTenantQuota(tenantID int, storageBytes int64) error
}

type cloudClient struct {
}

// CreateTenant adds a new tenant to the cloud.
func (c *cloudClient) CreateTenant(name string) (tenantID int, err error) {
	tenantID = rand.Int()
	fmt.Printf("created tenant %v\n", tenantID)
	return
}

// DeleteTenant removes the tenant from the cloud.
func (c *cloudClient) DeleteTenant(tenantID int) error {
	fmt.Printf("deleted tenant %v\n", tenantID)
	return nil
}

// SetTenantQuota updated storage quota for the tenant.
func (c *cloudClient) SetTenantQuota(tenantID int, storageBytes int64) error {
	fmt.Printf("set quota %v for tenant %v\n", storageBytes, tenantID)
	return nil
}

// NewClient returns the interface for tenant handling.
func NewClient(endpoint string) Cloud {
	return &cloudClient{}
}
