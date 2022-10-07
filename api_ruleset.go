package goldfish_re

import "sync"

// Ruleset interface to expose available actions to do with a ruleset
type Ruleset interface {
	AddRule(r *_rule)
	Context() *factContext
	//EvalFacts(ctx *factContext)
}

// ruleset wrapper to export methods
type ruleset struct {
	mtx       sync.Mutex
	rs        *_ruleset
	successFn func(string, Context)
	errorFn   func(error)
}

// AddRule adds a new rule to the ruleset.
func (rs *ruleset) AddRule(r *_rule) {
	rs.rs.addRule(r)
}

// Context returns a new fact context with the ruleset attached.
// Each time that a context.Update is called, the evaluation will be over this ruleset.
func (rs *ruleset) Context() *factContext {
	return newContext(rs)
}

// evalFacts thread-safe ruleset evaluation with the given context.
func (rs *ruleset) evalFacts(ctx *factContext) {
	rs.mtx.Lock()
	defer rs.mtx.Unlock()

	activated := rs.rs.evalFacts(ctx.iFactRef)
	for _, r := range activated {
		if r != nil {
			rs.successFn(r.then, ctx)
		}
	}
}

// evalFactsWithSkip thread-safe ruleset evaluation with the given context.
func (rs *ruleset) evalFactsWithSkip(ctx *factContext, skip map[string]struct{}) map[string]struct{} {
	rs.mtx.Lock()
	defer rs.mtx.Unlock()

	toSkip := map[string]struct{}{}
	activated := rs.rs.evalFacts(ctx.iFactRef)
	for _, r := range activated {
		if r != nil {
			if _, skipped := skip[r.then]; skipped {
				continue
			}
			toSkip[r.then] = struct{}{}
			rs.successFn(r.then, ctx)
		}
	}
	return toSkip
}
