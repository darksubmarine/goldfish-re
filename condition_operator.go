package goldfish_re

type tOperator uint8

const (
	_ tOperator = iota
	// commons
	opEquals

	// number and float
	opGreaterThan
	opGreaterThanOrEqual
	opLessThan
	opLessThanOrEqual

	// string
	opStarts
	opEnds
	opContains
	opIn

	// date
	opAfter
	opBefore
	opBetween
)

func (e tOperator) token() string {
	return e.String()
}

func (e tOperator) String() string {
	switch e {
	case opEquals:
		return "=="
	case opGreaterThan:
		return ">"
	case opGreaterThanOrEqual:
		return ">="
	case opLessThan:
		return "<"
	case opLessThanOrEqual:
		return "<="
	case opStarts:
		return "starts"
	case opEnds:
		return "ends"
	case opContains:
		return "contains"
	case opIn:
		return "in"
	default:
		return undefined
	}
}
