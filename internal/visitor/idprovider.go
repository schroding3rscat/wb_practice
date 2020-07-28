package visitor

import "fmt"

type tenantIDProvider struct{}

// VisitAdmin returns tenant ID for the current admin account.
func (t *tenantIDProvider) VisitAdmin(admin *Admin) (err error) {
	fmt.Printf("tenant id for admin %s is %v\n", admin.Name, admin.TenantID)
	return
}

// VisitUser returns tenant ID for the current user account.
func (t *tenantIDProvider) VisitUser(user *User) (err error) {
	fmt.Printf("tenant id for user %s is %v\n", user.Name, user.TenantID)
	return
}

// NewIDProvider returns new Visitor implementation providing tenant IDs for user accounts.
func NewIDProvider() Visitor {
	return &tenantIDProvider{}
}
