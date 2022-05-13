package main

import (
	"fmt"
)

type Value interface {
	Display() string
	DisplayCDR() string

	asInteger() (int, bool)
	asBoolean() (bool, bool)
	asString() (string, bool)
	asSymbol() (string, bool)
	asCons() (Value, Value, bool)
	asReference() (Value, func(Value), bool)
	setReference(Value) bool
	asArray() ([]Value, bool)
	asDict() (map[string]Value, bool)
	
	apply([]Value) (Value, error)
	str() string
	isAtom() bool
	isEmpty() bool
	isTrue() bool
	isNil() bool
	isFunction() bool
	//isEq() bool    -- don't think we need pointer equality for now - = is enough?
	isEqual(Value) bool
	typ() string

	// To shift out.
	boolValue() bool
	strValue() string
	headValue() Value
	tailValue() Value
	intValue() int
	
	isSymbol() bool
	isCons() bool
	isNumber() bool
	isBool() bool
	isString() bool
}

func boolValue(v Value) bool {
	b, ok := v.asBoolean()
	if ok {
		return b
	}
	panic(fmt.Sprintf("Value %s not a boolean", v.str()))
}

func intValue(v Value) int {
	b, ok := v.asInteger()
	if ok {
		return b
	}
	panic(fmt.Sprintf("Value %s not an integer", v.str()))
}

func strValue(v Value) string {
	b, ok := v.asString()
	if ok {
		return b
	}
	panic(fmt.Sprintf("Value %s not a string", v.str()))
}

func symbValue(v Value) string {
	b, ok := v.asSymbol()
	if ok {
		return b
	}
	panic(fmt.Sprintf("Value %s not a symbol", v.str()))
}

func headValue(v Value) Value {
	b, _, ok := v.asCons()
	if ok {
		return b
	}
	panic(fmt.Sprintf("Value %s not a cons cell", v.str()))
}

func tailValue(v Value) Value {
	_, c, ok := v.asCons()
	if ok {
		return c
	}
	panic(fmt.Sprintf("Value %s not a cons cell", v.str()))
}

func isArray(v Value) bool {
	_, ok := v.asArray()
	return ok
}

func getArray(v Value) []Value {
	b, ok := v.asArray()
	if ok {
		return b
	}
	panic(fmt.Sprintf("Value %s not an array", v.str()))
}

func isDict(v Value) bool {
	_, ok := v.asDict()
	return ok
}

func getDict(v Value) map[string]Value {
	b, ok := v.asDict()
	if ok {
		return b
	}
	panic(fmt.Sprintf("Value %s not a dict", v.str()))
}

func isRef(v Value) bool {
	_, _, ok := v.asReference()
	return ok
}

func getValue(v Value) Value {
	b, _, ok := v.asReference()
	if ok {
		return b
	}
	panic(fmt.Sprintf("Value %s not a ref", v.str()))
}

func setValue(v Value, nv Value) {
	_, update, ok := v.asReference()
	if ok {
		update(nv)
	} else {
		panic(fmt.Sprintf("Value %s not a ref", v.str()))
	}
}
	
