package main

import (
	"fmt"
	"time"
)

const milesSFO2NYC = 2580

// Flight Award program rules
//
// Ruleset User flyer gold award
//   Rule frequent flyer
//   When all
//     User.plan is gold
//     User.miles are greater than 3000
//   Then
//     ACTIVE_GOLD_AWARD
//
//   Rule user status update
//   When any
//     User.status could be one of active, referred, VIP
//   Then
//     ACTIVE_GOLD_AWARD_BY_STATUS_CHANGE

func main() {
	rs, err := createRuleset()
	if err != nil {
		panic(fmt.Sprintf("rulse cannot be created. Error=%s", err))
	}

	// Once that you have the ruleset context, you can inject it everywhere.
	ctx := rs.Context()
	ctx.WithMaxIterations(10)

	// User creation. See the NewUser function where all facts are registered into the given context
	usr, err := NewUser(1, "John", "jhon@gre.com", ctx)
	if err != nil {
		panic(err)
	}

	usr2, err := NewUser(2, "Phil", "phil@gre.com", rs.Context())
	if err != nil {
		panic(err)
	}

	go func() {
		time.Sleep(time.Second)
		if err := usr2.Set("silver", "active", milesSFO2NYC); err != nil {
			panic(err)
		}
	}()

	if err := usr.SetMiles(usr.Miles() + milesSFO2NYC); err != nil {
		panic(err)
	}

	time.Sleep(10 * time.Second)

	fmt.Println(usr)
}
