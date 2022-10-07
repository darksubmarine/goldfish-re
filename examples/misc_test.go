package examples

import (
	"fmt"
	gre "github.com/darksubmarine/goldfish-re"
	"testing"
	"time"
)

func createRuleset() gre.Ruleset {
	return gre.Builder().Ruleset().Build()
}

func TestPoc(t *testing.T) {

	// C1: User.Plan == "gold"
	c1 := gre.Builder().StringCondition().
		Term("User", "plan").
		Equal("gold").
		Build()

	// C2: User.miles > 300
	c2 := gre.Builder().NumberCondition().Term("User", "miles").GreaterThan(300).Build()

	// C3: User.status in ["active", "toReview", "VIP"]
	c3 := gre.Builder().StringCondition().
		Term("User", "status").
		In([]string{"active", "toReview", "VIP"}).
		Build()

	rs := gre.Builder().Ruleset().
		OnActivation(func(s string, ctx gre.Context) {
			fmt.Println("Evaluation rule: ", s)

			// context iteration
			ctx.ForEach(func(fact string, value interface{}) {
				fmt.Printf("%s=%v\n", fact, value)
			})

			// data type assertion
			if usrPlan, ok := ctx.Get("User.Plan"); ok {
				switch pln := usrPlan.(type) {
				case gre.String:
					fmt.Println("USER PLAN (string):", pln.Value())
				case gre.Number:
					fmt.Println("USER PLAN (number):", pln.Value())
				}
			}

		}).
		OnError(func(err error) {
			t.Error(err)
		}).
		Build()

	r1, _ := gre.Builder().Rule().AllOf(c1, c2, c3).Then("ACTIVATED").Build()
	rs.AddRule(r1)

	ctx := rs.Context()

	type child struct {
		Name gre.String `gre:"value=Steve"`
	}
	type User struct {
		Plan   gre.String `gre:"attribute=plan,value=gold"`
		Age    int
		Num    gre.Number `gre:"object=Usuario,value=789"` //if the name is not exported .. the tag has not effect. Must be initialized manually
		Leaf   *child
		Miles  gre.Number `gre:"attribute=miles,value=500"`
		Status gre.String `gre:"attribute=status,value=none"`
	}

	usr := User{}
	if err := ctx.Register(&usr); err != nil {
		t.Error(err)
	}

	start := time.Now()
	if err := ctx.SetString(usr.Status, "VIP"); err != nil {
		t.Log(err)
	}
	end := time.Now().Sub(start)
	fmt.Println("SET status= VIP:", end)

	start = time.Now()
	if err := ctx.Update(func(tx *gre.Tx) {
		tx.SetString(usr.Plan, "gold")
		tx.SetString(usr.Status, "active")
		tx.SetNumber(usr.Miles, 999)
		//tx.Error(errors.New("something wrong"))
	}); err != nil {
		t.Log(err)
	}
	end2 := time.Now().Sub(start)
	fmt.Println("SET:", end2)
}

func TestBoolean(t *testing.T) {

	// C1: User.active == true
	c1 := gre.Builder().BooleanCondition().Term("User", "active").Equal(true).Build()

	rs := gre.Builder().Ruleset().
		OnActivation(func(s string, ctx gre.Context) {
			fmt.Println("Evaluation rule: ", s)

			ctx.ForEach(func(fact string, value interface{}) {
				fmt.Printf("%s=%v\n", fact, value)
			})
		}).
		OnError(func(err error) {
			t.Error(err)
		}).
		Build()

	r1, _ := gre.Builder().Rule().AllOf(c1).Then("ACTIVE!!!").Build()
	rs.AddRule(r1)

	ctx := rs.Context()

	type User struct {
		Active gre.Boolean `gre:"attribute=active,value=true"`
	}

	usr := new(User)
	if err := ctx.Register(usr); err != nil {
		t.Error(err)
	}

	if err := ctx.SetBoolean(usr.Active, true); err != nil {
		t.Log(err)
	}
}

func TestFloat(t *testing.T) {

	// C1: User.nanos >= 123.22
	c1 := gre.Builder().FloatCondition().Term("User", "nanos").GreaterThan(123.22).Build()

	rs := gre.Builder().Ruleset().
		OnActivation(func(s string, ctx gre.Context) {
			fmt.Println("Evaluation rule: ", s)

			ctx.ForEach(func(fact string, value interface{}) {
				fmt.Printf("%s=%v\n", fact, value)
			})
		}).
		OnError(func(err error) {
			t.Error(err)
		}).
		Build()

	r1, _ := gre.Builder().Rule().AllOf(c1).Then("ACTIVE!!!").Build()
	rs.AddRule(r1)

	ctx := rs.Context()

	type User struct {
		Nanos gre.Float `gre:"attribute=nanos,value=98.2"`
	}

	usr := new(User)
	if err := ctx.Register(usr); err != nil {
		t.Error(err)
	}

	if err := ctx.SetFloat(usr.Nanos, 123.221); err != nil {
		t.Log(err)
	}
}

