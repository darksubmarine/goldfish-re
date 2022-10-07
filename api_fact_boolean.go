package goldfish_re

// Boolean data type used to declare user Facts. This is an alias to *booleanFact
type Boolean = *booleanFact

// booleanFact thread safe struct to work with Boolean facts
type booleanFact struct {
	*syncFact
}

// NewBoolean is the Boolean fact constructor
func NewBoolean(object, attribute string, value bool) Boolean {
	return &booleanFact{syncFact: &syncFact{fact: newFact(object, attribute, value)}}
}

// set the fact value. Not exported, user can set this value via a transactional context
func (f *booleanFact) set(v bool) {
	f.syncFact.set(v)
}

// Value fact value getter
func (f *booleanFact) Value() bool {
	return f.valueBoolean()
}
