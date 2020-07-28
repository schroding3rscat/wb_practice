package plan

import (
	"fmt"
	"math/rand"
)

// PlansProvider contains methods providing plans information.
type PlansProvider interface {
	GetQuota(plan string) (storageBytes int64, err error)
}

type plans struct{}

// GetQuota retuns storage quota for the plan.
func (p *plans) GetQuota(plan string) (storageBytes int64, err error) {
	storageBytes = rand.Int63()
	fmt.Printf("got quota for plan %s\n", plan)
	return
}

// NewProvider returns implementation of PlansProvider interface.
func NewProvider() PlansProvider {
	return &plans{}
}
