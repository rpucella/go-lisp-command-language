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

func (v *vNil) apply(args []Value) (Value, error) {
	return nil, fmt.Errorf("Value %s not applicable", v.str())
}

func (v *vNil) str() string {
	return fmt.Sprintf("VNil")
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

func (v *vNil) asInteger() (int, bool) {
	return 0, false
}

func (v *vNil) asBoolean() (bool, bool) {
	return false, false
}

func (v *vNil) asString() (string, bool) {
	return "", false
}

func (v *vNil) asSymbol() (string, bool) {
	return "", false
}

func (v *vNil) asCons() (Value, Value, bool) {
	return nil, nil, false
}

func (v *vNil) asReference() (Value, func(Value), bool) {
	return nil, nil, false
}

func (v *vNil) setReference(Value) bool {
	return false
}

func (v *vNil) asArray() ([]Value, bool) {
	return nil, false
}

func (v *vNil) asDict() (map[string]Value, bool) {
	return nil, false
}


func (v *vNil) intValue() int {
	return intValue(v)
}

func (v *vNil) strValue() string {
	return strValue(v)
}

func (v *vNil) boolValue() bool {
	return boolValue(v)
}

func (v *vNil) headValue() Value {
	return headValue(v)
}

func (v *vNil) tailValue() Value {
	return tailValue(v)
}


func (v *vNil) isArray() bool {
	return isArray(v)
}

func (v *vNil) getArray() []Value {
	return getArray(v)
}

func (v *vNil) isDict() bool {
	return isDict(v)
}

func (v *vNil) getDict() map[string]Value {
	return getDict(v)
}


func (v *vNil) isRef() bool {
	return isRef(v)
}

func (v *vNil) getValue() Value {
	return getValue(v)
}

func (v *vNil) setValue(cv Value) {
	setValue(v, cv)
}

