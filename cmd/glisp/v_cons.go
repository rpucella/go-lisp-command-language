package main

import (
	"fmt"
)

type vCons struct {
	head   Value
	tail   Value
	length int    // Doesn't appear used.
}

func NewCons(head Value, tail Value) Value {
	return &vCons{head: head, tail: tail}
}

type MutableCons = *vCons

func NewMutableCons(head Value, tail Value) MutableCons {
	return &vCons{head: head, tail: tail}
}

func (v MutableCons) setTail(tail Value) {
	v.tail = tail
}

func (v *vCons) Display() string {
	return "(" + v.head.Display() + v.tail.DisplayCDR()
}

func (v *vCons) DisplayCDR() string {
	return " " + v.head.Display() + v.tail.DisplayCDR()
}

func (v *vCons) apply(args []Value) (Value, error) {
	return nil, fmt.Errorf("Value %s not applicable", v.str())
}

func (v *vCons) str() string {
	return fmt.Sprintf("VCons[%s %s]", v.head.str(), v.tail.str())
}

func (v *vCons) isAtom() bool {
	return false
}

func (v *vCons) isSymbol() bool {
	return false
}

func (v *vCons) isCons() bool {
	return true
}

func (v *vCons) isEmpty() bool {
	return false
}

func (v *vCons) isNumber() bool {
	return false
}

func (v *vCons) isBool() bool {
	return false
}

func (v *vCons) isString() bool {
	return false
}

func (v *vCons) isFunction() bool {
	return false
}

func (v *vCons) isTrue() bool {
	return true
}

func (v *vCons) isNil() bool {
	return false
}

func (v *vCons) isEqual(vv Value) bool {
	if !vv.isCons() {
		return false
	}
	var curr1 Value = v
	var curr2 Value = vv
	for curr1.isCons() {
		if !curr2.isCons() {
			return false
		}
		if !curr1.headValue().isEqual(curr2.headValue()) {
			return false
		}
		curr1 = curr1.tailValue()
		curr2 = curr2.tailValue()
	}
	return curr1.isEqual(curr2) // should both be empty at the end
}

func (v *vCons) typ() string {
	return "list"
}

func (v *vCons) asInteger() (int, bool) {
	return 0, false
}

func (v *vCons) asBoolean() (bool, bool) {
	return false, false
}

func (v *vCons) asString() (string, bool) {
	return "", false
}

func (v *vCons) asSymbol() (string, bool) {
	return "", false
}

func (v *vCons) asCons() (Value, Value, bool) {
	return v.head, v.tail, true
}

func (v *vCons) asReference() (Value, func(Value), bool) {
	return nil, nil, false
}

func (v *vCons) setReference(Value) bool {
	return false
}

func (v *vCons) asArray() ([]Value, bool) {
	return nil, false
}

func (v *vCons) asDict() (map[string]Value, bool) {
	return nil, false
}


func (v *vCons) intValue() int {
	return intValue(v)
}

func (v *vCons) strValue() string {
	return strValue(v)
}

func (v *vCons) boolValue() bool {
	return boolValue(v)
}

func (v *vCons) headValue() Value {
	return headValue(v)
}

func (v *vCons) tailValue() Value {
	return tailValue(v)
}


func (v *vCons) isArray() bool {
	return isArray(v)
}

func (v *vCons) getArray() []Value {
	return getArray(v)
}

func (v *vCons) isDict() bool {
	return isDict(v)
}

func (v *vCons) getDict() map[string]Value {
	return getDict(v)
}


func (v *vCons) isRef() bool {
	return isRef(v)
}

func (v *vCons) getValue() Value {
	return getValue(v)
}

func (v *vCons) setValue(cv Value) {
	setValue(v, cv)
}