func TestDate(t *testing.T) {

	type User struct {
		Birthday gre.Date `gre:"attribute=birthday,value=1985-12-24T00:00:00"`
		Comp     gre.Date `gre:"value=1984-12-24T00:00:00"`
	}

	// C1: User.birthday before Year(1984)
	c1 := gre.Builder().DateCondition().
		Term("User", "birthday").
		//Before(gre.YearUTC(1984)).
		BeforeTerm("User", "comp").
		//Between(gre.CalendarDateUTC(1984, 3, 2), gre.CalendarDateUTC(1990, 12, 33)).
		Build()

	rs := gre.Builder().Ruleset().
		OnActivation(func(s string, ctx gre.Context) {
			fmt.Println("Evaluation rule: ", s)

			ctx.ForEach(func(fact string, value interface{}) {
				fmt.Printf("%s=%v\n", fact, value)
			})
		}).
		OnError(func(err error) {
			t.Error(err)
		}).
		Build()

	r1, _ := gre.Builder().Rule().AllOf(c1).Then("ACTIVE!!!").Build()
	rs.AddRule(r1)

	ctx := rs.Context()

	usr := new(User)
	if err := ctx.Register(usr); err != nil {
		t.Error(err)
	}

	if err := ctx.SetDate(usr.Birthday, gre.CalendarDateUTC(1988, 8, 14)); err != nil {
		t.Log(err)
	}

	if err := ctx.SetDate(usr.Comp, gre.CalendarDateUTC(1989, 8, 14)); err != nil {
		t.Log(err)
	}
}

func TestIndividualRegisters(t *testing.T) {

	c1 := gre.Builder().StringCondition().Term("User", "plan").Equal("gold").Build()
	c2 := gre.Builder().FloatCondition().Term("User", "nanos").GreaterThan(123.22).Build()
	c3 := gre.Builder().NumberCondition().Term("User", "age").GreaterThan(18).Build()
	c4 := gre.Builder().BooleanCondition().Term("User", "active").Equal(true).Build()
	c5 := gre.Builder().DateCondition().Term("User", "birthday").After(gre.YearUTC(1984)).Build()

	rs := gre.Builder().Ruleset().
		OnActivation(func(s string, ctx gre.Context) {
			fmt.Println("Evaluation rule: ", s)

			if usrPlan, ok := ctx.Get("User.plan"); ok {
				fmt.Println(usrPlan.(gre.String).Value())
			}

			ctx.ForEach(func(fact string, value interface{}) {
				fmt.Printf("%s=%v\n", fact, value)
			})
		}).
		OnError(func(err error) {
			t.Error(err)
		}).
		Build()

	r1, _ := gre.Builder().Rule().AllOf(c1, c2, c3, c4, c5).Then("ACTIVE!!!").Build()
	rs.AddRule(r1)

	ctx := rs.Context()

	type User struct {
		plan     gre.String
		nanos    gre.Float `gre:"attribute=nanos,value=98.2"`
		age      gre.Number
		active   gre.Boolean
		birthday gre.Date
	}

	usr := &User{
		plan:     gre.NewString("User", "plan", "gold"),
		nanos:    gre.NewFloat("User", "nanos", 90.2),
		age:      gre.NewNumber("User", "age", 22),
		active:   gre.NewBoolean("User", "active", false),
		birthday: gre.NewDate("User", "birthday", gre.CalendarDateUTC(1980, 10, 23)),
	}

	if err := ctx.RegisterString(usr, usr.plan); err != nil {
		t.Error(err)
	}

	if err := ctx.RegisterFloat(usr, usr.nanos); err != nil {
		t.Error(err)
	}

	if err := ctx.RegisterNumber(usr, usr.age); err != nil {
		t.Error(err)
	}

	if err := ctx.RegisterBoolean(usr, usr.active); err != nil {
		t.Error(err)
	}

	if err := ctx.RegisterDate(usr, usr.birthday); err != nil {
		t.Error(err)
	}

	if err := ctx.Update(func(tx *gre.Tx) {
		tx.SetString(usr.plan, "gold")
		tx.SetFloat(usr.nanos, 124.34)
		tx.SetNumber(usr.age, 46)
		tx.SetBoolean(usr.active, true)
		tx.SetDate(usr.birthday, gre.CalendarDateUTC(1990, 9, 20))
	}); err != nil {
		t.Error(err)
	}
}
