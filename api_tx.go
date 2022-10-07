package goldfish_re

import "time"

// Tx transaction struct
type Tx struct {
	err     error
	userErr error
	toApply map[interface{}]interface{}
}

// newTx transaction constructor
func newTx() *Tx {
	return &Tx{toApply: map[interface{}]interface{}{}}
}

// hasError checks if the tx has a user error or a lib error
func (tx *Tx) hasError() bool {
	return tx.err != nil || tx.userErr != nil
}

// Error exported method to let users add custom errors on context.Update method
func (tx *Tx) Error(err error) {
	tx.userErr = err
}

// commit apply the transaction operations on the target facts
func (tx *Tx) commit() {
	for obj, val := range tx.toApply {
		tx.set(obj, val)
	}
}

// set the value over the target fact
func (tx *Tx) set(object interface{}, value interface{}) {
	switch obj := object.(type) {
	case String:
		if str, ok := value.(string); ok {
			obj.set(str)
		} else {
			tx.err = ErrInvalidValueType
		}
	case Number:
		if num, ok := value.(int64); ok {
			obj.set(num)
		} else {
			tx.err = ErrInvalidValueType
		}
	case Float:
		if num, ok := value.(float64); ok {
			obj.set(num)
		} else {
			tx.err = ErrInvalidValueType
		}
	case Boolean:
		if b, ok := value.(bool); ok {
			obj.set(b)
		} else {
			tx.err = ErrInvalidValueType
		}
	case Date:
		if d, ok := value.(time.Time); ok {
			obj.set(d)
		} else {
			tx.err = ErrInvalidValueType
		}
	default:
		tx.err = ErrInvalidDataType
	}
}

// preset the values to the target facts
func (tx *Tx) preset(object interface{}, value interface{}) {
	switch obj := object.(type) {
	case String:
		if _, ok := value.(string); ok {
			tx.toApply[obj] = value
		} else {
			tx.err = ErrInvalidValueType
		}
	case Number:
		if _, ok := value.(int64); ok {
			tx.toApply[obj] = value
		} else {
			tx.err = ErrInvalidValueType
		}
	case Float:
		if _, ok := value.(float64); ok {
			tx.toApply[obj] = value
		} else {
			tx.err = ErrInvalidValueType
		}
	case Boolean:
		if _, ok := value.(bool); ok {
			tx.toApply[obj] = value
		} else {
			tx.err = ErrInvalidValueType
		}
	case Date:
		if _, ok := value.(time.Time); ok {
			tx.toApply[obj] = value
		} else {
			tx.err = ErrInvalidValueType
		}
	default:
		tx.err = ErrInvalidDataType
	}
}

// SetString preset the given fact with the given string value
func (tx *Tx) SetString(object String, value string) {
	tx.preset(object, value)
}

// SetNumber preset the given fact with the given int64 value
func (tx *Tx) SetNumber(object Number, value int64) {
	tx.preset(object, value)
}

// SetFloat preset the given fact with the given float64 value
func (tx *Tx) SetFloat(object Float, value float64) {
	tx.preset(object, value)
}

// SetBoolean preset the given fact with the given bool value
func (tx *Tx) SetBoolean(object Boolean, value bool) {
	tx.preset(object, value)
}

// SetDate preset the given fact with the given time.Time value
func (tx *Tx) SetDate(object Date, value time.Time) {
	tx.preset(object, value)
}
