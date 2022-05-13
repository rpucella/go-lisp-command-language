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

func (v *vFunction) intValue() int {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vFunction) strValue() string {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vFunction) boolValue() bool {
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

func (v *vFunction) headValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vFunction) tailValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
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

func (v *vFunction) isRef() bool {
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

func (v *vFunction) getValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vFunction) setValue(cv Value) {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vFunction) isArray() bool {
	return false
}

func (v *vFunction) getArray() []Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vFunction) isDict() bool {
	return false
}

func (v *vFunction) getDict() map[string]Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

