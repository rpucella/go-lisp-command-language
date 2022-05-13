package main

import "fmt"

func main() {
	fmt.Println("GoLisp Command Language Standalone Interpreter 1.0.0")
	//env := initialize()
	//shell(env)
	eng := NewEngine()
	eng.Repl("glisp")
}
