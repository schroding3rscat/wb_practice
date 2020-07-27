package main

import (
	"github.com/schroding3rscat/wb_practice/pkg/cloud"
	"github.com/schroding3rscat/wb_practice/pkg/db"
	"github.com/schroding3rscat/wb_practice/pkg/facade"
	"github.com/schroding3rscat/wb_practice/pkg/plan"
)

func main() {
	dbClient := db.NewClient("file:test.db?cache=shared&mode=memory")
	cloudClient := cloud.NewClient("http://localhost:8080")
	plans := plan.NewProvider("/tmp/config.bak")
	pr := facade.NewProvisioner(dbClient, cloudClient, plans)

	subscriptionID, err := pr.CreateSubscription("some name", "SKU123", 30)
	if err != nil {
		panic("cannot create a subscription")
	}

	err = pr.ChangeSubscription(subscriptionID, "SKU456")
	if err != nil {
		panic("cannot create a subscription")
	}

	err = pr.TerminateSubscription(subscriptionID)
	if err != nil {
		panic("cannot create a subscription")
	}

}
