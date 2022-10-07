package goldfish_re

type tBinaryOperator uint8

const (
	_ tBinaryOperator = iota
	opAnd
	opOr
)

func (e tBinaryOperator) token() string {
	return e.String()
}

func (e tBinaryOperator) String() string {
	switch e {
	case opAnd:
		return "&&"
	case opOr:
		return "||"
	default:
		return undefined
	}
}
