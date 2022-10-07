package goldfish_re

import "time"

// dateConditionBuilder builder struct
type dateConditionBuilder struct {
	c *conditionBuilder
}

// newDateConditionBuilder internal constructor
func newDateConditionBuilder() *dateConditionBuilder {
	return &dateConditionBuilder{c: newConditionBuilder()}
}

// right sets the condition right term and operation
func (cb *dateConditionBuilder) right(term iTerm, op tOperator) *finalConditionBuilder {
	cb.c.Right(term)
	cb.c.Operation(op)
	return newFinalConditionBuilder(cb.c)
}

// Term sets the left date condition term
func (cb *dateConditionBuilder) Term(object, attribute string) *dateConditionBuilder {
	cb.c.Left(newDateVarTerm(object, attribute))
	return cb
}

// Equal sets the right term value and equal operator
func (cb *dateConditionBuilder) Equal(d time.Time) *finalConditionBuilder {
	return cb.right(newDiscreteDateTerm(d), opEquals)
}

// EqualTerm sets the right term and equal operator
func (cb *dateConditionBuilder) EqualTerm(object, attribute string) *finalConditionBuilder {
	return cb.right(newDateVarTerm(object, attribute), opEquals)
}

// After sets the right term value and the after operator
func (cb *dateConditionBuilder) After(d time.Time) *finalConditionBuilder {
	return cb.right(newDiscreteDateTerm(d), opAfter)
}

// AfterTerm sets the right term and the after operator
func (cb *dateConditionBuilder) AfterTerm(object, attribute string) *finalConditionBuilder {
	return cb.right(newDateVarTerm(object, attribute), opAfter)
}

// Before sets the right term value and the before operator
func (cb *dateConditionBuilder) Before(d time.Time) *finalConditionBuilder {
	return cb.right(newDiscreteDateTerm(d), opBefore)
}

// BeforeTerm sets the right term and the before operator
func (cb *dateConditionBuilder) BeforeTerm(object, attribute string) *finalConditionBuilder {
	return cb.right(newDateVarTerm(object, attribute), opBefore)
}

// Between sets the right term value and the between operator
func (cb *dateConditionBuilder) Between(start time.Time, end time.Time) *finalConditionBuilder {
	cb.c.Right(newDateListTerm([]time.Time{start, end}))
	cb.c.Operation(opBetween)
	return newFinalConditionBuilder(cb.c)
}

// Not negates condition
func (cb *dateConditionBuilder) Not() *dateConditionBuilder {
	cb.c.Not()
	return cb
}
