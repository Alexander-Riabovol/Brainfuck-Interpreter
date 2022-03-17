package brainfuck

import "fmt"

type command interface {
	execute(mem *memmory)
}

type operation struct {
	mem *memmory
}

type incrementOperation struct{ operation }            // +
type decrementOperation struct{ operation }            // -
type incrementDataPointerOperation struct{ operation } // >
type decrementDataPointerOperation struct{ operation } // <
type outputOperation struct{ operation }               // .

func (op *incrementOperation) execute() {
	op.mem.cells[op.mem.pointer]++
}

func (op *decrementOperation) execute() {
	op.mem.cells[op.mem.pointer]--
}
func (op *incrementDataPointerOperation) execute() {
	op.mem.pointer++
}

func (op *decrementDataPointerOperation) execute() {
	op.mem.pointer--
}

func (op *outputOperation) execute() {
	fmt.Printf("%c", op.mem.cells[op.mem.pointer])
}
