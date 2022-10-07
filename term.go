package goldfish_re

import (
	"fmt"
	"time"
)

type tTerm uint8

const (
	termInvalid tTerm = iota
	termString
	termNumber
	termFloat
	termBoolean
	termDate
)

type iTerm interface {
	token() string
	path() string
	object() string
	attribute() string
	val() interface{}
}

type _term struct {
	isVar      bool
	object_    string
	attribute_ string
	value      interface{}
	kind       tTerm
}

func newTerm(isVar bool, object, attribute string, value interface{}, kind tTerm) *_term {
	return &_term{isVar: isVar, object_: object, attribute_: attribute, value: value, kind: kind}
}

func newDiscreteTerm(value interface{}) *_term {
	kind := termType(value)
	if kind == termInvalid {
		return nil
	}
	return newTerm(false, "", "", value, kind)
}

func newDiscreteBooleanTerm(value bool) *_term {
	return newDiscreteTerm(value)
}

func newDiscreteNumberTerm(value int64) *_term {
	return newDiscreteTerm(value)
}

func newDiscreteFloatTerm(value float64) *_term {
	return newDiscreteTerm(value)
}

func newDiscreteDateTerm(value time.Time) *_term {
	return newDiscreteTerm(value)
}

func newDateListTerm(value []time.Time) *_term {
	return newDiscreteTerm(value)
}

func newDiscreteStringTerm(value string) *_term {
	return newDiscreteTerm(value)
}

func newStringListTerm(value []string) *_term {
	return newDiscreteTerm(value)
}

func newVarTerm(object, attribute string, kind tTerm) *_term {
	return newTerm(true, object, attribute, nil, kind)
}

// --- Number

func newNumberVarTermWithValue(object, attribute string, value int64) *_term {
	return newTerm(true, object, attribute, value, termNumber)
}

func newNumberVarTerm(object, attribute string) *_term {
	return newNumberVarTermWithValue(object, attribute, 0)
}

// --- Float

func newFloatVarTermWithValue(object, attribute string, value float64) *_term {
	return newTerm(true, object, attribute, value, termFloat)
}

func newFloatVarTerm(object, attribute string) *_term {
	return newFloatVarTermWithValue(object, attribute, 0.0)
}

// --- String

func newStringVarTermWithValue(object, attribute string, value string) *_term {
	return newTerm(true, object, attribute, value, termString)
}

func newStringVarTerm(object, attribute string) *_term {
	return newStringVarTermWithValue(object, attribute, "")
}

// --- boolean

func newBooleanVarTerm(object, attribute string) *_term {
	return newBooleanVarTermWithValue(object, attribute, false)
}

func newBooleanVarTermWithValue(object, attribute string, value bool) *_term {
	return newTerm(true, object, attribute, value, termBoolean)
}

// --- date

func newDateVarTerm(object, attribute string) *_term {
	return newDateVarTermWithValue(object, attribute, zeroDate)
}

func newDateVarTermWithValue(object, attribute string, value time.Time) *_term {
	return newTerm(true, object, attribute, value, termDate)
}

func (t *_term) token() string {
	if t.isVar {
		return fmt.Sprintf("%s.%s", t.object_, t.attribute_)
	}

	return fmt.Sprintf("%v", t.value)
}

func (t *_term) path() string {
	if t.isVar {
		return fmt.Sprintf("/%s/%s/%v", t.object_, t.attribute_, t.value)
	}

	return fmt.Sprintf("%v", t.value)
}

func (t *_term) object() string    { return t.object_ }
func (t *_term) attribute() string { return t.attribute_ }
func (t *_term) val() interface{}  { return t.value }
