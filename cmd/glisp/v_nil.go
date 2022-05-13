package main

import (
	"fmt"
)

type vNil struct {
}

func NewNil() Value {
	return &vNil{}
}

func (v *vNil) Display() string {
	// figure out if this is the right thing?
	return "#nil"
}

func (v *vNil) DisplayCDR() string {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vNil) intValue() int {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vNil) strValue() string {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vNil) boolValue() bool {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vNil) apply(args []Value) (Value, error) {
	return nil, fmt.Errorf("Value %s not applicable", v.str())
}

func (v *vNil) str() string {
	return fmt.Sprintf("VNil")
}

func (v *vNil) headValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vNil) tailValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vNil) isAtom() bool {
	return false
}

func (v *vNil) isSymbol() bool {
	return false
}

func (v *vNil) isCons() bool {
	return false
}

func (v *vNil) isEmpty() bool {
	return false
}

func (v *vNil) isNumber() bool {
	return false
}

func (v *vNil) isBool() bool {
	return false
}

func (v *vNil) isRef() bool {
	return false
}

func (v *vNil) isString() bool {
	return false
}

func (v *vNil) isFunction() bool {
	return false
}

func (v *vNil) isTrue() bool {
	return false
}

func (v *vNil) isNil() bool {
	return true
}

func (v *vNil) isEqual(vv Value) bool {
	return vv.isNil()
}

func (v *vNil) typ() string {
	return "nil"
}

func (v *vNil) getValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vNil) setValue(cv Value) {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vNil) isArray() bool {
	return false
}

func (v *vNil) getArray() []Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vNil) isDict() bool {
	return false
}

func (v *vNil) getDict() map[string]Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

