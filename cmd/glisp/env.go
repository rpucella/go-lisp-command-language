package main

import "fmt"

type Env struct {
	bindings map[string]Value
	previous *Env
}

func (env *Env) find(name string) (Value, error) {
	current := env
	for current != nil {
		val, ok := current.bindings[name]
		if ok {
			return val, nil
		}
		current = current.previous
	}
	return nil, fmt.Errorf("no such identifier %s", name)
}

func (env *Env) lookup(module string, name string) (Value, error) {
	moduleEnv, ok := env, true
	v, ok := moduleEnv.bindings[name]
	if !ok {
		return nil, fmt.Errorf("no such identifier %s", name)
	}
	return v, nil
}
 
func (env *Env) update(name string, v Value) {
	env.bindings[name] = v
}

func (env *Env) layer(names []string, values []Value) *Env {
	// if values is nil or smaller than names, then
	// remaining names are bound to #nil
	bindings := map[string]Value{}
	for i, name := range names {
		if values != nil && i < len(values) {
			bindings[name] = values[i]
		} else {
			bindings[name] = &VNil{}
		}
	}
	return &Env{bindings: bindings, previous: env}
}

