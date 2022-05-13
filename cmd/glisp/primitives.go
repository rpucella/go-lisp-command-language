package main

import "fmt"
import "strings"

type Primitive struct {
	name string
	min  int
	max  int // <0 for no max #
	prim func(string, []Value) (Value, error)
}

func listLength(v Value) int {
	current := v
	result := 0
	for current.isCons() {
		result += 1
		current = current.tailValue()
	}
	return result
}

func listAppend(v1 Value, v2 Value) Value {
	current := v1
	var result Value = nil
	var current_result MutableCons = nil
	for current.isCons() {
		cell := NewMutableCons(current.headValue(), nil)
		current = current.tailValue()
		if current_result == nil {
			result = cell
		} else {
			current_result.setTail(cell)
		}
		current_result = cell
	}
	if current_result == nil {
		return v2
	}
	current_result.setTail(v2)
	return result
}

func allConses(vs []Value) bool {
	for _, v := range vs {
		if !v.isCons() {
			return false
		}
	}
	return true
}

func corePrimitives() map[string]Value {
	bindings := map[string]Value{}
	for _, d := range CORE_PRIMITIVES {
		bindings[d.name] = NewPrimitive(d.name, MakePrimitive(d))
	}
	return bindings
}

func MakePrimitive(d Primitive) func([]Value) (Value, error) {
	f := func(args []Value) (Value, error) {
		if err := checkMinArgs(d.name, args, d.min); err != nil {
			return nil, err
		}
		if d.max >= 0 {
			if err := checkMaxArgs(d.name, args, d.max); err != nil {
				return nil, err
			}
		}
		return d.prim(d.name, args)
	}
	return f
}

func checkArgType(name string, arg Value, pred func(Value) bool) error {
	if !pred(arg) {
		return fmt.Errorf("%s - wrong argument type %s", name, arg.typ())
	}
	return nil
}

func checkMinArgs(name string, args []Value, n int) error {
	if len(args) < n {
		return fmt.Errorf("%s - too few arguments %d", name, len(args))
	}
	return nil
}

func checkMaxArgs(name string, args []Value, n int) error {
	if len(args) > n {
		return fmt.Errorf("%s - too many arguments %d", name, len(args))
	}
	return nil
}

func checkExactArgs(name string, args []Value, n int) error {
	if len(args) != n {
		return fmt.Errorf("%s - wrong number of arguments %d", name, len(args))
	}
	return nil
}

func isInt(v Value) bool {
	return v.isNumber()
}

func isString(v Value) bool {
	return v.isString()
}

func isSymbol(v Value) bool {
	return v.isSymbol()
}

func isFunction(v Value) bool {
	return v.isFunction()
}

func isList(v Value) bool {
	return v.isCons() || v.isEmpty()
}

func isReference(v Value) bool {
	return v.isRef()
}

func mkNumPredicate(pred func(int, int) bool) func(string, []Value) (Value, error) {
	return func(name string, args []Value) (Value, error) {
		if err := checkExactArgs(name, args, 2); err != nil {
			return nil, err
		}
		if err := checkArgType(name, args[0], isInt); err != nil {
			return nil, err
		}
		if err := checkArgType(name, args[1], isInt); err != nil {
			return nil, err
		}
		return NewBoolean(pred(args[0].intValue(), args[1].intValue())), nil
	}
}

