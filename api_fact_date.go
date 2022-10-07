package goldfish_re

import "time"

// Date data type used to declare user Facts. This is an alias to *dateFact
// the data type is a wrapper of time.Time
type Date = *dateFact

// dateFact thread safe struct to work with Date facts
type dateFact struct {
	*syncFact
}

// NewDate is the Date fact constructor
func NewDate(object, attribute string, value time.Time) Date {
	return &dateFact{syncFact: &syncFact{fact: newFact(object, attribute, value)}}
}

// set the fact value. Not exported, user can set this value via a transactional context
func (f *dateFact) set(v time.Time) {
	f.syncFact.set(v)
}

// Value fact value getter
func (f *dateFact) Value() time.Time {
	return f.valueDate()
}
