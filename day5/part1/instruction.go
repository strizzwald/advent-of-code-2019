
package main

type instruction interface {
	OpCode() int
	Pointer() int
	Offset() int
	Execute(memory []int)
}