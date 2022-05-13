package main

import (
	"fmt"
	"strings"
)

type vDict struct {
	content map[string]Value
}

func NewDict(vs map[string]Value) Value {
	return &vDict{vs}
}

func (v *vDict) Display() string {
	s := make([]string, len(v.content))
	i := 0
	for k, vv := range v.content {
		s[i] = fmt.Sprintf("(%s %s)", k, vv.Display())
		i++
	}
	return fmt.Sprintf("#(%s)", strings.Join(s, " "))
}

func (v *vDict) DisplayCDR() string {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vDict) apply(args []Value) (Value, error) {
	if len(args) < 1 || !args[0].isSymbol() {
		return nil, fmt.Errorf("dict indexing requires a key")
	}
	if len(args) > 2 {
		return nil, fmt.Errorf("too many arguments %d to dict update", len(args))
	}
	key := args[0].strValue()
	if len(args) == 2 {
		v.content[key] = args[1]
		return &vNil{}, nil
	}
	result, ok := v.content[key]
	if !ok {
		return nil, fmt.Errorf("key %s not in dict", key)
	}
	return result, nil
}

func (v *vDict) str() string {
	s := make([]string, len(v.content))
	i := 0
	for k, vv := range v.content {
		s[i] = fmt.Sprintf("[%s %s]", k, vv.str())
		i++
	}
	return fmt.Sprintf("VDict[%s]", strings.Join(s, " "))
}

func (v *vDict) isAtom() bool {
	return false // ?
}

func (v *vDict) isSymbol() bool {
	return false
}

func (v *vDict) isCons() bool {
	return false
}

func (v *vDict) isEmpty() bool {
	return false
}

func (v *vDict) isNumber() bool {
	return false
}

func (v *vDict) isBool() bool {
	return false
}

func (v *vDict) isString() bool {
	return false
}

func (v *vDict) isFunction() bool {
	return false
}

func (v *vDict) isTrue() bool {
	return false
}

func (v *vDict) isNil() bool {
	return false
}

func (v *vDict) isEqual(vv Value) bool {
	return v == vv // pointer equality due to mutability
}

func (v *vDict) typ() string {
	return "dict"
}

func (v *vDict) asInteger() (int, bool) {
	return 0, false
}

func (v *vDict) asBoolean() (bool, bool) {
	return false, false
}

func (v *vDict) asString() (string, bool) {
	return "", false
}

func (v *vDict) asSymbol() (string, bool) {
	return "", false
}

func (v *vDict) asCons() (Value, Value, bool) {
	return nil, nil, false
}

func (v *vDict) asReference() (Value, func(Value), bool) {
	return nil, nil, false
}

func (v *vDict) setReference(Value) bool {
	return false
}

func (v *vDict) asArray() ([]Value, bool) {
	return nil, false
}

func (v *vDict) asDict() (map[string]Value, bool) {
	return v.content, true
}


func (v *vDict) intValue() int {
	return intValue(v)
}

func (v *vDict) strValue() string {
	return strValue(v)
}

func (v *vDict) boolValue() bool {
	return boolValue(v)
}

func (v *vDict) headValue() Value {
	return headValue(v)
}

func (v *vDict) tailValue() Value {
	return tailValue(v)
}


func (v *vDict) isArray() bool {
	return isArray(v)
}

func (v *vDict) getArray() []Value {
	return getArray(v)
}

func (v *vDict) isDict() bool {
	return isDict(v)
}

func (v *vDict) getDict() map[string]Value {
	return getDict(v)
}


func (v *vDict) isRef() bool {
	return isRef(v)
}

func (v *vDict) getValue() Value {
	return getValue(v)
}

func (v *vDict) setValue(cv Value) {
	setValue(v, cv)
}

