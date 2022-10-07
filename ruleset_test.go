package goldfish_re

import (
	"github.com/stretchr/testify/assert"
	"math"
	"sync"
	"testing"
)

func Test_ruleset_nextRuid(t *testing.T) {
	mruid := sync.Map{}
	mcuid := sync.Map{}
	rs := newRuleset()
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				lruid := rs.nextRuid()
				lcuid := rs.nextCuid()

				if vruid, ok := mruid.Load(lruid); ok {
					vruid = vruid.(int) + 1
					mruid.Store(lruid, vruid)
				} else {
					mruid.Store(lruid, 1)
				}

				if vcuid, ok := mcuid.Load(lcuid); ok {
					vcuid = vcuid.(int) + 1
					mcuid.Store(lcuid, vcuid)
				} else {
					mcuid.Store(lcuid, 1)
				}

			}
			wg.Done()
		}()
	}
	wg.Wait()
	assert.EqualValues(t, int64(50001), int64(rs.nextRuid()))
	assert.EqualValues(t, int64(50001), int64(rs.nextCuid()))

	mruid.Range(func(key, value interface{}) bool {
		if v, ok := value.(int); ok {
			assert.EqualValues(t, 1, v)
		} else {
			assert.Fail(t, "invalid data type")
		}
		return true
	})

	mcuid.Range(func(key, value interface{}) bool {
		if v, ok := value.(int); ok {
			assert.EqualValues(t, 1, v)
		} else {
			assert.Fail(t, "invalid data type")
		}
		return true
	})
}

func Test_ruleset_addRule(t *testing.T) {

	rs := newRuleset()

	totalRules := 150
	expLenr := int(math.Ceil(float64(totalRules)/float64(defaultRules))) * defaultRules
	expLenc := 3

	for i := 0; i < totalRules; i++ {

		r1 := newRule(1, opAnd, "apply")

		c1 := newCondition(1, newStringVarTerm("User", "plan"), newDiscreteTerm("gold"), opEquals)
		c2 := newCondition(2, newNumberVarTerm("Trip", "miles"), newDiscreteTerm(1000), opGreaterThan)
		c3 := newCondition(3, newNumberVarTerm("User", "miles"), newDiscreteTerm(300), opGreaterThan)

		assert.Nil(t, r1.addCondition(c1))
		assert.Nil(t, r1.addCondition(c2))
		assert.Nil(t, r1.addCondition(c3))

		rs.addRule(r1)
	}

	assert.EqualValues(t, totalRules, rs.lenr())
	assert.EqualValues(t, expLenr, len(rs.rules))
	assert.EqualValues(t, expLenc, len(rs.conditionRef))
	assert.EqualValues(t, expLenc, rs.lenc())
}

func Test_ruleset_wme(t *testing.T) {

	rs := newRuleset()

	r1 := newRule(1, opAnd, "apply")

	c1 := newCondition(1, newStringVarTerm("User", "plan"), newDiscreteStringTerm("gold"), opEquals)
	c2 := newCondition(2, newNumberVarTerm("Trip", "miles"), newDiscreteNumberTerm(1000), opGreaterThan)
	c3 := newCondition(3, newNumberVarTerm("User", "miles"), newDiscreteNumberTerm(300), opGreaterThan)

	c4 := newCondition(4, newNumberVarTerm("Trip", "miles"), newNumberVarTerm("User", "miles"), opGreaterThan)

	assert.Nil(t, r1.addCondition(c1))
	assert.Nil(t, r1.addCondition(c2))
	assert.Nil(t, r1.addCondition(c3))
	assert.Nil(t, r1.addCondition(c4))

	rs.addRule(r1)

	factCtx := _factContext{}

	fact := newString("User", "plan", "gold")
	factCtx.set(fact)
	rs.wme(fact, factCtx)
	assert.NotNil(t, rs.idx.Get("/User/plan/gold"))

	fact2 := newNumber("User", "miles", 500)
	factCtx.set(fact2)
	rs.wme(fact2, factCtx)
	assert.NotNil(t, rs.idx.Get("/User/miles/500"))

	fact3 := newNumber("Trip", "miles", 1300)
	factCtx.set(fact3)
	rs.wme(fact3, factCtx)
	assert.NotNil(t, rs.idx.Get("/Trip/miles/1300"))

	fact4 := newNumber("Trip", "miles", 1600)
	factCtx.set(fact4)
	rs.wme(fact4, factCtx)
	assert.NotNil(t, rs.idx.Get("/Trip/miles/1600"))
}

func Test_ruleset_evalFact(t *testing.T) {

	rs := newRuleset()

	//R1 (all): User.plan == "gold" && Trip.miles > 1000 && User.miles > 300
	r1 := newRule(1, opAnd, "apply")

	//R2 (all): Trip.miles > User.miles
	r2 := newRule(2, opAnd, "apply")

	c1 := newCondition(1, newStringVarTerm("User", "plan"), newDiscreteStringTerm("gold"), opEquals)
	//c2 := newCondition(2, newNumberVarTerm("Trip", "miles"), newDiscreteNumberTerm(1000), opGreaterThan)
	c3 := newCondition(3, newNumberVarTerm("User", "miles"), newDiscreteNumberTerm(300), opGreaterThan)

	assert.Nil(t, r1.addCondition(c1))
	//assert.Nil(t, r1.addCondition(c2))
	assert.Nil(t, r1.addCondition(c3))

	c4 := newCondition(4, newNumberVarTerm("Trip", "miles"), newNumberVarTerm("User", "miles"), opGreaterThan)

	assert.Nil(t, r2.addCondition(c4))

	rs.addRule(r1)
	rs.addRule(r2)

	factCtx := _factContext{}

	fact := newString("User", "plan", "gold")
	factCtx.set(fact)
	rs.wme(fact, factCtx)
	assert.NotNil(t, rs.idx.Get("/User/plan/gold"))

	fact2 := newNumber("User", "miles", 500)
	factCtx.set(fact2)
	rs.wme(fact2, factCtx)
	assert.NotNil(t, rs.idx.Get("/User/miles/500"))

	fact3 := newNumber("Trip", "miles", 1300)
	factCtx.set(fact3)
	rs.wme(fact3, factCtx)
	assert.NotNil(t, rs.idx.Get("/Trip/miles/1300"))

	fact4 := newNumber("Trip", "miles", 1600)
	factCtx.set(fact4)
	rs.wme(fact4, factCtx)
	assert.NotNil(t, rs.idx.Get("/Trip/miles/1600"))

	factToEval := newNumber("Trip", "miles", 1200)
	factCtx.set(factToEval)
	rs.wme(factToEval, factCtx)

	rules := rs.evalFact(factToEval, factCtx)
	//rules := rs.evalFacts(factCtx)
	var activated int
	for _, r := range rules {
		if r != nil {
			activated++
		}
	}

	assert.EqualValues(t, 1, activated)
}
