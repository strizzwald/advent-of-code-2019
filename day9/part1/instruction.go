package main

type instruction interface {
	OpCode() int
	Pointer() int64
	Offset() int64
	Execute(memory []int64, relativeOffset int64)
}
