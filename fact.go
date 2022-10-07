package goldfish_re

import (
	"fmt"
	"time"
)

// iFact fact interface
type iFact interface {
	object() string
	attribute() string
	value() interface{}

	valueDate() time.Time
	valueBoolean() bool
	valueString() string
	valueNumber() int64
	valueFloat() float64
	isString() bool
	isNumber() bool
	isFloat() bool
	isBoolean() bool
	isDate() bool
	token() string
	fullToken() string
}

// _fact implements iFact
type _fact struct {
	obj  string
	attr string
	val  interface{}
}

func newFact(object, attribute string, value interface{}) *_fact {
	return &_fact{obj: object, attr: attribute, val: value}
}

func newNumber(object, attribute string, value int64) *_fact {
	return newFact(object, attribute, value)
}

func newFloat(object, attribute string, value float64) *_fact {
	return newFact(object, attribute, value)
}

func newString(object, attribute string, value string) *_fact {
	return newFact(object, attribute, value)
}

func newBoolean(object, attribute string, value bool) *_fact {
	return newFact(object, attribute, value)
}

func newDate(object, attribute string, value time.Time) *_fact {
	return newFact(object, attribute, value)
}

func (f *_fact) object() string {
	return f.obj
}

func (f *_fact) attribute() string {
	return f.attr
}

func (f *_fact) value() interface{} {
	return f.val
}

func (f *_fact) token() string {
	return fmt.Sprintf("%s.%s", f.obj, f.attr)
}

func (f *_fact) fullToken() string {
	return fmt.Sprintf("%s.%s=%v", f.obj, f.attr, f.val)
}

func (f *_fact) valueDate() time.Time { return f.val.(time.Time) }
func (f *_fact) valueBoolean() bool   { return f.val.(bool) }
func (f *_fact) valueString() string  { return f.val.(string) }
func (f *_fact) valueNumber() int64   { return f.val.(int64) }
func (f *_fact) valueFloat() float64  { return f.val.(float64) }
func (f *_fact) isString() bool {
	if termType(f.val) == termString {
		return true
	}
	return false
}
func (f *_fact) isNumber() bool {
	if termType(f.val) == termNumber {
		return true
	}
	return false
}

func (f *_fact) isFloat() bool {
	if termType(f.val) == termFloat {
		return true
	}
	return false
}

func (f *_fact) isBoolean() bool {
	if termType(f.val) == termBoolean {
		return true
	}
	return false
}

func (f *_fact) isDate() bool {
	if termType(f.val) == termDate {
		return true
	}
	return false
}
