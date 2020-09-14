package main

import (
	"github.com/schroding3rscat/wb_practice/internal/cloud"
	"github.com/schroding3rscat/wb_practice/internal/db"
	"github.com/schroding3rscat/wb_practice/internal/facade"
	"github.com/schroding3rscat/wb_practice/internal/plan"
)

func main() {
	dbClient := db.NewClient("file:test.db?cache=shared&mode=memory")
	cloudClient := cloud.NewClient("http://localhost:8080")
	plans := plan.NewProvider()
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
