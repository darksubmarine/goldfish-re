package goldfish_re

// numberConditionRightBuilder right condition part builder struct
type numberConditionRightBuilder struct {
	_cb *numberConditionBuilder
}

// numberConditionBuilder condition builder
type numberConditionBuilder struct {
	c *conditionBuilder
}

// newNumberConditionBuilder constructor
func newNumberConditionBuilder() *numberConditionBuilder {
	return &numberConditionBuilder{c: newConditionBuilder()}
}

// Term sets the left condition term
func (cb *numberConditionBuilder) Term(object, attribute string) *numberConditionRightBuilder {
	cb.c.Left(newNumberVarTerm(object, attribute))
	return &numberConditionRightBuilder{_cb: cb}
}

// right sets the condition right term and operation
func (cb *numberConditionRightBuilder) right(term iTerm, op tOperator) *finalConditionBuilder {
	cb._cb.c.Operation(op)
	cb._cb.c.Right(term)
	return newFinalConditionBuilder(cb._cb.c)
}

// Equal sets the right term value and equal operator
func (cb *numberConditionRightBuilder) Equal(n int64) *finalConditionBuilder {
	return cb.right(newDiscreteNumberTerm(n), opEquals)
}

// EqualTerm sets the right term and equal operator
func (cb *numberConditionRightBuilder) EqualTerm(object, attribute string) *finalConditionBuilder {
	return cb.right(newNumberVarTerm(object, attribute), opEquals)
}

// GreaterThan sets the right term value and greater than operator
func (cb *numberConditionRightBuilder) GreaterThan(n int64) *finalConditionBuilder {
	return cb.right(newDiscreteNumberTerm(n), opGreaterThan)
}

// GreaterThanTerm sets the right term and greater than operator
func (cb *numberConditionRightBuilder) GreaterThanTerm(object, attribute string) *finalConditionBuilder {
	return cb.right(newNumberVarTerm(object, attribute), opGreaterThan)
}

// GreaterThanOrEqual sets the right term value and greater than or equal operator
func (cb *numberConditionRightBuilder) GreaterThanOrEqual(n int64) *finalConditionBuilder {
	return cb.right(newDiscreteNumberTerm(n), opGreaterThanOrEqual)
}

// GreaterThanOrEqualTerm sets the right term and greater than or equal operator
func (cb *numberConditionRightBuilder) GreaterThanOrEqualTerm(object, attribute string) *finalConditionBuilder {
	return cb.right(newNumberVarTerm(object, attribute), opGreaterThanOrEqual)
}

// LessThan sets the right term value and less than operator
func (cb *numberConditionRightBuilder) LessThan(n int64) *finalConditionBuilder {
	return cb.right(newDiscreteNumberTerm(n), opLessThan)
}

// LessThanTerm sets the right term and less than operator
func (cb *numberConditionRightBuilder) LessThanTerm(object, attribute string) *finalConditionBuilder {
	return cb.right(newNumberVarTerm(object, attribute), opLessThan)
}

// LessThanOrEqual sets the right term value and less than or equal operator
func (cb *numberConditionRightBuilder) LessThanOrEqual(n int64) *finalConditionBuilder {
	return cb.right(newDiscreteNumberTerm(n), opLessThanOrEqual)
}

// LessThanOrEqualTerm sets the right term and less than or equal operator
func (cb *numberConditionRightBuilder) LessThanOrEqualTerm(object, attribute string) *finalConditionBuilder {
	return cb.right(newNumberVarTerm(object, attribute), opLessThanOrEqual)
}

// Not negates the condition
func (cb *numberConditionBuilder) Not() *numberConditionBuilder {
	cb.c.Not()
	return cb
}
