package main

import (
	"fmt"
)

type vBoolean struct {
	val bool
}

func NewBoolean(v bool) Value {
	return &vBoolean{v}
}

func (v *vBoolean) Display() string {
	if v.val {
		return "#t"
	} else {
		return "#f"
	}
}

func (v *vBoolean) DisplayCDR() string {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vBoolean) apply(args []Value) (Value, error) {
	return nil, fmt.Errorf("Value %s not applicable", v.str())
}

func (v *vBoolean) str() string {
	if v.val {
		return "VBoolean[true]"
	} else {
		return "VBoolean[false]"
	}
}

func (v *vBoolean) isAtom() bool {
	return true
}

func (v *vBoolean) isSymbol() bool {
	return false
}

func (v *vBoolean) isCons() bool {
	return false
}

func (v *vBoolean) isEmpty() bool {
	return false
}

func (v *vBoolean) isNumber() bool {
	return false
}

func (v *vBoolean) isBool() bool {
	return true
}

func (v *vBoolean) isString() bool {
	return false
}

func (v *vBoolean) isFunction() bool {
	return false
}

func (v *vBoolean) isTrue() bool {
	return v.val
}

func (v *vBoolean) isNil() bool {
	return false
}

func (v *vBoolean) isEqual(vv Value) bool {
	return vv.isBool() && v.boolValue() == vv.boolValue()
}

func (v *vBoolean) typ() string {
	return "bool"
}

func (v *vBoolean) asInteger() (int, bool) {
	return 0, false
}

func (v *vBoolean) asBoolean() (bool, bool) {
	return v.val, true
}

func (v *vBoolean) asString() (string, bool) {
	return "", false
}

func (v *vBoolean) asSymbol() (string, bool) {
	return "", false
}

func (v *vBoolean) asCons() (Value, Value, bool) {
	return nil, nil, false
}

func (v *vBoolean) asReference() (Value, func(Value), bool) {
	return nil, nil, false
}

func (v *vBoolean) setReference(Value) bool {
	return false
}

func (v *vBoolean) asArray() ([]Value, bool) {
	return nil, false
}

func (v *vBoolean) asDict() (map[string]Value, bool) {
	return nil, false
}


func (v *vBoolean) intValue() int {
	return intValue(v)
}

func (v *vBoolean) strValue() string {
	return strValue(v)
}

func (v *vBoolean) boolValue() bool {
	return boolValue(v)
}

func (v *vBoolean) headValue() Value {
	return headValue(v)
}

func (v *vBoolean) tailValue() Value {
	return tailValue(v)
}


func (v *vBoolean) isArray() bool {
	return isArray(v)
}

func (v *vBoolean) getArray() []Value {
	return getArray(v)
}

func (v *vBoolean) isDict() bool {
	return isDict(v)
}

func (v *vBoolean) getDict() map[string]Value {
	return getDict(v)
}


func (v *vBoolean) isRef() bool {
	return isRef(v)
}

func (v *vBoolean) getValue() Value {
	return getValue(v)
}

func (v *vBoolean) setValue(cv Value) {
	setValue(v, cv)
}

