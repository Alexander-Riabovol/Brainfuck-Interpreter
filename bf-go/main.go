package main

import (
	"github.com/Alexander-Riabovol/Brainfuck-Interpreter/brainfuck"
)

func main() {
	brainfuck.Interpret("++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.c>p.")
}
