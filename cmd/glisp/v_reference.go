package main

import (
	"fmt"
)

type vReference struct {
	content Value
}

func NewReference(v Value) Value {
	return &vReference{v}
}

func (v *vReference) Display() string {
	return fmt.Sprintf("#<ref %s>", v.content.Display())
}

func (v *vReference) DisplayCDR() string {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vReference) intValue() int {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vReference) strValue() string {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vReference) boolValue() bool {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vReference) apply(args []Value) (Value, error) {
	if len(args) > 1 {
		return nil, fmt.Errorf("too many arguments %d to ref update", len(args))
	}
	if len(args) == 1 {
		v.content = args[0]
		return &vNil{}, nil
	}
	return v.content, nil
}

func (v *vReference) str() string {
	return fmt.Sprintf("VReference[%s]", v.content.str())
}

func (v *vReference) headValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vReference) tailValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vReference) isAtom() bool {
	return false // ?
}

func (v *vReference) isSymbol() bool {
	return false
}

func (v *vReference) isCons() bool {
	return false
}

func (v *vReference) isEmpty() bool {
	return false
}

func (v *vReference) isNumber() bool {
	return false
}

func (v *vReference) isBool() bool {
	return false
}

func (v *vReference) isRef() bool {
	return true
}

func (v *vReference) isString() bool {
	return false
}

func (v *vReference) isFunction() bool {
	return false
}

func (v *vReference) isTrue() bool {
	return false
}

func (v *vReference) isNil() bool {
	return false
}

func (v *vReference) isEqual(vv Value) bool {
	return v == vv // pointer equality
}

func (v *vReference) typ() string {
	return "reference"
}

func (v *vReference) getValue() Value {
	return v.content
}

func (v *vReference) setValue(cv Value) {
	v.content = cv
}

func (v *vReference) isArray() bool {
	return false
}

func (v *vReference) getArray() []Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vReference) isDict() bool {
	return false
}

func (v *vReference) getDict() map[string]Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

