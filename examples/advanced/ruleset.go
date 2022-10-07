package main

import (
	"fmt"
	gre "github.com/darksubmarine/goldfish-re"
	"log"
)

const thenActiveGoldAward = "ACTIVE_GOLD_AWARD"
const thenActiveGoldAwardByStatusChange = "ACTIVE_GOLD_AWARD_BY_STATUS_CHANGE"

func createRuleset() (gre.Ruleset, error) {
	rs := gre.Builder().Ruleset().
		OnActivation(func(then string, context gre.Context) {

			// Print the THEN variation
			fmt.Println("Activated rule:", then)
			switch then {
			case thenActiveGoldAward:
				// rule activation by gold award

				// Sometimes when we are working with a ruleset where some facts depend on other facts is needed an
				// inference mechanism.
				// In this case we called Feedback due to each feedback will trigger a ruleset re-evaluation.
				// More info at: https://darksubmarine.com/docs/goldfish-re/starthere.html#context-into-onactivation-handler
				context.Feedback(func(tx *gre.Tx) {
					log.Println("Feedback")
					if userStatusFact, err := context.GetString("User.status"); err == nil {
						tx.SetString(userStatusFact, "VIP")
					}
				})

			case thenActiveGoldAwardByStatusChange:
				// rule activated by status change
			}

			// Iterates over all context facts
			context.ForEach(func(fact string, value interface{}) {
				fmt.Printf("fact=%s val=%#v\n", fact, value)
			})

			// Fetch an individual fact
			if obj, ok := context.Get("User.plan"); ok {
				fmt.Println("-- context.Get(User.plan)=", obj.(gre.String).Value())
			}

			// Fetch the fact's parent object to access other fields that are not facts
			if obj, ok := context.GetObject("User"); ok {
				fmt.Println("-- context.GetObject(User).Email()=", obj.(*User).Email())

				// fatal error: all goroutines are asleep - deadlock!
				// the User).SetMiles() has a reference to the same context that triggered the previous update.
				// To cover this use case please see the previous Feedback function.
				//if err := obj.(*User).SetMiles(5000); err != nil {
				//	fmt.Println("**** DEADLOCK!", err)
				//}
			}
		}).
		OnError(func(err error) {
			log.Println(err)
		}).Build()

	//User.plan is gold
	condUserPlanGold := gre.Builder().StringCondition().Term("User", "plan").Equal("gold").Build()

	//User.miles are greater than 3000
	condUserMiles3000 := gre.Builder().NumberCondition().Term("User", "miles").GreaterThan(3000).Build()

	//User.status could be one of active, referred, VIP
	condUserStatus := gre.Builder().StringCondition().Term("User", "status").In([]string{"active", "referred", "VIP"}).Build()

	// Rule frequent flyer
	ruleFrequentFlyer, err := gre.Builder().Rule().AllOf(condUserPlanGold, condUserMiles3000).Then(thenActiveGoldAward).Build()
	if err != nil {
		return nil, err
	}

	// Rule user status update
	ruleUserStatusUpdate, err := gre.Builder().Rule().AnyOf(condUserStatus).Then(thenActiveGoldAwardByStatusChange).Build()
	if err != nil {
		return nil, err
	}

	rs.AddRule(ruleFrequentFlyer)
	rs.AddRule(ruleUserStatusUpdate)

	return rs, nil
}
