package main

import (
	gre "github.com/darksubmarine/goldfish-re"
	"log"
)

// UserModel object with facts
type UserModel struct {
	// usual fields. Not included as facts in context
	id    int
	name  string
	email string

	// Facts
	Plan  gre.String `gre:"object=User,attribute=plan,value=silver"`
	Miles gre.Number `gre:"object=User,attribute=miles,value=500"`
}

func main() {

	// Condition 1:
	//		User.plan in ["gold", "platinum"]
	c1 := gre.Builder().StringCondition().Term("User", "plan").In([]string{"gold", "platinum"}).Build()

	// Condition 2:
	//		User.miles > 1000
	c2 := gre.Builder().NumberCondition().Term("User", "miles").GreaterThan(1000).Build()

	// Rule 1:
	//	User.plan in ["gold", "platinum"] || User.miles > 1000
	r1, err := gre.Builder().Rule().AnyOf(c1, c2).Then("APPLY_RULE_1").Build()
	if err != nil {
		panic(err)
	}

	// Ruleset build
	rs := gre.Builder().Ruleset().
		OnActivation(func(then string, context gre.Context) {
			log.Println(then)
		}).
		OnError(func(err error) {
			log.Println(err)
		}).
		Build()

	// Adding rule 1 into the ruleset
	rs.AddRule(r1)

	// Getting a ruleset context to work with facts
	ctx := rs.Context()

	// New user object
	usr := new(UserModel)
	if err := ctx.Register(usr); err != nil {
		log.Println("error registering user into ruleset context", err)
	}

	// Updating facts into a transactional function
	if err := ctx.Update(func(tx *gre.Tx) {
		tx.SetNumber(usr.Miles, 2500)
	}); err != nil {
		log.Println(err)
	}
}
