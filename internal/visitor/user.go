package visitor

// User represents regular user account.
type User struct {
	ID       int
	Name     string
	TenantID int
	ReadOnly bool
}

// GetID returns ID of the current user.
func (u *User) GetID() (id int, err error) {
	return u.ID, nil
}

// GetName returns name of the current user.
func (u *User) GetName() (name string, err error) {
	return u.Name, nil
}

// SetName sets the name for the current user.
func (u *User) SetName(name string) (err error) {
	u.Name = name
	return
}

// Accept extend functionality for the current user.
func (u *User) Accept(v Visitor) (err error) {
	return v.VisitUser(u)
}
