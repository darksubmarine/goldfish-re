package goldfish_re

// finalConditionBuilder internal builder to construct any condition type
type finalConditionBuilder struct {
	cond *conditionBuilder
}

// newFinalConditionBuilder constructor function
func newFinalConditionBuilder(cb *conditionBuilder) *finalConditionBuilder {
	return &finalConditionBuilder{cond: cb}
}

// Build _condition build method
func (fcb *finalConditionBuilder) Build() *_condition {
	if fcb.cond.negated {
		return newNegatedCondition(0, fcb.cond.left, fcb.cond.right, fcb.cond.op)
	}
	return newCondition(0, fcb.cond.left, fcb.cond.right, fcb.cond.op)
}

// conditionBuilder basic condition builder
type conditionBuilder struct {
	negated bool
	left    iTerm
	right   iTerm
	op      tOperator
}

// newConditionBuilder constructor of conditionBuilder
func newConditionBuilder() *conditionBuilder { return &conditionBuilder{} }

// Build _condition build method
func (cb *conditionBuilder) Build() *_condition {
	if cb.negated {
		return newNegatedCondition(0, cb.left, cb.right, cb.op)
	}
	return newCondition(0, cb.left, cb.right, cb.op)
}

// Left sets the condition left term
func (cb *conditionBuilder) Left(term iTerm) *conditionBuilder {
	cb.left = term
	return cb
}

// Right sets the condition right term
func (cb *conditionBuilder) Right(term iTerm) *conditionBuilder {
	cb.right = term
	return cb
}

// Operation sets the condition operator
func (cb *conditionBuilder) Operation(op tOperator) *conditionBuilder {
	cb.op = op
	return cb
}

// Not negates condition
func (cb *conditionBuilder) Not() *conditionBuilder {
	cb.negated = true
	return cb
}
