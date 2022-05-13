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

func (v *vEmpty) apply(args []Value) (Value, error) {
	return nil, fmt.Errorf("Value %s not applicable", v.str())
}

func (v *vEmpty) str() string {
	return fmt.Sprintf("VEmpty")
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

func (v *vEmpty) asInteger() (int, bool) {
	return 0, false
}

func (v *vEmpty) asBoolean() (bool, bool) {
	return false, false
}

func (v *vEmpty) asString() (string, bool) {
	return "", false
}

func (v *vEmpty) asSymbol() (string, bool) {
	return "", false
}

func (v *vEmpty) asCons() (Value, Value, bool) {
	return nil, nil, false
}

func (v *vEmpty) asReference() (Value, func(Value), bool) {
	return nil, nil, false
}

func (v *vEmpty) setReference(Value) bool {
	return false
}

func (v *vEmpty) asArray() ([]Value, bool) {
	return nil, false
}

func (v *vEmpty) asDict() (map[string]Value, bool) {
	return nil, false
}


func (v *vEmpty) intValue() int {
	return intValue(v)
}

func (v *vEmpty) strValue() string {
	return strValue(v)
}

func (v *vEmpty) boolValue() bool {
	return boolValue(v)
}

func (v *vEmpty) headValue() Value {
	return headValue(v)
}

func (v *vEmpty) tailValue() Value {
	return tailValue(v)
}


func (v *vEmpty) isArray() bool {
	return isArray(v)
}

func (v *vEmpty) getArray() []Value {
	return getArray(v)
}

func (v *vEmpty) isDict() bool {
	return isDict(v)
}

func (v *vEmpty) getDict() map[string]Value {
	return getDict(v)
}


func (v *vEmpty) isRef() bool {
	return isRef(v)
}

func (v *vEmpty) getValue() Value {
	return getValue(v)
}

func (v *vEmpty) setValue(cv Value) {
	setValue(v, cv)
}