var CORE_PRIMITIVES = []Primitive{

	Primitive{
		"type", 1, 1,
		func(name string, args []Value) (Value, error) {
			return NewSymbol(args[0].typ()), nil
		},
	},

	Primitive{
		"+", 0, -1,
		func(name string, args []Value) (Value, error) {
			v := 0
			for _, arg := range args {
				if err := checkArgType(name, arg, isInt); err != nil {
					return nil, err
				}
				v += arg.intValue()
			}
			return NewInteger(v), nil
		},
	},

	Primitive{
		"*", 0, -1,
		func(name string, args []Value) (Value, error) {
			v := 1
			for _, arg := range args {
				if err := checkArgType(name, arg, isInt); err != nil {
					return nil, err
				}
				v *= arg.intValue()
			}
			return NewInteger(v), nil
		},
	},

	Primitive{
		"-", 1, -1,
		func(name string, args []Value) (Value, error) {
			v := args[0].intValue()
			if len(args) > 1 {
				for _, arg := range args[1:] {
					if err := checkArgType(name, arg, isInt); err != nil {
						return nil, err
					}
					v -= arg.intValue()
				}
			} else {
				v = -v
			}
			return NewInteger(v), nil
		},
	},

	Primitive{"=", 2, -1,
		func(name string, args []Value) (Value, error) {
			var reference Value = args[0]
			for _, v := range args[1:] {
				if !reference.isEqual(v) {
					return NewBoolean(false), nil
				}
			}
			return NewBoolean(true), nil
		},
	},

	Primitive{"<", 2, 2,
		mkNumPredicate(func(n1 int, n2 int) bool { return n1 < n2 }),
	},

	Primitive{"<=", 2, 2,
		mkNumPredicate(func(n1 int, n2 int) bool { return n1 <= n2 }),
	},

	Primitive{">", 2, 2,
		mkNumPredicate(func(n1 int, n2 int) bool { return n1 > n2 }),
	},

	Primitive{">=", 2, 2,
		mkNumPredicate(func(n1 int, n2 int) bool { return n1 >= n2 }),
	},

	Primitive{"not", 1, 1,
		func(name string, args []Value) (Value, error) {
			return NewBoolean(!args[0].isTrue()), nil
		},
	},

	Primitive{
		"string-append", 0, -1,
		func(name string, args []Value) (Value, error) {
			v := ""
			for _, arg := range args {
				if err := checkArgType(name, arg, isString); err != nil {
					return nil, err
				}
				v += arg.strValue()
			}
			return NewSymbol(v), nil
		},
	},

	Primitive{"string-length", 1, 1,
		func(name string, args []Value) (Value, error) {
			if err := checkArgType(name, args[0], isString); err != nil {
				return nil, err
			}
			return NewInteger(len(args[0].strValue())), nil
		},
	},

	Primitive{"string-lower", 1, 1,
		func(name string, args []Value) (Value, error) {
			if err := checkArgType(name, args[0], isString); err != nil {
				return nil, err
			}
			return NewSymbol(strings.ToLower(args[0].strValue())), nil
		},
	},

	Primitive{"string-upper", 1, 1,
		func(name string, args []Value) (Value, error) {
			if err := checkArgType(name, args[0], isString); err != nil {
				return nil, err
			}
			return NewSymbol(strings.ToUpper(args[0].strValue())), nil
		},
	},

	Primitive{"string-substring", 1, 3,
		func(name string, args []Value) (Value, error) {
			if err := checkArgType(name, args[0], isString); err != nil {
				return nil, err
			}
			start := 0
			end := len(args[0].strValue())
			if len(args) > 2 {
				if err := checkArgType(name, args[2], isInt); err != nil {
					return nil, err
				}
				end = min(args[2].intValue(), end)
			}
			if len(args) > 1 {
				if err := checkArgType(name, args[1], isInt); err != nil {
					return nil, err
				}
				start = max(args[1].intValue(), start)
			}
			// or perhaps raise an exception
			if end < start {
				return NewSymbol(""), nil
			}
			return NewSymbol(args[0].strValue()[start:end]), nil
		},
	},

	Primitive{"apply", 2, 2,
		func(name string, args []Value) (Value, error) {
			if err := checkArgType(name, args[0], isFunction); err != nil {
				return nil, err
			}
			if err := checkArgType(name, args[1], isList); err != nil {
				return nil, err
			}
			arguments := make([]Value, listLength(args[1]))
			current := args[1]
			for i := range arguments {
				arguments[i] = current.headValue()
				current = current.tailValue()
			}
			if !current.isEmpty() {
				return nil, fmt.Errorf("%s - malformed list", name)
			}
			return args[0].apply(arguments)
		},
	},

	Primitive{"cons", 2, 2,
		func(name string, args []Value) (Value, error) {
			if err := checkArgType(name, args[1], isList); err != nil {
				return nil, err
			}
			return NewCons(args[0], args[1]), nil
		},
	},

	Primitive{
		"append", 0, -1,
		func(name string, args []Value) (Value, error) {
			if len(args) == 0 {
				return NewEmpty(), nil
			}
			if err := checkArgType(name, args[len(args)-1], isList); err != nil {
				return nil, err
			}
			result := args[len(args)-1]
			for i := len(args) - 2; i >= 0; i -= 1 {
				if err := checkArgType(name, args[i], isList); err != nil {
					return nil, err
				}
				result = listAppend(args[i], result)
			}
			return result, nil
		},
	},

	Primitive{"reverse", 1, 1,
		func(name string, args []Value) (Value, error) {
			if err := checkArgType(name, args[0], isList); err != nil {
				return nil, err
			}
			var result Value = NewEmpty()
			current := args[0]
			for current.isCons() {
				result = NewCons(current.headValue(), result)
				current = current.tailValue()
			}
			if !current.isEmpty() {
				return nil, fmt.Errorf("%s - malformed list", name)
			}
			return result, nil
		},
	},

	Primitive{"head", 1, 1,
		func(name string, args []Value) (Value, error) {
			if err := checkArgType(name, args[0], isList); err != nil {
				return nil, err
			}
			if args[0].isEmpty() {
				return nil, fmt.Errorf("%s - empty list argument", name)
			}
			return args[0].headValue(), nil
		},
	},

	Primitive{"tail", 1, 1,
		func(name string, args []Value) (Value, error) {
			if err := checkArgType(name, args[0], isList); err != nil {
				return nil, err
			}
			if args[0].isEmpty() {
				return nil, fmt.Errorf("%s - empty list argument", name)
			}
			return args[0].tailValue(), nil
		},
	},

	Primitive{"list", 0, -1,
		func(name string, args []Value) (Value, error) {
			var result Value = NewEmpty()
			for i := len(args) - 1; i >= 0; i -= 1 {
				result = NewCons(args[i], result)
			}
			return result, nil
		},
	},

	Primitive{"length", 1, 1,
		func(name string, args []Value) (Value, error) {
			if err := checkArgType(name, args[0], isList); err != nil {
				return nil, err
			}
			count := 0
			current := args[0]
			for current.isCons() {
				count += 1
				current = current.tailValue()
			}
			if !current.isEmpty() {
				return nil, fmt.Errorf("%s - malformed list", name)
			}
			return NewInteger(count), nil
		},
	},

	Primitive{"nth", 2, 2,
		func(name string, args []Value) (Value, error) {
			if err := checkArgType(name, args[0], isList); err != nil {
				return nil, err
			}
			if err := checkArgType(name, args[1], isInt); err != nil {
				return nil, err
			}
			idx := args[1].intValue()
			if idx >= 0 {
				current := args[0]
				for current.isCons() {
					if idx == 0 {
						return current.headValue(), nil
					} else {
						idx -= 1
						current = current.tailValue()
					}
				}
			}
			return nil, fmt.Errorf("%s - index %d out of bound", name, args[1].intValue())
		},
	},

	Primitive{"map", 2, -1,
		func(name string, args []Value) (Value, error) {
			if err := checkArgType(name, args[0], isFunction); err != nil {
				return nil, err
			}
			for i := range args[1:] {
				if err := checkArgType(name, args[i+1], isList); err != nil {
					return nil, err
				}
			}
			var result Value = nil
			var current_result MutableCons = nil
			currents := make([]Value, len(args)-1)
			firsts := make([]Value, len(args)-1)
			for i := range args[1:] {
				currents[i] = args[i+1]
			}
			for allConses(currents) {
				for i := range currents {
					firsts[i] = currents[i].headValue()
				}
				v, err := args[0].apply(firsts)
				if err != nil {
					return nil, err
				}
				cell := NewMutableCons(v, nil)
				if current_result == nil {
					result = cell
				} else {
					current_result.setTail(cell)
				}
				current_result = cell
				for i := range currents {
					currents[i] = currents[i].tailValue()
				}
			}
			if current_result == nil {
				return NewEmpty(), nil
			}
			current_result.setTail(NewEmpty())
			return result, nil
		},
	},

	Primitive{"for", 2, -1,
		func(name string, args []Value) (Value, error) {
			if err := checkArgType(name, args[0], isFunction); err != nil {
				return nil, err
			}
			// TODO - allow different types in the same iteration!
			for i := range args[1:] {
				if err := checkArgType(name, args[i+1], isList); err != nil {
					return nil, err
				}
			}
			currents := make([]Value, len(args)-1)
			firsts := make([]Value, len(args)-1)
			for i := range args[1:] {
				currents[i] = args[i+1]
			}
			for allConses(currents) {
				for i := range currents {
					firsts[i] = currents[i].headValue()
				}
				_, err := args[0].apply(firsts)
				if err != nil {
					return nil, err
				}
				for i := range currents {
					currents[i] = currents[i].tailValue()
				}
			}
			return NewNil(), nil
		},
	},

	Primitive{"filter", 2, 2,
		func(name string, args []Value) (Value, error) {
			if err := checkArgType(name, args[0], isFunction); err != nil {
				return nil, err
			}
			if err := checkArgType(name, args[1], isList); err != nil {
				return nil, err
			}
			var result Value = nil
			var current_result MutableCons = nil
			current := args[1]
			for current.isCons() {
				v, err := args[0].apply([]Value{current.headValue()})
				if err != nil {
					return nil, err
				}
				if v.isTrue() {
					cell := NewMutableCons(current.headValue(), nil)
					if current_result == nil {
						result = cell
					} else {
						current_result.setTail(cell)
					}
					current_result = cell
				}
				current = current.tailValue()
			}
			if !current.isEmpty() {
				return nil, fmt.Errorf("%s - malformed list", name)
			}
			if current_result == nil {
				return NewEmpty(), nil
			}
			current_result.setTail(NewEmpty())
			return result, nil
		},
	},

	Primitive{"foldr", 3, 3,
		func(name string, args []Value) (Value, error) {
			if err := checkArgType(name, args[0], isFunction); err != nil {
				return nil, err
			}
			if err := checkArgType(name, args[1], isList); err != nil {
				return nil, err
			}
			var temp Value = NewEmpty()
			// first reverse the list
			current := args[1]
			for current.isCons() {
				temp = NewCons(current.headValue(), temp)
				current = current.tailValue()
			}
			if !current.isEmpty() {
				return nil, fmt.Errorf("%s - malformed list", name)
			}
			// then fold it
			result := args[2]
			current = temp
			for current.isCons() {
				v, err := args[0].apply([]Value{current.headValue(), result})
				if err != nil {
					return nil, err
				}
				result = v
				current = current.tailValue()
			}
			if !current.isEmpty() {
				return nil, fmt.Errorf("%s - malformed list", name)
			}
			return result, nil
		},
	},

	Primitive{"foldl", 3, 3,
		func(name string, args []Value) (Value, error) {
			if err := checkArgType(name, args[0], isFunction); err != nil {
				return nil, err
			}
			if err := checkArgType(name, args[1], isList); err != nil {
				return nil, err
			}
			result := args[2]
			current := args[1]
			for current.isCons() {
				v, err := args[0].apply([]Value{result, current.headValue()})
				if err != nil {
					return nil, err
				}
				result = v
				current = current.tailValue()
			}
			if !current.isEmpty() {
				return nil, fmt.Errorf("%s - malformed list", name)
			}
			return result, nil
		},
	},

	Primitive{"ref", 1, 1,
		func(name string, args []Value) (Value, error) {
			return NewReference(args[0]), nil
		},
	},

	// set should probably be a special form
	// (set (x) 10)
	// (set (arr 1) 10)
	// (set (dict 'a) 10)
	// like setf in CLISP

	// Primitive{"set", 2, 2,
	// 	func(name string, args []Value) (Value, error) {
	// 		if err := checkArgType(name, args[0], isReference); err != nil {
	// 			return nil, err
	// 		}
	// 		args[0].setValue(args[1])
	// 		return NewNil(), nil
	// 	},
	// },

	Primitive{"empty?", 1, 1,
		func(name string, args []Value) (Value, error) {
			return NewBoolean(args[0].isEmpty()), nil
		},
	},

	Primitive{"cons?", 1, 1,
		func(name string, args []Value) (Value, error) {
			return NewBoolean(args[0].isCons()), nil
		},
	},

	Primitive{"list?", 1, 1,
		func(name string, args []Value) (Value, error) {
			return NewBoolean(args[0].isCons() || args[0].isEmpty()), nil
		},
	},

	Primitive{"number?", 1, 1,
		func(name string, args []Value) (Value, error) {
			return NewBoolean(args[0].isNumber()), nil
		},
	},

	Primitive{"ref?", 1, 1,
		func(name string, args []Value) (Value, error) {
			return NewBoolean(args[0].isRef()), nil
		},
	},

	Primitive{"boolean?", 1, 1,
		func(name string, args []Value) (Value, error) {
			return NewBoolean(args[0].isBool()), nil
		},
	},

	Primitive{"string?", 1, 1,
		func(name string, args []Value) (Value, error) {
			return NewBoolean(args[0].isString()), nil
		},
	},

	Primitive{"symbol?", 1, 1,
		func(name string, args []Value) (Value, error) {
			return NewBoolean(args[0].isSymbol()), nil
		},
	},

	Primitive{"function?", 1, 1,
		func(name string, args []Value) (Value, error) {
			return NewBoolean(args[0].isFunction()), nil
		},
	},

	Primitive{"nil?", 1, 1,
		func(name string, args []Value) (Value, error) {
			return NewBoolean(args[0].isNil()), nil
		},
	},

	Primitive{"array", 0, -1,
		func(name string, args []Value) (Value, error) {
			content := make([]Value, len(args))
			for i, v := range args {
				content[i] = v
			}
			return NewArray(content), nil
		},
	},

	Primitive{"array?", 1, 1,
		func(name string, args []Value) (Value, error) {
			return NewBoolean(args[0].isArray()), nil
		},
	},

	Primitive{"dict", 0, -1,
		func(name string, args []Value) (Value, error) {
			content := make(map[string]Value, len(args))
			for _, v := range args {
				if !v.isCons() || !v.tailValue().isCons() || !v.tailValue().tailValue().isEmpty() {
					return nil, fmt.Errorf("dict item not a pair - %s", v.Display())
				}
				if !v.headValue().isSymbol() {
					return nil, fmt.Errorf("dict key is not a symbol - %s", v.headValue().Display())
				}
				content[v.headValue().strValue()] = v.tailValue().headValue()
			}
			return NewDict(content), nil
		},
	},

	Primitive{"dict?", 1, 1,
		func(name string, args []Value) (Value, error) {
			return NewBoolean(args[0].isDict()), nil
		},
	},

	Primitive{
		"quit", 0, 0,
		func(name string, args []Value) (Value, error) {
			bail()
			return NewNil(), nil
		},
	},
}

// left:
//
// dictionaries #((a 1) (b 2))  (dict '(a 10) '(b 20) '(c 30))  vs (apply dict '((a 10) (b 20) (c 30)))?
// arrays #[a b c]
