
## Installation
To install the library into your project please run:

```text
go install github.com/darksubmarine/goldfish-re
```

## Quick start

The library exposes a simple API to create conditions, rules, ruleset, facts and run the evaluation.

### Ruleset definition
The following snippets illustrates how to create a ruleset for the next definition:

```text
Ruleset User flyer gold award
    Rule frequent flyer
    When
        User.plan is gold
        User.miles are greater than 300
    Then
        return ACTIVE_GOLD_AWARD
        
    Rule user status update
    When
        User.status could be one of active, referred, VIP
    Then
        return ACTIVE_GOLD_AWARD_BY_STATUS_CHANGE
```

!!! info "Termination rules"
    Note that so far the engine only supports termination rules, that means: the THEN part of the rule only can return
    a string that will be sent to the activation function

The code should be written like:

```go
package main

import (
	"fmt"
	
	gre "github.com/darksubmarine/goldfish-re"
)

func createRuleset() gre.Ruleset {
	// C1: User.Plan == "gold"
	c1 := gre.Builder().StringCondition().
		Term("User", "plan").
		Equal("gold").
		Build()

	// C2: User.miles > 300
	c2 := gre.Builder().NumberCondition().Term("User", "miles").GreaterThan(300).Build()

	// C3: User.status in ["active", "referred", "VIP"]
	c3 := gre.Builder().StringCondition().
		Term("User", "status").
		In([]string{"active", "referred", "VIP"}).
		Build()

	var frequentFlyer = gre.Builder().Rule().AllOf(c1, c2).Then("ACTIVE_GOLD_AWARD").Build()
	var userStatusUpdate = gre.Builder().Rule().AllOf(c3).Then("ACTIVE_GOLD_AWARD_BY_STATUS_CHANGE").Build()
	
	rs := gre.Builder().Ruleset().
		OnActivation(func(then string, ctx gre.Context) {
            // Do soomething on activation (1)
		}).
		OnError(func(err error) {
			// handle the error at your convenience (2)
		}).
		Build()

	rs.AddRule(frequentFlyer)
	rs.AddRule(userStatusUpdate)
	
	return rs
}
```

1. Activation function that receives the `then` value of the activated rule and the context `ctx` with all facts.
2. Error handler function receives the error and you can log it or do something else based on it.

#### Activation function
This function is a callback function that will be called each time that a rule returns `true` when a fact or facts are
updated.

The function receives the `then` value of the activated rule and a pointer to the evaluated context having available all
registered facts and objects into this one.

```go
package main

import (
	"fmt"

	gre "github.com/darksubmarine/goldfish-re"
)

func yourActivationHandler(then string, ctx gre.Context) {
	
	switch then {
	case "ACTIVE_GOLD_AWARD":
		// do something 
    case "ACTIVE_GOLD_AWARD_BY_STATUS_CHANGE":
		// do something relate with user status
    }

	// iterate over the context all registered facts
	ctx.ForEach(func(fact string, value interface{}) {
		fmt.Printf("%s=%v\n", fact, value) // will print User.plan=gold or User.miles=500 or User.status=active
	})

	// get the fact that you would like to work with
	if usrPlan, ok := ctx.Get("User.plan"); ok {
		switch pln := usrPlan.(type) {
		case gre.String:
			fmt.Println("fact type (string):", pln.Value())
		case gre.Number:
			fmt.Println("fact type (number):", pln.Value())
		}
	}

	// get the fact's parent object to access it and use other fields that are not facts. 
	if usr, ok := ctx.GetObject("User"); ok {
		fmt.Println(obj.(*User).Email)
    }

}
```

