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

func (v *vSymbol) apply(args []Value) (Value, error) {
	return nil, fmt.Errorf("Value %s not applicable", v.str())
}

func (v *vSymbol) str() string {
	return fmt.Sprintf("VSymbol[%s]", v.name)
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

func (v *vSymbol) asInteger() (int, bool) {
	return 0, false
}

func (v *vSymbol) asBoolean() (bool, bool) {
	return false, false
}

func (v *vSymbol) asString() (string, bool) {
	return "", false
}

func (v *vSymbol) asSymbol() (string, bool) {
	return v.name, true
}

func (v *vSymbol) asCons() (Value, Value, bool) {
	return nil, nil, false
}

func (v *vSymbol) asReference() (Value, func(Value), bool) {
	return nil, nil, false
}

func (v *vSymbol) setReference(Value) bool {
	return false
}

func (v *vSymbol) asArray() ([]Value, bool) {
	return nil, false
}

func (v *vSymbol) asDict() (map[string]Value, bool) {
	return nil, false
}


func (v *vSymbol) intValue() int {
	return intValue(v)
}

func (v *vSymbol) strValue() string {
	return strValue(v)
}

func (v *vSymbol) boolValue() bool {
	return boolValue(v)
}

func (v *vSymbol) headValue() Value {
	return headValue(v)
}

func (v *vSymbol) tailValue() Value {
	return tailValue(v)
}


func (v *vSymbol) isArray() bool {
	return isArray(v)
}

func (v *vSymbol) getArray() []Value {
	return getArray(v)
}

func (v *vSymbol) isDict() bool {
	return isDict(v)
}

func (v *vSymbol) getDict() map[string]Value {
	return getDict(v)
}


func (v *vSymbol) isRef() bool {
	return isRef(v)
}

func (v *vSymbol) getValue() Value {
	return getValue(v)
}

func (v *vSymbol) setValue(cv Value) {
	setValue(v, cv)
}

