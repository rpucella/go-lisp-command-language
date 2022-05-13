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

func (v *vCons) intValue() int {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vCons) strValue() string {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vCons) boolValue() bool {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vCons) apply(args []Value) (Value, error) {
	return nil, fmt.Errorf("Value %s not applicable", v.str())
}

func (v *vCons) str() string {
	return fmt.Sprintf("VCons[%s %s]", v.head.str(), v.tail.str())
}

func (v *vCons) headValue() Value {
	return v.head
}

func (v *vCons) tailValue() Value {
	return v.tail
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

func (v *vCons) isRef() bool {
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

func (v *vCons) getValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vCons) setValue(cv Value) {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vCons) isArray() bool {
	return false
}

func (v *vCons) getArray() []Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vCons) isDict() bool {
	return false
}

func (v *vCons) getDict() map[string]Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

