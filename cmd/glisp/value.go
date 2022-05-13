package main

type Value interface {
	Display() string
	DisplayCDR() string

	/*
	asInteger() (int, bool)
	asBoolean() (bool, bool)
	asString() (string, bool)
	asSymbol() (string, bool)
	asCons() (Value, Value, bool)
	asReference() (Value, bool)
	setReference(Value) bool
	asArray() ([]Value, bool)
	asDict() (map[string]Value, bool)
	*/
	
	apply([]Value) (Value, error)
	str() string
	isAtom() bool
	isEmpty() bool
	isTrue() bool
	isNil() bool
	//isEq() bool    -- don't think we need pointer equality for now - = is enough?
	isEqual(Value) bool
	typ() string

	// To shift out.
	boolValue() bool
	strValue() string
	headValue() Value
	tailValue() Value
	isSymbol() bool
	isCons() bool
	intValue() int
	isNumber() bool
	isBool() bool
	isRef() bool
	isString() bool
	isFunction() bool
	getValue() Value
	setValue(Value)
	isArray() bool
	getArray() []Value
	isDict() bool
	getDict() map[string]Value
}

