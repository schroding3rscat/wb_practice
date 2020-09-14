package visitor

// Admin represents admin account.
type Admin struct {
	ID       int
	Name     string
	TenantID int
}

// GetID returns ID of the current admin.
func (a *Admin) GetID() (id int, err error) {
	return a.ID, nil
}

// GetName returns name of the current admin.
func (a *Admin) GetName() (name string, err error) {
	return a.Name, nil
}

// SetName sets the name for the current admin.
func (a *Admin) SetName(name string) (err error) {
	a.Name = name
	return
}

// Accept extend functionality for the current admin.
func (a *Admin) Accept(v Visitor) (err error) {
	return v.VisitAdmin(a)
}
