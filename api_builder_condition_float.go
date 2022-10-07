package goldfish_re

// floatConditionBuilder builder struct
type floatConditionBuilder struct {
	c *conditionBuilder
}

// newFloatConditionBuilder constructor function
func newFloatConditionBuilder() *floatConditionBuilder {
	return &floatConditionBuilder{c: newConditionBuilder()}
}

// right sets the condition right term and operation
func (cb *floatConditionBuilder) right(term iTerm, op tOperator) *finalConditionBuilder {
	cb.c.Operation(op)
	cb.c.Right(term)
	return newFinalConditionBuilder(cb.c)
}

// Term sets the left condition term
func (cb *floatConditionBuilder) Term(object, attribute string) *floatConditionBuilder {
	cb.c.Left(newFloatVarTerm(object, attribute))
	return cb
}

// Equal sets the right term value and equal operator
func (cb *floatConditionBuilder) Equal(n float64) *finalConditionBuilder {
	return cb.right(newDiscreteFloatTerm(n), opEquals)
}

// EqualTerm sets the right term and equal operator
func (cb *floatConditionBuilder) EqualTerm(object, attribute string) *finalConditionBuilder {
	return cb.right(newFloatVarTerm(object, attribute), opEquals)
}

// GreaterThan sets the right term value and greater than operator
func (cb *floatConditionBuilder) GreaterThan(n float64) *finalConditionBuilder {
	return cb.right(newDiscreteFloatTerm(n), opGreaterThan)
}

// GreaterThanTerm sets the right term and greater than operator
func (cb *floatConditionBuilder) GreaterThanTerm(object, attribute string) *finalConditionBuilder {
	return cb.right(newFloatVarTerm(object, attribute), opGreaterThan)
}

// GreaterThanOrEqual sets the right term value and greater than or equal operator
func (cb *floatConditionBuilder) GreaterThanOrEqual(n float64) *finalConditionBuilder {
	return cb.right(newDiscreteFloatTerm(n), opGreaterThanOrEqual)
}

// GreaterThanOrEqualTerm sets the right term and greater than or equal operator
func (cb *floatConditionBuilder) GreaterThanOrEqualTerm(object, attribute string) *finalConditionBuilder {
	return cb.right(newFloatVarTerm(object, attribute), opGreaterThanOrEqual)
}

// LessThan sets the right term value and less than operator
func (cb *floatConditionBuilder) LessThan(n float64) *finalConditionBuilder {
	return cb.right(newDiscreteFloatTerm(n), opLessThan)
}

// LessThanTerm sets the right term and less than operator
func (cb *floatConditionBuilder) LessThanTerm(object, attribute string) *finalConditionBuilder {
	return cb.right(newFloatVarTerm(object, attribute), opLessThan)
}

// LessThanOrEqual sets the right term value and less than or equal operator
func (cb *floatConditionBuilder) LessThanOrEqual(n float64) *finalConditionBuilder {
	return cb.right(newDiscreteFloatTerm(n), opLessThanOrEqual)
}

// LessThanOrEqualTerm sets the right term and less than or equal operator
func (cb *floatConditionBuilder) LessThanOrEqualTerm(object, attribute string) *finalConditionBuilder {
	return cb.right(newFloatVarTerm(object, attribute), opLessThanOrEqual)
}

// Not negates the condition
func (cb *floatConditionBuilder) Not() *floatConditionBuilder {
	cb.c.Not()
	return cb
}
