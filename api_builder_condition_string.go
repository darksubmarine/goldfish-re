package goldfish_re

// stringConditionRightBuilder string condition right term builder
type stringConditionRightBuilder struct {
	scb *stringConditionBuilder
}

// stringConditionBuilder string condition builder
type stringConditionBuilder struct {
	c *conditionBuilder
}

// newStringConditionBuilder constructor function
func newStringConditionBuilder() *stringConditionBuilder {
	return &stringConditionBuilder{c: newConditionBuilder()}
}

// Term sets the left string condition term
func (cb *stringConditionBuilder) Term(object, attribute string) *stringConditionRightBuilder {
	cb.c.Left(newStringVarTerm(object, attribute))
	return &stringConditionRightBuilder{scb: cb}
}

// Not negates the parent condition
func (cb *stringConditionBuilder) Not() *stringConditionBuilder {
	cb.c.Not()
	return cb
}

// right sets the condition right term and operation
func (cb *stringConditionRightBuilder) right(term iTerm, op tOperator) *finalConditionBuilder {
	cb.scb.c.Right(term)
	cb.scb.c.Operation(op)
	return newFinalConditionBuilder(cb.scb.c)
}

// Equal sets the right term string value and equal operator
func (cb *stringConditionRightBuilder) Equal(str string) *finalConditionBuilder {
	return cb.right(newDiscreteStringTerm(str), opEquals)
}

// EqualTerm sets the right term and equal operator
func (cb *stringConditionRightBuilder) EqualTerm(object, attribute string) *finalConditionBuilder {
	return cb.right(newStringVarTerm(object, attribute), opEquals)
}

// In sets the right term as a list of string and the IN operator
func (cb *stringConditionRightBuilder) In(list []string) *finalConditionBuilder {
	return cb.right(newStringListTerm(list), opIn)
}

// Contains sets the right term as a value and the contains operation
func (cb *stringConditionRightBuilder) Contains(str string) *finalConditionBuilder {
	return cb.right(newDiscreteStringTerm(str), opContains)
}

// ContainsTerm sets the right term and the contains operator
func (cb *stringConditionRightBuilder) ContainsTerm(object, attribute string) *finalConditionBuilder {
	return cb.right(newStringVarTerm(object, attribute), opContains)
}

// Starts sets the right term value and the starts operation
func (cb *stringConditionRightBuilder) Starts(str string) *finalConditionBuilder {
	return cb.right(newDiscreteStringTerm(str), opStarts)
}

// StartsTerm sets the right term and the starts operation
func (cb *stringConditionRightBuilder) StartsTerm(object, attribute string) *finalConditionBuilder {
	return cb.right(newStringVarTerm(object, attribute), opStarts)
}

// Ends sets the right term value and the ends operation
func (cb *stringConditionRightBuilder) Ends(str string) *finalConditionBuilder {
	return cb.right(newDiscreteStringTerm(str), opEnds)
}

// EndsTerm sets the right term and the ends operation
func (cb *stringConditionRightBuilder) EndsTerm(object, attribute string) *finalConditionBuilder {
	return cb.right(newStringVarTerm(object, attribute), opEnds)
}
