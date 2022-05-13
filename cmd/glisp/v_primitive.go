package main

import (
	"fmt"
)

type vPrimitive struct {
	name      string
	primitive func([]Value) (Value, error)
}

func NewPrimitive(name string, prim func([]Value) (Value, error)) Value {
	return &vPrimitive{name, prim}
}

func (v *vPrimitive) Display() string {
	return fmt.Sprintf("#<prim %s>", v.name)
}

func (v *vPrimitive) DisplayCDR() string {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vPrimitive) intValue() int {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vPrimitive) strValue() string {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vPrimitive) boolValue() bool {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vPrimitive) apply(args []Value) (Value, error) {
	return v.primitive(args)
}

func (v *vPrimitive) str() string {
	return fmt.Sprintf("VPrimitive[%s]", v.name)
}

func (v *vPrimitive) headValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vPrimitive) tailValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vPrimitive) isAtom() bool {
	return false
}

func (v *vPrimitive) isSymbol() bool {
	return false
}

func (v *vPrimitive) isCons() bool {
	return false
}

func (v *vPrimitive) isEmpty() bool {
	return false
}

func (v *vPrimitive) isNumber() bool {
	return false
}

func (v *vPrimitive) isBool() bool {
	return false
}

func (v *vPrimitive) isRef() bool {
	return false
}

func (v *vPrimitive) isString() bool {
	return false
}

func (v *vPrimitive) isFunction() bool {
	return true
}

func (v *vPrimitive) isTrue() bool {
	return true
}

func (v *vPrimitive) isNil() bool {
	return false
}

func (v *vPrimitive) isEqual(vv Value) bool {
	return v == vv // pointer equality
}

func (v *vPrimitive) typ() string {
	return "fun"
}

func (v *vPrimitive) getValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vPrimitive) setValue(cv Value) {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vPrimitive) isArray() bool {
	return false
}

func (v *vPrimitive) getArray() []Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vPrimitive) isDict() bool {
	return false
}

func (v *vPrimitive) getDict() map[string]Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

