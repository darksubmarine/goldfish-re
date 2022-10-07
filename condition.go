package goldfish_re

import (
	"strings"
	"time"
)

// _condition internal condition representation
type _condition struct {
	id        cuid
	operator  tOperator
	lTerm     iTerm
	rTerm     iTerm
	token_    string
	ruleSlice []*_rule
	negated   bool
}

// _newCondition constructor function
func _newCondition(id cuid, left iTerm, right iTerm, operator tOperator, negated bool) *_condition {
	// because this is internal object_ all checks like terms type matches happens before construction
	tkn := conditionToken(left.token(), right.token(), operator.token(), negated)
	return &_condition{id: id, token_: tkn, operator: operator, lTerm: left, rTerm: right, negated: negated}
}

// newCondition non negated constructor
func newCondition(id cuid, left iTerm, right iTerm, operator tOperator) *_condition {
	return _newCondition(id, left, right, operator, false)
}

// newNegatedCondition negated constructor
func newNegatedCondition(id cuid, left iTerm, right iTerm, operator tOperator) *_condition {
	return _newCondition(id, left, right, operator, true)
}

// cloneWithId clone a condition with a given ID
func (c *_condition) cloneWithId(id cuid) *_condition {
	return &_condition{
		id:        id,
		operator:  c.operator,
		lTerm:     c.lTerm,
		rTerm:     c.rTerm,
		token_:    c.token_,
		ruleSlice: c.ruleSlice,
		negated:   c.negated,
	}
}

// token condition string representation
func (c *_condition) token() string {
	return c.token_
}

// addRule link the condition with a given rule
func (c *_condition) addRule(r *_rule) {
	c.ruleSlice = append(c.ruleSlice, r) // TODO set slice capacity and use RUID as slice index
}

// eval evaluate the condition with a given fact and context
func (c *_condition) eval(fact iFact, ctx _factContext) (bool, iFact) {

	var left, right iFact
	var result = false
	var againstFact = false
	var evalLeft = true

	if fact.object() == c.lTerm.object() && fact.attribute() == c.lTerm.attribute() {
		left = fact
		if fc, ok := ctx.get(c.rTerm.token()); ok {
			right = fc
			againstFact = true
		} else {
			right = newFact(c.rTerm.object(), c.rTerm.attribute(), c.rTerm.val())
		}
	} else if fact.object() == c.rTerm.object() && fact.attribute() == c.rTerm.attribute() {
		evalLeft = false
		right = fact
		if fc, ok := ctx.get(c.lTerm.token()); ok {
			left = fc
			againstFact = true
		} else {
			left = newFact(c.lTerm.object(), c.lTerm.attribute(), c.lTerm.val())
		}
	} else {
		return false, nil
	}

	if fact.isString() {
		result = evalString(left, right, c.operator)
	} else if fact.isNumber() {
		result = evalNumber(left, right, c.operator)
	} else if fact.isFloat() {
		result = evalFloat(left, right, c.operator)
	} else if fact.isBoolean() {
		result = evalBoolean(left, right, c.operator)
	} else if fact.isDate() {
		result = evalDate(left, right, c.operator)
	}

	if againstFact {
		if evalLeft {
			return c.negated != result, right
		} else {
			return c.negated != result, left
		}
	}

	return c.negated != result, nil
}

// evalString eval string data type
func evalString(fact iFact, rFact iFact, op tOperator) bool {
	if op == opIn {
		factValue := fact.valueString()
		if values, ok := rFact.value().([]string); ok {
			for _, val := range values {
				if factValue == val {
					return true
				}
			}
		}
		return false
	}

	switch op {
	case opEquals:
		return fact.valueString() == rFact.valueString()
	case opContains:
		return strings.Contains(fact.valueString(), rFact.valueString())
	case opStarts:
		return strings.HasPrefix(fact.valueString(), rFact.valueString())
	case opEnds:
		return strings.HasSuffix(fact.valueString(), rFact.valueString())
	}

	return false
}

// evalNumber eval number data type
func evalNumber(fact iFact, rFact iFact, op tOperator) bool {
	switch op {
	case opEquals:
		return fact.valueNumber() == rFact.valueNumber()
	case opGreaterThan:
		return fact.valueNumber() > rFact.valueNumber()
	case opGreaterThanOrEqual:
		return fact.valueNumber() >= rFact.valueNumber()
	case opLessThan:
		return fact.valueNumber() < rFact.valueNumber()
	case opLessThanOrEqual:
		return fact.valueNumber() <= rFact.valueNumber()
	}

	return false
}

// evalFloat eval float data type
func evalFloat(fact iFact, rFact iFact, op tOperator) bool {
	switch op {
	case opEquals:
		return fact.valueFloat() == rFact.valueFloat()
	case opGreaterThan:
		return fact.valueFloat() > rFact.valueFloat()
	case opGreaterThanOrEqual:
		return fact.valueFloat() >= rFact.valueFloat()
	case opLessThan:
		return fact.valueFloat() < rFact.valueFloat()
	case opLessThanOrEqual:
		return fact.valueFloat() <= rFact.valueFloat()
	}

	return false
}

// evalDate eval date data type
func evalDate(fact iFact, rFact iFact, op tOperator) bool {

	if op == opBetween {
		factValue := fact.valueDate()
		if values, ok := rFact.value().([]time.Time); ok {
			if len(values) != 2 {
				return false
			}

			if factValue.After(values[0]) && factValue.Before(values[1]) {
				return true
			}
		}
		return false
	}

	switch op {
	case opEquals:
		return fact.valueDate() == rFact.valueDate()
	case opAfter:
		return fact.valueDate().After(rFact.valueDate())
	case opBefore:
		return fact.valueDate().Before(rFact.valueDate())
	}

	return false
}

// evalBoolean eval bool data type
func evalBoolean(fact iFact, rFact iFact, op tOperator) bool {
	switch op {
	case opEquals:
		return fact.valueBoolean() == rFact.valueBoolean()
	}
	return false
}
