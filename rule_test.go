package goldfish_re

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_rule_addCondition(t *testing.T) {
	//R1 (all): User.plan == "gold" && Trip.miles > 1000 && User.miles > 300
	//R2 (any): Trip.miles > 2000 || Trip.miles > User.miles

	c1 := newCondition(1, newStringVarTerm("User", "plan"), newDiscreteTerm("gold"), opEquals)
	c2 := newCondition(2, newNumberVarTerm("Trip", "miles"), newDiscreteTerm(1000), opGreaterThan)
	c3 := newCondition(3, newNumberVarTerm("User", "miles"), newDiscreteTerm(300), opGreaterThan)

	c4 := newCondition(4, newNumberVarTerm("Trip", "miles"), newDiscreteTerm(2000), opGreaterThan)
	c5 := newCondition(5, newNumberVarTerm("Trip", "miles"), newNumberVarTerm("User", "miles"), opGreaterThan)

	r1 := newRule(1, opAnd, "apply")
	assert.Nil(t, r1.addCondition(c1))
	assert.Nil(t, r1.addCondition(c2))
	assert.Nil(t, r1.addCondition(c3))
	assert.Error(t, r1.addCondition(c3))

	r2 := newRule(2, opOr, "apply")
	assert.Nil(t, r2.addCondition(c4))
	assert.Nil(t, r2.addCondition(c5))
}
