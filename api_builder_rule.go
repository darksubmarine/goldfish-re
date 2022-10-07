package goldfish_re

// ruleBuilder builder struct
type ruleBuilder struct {
	op         tBinaryOperator
	conditions []*_condition
	then       string
}

// newRuleBuilder ruleBuilder constructor
func newRuleBuilder() *ruleBuilder { return &ruleBuilder{conditions: []*_condition{}} }

// AllOf sets the rule operation as "All given conditions must be true to activate this rule"
func (rb *ruleBuilder) AllOf(conditions ...*_condition) *ruleBuilder {
	rb.op = opAnd
	rb.conditions = conditions
	return rb
}

// AnyOf sets the rule operation as "At least one given conditions must be true to activate this rule"
func (rb *ruleBuilder) AnyOf(conditions ...*_condition) *ruleBuilder {
	rb.op = opOr
	rb.conditions = conditions
	return rb
}

// Then value to return when the rule is activated
func (rb *ruleBuilder) Then(s string) *ruleBuilder {
	rb.then = s
	return rb
}

// Build rule builder method
func (rb *ruleBuilder) Build() (*_rule, error) {
	if rb.then == emptyStr {
		return nil, ErrEmptyThenSentence
	}

	if len(rb.conditions) == 0 {
		return nil, ErrEmptyConditionList
	}

	r := newRule(0, rb.op, rb.then)
	for i, c := range rb.conditions {
		c.id = cuid(i)
		if err := r.addCondition(c); err != nil {
			return nil, err
		}
	}
	return r, nil
}
