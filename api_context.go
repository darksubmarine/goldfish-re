package goldfish_re

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"
)

const reflectiveStringType = "*goldfish_re.stringFact"
const reflectiveNumberType = "*goldfish_re.numberFact"
const reflectiveFloatType = "*goldfish_re.floatFact"
const reflectiveBooleanType = "*goldfish_re.booleanFact"
const reflectiveDateType = "*goldfish_re.dateFact"

const maxIterations = 100

// Context interface that can be accessed from onActivation function
type Context interface {
	GetString(fact string) (String, error)
	GetNumber(fact string) (Number, error)
	GetFloat(fact string) (Float, error)
	GetBoolean(fact string) (Boolean, error)
	GetDate(fact string) (Date, error)
	Get(fact string) (interface{}, bool)
	GetObject(object string) (interface{}, bool)
	ForEach(fn func(fact string, value interface{}))
	Feedback(func(tx *Tx))
}

// FactsContext interface that is returned when a Context is created from a ruleset
type FactsContext interface {
	WithMaxIterations(i int)
	Register(object interface{}) error
	RegisterString(object interface{}, attribute String) error
	RegisterNumber(object interface{}, attribute Number) error
	RegisterFloat(object interface{}, attribute Float) error
	RegisterBoolean(object interface{}, attribute Boolean) error
	RegisterDate(object interface{}, attribute Date) error
	SetString(attribute interface{}, value string) error
	SetNumber(attribute interface{}, value int64) error
	SetFloat(attribute interface{}, value float64) error
	SetBoolean(attribute interface{}, value bool) error
	SetDate(attribute interface{}, value time.Time) error
	Update(fn func(tx *Tx)) error
}

// factContext internal context
type factContext struct {
	mt                sync.Mutex
	registeredFacts   map[string]interface{}
	registeredObjects map[string]interface{}
	iFactRef          _factContext
	rs                *ruleset

	feedback      bool
	feedbackFn    func(tx *Tx)
	maxIterations int
}

// newContext internal context constructor
func newContext(rs *ruleset) *factContext {
	return &factContext{registeredFacts: map[string]interface{}{}, registeredObjects: map[string]interface{}{},
		iFactRef: _factContext{}, rs: rs, maxIterations: maxIterations}
}

func (ctx *factContext) WithMaxIterations(i int) {
	ctx.maxIterations = i
}

// register internal method to register a fact and its parent object into the context
func (ctx *factContext) register(key string, obj interface{}, attr interface{}, ref iFact) {
	objKey := objectName(key)
	if _, ok := ctx.registeredObjects[objKey]; !ok {
		ctx.registeredObjects[objKey] = obj
	}

	ctx.registeredFacts[key] = attr
	ctx.iFactRef.set(ref)
}

