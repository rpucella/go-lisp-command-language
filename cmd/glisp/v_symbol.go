package main

import (
	"fmt"
)

type vSymbol struct {
	name string
}

func NewSymbol(name string) Value {
	return &vSymbol{name}
}

func (v *vSymbol) Display() string {
	return v.name
}

func (v *vSymbol) DisplayCDR() string {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vSymbol) intValue() int {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vSymbol) strValue() string {
	return v.name
}

func (v *vSymbol) boolValue() bool {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vSymbol) apply(args []Value) (Value, error) {
	return nil, fmt.Errorf("Value %s not applicable", v.str())
}

func (v *vSymbol) str() string {
	return fmt.Sprintf("VSymbol[%s]", v.name)
}

func (v *vSymbol) headValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vSymbol) tailValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vSymbol) isAtom() bool {
	return true
}

func (v *vSymbol) isSymbol() bool {
	return true
}

func (v *vSymbol) isCons() bool {
	return false
}

func (v *vSymbol) isEmpty() bool {
	return false
}

func (v *vSymbol) isNumber() bool {
	return false
}

func (v *vSymbol) isBool() bool {
	return false
}

func (v *vSymbol) isRef() bool {
	return false
}

func (v *vSymbol) isString() bool {
	return false
}

func (v *vSymbol) isFunction() bool {
	return false
}

func (v *vSymbol) isTrue() bool {
	return true
}

func (v *vSymbol) isNil() bool {
	return false
}

func (v *vSymbol) isEqual(vv Value) bool {
	return vv.isSymbol() && v.strValue() == vv.strValue()
}

func (v *vSymbol) typ() string {
	return "symbol"
}

func (v *vSymbol) getValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vSymbol) setValue(cv Value) {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vSymbol) isArray() bool {
	return false
}

func (v *vSymbol) getArray() []Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vSymbol) isDict() bool {
	return false
}

func (v *vSymbol) getDict() map[string]Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

