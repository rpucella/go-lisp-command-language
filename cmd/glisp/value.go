package main

import "fmt"
import "strings"

type Value interface {
	display() string
	displayCDR() string
	intValue() int
	boolValue() bool
	strValue() string
	headValue() Value
	tailValue() Value
	apply([]Value) (Value, error)
	str() string
	isAtom() bool
	isSymbol() bool
	isCons() bool
	isEmpty() bool
	isNumber() bool
	isBool() bool
	isRef() bool
	isString() bool
	isFunction() bool
	isTrue() bool
	isNil() bool
	//isEq() bool    -- don't think we need pointer equality for now - = is enough?
	isEqual(Value) bool
	typ() string
	getValue() Value
	setValue(Value)
	isArray() bool
	getArray() []Value
	isDict() bool
	getDict() map[string]Value
}

type vInteger struct {
	val int
}

func NewInteger(v int) Value {
	return &vInteger{v}
}

type vBoolean struct {
	val bool
}

func NewBoolean(v bool) Value {
	return &vBoolean{v}
}

type vPrimitive struct {
	name      string
	primitive func([]Value) (Value, error)
}

func NewPrimitive(name string, prim func([]Value) (Value, error)) Value {
	return &vPrimitive{name, prim}
}

type vEmpty struct {
}

func NewEmpty() Value {
	return &vEmpty{}
}

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

type vSymbol struct {
	name string
}

func NewSymbol(name string) Value {
	return &vSymbol{name}
}

type vFunction struct {
	params []string
	body   ast
	env    *Env
}

func NewFunction(params []string, body ast, env *Env) Value {
	return &vFunction{params, body, env}
}

type vString struct {
	val string
}

func NewString(v string) Value {
	return &vString{v}
}

type vNil struct {
}

func NewNil() Value {
	return &vNil{}
}

type vReference struct {
	content Value
}

func NewReference(v Value) Value {
	return &vReference{v}
}

type vArray struct {
	content []Value
}

func NewArray(vs []Value) Value {
	return &vArray{vs}
}

type vDict struct {
	content map[string]Value
}

func NewDict(vs map[string]Value) Value {
	return &vDict{vs}
}

func (v *vInteger) display() string {
	return fmt.Sprintf("%d", v.val)
}

func (v *vInteger) displayCDR() string {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vInteger) intValue() int {
	return v.val
}

func (v *vInteger) strValue() string {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vInteger) boolValue() bool {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vInteger) apply(args []Value) (Value, error) {
	return nil, fmt.Errorf("Value %s not applicable", v.str())
}

func (v *vInteger) str() string {
	return fmt.Sprintf("VInteger[%d]", v.val)
}

func (v *vInteger) headValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vInteger) tailValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vInteger) isAtom() bool {
	return true
}

func (v *vInteger) isSymbol() bool {
	return false
}

func (v *vInteger) isCons() bool {
	return false
}

func (v *vInteger) isEmpty() bool {
	return false
}

func (v *vInteger) isNumber() bool {
	return true
}

func (v *vInteger) isBool() bool {
	return false
}

func (v *vInteger) isRef() bool {
	return false
}

func (v *vInteger) isString() bool {
	return false
}

func (v *vInteger) isFunction() bool {
	return false
}

func (v *vInteger) isTrue() bool {
	return v.val != 0
}

func (v *vInteger) isNil() bool {
	return false
}

func (v *vInteger) isEqual(vv Value) bool {
	return vv.isNumber() && v.intValue() == vv.intValue()
}

func (v *vInteger) typ() string {
	return "int"
}

func (v *vInteger) getValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vInteger) setValue(cv Value) {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vInteger) isArray() bool {
	return false
}

func (v *vInteger) getArray() []Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vInteger) isDict() bool {
	return false
}

func (v *vInteger) getDict() map[string]Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vBoolean) display() string {
	if v.val {
		return "#t"
	} else {
		return "#f"
	}
}

func (v *vBoolean) displayCDR() string {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vBoolean) intValue() int {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vBoolean) strValue() string {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vBoolean) boolValue() bool {
	return v.val
}

func (v *vBoolean) apply(args []Value) (Value, error) {
	return nil, fmt.Errorf("Value %s not applicable", v.str())
}

func (v *vBoolean) str() string {
	if v.val {
		return "VBoolean[true]"
	} else {
		return "VBoolean[false]"
	}
}

func (v *vBoolean) headValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vBoolean) tailValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vBoolean) isAtom() bool {
	return true
}

func (v *vBoolean) isSymbol() bool {
	return false
}

