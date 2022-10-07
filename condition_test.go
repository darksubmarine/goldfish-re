package goldfish_re

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_condition_token(t *testing.T) {
	left := newStringVarTerm("User", "plan")
	right := newDiscreteTerm("gold")
	c := newCondition(1, left, right, opEquals)
	assert.EqualValues(t, "User.plan_==_gold", c.token())
}

func Test_condition_eval(t *testing.T) {
	left := newStringVarTerm("User", "plan")
	right := newDiscreteTerm("gold")
	c := newCondition(1, left, right, opEquals)

	fact := newFact("User", "plan", "silver")
	val, rf := c.eval(fact, _factContext{})
	assert.False(t, val)
	assert.Nil(t, rf)

	fact.val = "gold"
	val, rf = c.eval(fact, _factContext{})
	assert.True(t, val)
}

func TestXOR(t *testing.T) {

	negated := false
	eval := false
	assert.False(t, negated != eval)

	negated = false
	eval = true
	assert.True(t, negated != eval)

	negated = true
	eval = false
	assert.True(t, negated != eval)

	negated = true
	eval = true
	assert.False(t, negated != eval)
}

func Test_condition_negated_eval(t *testing.T) {
	left := newStringVarTerm("User", "plan")
	right := newDiscreteTerm("gold")
	c := newNegatedCondition(1, left, right, opEquals)

	fact := newFact("User", "plan", "silver")
	val, rf := c.eval(fact, _factContext{})
	assert.True(t, val)
	assert.Nil(t, rf)

	fact.val = "gold"
	val, rf = c.eval(fact, _factContext{})
	assert.False(t, val)
}
