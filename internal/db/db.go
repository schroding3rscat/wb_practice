package db

import (
	"fmt"
	"math/rand"

	"github.com/schroding3rscat/wb_practice/internal/models"
)

// Subscriptioner describes all methods for subscription handling.
type Subscriptioner interface {
	GetTenant(subscriptionID int) (tenant models.Tenant, err error)
	CreateSubscription(subscription models.Subscription) (subscriptionID int, err error)
	DeleteSubscription(subscriptionID int) (err error)
	UpdateSubscription(subscriptionID int, plan string) (err error)
}

type dbClient struct{}

func (c *dbClient) GetTenant(subscriptionID int) (tenant models.Tenant, err error) {
	tenant = models.Tenant{ID: rand.Int(), Name: "some name"}
	fmt.Printf("got tenant id %+v\n", tenant)
	return
}

// CreateSubscription adds a new subscription to the database.
func (c *dbClient) CreateSubscription(subscription models.Subscription) (subscriptionID int, err error) {
	subscriptionID = rand.Int()
	fmt.Printf("inserted subscription %v\n", subscriptionID)
	return
}

// DeleteSubscription removes subscription from the database.
func (c *dbClient) DeleteSubscription(subscriptionID int) (err error) {
	fmt.Printf("deleted subscription %v\n", subscriptionID)
	return
}

// UpdateSubscription sets a new plan for the existing subscription.
func (c *dbClient) UpdateSubscription(subscriptionID int, plan string) (err error) {
	fmt.Printf("updated subscription %v\n", subscriptionID)
	return
}

// NewClient returns the interface for subscription handling.
func NewClient(cs string) Subscriptioner {
	return &dbClient{}
}
