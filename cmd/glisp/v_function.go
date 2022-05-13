package main

import (
	"fmt"
	"strings"
)

type vFunction struct {
	params []string
	body   ast
	env    *Env
}

func NewFunction(params []string, body ast, env *Env) Value {
	return &vFunction{params, body, env}
}

func (v *vFunction) Display() string {
	return fmt.Sprintf("#<fun %s ...>", strings.Join(v.params, " "))
}

func (v *vFunction) DisplayCDR() string {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vFunction) apply(args []Value) (Value, error) {
	if len(v.params) != len(args) {
		return nil, fmt.Errorf("Wrong number of arguments to application to %s", v.str())
	}
	newEnv := layer(v.env, v.params, args)
	return v.body.eval(newEnv)
}

func (v *vFunction) str() string {
	return fmt.Sprintf("VFunction[[%s] %s]", strings.Join(v.params, " "), v.body.str())
}

func (v *vFunction) isAtom() bool {
	return false
}

func (v *vFunction) isSymbol() bool {
	return false
}

func (v *vFunction) isCons() bool {
	return false
}

func (v *vFunction) isEmpty() bool {
	return false
}

func (v *vFunction) isNumber() bool {
	return false
}

func (v *vFunction) isBool() bool {
	return false
}

func (v *vFunction) isString() bool {
	return false
}

func (v *vFunction) isFunction() bool {
	return true
}

func (v *vFunction) isTrue() bool {
	return true
}

func (v *vFunction) isNil() bool {
	return false
}

func (v *vFunction) isEqual(vv Value) bool {
	return v == vv // pointer equality
}

func (v *vFunction) typ() string {
	return "fun"
}

func (v *vFunction) asInteger() (int, bool) {
	return 0, false
}

func (v *vFunction) asBoolean() (bool, bool) {
	return false, false
}

func (v *vFunction) asString() (string, bool) {
	return "", false
}

func (v *vFunction) asSymbol() (string, bool) {
	return "", false
}

func (v *vFunction) asCons() (Value, Value, bool) {
	return nil, nil, false
}

func (v *vFunction) asReference() (Value, func(Value), bool) {
	return nil, nil, false
}

func (v *vFunction) setReference(Value) bool {
	return false
}

func (v *vFunction) asArray() ([]Value, bool) {
	return nil, false
}

func (v *vFunction) asDict() (map[string]Value, bool) {
	return nil, false
}


func (v *vFunction) intValue() int {
	return intValue(v)
}

func (v *vFunction) strValue() string {
	return strValue(v)
}

func (v *vFunction) boolValue() bool {
	return boolValue(v)
}

func (v *vFunction) headValue() Value {
	return headValue(v)
}

func (v *vFunction) tailValue() Value {
	return tailValue(v)
}


func (v *vFunction) isArray() bool {
	return isArray(v)
}

func (v *vFunction) getArray() []Value {
	return getArray(v)
}

func (v *vFunction) isDict() bool {
	return isDict(v)
}

func (v *vFunction) getDict() map[string]Value {
	return getDict(v)
}


func (v *vFunction) isRef() bool {
	return isRef(v)
}

func (v *vFunction) getValue() Value {
	return getValue(v)
}

func (v *vFunction) setValue(cv Value) {
	setValue(v, cv)
}

