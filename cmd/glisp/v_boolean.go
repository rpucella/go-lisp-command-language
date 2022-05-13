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

func (v *vBoolean) intValue() int {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vBoolean) strValue() string {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vBoolean) boolValue() bool {
	return v.val
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

func (v *vBoolean) headValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vBoolean) tailValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
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

func (v *vBoolean) isRef() bool {
	return false
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

func (v *vBoolean) getValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vBoolean) setValue(cv Value) {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vBoolean) isArray() bool {
	return false
}

func (v *vBoolean) getArray() []Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vBoolean) isDict() bool {
	return false
}

func (v *vBoolean) getDict() map[string]Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

