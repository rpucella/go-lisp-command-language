package main

import (
	"fmt"
)

type vEmpty struct {
}

func NewEmpty() Value {
	return &vEmpty{}
}

func (v *vEmpty) Display() string {
	return "()"
}

func (v *vEmpty) DisplayCDR() string {
	return ")"
}

func (v *vEmpty) intValue() int {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vEmpty) strValue() string {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vEmpty) boolValue() bool {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vEmpty) apply(args []Value) (Value, error) {
	return nil, fmt.Errorf("Value %s not applicable", v.str())
}

func (v *vEmpty) str() string {
	return fmt.Sprintf("VEmpty")
}

func (v *vEmpty) headValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vEmpty) tailValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vEmpty) isAtom() bool {
	return false
}

func (v *vEmpty) isSymbol() bool {
	return false
}

func (v *vEmpty) isCons() bool {
	return false
}

func (v *vEmpty) isEmpty() bool {
	return true
}

func (v *vEmpty) isNumber() bool {
	return false
}

func (v *vEmpty) isBool() bool {
	return false
}

func (v *vEmpty) isRef() bool {
	return false
}

func (v *vEmpty) isString() bool {
	return false
}

func (v *vEmpty) isFunction() bool {
	return false
}

func (v *vEmpty) isTrue() bool {
	return false
}

func (v *vEmpty) isNil() bool {
	return false
}

func (v *vEmpty) isEqual(vv Value) bool {
	return vv.isEmpty()
}

func (v *vEmpty) typ() string {
	return "list"
}

func (v *vEmpty) getValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vEmpty) setValue(cv Value) {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vEmpty) isArray() bool {
	return false
}

func (v *vEmpty) getArray() []Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vEmpty) isDict() bool {
	return false
}

func (v *vEmpty) getDict() map[string]Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}