func (v *vBoolean) isCons() bool {
	return false
}

func (v *vBoolean) isEmpty() bool {
	return false
}

func (v *vBoolean) isNumber() bool {
	return false
}

func (v *vBoolean) isBool() bool {
	return true
}

func (v *vBoolean) isRef() bool {
	return false
}

func (v *vBoolean) isString() bool {
	return false
}

func (v *vBoolean) isFunction() bool {
	return false
}

func (v *vBoolean) isTrue() bool {
	return v.val
}

func (v *vBoolean) isNil() bool {
	return false
}

func (v *vBoolean) isEqual(vv Value) bool {
	return vv.isBool() && v.boolValue() == vv.boolValue()
}

func (v *vBoolean) typ() string {
	return "bool"
}

func (v *vBoolean) getValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vBoolean) setValue(cv Value) {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vBoolean) isArray() bool {
	return false
}

func (v *vBoolean) getArray() []Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vBoolean) isDict() bool {
	return false
}

func (v *vBoolean) getDict() map[string]Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vPrimitive) display() string {
	return fmt.Sprintf("#<prim %s>", v.name)
}

func (v *vPrimitive) displayCDR() string {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vPrimitive) intValue() int {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vPrimitive) strValue() string {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vPrimitive) boolValue() bool {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vPrimitive) apply(args []Value) (Value, error) {
	return v.primitive(args)
}

func (v *vPrimitive) str() string {
	return fmt.Sprintf("VPrimitive[%s]", v.name)
}

func (v *vPrimitive) headValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vPrimitive) tailValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vPrimitive) isAtom() bool {
	return false
}

func (v *vPrimitive) isSymbol() bool {
	return false
}

func (v *vPrimitive) isCons() bool {
	return false
}

func (v *vPrimitive) isEmpty() bool {
	return false
}

func (v *vPrimitive) isNumber() bool {
	return false
}

func (v *vPrimitive) isBool() bool {
	return false
}

func (v *vPrimitive) isRef() bool {
	return false
}

func (v *vPrimitive) isString() bool {
	return false
}

func (v *vPrimitive) isFunction() bool {
	return true
}

func (v *vPrimitive) isTrue() bool {
	return true
}

func (v *vPrimitive) isNil() bool {
	return false
}

func (v *vPrimitive) isEqual(vv Value) bool {
	return v == vv // pointer equality
}

func (v *vPrimitive) typ() string {
	return "fun"
}

func (v *vPrimitive) getValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vPrimitive) setValue(cv Value) {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vPrimitive) isArray() bool {
	return false
}

func (v *vPrimitive) getArray() []Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vPrimitive) isDict() bool {
	return false
}

func (v *vPrimitive) getDict() map[string]Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vEmpty) display() string {
	return "()"
}

func (v *vEmpty) displayCDR() string {
	return ")"
}

func (v *vEmpty) intValue() int {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vEmpty) strValue() string {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vEmpty) boolValue() bool {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vEmpty) apply(args []Value) (Value, error) {
	return nil, fmt.Errorf("Value %s not applicable", v.str())
}

func (v *vEmpty) str() string {
	return fmt.Sprintf("VEmpty")
}

func (v *vEmpty) headValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vEmpty) tailValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
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

func (v *vEmpty) isRef() bool {
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

func (v *vEmpty) getValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vEmpty) setValue(cv Value) {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vEmpty) isArray() bool {
	return false
}

func (v *vEmpty) getArray() []Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vEmpty) isDict() bool {
	return false
}

func (v *vEmpty) getDict() map[string]Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vCons) display() string {
	return "(" + v.head.display() + v.tail.displayCDR()
}

func (v *vCons) displayCDR() string {
	return " " + v.head.display() + v.tail.displayCDR()
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

func (v *vSymbol) display() string {
	return v.name
}

func (v *vSymbol) displayCDR() string {
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

func (v *vFunction) display() string {
	return fmt.Sprintf("#<fun %s ...>", strings.Join(v.params, " "))
}

func (v *vFunction) displayCDR() string {
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

func (v *vString) display() string {
	return "\"" + v.val + "\""
}

func (v *vString) displayCDR() string {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vString) intValue() int {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vString) strValue() string {
	return v.val
}

func (v *vString) boolValue() bool {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vString) apply(args []Value) (Value, error) {
	return nil, fmt.Errorf("Value %s not applicable", v.str())
}

func (v *vString) str() string {
	return fmt.Sprintf("VString[%s]", v.str())
}

func (v *vString) headValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vString) tailValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vString) isAtom() bool {
	return true
}

func (v *vString) isSymbol() bool {
	return false
}

func (v *vString) isCons() bool {
	return false
}

func (v *vString) isEmpty() bool {
	return false
}

func (v *vString) isNumber() bool {
	return false
}

func (v *vString) isBool() bool {
	return false
}

func (v *vString) isRef() bool {
	return false
}

func (v *vString) isString() bool {
	return true
}

func (v *vString) isFunction() bool {
	return false
}

func (v *vString) isTrue() bool {
	return (v.val != "")
}

func (v *vString) isNil() bool {
	return false
}

func (v *vString) isEqual(vv Value) bool {
	return vv.isString() && v.strValue() == vv.strValue()
}

func (v *vString) typ() string {
	return "string"
}

func (v *vString) getValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vString) setValue(cv Value) {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vString) isArray() bool {
	return false
}

