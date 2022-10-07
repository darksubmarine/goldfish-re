package goldfish_re

// booleanConditionBuilder boolean condition builder struct
type booleanConditionBuilder struct {
	c *conditionBuilder
}

// newBooleanConditionBuilder constructor
func newBooleanConditionBuilder() *booleanConditionBuilder {
	return &booleanConditionBuilder{c: newConditionBuilder()}
}

// Term sets the left boolean condition term
func (cb *booleanConditionBuilder) Term(object, attribute string) *booleanConditionBuilder {
	cb.c.Left(newBooleanVarTerm(object, attribute))
	return cb
}

// right sets the condition right term and operation
func (cb *booleanConditionBuilder) right(term iTerm, op tOperator) *finalConditionBuilder {
	cb.c.Operation(op)
	cb.c.Right(term)
	return newFinalConditionBuilder(cb.c)
}

// Equal sets the right term value and equal operator
func (cb *booleanConditionBuilder) Equal(val bool) *finalConditionBuilder {
	return cb.right(newDiscreteBooleanTerm(val), opEquals)
}

// EqualTerm sets the right term and equal operator
func (cb *booleanConditionBuilder) EqualTerm(object, attribute string) *finalConditionBuilder {
	return cb.right(newBooleanVarTerm(object, attribute), opEquals)
}

// IsTrue sets the right term value as true and equal operator
func (cb *booleanConditionBuilder) IsTrue() *finalConditionBuilder {
	return cb.right(newDiscreteBooleanTerm(true), opEquals)
}

// IsFalse sets the right term value as false and equal operator
func (cb *booleanConditionBuilder) IsFalse() *finalConditionBuilder {
	return cb.right(newDiscreteBooleanTerm(false), opEquals)
}

// Not negates condition
func (cb *booleanConditionBuilder) Not() *booleanConditionBuilder {
	cb.c.Not()
	return cb
}
