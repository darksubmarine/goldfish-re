package goldfish_re

import "errors"

// noErr is the same that a nil, only adds semantic
var noErr error = nil

var (
	// errConditionAddedPreviously condition added previously
	errConditionAddedPreviously = errors.New("condition added previously")

	// ErrInvalidDataType invalid data type
	ErrInvalidDataType = errors.New("invalid data type")

	// ErrInvalidValueType invalid value type
	ErrInvalidValueType = errors.New("invalid value type")

	// ErrContextUpdateRecovered recovered context update after panic
	ErrContextUpdateRecovered = errors.New("recovered context update after panic")

	// ErrRegisteredObjectMustBePointer error registering object must be a pointer
	ErrRegisteredObjectMustBePointer = errors.New("error registering object must be a pointer")

	// ErrMalformedTag malformed tag
	ErrMalformedTag = errors.New("malformed tag")

	// ErrEmptyThenSentence empty 'then' sentence in rule
	ErrEmptyThenSentence = errors.New("empty 'then' sentence in rule")

	// ErrEmptyConditionList the rule must contains at least one condition
	ErrEmptyConditionList = errors.New("the rule must contains at least one condition")

	// ErrNilObject nil object
	ErrNilObject = errors.New("nil object")

	// ErrFactNotFound fact not found
	ErrFactNotFound = errors.New("fact not found")

	// ErrFactInvalidType fact is registered with different data type
	ErrFactInvalidType = errors.New("fact is registered with different data type")
)
