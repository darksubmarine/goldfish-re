package goldfish_re

import "github.com/kelindar/bitmap"

const (
	all = true
	any = false
)

type _rule struct {
	id         ruid
	token      string
	operator   tBinaryOperator
	conditions map[cuid]*_condition
	condBitmap *bitmap.Bitmap

	then string
}

func newRule(id ruid, operator tBinaryOperator, then string) *_rule {
	return &_rule{id: id, operator: operator, conditions: map[cuid]*_condition{}, condBitmap: &bitmap.Bitmap{}, then: then}
}

func (r *_rule) matchAll(bm bitmap.Bitmap) bool {
	rMem := &bitmap.Bitmap{}
	r.condBitmap.Clone(rMem)
	rMem.And(bm)
	if rMem.Count() == r.condBitmap.Count() {
		return true
	}

	return false
}

func (r *_rule) matchAny(bm bitmap.Bitmap) bool {
	rMem := &bitmap.Bitmap{}
	r.condBitmap.Clone(rMem)
	rMem.And(bm)
	if rMem.Count() > 0 {
		return true
	}

	return false
}

func (r *_rule) addCondition(c *_condition) error {
	if _, exists := r.conditions[c.id]; exists {
		return errConditionAddedPreviously
	}

	r.conditions[c.id] = c
	r.condBitmap.Set(c.id)
	return noErr
}