// Register generic method to register an object with its facts.
// Also supports Go tags and is a recursive method to initialize/register nested structs
func (ctx *factContext) Register(object interface{}) error {
	if reflect.ValueOf(object).Kind() != reflect.Ptr {
		return ErrRegisteredObjectMustBePointer
	}

	var err error

	indirectVal := reflect.Indirect(reflect.ValueOf(object))
	typeOf := reflect.TypeOf(reflect.ValueOf(object).Elem().Interface())
	numFields := indirectVal.NumField()

	for i := 0; i < numFields; i++ {
		field := indirectVal.Field(i)
		ftype := field.Type().String()

		switch ftype {
		case reflectiveStringType, reflectiveNumberType, reflectiveBooleanType, reflectiveFloatType, reflectiveDateType:
			// pre-allocate pointer fields
			if field.Kind() == reflect.Ptr && field.IsNil() {
				if field.CanSet() {
					// memory allocation
					field.Set(reflect.New(field.Type().Elem()))

					// metadata from tag
					vt := typeOf.Field(i)
					var obj = typeOf.Name()
					var attr = strings.ToLower(vt.Name)
					var val = emptyStr
					if tag, ok := vt.Tag.Lookup(tag_); ok {
						if o, a, v, err := parseTag(tag); err == nil {
							if o != emptyStr {
								obj = o
							}

							if a != emptyStr {
								attr = a
							}

							if v != emptyStr {
								val = v
							}
						} else {
							return err
						}
					}

					// initializing value
					if ftype == reflectiveStringType {
						elem := NewString(obj, attr, val)
						field.Elem().Set(reflect.Indirect(reflect.ValueOf(elem)))
						f := field.Interface().(String)
						ctx.register(f.token(), object, f, f.fact)
					} else if ftype == reflectiveNumberType {
						num := parseIntOrDefault(val, 0)
						elem := NewNumber(obj, attr, num)
						field.Elem().Set(reflect.Indirect(reflect.ValueOf(elem)))
						f := field.Interface().(Number)
						ctx.register(f.token(), object, f, f.fact)
					} else if ftype == reflectiveBooleanType {
						b := parseBooleanOrDefault(val, false)
						elem := NewBoolean(obj, attr, b)
						field.Elem().Set(reflect.Indirect(reflect.ValueOf(elem)))
						f := field.Interface().(Boolean)
						ctx.register(f.token(), object, f, f.fact)
					} else if ftype == reflectiveFloatType {
						num := parseFloatOrDefault(val, 0.0)
						elem := NewFloat(obj, attr, num)
						field.Elem().Set(reflect.Indirect(reflect.ValueOf(elem)))
						f := field.Interface().(Float)
						ctx.register(f.token(), object, f, f.fact)
					} else if ftype == reflectiveDateType {
						date := parseDateOrDefault(val, zeroDate)
						elem := NewDate(obj, attr, date)
						field.Elem().Set(reflect.Indirect(reflect.ValueOf(elem)))
						f := field.Interface().(Date)
						ctx.register(f.token(), object, f, f.fact)
					}
				}
			}
			break
		default:

			// pre-allocate pointer fields
			if field.Kind() == reflect.Ptr && field.IsNil() {
				if field.CanSet() {
					// memory allocation
					field.Set(reflect.New(field.Type().Elem()))
				}
			}

			indirectField := reflect.Indirect(field)
			switch indirectField.Kind() {
			case reflect.Struct:
				// recursively allocate each of the structs embedded fields
				if field.Kind() == reflect.Ptr {
					err = ctx.Register(field.Interface())
				} else {
					// field of Struct can always use field.Addr()
					fieldAddr := field.Addr()
					if fieldAddr.CanInterface() {
						err = ctx.Register(fieldAddr.Interface())
					} else {
						err = fmt.Errorf("struct field can't interface, %#v", fieldAddr)
					}
				}

			}
		}
	}
	return err
}

// registerFact register a fact internally by data type
func (ctx *factContext) registerFact(object interface{}, attr interface{}) error {
	if attr == nil || object == nil {
		return ErrNilObject
	}

	switch f := attr.(type) {
	case String:
		ctx.register(f.token(), object, f, f.fact)
	case Number:
		ctx.register(f.token(), object, f, f.fact)
	case Float:
		ctx.register(f.token(), object, f, f.fact)
	case Date:
		ctx.register(f.token(), object, f, f.fact)
	case Boolean:
		ctx.register(f.token(), object, f, f.fact)
	default:
		return ErrInvalidDataType
	}

	return nil
}

// RegisterString registers String facts
func (ctx *factContext) RegisterString(object interface{}, attribute String) error {
	return ctx.registerFact(object, attribute)
}

// RegisterNumber registers Number facts
func (ctx *factContext) RegisterNumber(object interface{}, attribute Number) error {
	return ctx.registerFact(object, attribute)
}

// RegisterFloat registers Float facts
func (ctx *factContext) RegisterFloat(object interface{}, attribute Float) error {
	return ctx.registerFact(object, attribute)
}

// RegisterBoolean registers Boolean facts
func (ctx *factContext) RegisterBoolean(object interface{}, attribute Boolean) error {
	return ctx.registerFact(object, attribute)
}

// RegisterDate registers Date facts
func (ctx *factContext) RegisterDate(object interface{}, attribute Date) error {
	return ctx.registerFact(object, attribute)
}

// set internal fact update
func (ctx *factContext) set(object interface{}, value interface{}) error {
	return ctx.Update(func(tx *Tx) { tx.preset(object, value) })
}

// SetString sets the string value into a given fact via a transaction
func (ctx *factContext) SetString(attribute interface{}, value string) error {
	return ctx.set(attribute, value)
}

// SetNumber sets the number value into a given fact via a transaction
func (ctx *factContext) SetNumber(attribute interface{}, value int64) error {
	return ctx.set(attribute, value)
}

// SetFloat sets the float value into a given fact via a transaction
func (ctx *factContext) SetFloat(attribute interface{}, value float64) error {
	return ctx.set(attribute, value)
}

