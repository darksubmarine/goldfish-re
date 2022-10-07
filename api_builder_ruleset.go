package goldfish_re

// rulesetBuilder ruleset build object
type rulesetBuilder struct {
	successFn func(string, Context)
	errorFn   func(error)
}

// OnActivation sets the user function to call when a rule is activated
func (rb *rulesetBuilder) OnActivation(fn func(string, Context)) *rulesetBuilder {
	rb.successFn = fn
	return rb
}

// OnError sets the user error handler to call when a ruleset evaluation runs an error
func (rb *rulesetBuilder) OnError(fn func(err error)) *rulesetBuilder {
	rb.errorFn = fn
	return rb
}

// Build rulset build method.
// Panic if some one of required user handlers are not provided
func (rb *rulesetBuilder) Build() *ruleset {
	if rb.successFn == nil || rb.errorFn == nil {
		panic("Success function and Error function must be provided")
	}
	return &ruleset{rs: newRuleset(), successFn: rb.successFn, errorFn: rb.errorFn}
}

// newRulesetBuilder rulesetBuilder constructor function
func newRulesetBuilder() *rulesetBuilder {
	return &rulesetBuilder{}
}
