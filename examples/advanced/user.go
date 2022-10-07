package main

import (
	"fmt"
	gre "github.com/darksubmarine/goldfish-re"
)

// User object which is linked to the ruleset.
type User struct {
	// usual fields. Not included as facts in context
	id    int
	name  string
	email string

	// Facts
	plan   gre.String
	miles  gre.Number
	status gre.String

	// Reactive context
	ctx gre.FactsContext
}

// NewUser returns a User pointer or error if facts were not registered into the given context.
// Initializes and registers each fact into the given context.
func NewUser(id int, name string, email string, ctx gre.FactsContext) (*User, error) {
	usr := &User{
		id:     id,
		name:   name,
		email:  email,
		plan:   gre.NewString("User", "plan", "gold"),
		status: gre.NewString("User", "status", "pending"),
		miles:  gre.NewNumber("User", "miles", 1000),
		ctx:    ctx,
	}

	if err := ctx.RegisterString(usr, usr.plan); err != nil {
		return nil, err
	}

	if err := ctx.RegisterString(usr, usr.status); err != nil {
		return nil, err
	}

	if err := ctx.RegisterNumber(usr, usr.miles); err != nil {
		return nil, err
	}
	return usr, nil
}

// Set sets all attributes into an update transaction to trigger only once the ruleset evaluation (blocking method).
func (u *User) Set(plan, status string, miles int64) error {
	return u.ctx.Update(func(tx *gre.Tx) {
		tx.SetString(u.plan, plan)
		tx.SetString(u.status, status)
		tx.SetNumber(u.miles, miles)
	})
}

// SetPlan sets the plan attribute via the given context.
// Triggers a ruleset evaluation due to is set into a single transaction (blocking method).
func (u *User) SetPlan(plan string) error {
	return u.ctx.SetString(u.plan, plan)
}

// SetMiles sets the miles attribute via the given context.
// Triggers a ruleset evaluation due to is set into a single transaction (blocking method).
func (u *User) SetMiles(miles int64) error {
	return u.ctx.SetNumber(u.miles, miles)
}

// SetStatus sets the status attribute via the given context.
// Triggers a ruleset evaluation due to is set into a single transaction (blocking method).
func (u *User) SetStatus(status string) error {
	return u.ctx.SetString(u.status, status)
}

// Plan returns the fact value.
// This method locks the fact to avoid changes when it is read.
func (u *User) Plan() string {
	return u.plan.Value()
}

// Miles returns the fact value.
// This method locks the fact to avoid changes when it is read.
func (u *User) Miles() int64 {
	return u.miles.Value()
}

// Status returns the fact value.
// This method locks the fact to avoid changes when it is read.
func (u *User) Status() string {
	return u.status.Value()
}

// Id getter for id field
func (u *User) Id() int { return u.id }

// Name getter for name field
func (u *User) Name() string { return u.name }

// Email getter for email field
func (u *User) Email() string { return u.email }

func (u *User) String() string {
	return fmt.Sprintf("User:{id=%d name=%s email=%s plam=%s miles=%d status=%s}", u.id, u.name, u.email, u.Plan(), u.Miles(), u.Status())
}
