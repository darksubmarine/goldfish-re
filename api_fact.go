package goldfish_re

import (
	"sync"
	"time"
)

// syncFact Synchronous wrapper to work with a _fact struct
type syncFact struct {
	mt   sync.Mutex
	fact *_fact
}

// token calls the fact token
func (f *syncFact) token() string {
	return f.fact.token()
}

// set the fact val locking it
func (f *syncFact) set(v interface{}) {
	f.mt.Lock()
	defer f.mt.Unlock()

	f.fact.val = v
}

// valueNumber gets the fact value number with lock
func (f *syncFact) valueNumber() int64 {
	f.mt.Lock()
	defer f.mt.Unlock()

	return f.fact.valueNumber()
}

// valueFloat gets the fact value float with lock
func (f *syncFact) valueFloat() float64 {
	f.mt.Lock()
	defer f.mt.Unlock()

	return f.fact.valueFloat()
}

// valueString gets the fact value string with lock
func (f *syncFact) valueString() string {
	f.mt.Lock()
	defer f.mt.Unlock()

	return f.fact.valueString()
}

// valueBoolean gets the fact value bool with lock
func (f *syncFact) valueBoolean() bool {
	f.mt.Lock()
	defer f.mt.Unlock()

	return f.fact.valueBoolean()
}

// valueDate gets the fact value date (time.Time) with lock
func (f *syncFact) valueDate() time.Time {
	f.mt.Lock()
	defer f.mt.Unlock()

	return f.fact.valueDate()
}