!!! Tip "Facts inference"
    Sometimes is needed performing an inference due to a facts change. The given `context` into the `onActivation` handler
    will let you update facts in a thread-safe manner calling the cnotext method `Feedback`.
    Please refer to [Context section](#context-into-onactivation-handler) for more details.

#### Error handler
If for some reason the evaluation ran into an error this callback function will be called letting you know about the error.

```go
package main

import (
	"fmt"
	"log"

	gre "github.com/darksubmarine/goldfish-re"
)

func errorHandler(err error) {
	// do something with the error
	log.Println(err)
}
```

### Facts as struct fields

Following the previous example we have an object `User` with 3 fields `plan` `miles` and `status`. 
Those attributes are the facts and can be defined into your go program as:

```go
package main

import gre "github.com/darksubmarine/goldfish-re" 

type User struct {
    plan   gre.String
    miles  gre.Number
    status gre.String
}
```

Where the struct name (User) matches with the fact object and the field name matches with the fact attribute. Also the
data type must be set using the Goldfish-RE data types in order to trigger an evaluation each time that a value is updated.


#### Go annotation
The library also supports Go tags to configure your facts. The tag must be written as:

```text
gre:"object=User,attribute=plan,value=silver"
```
where this is useful when you need to specify some custom values:

 - **object**: defines the object name for the given fact, useful to overwrite the struct name.
 - **attribute**: sets the attribute name for the given fact if you need a name different to the struct field
 - **value**: the default value which the fact will be initialized. Possible values
    - _String_:  `value=some string value`
    - _Number_: `value=234`
    - _Float_: `value=73.2`
    - _Boolean_: `value=true`
    - _Date_: `value=1984-12-24T00:00:00`



```go
package main

import gre "github.com/darksubmarine/goldfish-re" 

type User struct {
    Plan   gre.String `gre:"attribute=plan,value=silver"`
    Miles  gre.Number `gre:"attribute=miles,value=500"`
    Status gre.String `gre:"attribute=status,value=none"`
}
```

!!! Danger "Exported fields"
    In order to use the Go tags to configure the facts, is mandatory exporting the fields otherwise the
    context registration will fail.

### Context
Once that you have a `ruleset` what do you need is a `context` to eval your facts against the ruleset. That means:

 - A context is created from a ruleset: `rs.Context()`
 - Each ruleset has only one context
 - The facts are thread-safe into the same context via `ctx.Update` method
 - The facts are not thread-safe between different context

#### Register facts
In order to run evaluations against the ruleset each time that a Fact is updated, is required to register the facts into a Context.

!!! info ""
    Register methods do not run a ruleset evaluation. Only register the facts with their `zero`/`default` values into the context.
    The evaluation happens when a `context.Update` is called and the facts are modified via a transaction `Tx`

To do this the Context struct has some methods:

##### Register a full struct

`ctx.Register(obj interface{})`: This method is the most useful when you have a struct with facts to be registered.
 Also fetchs the tags and applies its configuration. Returns an error if the registration process fails.

For instance:
```go
type User struct {
	Plan   gre.String `gre:"attribute=plan,value=silver"`
	Miles  gre.Number `gre:"attribute=miles,value=500"`
	Status gre.String `gre:"attribute=status,value=none"`
}

usr := new(User)
if err := ctx.Register(usr); err != nil {
    // error on registration
}
```

##### Register field by field

In addition to the `Register` method it is possible registering individual field with its data type:

 - `RegisterString(object interface{}, attribute String)`
 - `RegisterNumber(object interface{}, attribute Number)`
 - `RegisterFloat(object interface{}, attribute Float)`
 - `RegisterBoolean(object interface{}, attribute Boolean)`
 - `RegisterDate(object interface{}, attribute Date)`

For instance:
```go
type User struct {
	Plan   gre.String
}

usr := User{ Plan: gre.NewString("User", "plan", "gold")}
ctx.RegisterString(usr, usr.Plan)
```



#### Update facts
Each time that a fact or facts are updated a ruleset evaluation must be run in order to check if some variation activates any rule.

The context object exposes a method to execute fact updates in a thread-safe mode. Each call is a locking call and returns the control
when the success function (`onActivation`) has finished.  

!!! Tip "Improving performance"
    To improve performance your success function could be executed as a go routine.
    Have into account to copy the context facts if you need it because, to avoid degrade performance, are pointers to the main ruleset context.

Calling `context.Update` you can update multiples facts at once in a transactional way and after that the evaluation will be trigger.

```go
package main

import (
	gre "github.com/darksubmarine/goldfish-re"
	"log"
)

func updateFacts() {
	if err := ctx.Update(func(tx *gre.Tx) {
		tx.SetString(usr.Status, "active")
		tx.SetNumber(usr.Miles, 999)

		tx.Error(errors.New("something was wrong"))
	}); err != nil {
		log.Println(err)
	}
}
```

The `ctx.Update` executes a function where a `tx *gre.Tx` is a transaction object and let you perform transactional 
updates to your facts and exposes a `tx.Error` method in case that you would like to return a custom error to avoid run 
the ruleset evaluation.

When an `error` is returned the transaction is not applied that means: fact changes are not updated (committed).

If you need updating only one fact, the context object exposes individual methods for each data type.

 - `SetString(attribute interface{}, value string) error`
 - `SetNumber(attribute interface{}, value int64) error`
 - `SetFloat(attribute interface{}, value float64) error`
 - `SetBoolean(attribute interface{}, value bool) error`
 - `SetDate(attribute interface{}, value time.Time) error`

For instance: 
```go
if err := ctx.SetString(usr.Status, "VIP"); err != nil {
    log(err)
}
```

!!! warning "Blocking method"
    The previous methods (`SetString, SetNumber, SetFloat, SetBoolean, SetDate`) are blocking methods that executes into a 
    transaction meaning that in case of error the new value is not applied.

#### Context into onActivation handler

The context into the activation handler contains all the previous registered facts, so all facts are accessible to read it or to write it.
Also, the parent object to each fact can be accessed via the context. 

The exposed API methods are described below:

**Fact Getters:**

These methods are useful to fetch a fact object given the fact name.

 - `Get(fact string) (interface{}, bool)`
 - `GetString(fact string) (String, error)`
 - `GetNumber(fact string) (Number, error)`
 - `GetFloat(fact string) (Float, error)`
 - `GetBoolean(fact string) (Boolean, error)`
 - `GetDate(fact string) (Date, error)`

**Iteration over all facts into the context**

This method is useful when you need to run an iteration over all facts into the context. An iteration function must be provided with
2 parameters: `fact` which is the fact name like `User.plan` and its value which in this case is an interface to match with all possible data types.

 - `ForEach(fn func(fact string, value interface{}))`
   
**Parent Object Getter**

This method fetch the parent object that contains a fact.

 - `GetObject(object string) (interface{}, bool)`

**Fact Inference (feedback)**    

Sometimes when we are working with a ruleset where some facts depend on other facts is needed an inference mechanism. In this case we called 
`Feedback` due to each feedback will trigger a ruleset re-evaluation. The `Feedback` function is transactional, thread-safe, and it is called
from the `context.Update` life cycle. If the `tx.Error()` is executed into the `Feedback method` the error will be accessible as outcome of `context.Update`.

 - `Feedback(func(tx *Tx))`

!!! example "Advanced example"
    Please check the advanced example app into the [goldfish-re](https://github.com/darksubmarine/goldfish-re) repo to see it in action!