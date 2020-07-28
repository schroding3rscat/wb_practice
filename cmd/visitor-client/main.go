package main

import (
	"fmt"

	"github.com/schroding3rscat/wb_practice/internal/visitor"
)

func main() {
	admin := visitor.Admin{
		ID:       1,
		Name:     "root",
		TenantID: 100,
	}

	user := visitor.User{
		ID:       2,
		Name:     "john",
		ReadOnly: true,
		TenantID: 100,
	}

	name, _ := admin.GetName()
	fmt.Printf("admin name is %s\n", name)

	name, _ = user.GetName()
	fmt.Printf("user name is %s\n", name)

	idProvider := visitor.NewIDProvider()

	_ = admin.Accept(idProvider)
	_ = user.Accept(idProvider)
}
