package main

import (
	"fmt"
)

type vString struct {
	val string
}

func NewString(v string) Value {
	return &vString{v}
}

func (v *vString) Display() string {
	return "\"" + v.val + "\""
}

func (v *vString) DisplayCDR() string {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vString) intValue() int {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vString) strValue() string {
	return v.val
}

func (v *vString) boolValue() bool {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vString) apply(args []Value) (Value, error) {
	return nil, fmt.Errorf("Value %s not applicable", v.str())
}

func (v *vString) str() string {
	return fmt.Sprintf("VString[%s]", v.str())
}

func (v *vString) headValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vString) tailValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vString) isAtom() bool {
	return true
}

func (v *vString) isSymbol() bool {
	return false
}

func (v *vString) isCons() bool {
	return false
}

func (v *vString) isEmpty() bool {
	return false
}

func (v *vString) isNumber() bool {
	return false
}

func (v *vString) isBool() bool {
	return false
}

func (v *vString) isRef() bool {
	return false
}

func (v *vString) isString() bool {
	return true
}

func (v *vString) isFunction() bool {
	return false
}

func (v *vString) isTrue() bool {
	return (v.val != "")
}

func (v *vString) isNil() bool {
	return false
}

func (v *vString) isEqual(vv Value) bool {
	return vv.isString() && v.strValue() == vv.strValue()
}

func (v *vString) typ() string {
	return "string"
}

func (v *vString) getValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vString) setValue(cv Value) {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vString) isArray() bool {
	return false
}

func (v *vString) getArray() []Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vString) isDict() bool {
	return false
}

func (v *vString) getDict() map[string]Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

