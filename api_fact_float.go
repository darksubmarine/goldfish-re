package goldfish_re

// Float data type used to declare user Facts. This is an alias to *floatFact
type Float = *floatFact

// floatFact thread safe struct to work with Float facts
type floatFact struct {
	*syncFact
}

// NewFloat is the Float fact constructor
func NewFloat(object, attribute string, value float64) Float {
	return &floatFact{syncFact: &syncFact{fact: newFact(object, attribute, value)}}
}

// set the fact value. Not exported, user can set this value via a transactional context
func (f *floatFact) set(n float64) {
	f.syncFact.set(n)
}

// Value fact value getter
func (f *floatFact) Value() float64 {
	return f.valueFloat()
}