func (v *vString) getArray() []Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vString) isDict() bool {
	return false
}

func (v *vString) getDict() map[string]Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vNil) display() string {
	// figure out if this is the right thing?
	return "#nil"
}

func (v *vNil) displayCDR() string {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vNil) intValue() int {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vNil) strValue() string {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vNil) boolValue() bool {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vNil) apply(args []Value) (Value, error) {
	return nil, fmt.Errorf("Value %s not applicable", v.str())
}

func (v *vNil) str() string {
	return fmt.Sprintf("VNil")
}

func (v *vNil) headValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vNil) tailValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
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

func (v *vNil) isRef() bool {
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

func (v *vNil) getValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vNil) setValue(cv Value) {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vNil) isArray() bool {
	return false
}

func (v *vNil) getArray() []Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vNil) isDict() bool {
	return false
}

func (v *vNil) getDict() map[string]Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vReference) display() string {
	return fmt.Sprintf("#<ref %s>", v.content.display())
}

func (v *vReference) displayCDR() string {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vReference) intValue() int {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vReference) strValue() string {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vReference) boolValue() bool {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vReference) apply(args []Value) (Value, error) {
	if len(args) > 1 {
		return nil, fmt.Errorf("too many arguments %d to ref update", len(args))
	}
	if len(args) == 1 {
		v.content = args[0]
		return &vNil{}, nil
	}
	return v.content, nil
}

func (v *vReference) str() string {
	return fmt.Sprintf("VReference[%s]", v.content.str())
}

func (v *vReference) headValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vReference) tailValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vReference) isAtom() bool {
	return false // ?
}

func (v *vReference) isSymbol() bool {
	return false
}

func (v *vReference) isCons() bool {
	return false
}

func (v *vReference) isEmpty() bool {
	return false
}

func (v *vReference) isNumber() bool {
	return false
}

func (v *vReference) isBool() bool {
	return false
}

func (v *vReference) isRef() bool {
	return true
}

func (v *vReference) isString() bool {
	return false
}

func (v *vReference) isFunction() bool {
	return false
}

func (v *vReference) isTrue() bool {
	return false
}

func (v *vReference) isNil() bool {
	return false
}

func (v *vReference) isEqual(vv Value) bool {
	return v == vv // pointer equality
}

func (v *vReference) typ() string {
	return "reference"
}

func (v *vReference) getValue() Value {
	return v.content
}

func (v *vReference) setValue(cv Value) {
	v.content = cv
}

func (v *vReference) isArray() bool {
	return false
}

func (v *vReference) getArray() []Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vReference) isDict() bool {
	return false
}

func (v *vReference) getDict() map[string]Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vArray) display() string {
	s := make([]string, len(v.content))
	for i, vv := range v.content {
		s[i] = vv.display()
	}
	return fmt.Sprintf("#[%s]", strings.Join(s, " "))
}

func (v *vArray) displayCDR() string {
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

func (v *vDict) display() string {
	s := make([]string, len(v.content))
	i := 0
	for k, vv := range v.content {
		s[i] = fmt.Sprintf("(%s %s)", k, vv.display())
		i++
	}
	return fmt.Sprintf("#(%s)", strings.Join(s, " "))
}

func (v *vDict) displayCDR() string {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vDict) intValue() int {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vDict) strValue() string {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vDict) boolValue() bool {
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

func (v *vDict) headValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vDict) tailValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
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

func (v *vDict) isRef() bool {
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

func (v *vDict) getValue() Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vDict) setValue(cv Value) {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vDict) isArray() bool {
	return false
}

func (v *vDict) getArray() []Value {
	panic(fmt.Sprintf("unchecked access to %s", v.str()))
}

func (v *vDict) isDict() bool {
	return true
}

func (v *vDict) getDict() map[string]Value {
	return v.content
}
