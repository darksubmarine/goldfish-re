package goldfish_re

// String data type used to declare user Facts. This is an alias to *stringFact
type String = *stringFact

// thread safe struct to work with String facts
type stringFact struct {
	*syncFact
}

// NewString is the String fact constructor
func NewString(object, attribute, value string) String {
	return &stringFact{syncFact: &syncFact{fact: newFact(object, attribute, value)}}
}

// set the fact value. Not exported, user can set this value via a transactional context
func (f *stringFact) set(s string) {
	f.syncFact.set(s)
}

// Value fact value getter
func (f *stringFact) Value() string {
	return f.valueString()
}
