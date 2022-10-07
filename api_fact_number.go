package goldfish_re

// Number data type used to declare user Facts. This is an alias to *numberFact
type Number = *numberFact

// numberFact thread safe struct to work with Number facts
type numberFact struct {
	*syncFact
}

// NewNumber is the Number fact constructor
func NewNumber(object, attribute string, value int64) Number {
	return &numberFact{syncFact: &syncFact{fact: newFact(object, attribute, value)}}
}

// set the fact value. Not exported, user can set this value via a transactional context
func (f *numberFact) set(n int64) {
	f.syncFact.set(n)
}

// Value fact value getter
func (f *numberFact) Value() int64 {
	return f.valueNumber()
}
