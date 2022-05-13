package main

import "fmt"
import "bufio"
import "os"
import "strings"
import "io"

type Engine struct {
	env *Env
}

func NewEngine() Engine {
	coreBindings := corePrimitives()
	coreBindings["true"] = &VBoolean{true}
	coreBindings["false"] = &VBoolean{false}
	env := &Env{bindings: coreBindings, previous: nil}
	return Engine{env}
}

// TODO: engine.Read()
// TODO: engine.Eval()
// TODO: engine.ReadEval()

// TODO: engine.DefineConstant()
// TODO: engine.DefineFunction()
// TODO: engine.DefineMacro()

// TODO: make prompt a function (of what?)

// TODO: what do we export? Engine, Value

func (e Engine) Repl(prompt string) {
	env := e.env
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s> ", prompt)
		text, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println()
				bail()
			}
			fmt.Println("IO ERROR - ", err.Error())
		}
		if strings.TrimSpace(text) == "" {
			continue
		}
		v, _, err := read(text)
		if err != nil {
			fmt.Println("READ ERROR -", err.Error())
			continue
		}
		// check if it's a declaration
		d, err := parseDef(v)
		if err != nil {
			fmt.Println("PARSE ERROR -", err.Error())
			continue
		}
		if d != nil {
			if d.typ == DEF_FUNCTION {
				update(env, d.name, &VFunction{d.params, d.body, env})
				fmt.Println(d.name)
				continue
			}
			if d.typ == DEF_VALUE {
				v, err := d.body.eval(env)
				if err != nil {
					fmt.Println("EVAL ERROR -", err.Error())
					continue
				}
				update(env, d.name, v)
				fmt.Println(d.name)
				continue
			}
			fmt.Println("DECLARE ERROR - unknow declaration type", d.typ)
			continue
		}
		// check if it's an expression
		e, err := parseExpr(v)
		if err != nil {
			fmt.Println("PARSE ERROR -", err.Error())
			continue
		}
		///fmt.Println("expr =", e.str())
		v, err = e.eval(env)
		if err != nil {
			fmt.Println("EVAL ERROR -", err.Error())
			continue
		}
		if !v.isNil() {
			fmt.Println(v.display())
		}
	}
}

func bail() {
	os.Exit(0)
}
