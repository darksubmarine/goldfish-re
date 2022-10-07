package goldfish_re

import (
	"github.com/darksubmarine/goldfish-re/trie"
	"github.com/kelindar/bitmap"
	"sync"
	"sync/atomic"
)

const (
	defaultConditions = 100
	defaultRules      = 100
)

type _ruleset struct {
	mtx           sync.Mutex
	ctrRules      uint32
	ctrConditions uint32

	lruid ruid // atomic ID for rules
	lcuid cuid // atomic ID for conditions

	conditions []*_condition
	rules      []*_rule

	conditionRef map[string]*_condition
	idx          *trie.PathTrie
}

func newRuleset() *_ruleset {
	return &_ruleset{
		conditions:   make([]*_condition, defaultConditions),
		rules:        make([]*_rule, defaultRules),
		conditionRef: map[string]*_condition{},
		idx:          trie.NewPathTrie(),
	}
}

func (rs *_ruleset) lenr() int {
	return int(rs.ctrRules)
}

func (rs *_ruleset) lenc() int {
	return int(rs.ctrConditions)
}

func (rs *_ruleset) nextRuid() ruid {
	return atomic.AddUint32(&rs.lruid, 1)
}

func (rs *_ruleset) nextCuid() cuid {
	return atomic.AddUint32(&rs.lcuid, 1)
}

func (rs *_ruleset) addRule(rule *_rule) {
	// TODO check if given rule is not present before to add it
	rs.mtx.Lock()
	defer rs.mtx.Unlock()

	// cloning rule
	ruleToAdd := newRule(rs.nextRuid(), rule.operator, rule.then)

	for _, c := range rule.conditions {
		if cond, existsInRuleset := rs.conditionRef[c.token_]; existsInRuleset {

			ruleToAdd.addCondition(cond)
			cond.addRule(ruleToAdd) // cycle ref

		} else {
			condToAdd := c.cloneWithId(rs.nextCuid())
			ruleToAdd.addCondition(condToAdd)
			condToAdd.addRule(ruleToAdd) // cycle ref

			// check slice size and growth if needed
			if cuid(len(rs.conditions)) <= condToAdd.id {
				rs.conditions = growthSlice[*_condition](rs.conditions, defaultConditions)
			}

			// add new condition to ruleset
			rs.conditions[condToAdd.id] = condToAdd
			rs.conditionRef[condToAdd.token_] = condToAdd
			rs.ctrConditions++
		}
	}

	// check slice size and growth if needed
	if ruid(len(rs.rules)) <= ruleToAdd.id {
		rs.rules = growthSlice[*_rule](rs.rules, defaultRules)
	}

	rs.rules[ruleToAdd.id] = ruleToAdd
	rs.ctrRules++
}

func (rs *_ruleset) wme(fact iFact, ctx _factContext) {

	path := indexPath(fact.object(), fact.attribute(), fact.value())
	if node := rs.idx.Get(path); node != nil {
		return
	}

	var activeConditions = make([]*_condition, 0)
	betaNodes := make(map[cuid]_beta)
	for _, c := range rs.conditions {
		if c == nil {
			continue
		}
		if ok, relFact := c.eval(fact, ctx); ok {
			if relFact == nil {
				activeConditions = append(activeConditions, c)
			} else {
				if betaNodes[c.id] == nil {
					betaNodes[c.id] = _beta{}
				}
				betaNodes[c.id][indexPathFact(relFact)] = struct{}{}
			}
		}
	}

	rs.idx.Put(path, _alpha{active: activeConditions, rel: betaNodes})

	for cid, mm := range betaNodes {
		for relPath, _ := range mm {
			if node := rs.idx.Get(relPath); node != nil {
				if _, ok := node.(_alpha).rel[cid]; !ok {
					node.(_alpha).rel[cid] = _beta{}
				}
				node.(_alpha).rel[cid][path] = struct{}{}
			}
		}
	}
}

func (rs *_ruleset) evalFacts(ctx _factContext) []*_rule {
	activatedBm := bitmap.Bitmap{}
	partialActivation := []*_rule{}

	for _, fact := range ctx {
		path := indexPathFact(fact)
		node := rs.idx.Get(path)
		if node == nil { // if we don't have node yet.. just add it!
			rs.wme(fact, ctx)
			node = rs.idx.Get(path)
		}

		nAlpha := node.(_alpha)

		for _, cond := range nAlpha.active {
			if activatedBm.Contains(cond.id) {
				continue
			}
			activatedBm.Set(cond.id)
			partialActivation = append(partialActivation, cond.ruleSlice...) //TODO reduce partialActivation per rule ID
		}

		if len(nAlpha.rel) > 0 {
			factToken := fact.token()
			for condId, beta := range nAlpha.rel {
				if activatedBm.Contains(condId) {
					continue
				}

				for _, fCtx := range ctx {
					if factToken != fCtx.token() {
						if _, exists := beta[indexPathFact(fCtx)]; exists {
							cond := rs.conditions[condId]
							activatedBm.Set(condId)
							partialActivation = append(partialActivation, cond.ruleSlice...)
						}
					}
				}
			}
		}
	}

	_activeSlice := make([]*_rule, len(rs.rules))
	for _, r := range partialActivation {
		if _activeSlice[r.id] != nil {
			continue
		}

		if r.operator == opAnd {
			if r.matchAll(activatedBm) {
				_activeSlice[r.id] = r
			}
		} else {
			if r.matchAny(activatedBm) {
				_activeSlice[r.id] = r
			}
		}
	}

	return _activeSlice
}

func (rs *_ruleset) evalFact(fact iFact, ctx _factContext) []*_rule {
	activatedBm := bitmap.Bitmap{}
	partialActivation := []*_rule{}

	path := indexPathFact(fact)
	node := rs.idx.Get(path)
	if node != nil { // if we don't have node yet... just add it!
		rs.wme(fact, ctx)
	}

	nAlpha := node.(_alpha)

	for _, cond := range nAlpha.active {
		if activatedBm.Contains(cond.id) {
			continue
		}
		activatedBm.Set(cond.id)
		partialActivation = append(partialActivation, cond.ruleSlice...) //TODO reduce partialActivation per rule ID
	}

	if len(nAlpha.rel) > 0 {
		factToken := fact.token()
		for condId, beta := range nAlpha.rel {
			if activatedBm.Contains(condId) {
				continue
			}

			for _, fCtx := range ctx {
				if factToken != fCtx.token() {
					if _, exists := beta[indexPathFact(fCtx)]; exists {
						cond := rs.conditions[condId]
						activatedBm.Set(condId)
						partialActivation = append(partialActivation, cond.ruleSlice...)
					}
				}
			}
		}
	}

	allActiveRules := rs.evalFacts(ctx)

	_activeSlice := make([]*_rule, len(rs.rules))

	for _, r := range partialActivation {
		if _activeSlice[r.id] != nil {
			continue
		}

		for _, v := range allActiveRules {
			if v == nil {
				continue
			}

			if r.id == v.id {
				_activeSlice[r.id] = r
			}

		}
	}

	return _activeSlice
}
