package main

import (
	"fmt"
	"strings"
)

type vArray struct {
	content []Value
}

func NewArray(vs []Value) Value {
	return &vArray{vs}
}

func (v *vArray) Display() string {
	s := make([]string, len(v.content))
	for i, vv := range v.content {
		s[i] = vv.Display()
	}
	return fmt.Sprintf("#[%s]", strings.Join(s, " "))
}

func (v *vArray) DisplayCDR() string {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vArray) intValue() int {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vArray) strValue() string {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vArray) boolValue() bool {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vArray) apply(args []Value) (Value, error) {
	if len(args) < 1 || !args[0].isNumber() {
		return nil, fmt.Errorf("array indexing requires an index")
	}
	if len(args) > 2 {
		return nil, fmt.Errorf("too many arguments %d to array update", len(args))
	}
	idx := args[0].intValue()
	if idx < 0 || idx >= len(v.content) {
		return nil, fmt.Errorf("array index out of bounds %d", idx)
	}
	if len(args) == 2 {
		v.content[idx] = args[1]
		return &vNil{}, nil
	}
	return v.content[idx], nil
}

func (v *vArray) str() string {
	s := make([]string, len(v.content))
	for i, vv := range v.content {
		s[i] = vv.str()
	}
	return fmt.Sprintf("VArray[%s]", strings.Join(s, " "))
}

func (v *vArray) headValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vArray) tailValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vArray) isAtom() bool {
	return false // ?
}

func (v *vArray) isSymbol() bool {
	return false
}

func (v *vArray) isCons() bool {
	return false
}

func (v *vArray) isEmpty() bool {
	return false
}

func (v *vArray) isNumber() bool {
	return false
}

func (v *vArray) isBool() bool {
	return false
}

func (v *vArray) isRef() bool {
	return false
}

func (v *vArray) isString() bool {
	return false
}

func (v *vArray) isFunction() bool {
	return false
}

func (v *vArray) isTrue() bool {
	return false
}

func (v *vArray) isNil() bool {
	return false
}

func (v *vArray) isEqual(vv Value) bool {
	return v == vv // pointer equality because arrays will be mutable
	/*
		if !vv.isArray() || len(v.content) != len(vv.getArray()) {
			return false}
		vvcontent := vv.getArray()
		for i := range(v.content) {
			if !v.content[i].isEqual(vvcontent[i]) {
				return false
			}
		}
		return true
	*/
}

func (v *vArray) typ() string {
	return "array"
}

func (v *vArray) getValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vArray) setValue(cv Value) {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vArray) isArray() bool {
	return true
}

func (v *vArray) getArray() []Value {
	return v.content
}

func (v *vArray) isDict() bool {
	return false
}

func (v *vArray) getDict() map[string]Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}
