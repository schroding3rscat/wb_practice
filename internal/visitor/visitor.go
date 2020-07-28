package visitor

// Users provides methods to work with various user accounts.
type Users interface {
	GetID() (id int, err error)
	GetName() (name string, err error)
	SetName(name string) (err error)
	Accept(v Visitor) (err error)
}

// Visitor is an interface extending common users operations.
type Visitor interface {
	VisitAdmin(admin *Admin) (err error)
	VisitUser(user *User) (err error)
}