// SetBoolean sets the bool value into a given fact via a transaction
func (ctx *factContext) SetBoolean(attribute interface{}, value bool) error {
	return ctx.set(attribute, value)
}

// SetDate sets the date value into a given fact via a transaction
func (ctx *factContext) SetDate(attribute interface{}, value time.Time) error {
	return ctx.set(attribute, value)
}

func (ctx *factContext) update(fn func(tx *Tx), skip map[string]struct{}) (toSkip map[string]struct{}, finalErr error) {
	toSkip = map[string]struct{}{}
	// catch possible custom user errors into update function
	defer func() {
		if r := recover(); r != nil {
			finalErr = ErrContextUpdateRecovered
		}
	}()

	tx := newTx()

	fn(tx)

	if !tx.hasError() { // TODO if performance is poor... run evaluation async (use mutex to ensure the context data)
		tx.commit()
		//ctx.rs.EvalFacts(ctx)
		toSkip = ctx.rs.evalFactsWithSkip(ctx, skip)
	}

	if tx.err != nil {
		return toSkip, tx.err
	}

	return toSkip, tx.userErr
}

// Update run a thread-safe facts/context update via a transaction
func (ctx *factContext) Update(fn func(tx *Tx)) (finalErr error) {
	ctx.mt.Lock()
	defer ctx.mt.Unlock()

	var toSkip map[string]struct{}
	if skp, err := ctx.update(fn, map[string]struct{}{}); err != nil {
		return err
	} else {
		toSkip = skp
	}

	for i := 0; ctx.feedback && i < ctx.maxIterations; i++ {
		ctx.feedback = false
		if skp, err := ctx.update(ctx.feedbackFn, toSkip); err != nil {
			return err
		} else {
			toSkip = skp
		}
	}

	return nil
}

// Feedback run a thread-safe facts/context update via a transaction
func (ctx *factContext) Feedback(fn func(tx *Tx)) {
	if fn == nil {
		return
	}

	ctx.feedbackFn = fn
	ctx.feedback = true
}

// GetObject returns a fact's parent object
func (ctx *factContext) GetObject(object string) (interface{}, bool) {
	f, ok := ctx.registeredObjects[object]
	return f, ok
}

// Get returns a registered fact
func (ctx *factContext) Get(fact string) (interface{}, bool) {
	f, ok := ctx.registeredFacts[fact]
	return f, ok
}

// GetString gets a String fact or error it
func (ctx *factContext) GetString(fact string) (String, error) {
	if iobj, ok := ctx.Get(fact); !ok {
		return nil, ErrFactNotFound
	} else {
		if obj, ok := iobj.(String); ok {
			return obj, nil
		} else {
			return nil, ErrFactInvalidType
		}
	}
}

// GetNumber gets a Number fact or error it
func (ctx *factContext) GetNumber(fact string) (Number, error) {
	if iobj, ok := ctx.Get(fact); !ok {
		return nil, ErrFactNotFound
	} else {
		if obj, ok := iobj.(Number); ok {
			return obj, nil
		} else {
			return nil, ErrFactInvalidType
		}
	}
}

// GetFloat gets a Float fact or error it
func (ctx *factContext) GetFloat(fact string) (Float, error) {
	if iobj, ok := ctx.Get(fact); !ok {
		return nil, ErrFactNotFound
	} else {
		if obj, ok := iobj.(Float); ok {
			return obj, nil
		} else {
			return nil, ErrFactInvalidType
		}
	}
}

// GetBoolean gets a Boolean fact or error it
func (ctx *factContext) GetBoolean(fact string) (Boolean, error) {
	if iobj, ok := ctx.Get(fact); !ok {
		return nil, ErrFactNotFound
	} else {
		if obj, ok := iobj.(Boolean); ok {
			return obj, nil
		} else {
			return nil, ErrFactInvalidType
		}
	}
}

// GetDate gets a Date fact or error it
func (ctx *factContext) GetDate(fact string) (Date, error) {
	if iobj, ok := ctx.Get(fact); !ok {
		return nil, ErrFactNotFound
	} else {
		if obj, ok := iobj.(Date); ok {
			return obj, nil
		} else {
			return nil, ErrFactInvalidType
		}
	}
}

// ForEach iterates over all registered facts
func (ctx *factContext) ForEach(fn func(fact string, value interface{})) {
	for k, v := range ctx.iFactRef {
		fn(k, v.value())
	}
}
