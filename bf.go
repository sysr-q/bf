package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"container/list"
)

type Brainfuck struct {
	Instructions []byte
	Pointer, At int
	Memory []byte
	Loops *list.List
}

func (bf *Brainfuck) PeekInstruction() (next byte) {
	if bf.Pointer >= len(bf.Instructions) {
		next = 0
	} else {
		next = bf.Instructions[bf.Pointer]
	}
	return
}

func (bf *Brainfuck) NextInstruction() (next byte) {
	next = bf.PeekInstruction()
	bf.Pointer++
	return
}

func (bf *Brainfuck) IsZero() (bool) {
	return bf.Memory[bf.At] == 0
}

//// BF instructions as functions!
// +
func (bf *Brainfuck) Plus() {
	*&bf.Memory[bf.At]++
}

// -
func (bf *Brainfuck) Minus() {
	*&bf.Memory[bf.At]--
}

// >
func (bf *Brainfuck) Next() {
	if bf.At >= len(bf.Memory)-1 {
		*&bf.At = 0
	} else {
		*&bf.At++
	}
}

// <
func (bf *Brainfuck) Last() {
	if bf.At <= 0 {
		*&bf.At = len(bf.Memory) - 1
	} else {
		*&bf.At--
	}
}

// [
func (bf *Brainfuck) Open() {
	if bf.IsZero() {
		// TODO: Skip to closing bracket?
	} else {
		// Push pointer to loop stack
		(*&bf.Loops).PushBack(*&bf.Pointer)
	}
}

// ]
func (bf *Brainfuck) Close() {
	if bf.IsZero() {
		// Pop off the loop stack and continue.
		(*&bf.Loops).Remove((*&bf.Loops).Back())
	} else {
		// Go back to the opening.
		*&bf.Pointer = (*&bf.Loops).Back().Value.(int)
	}
}

// .
func (bf *Brainfuck) Print() {
	fmt.Printf("%c", bf.Memory[bf.At])
}

// ,
func (bf *Brainfuck) ReadByte() {
	b := make([]byte, 1)
	if _, err := os.Stdin.Read(b); err != nil {
		panic(err) // programming
	} else {
		*&bf.Memory[bf.At] = b[0]
	}
}


func main() {
	var bf Brainfuck
	if text, err := ioutil.ReadFile(os.Args[1]); err == nil {
		bf = Brainfuck{
			// Strip the newline Stdin gives us.
			Instructions: text[:len(text)-1],
			Pointer: 0,
			At: 0,
			Memory: make([]byte, 30),
			Loops: list.New(),
		}
	} else {
		panic(err) // programming
	}
	for next := bf.NextInstruction(); next != 0; next = bf.NextInstruction() {
		inst := string(next)
		// TODO: less wasteful to make a map[string]func?
		if inst == "+" {
			bf.Plus()
		} else if inst == "-" {
			bf.Minus()
		} else if inst == ">" {
			bf.Next()
		} else if inst == "<" {
			bf.Last()
		} else if inst == "[" {
			bf.Open()
		} else if inst == "]" {
			bf.Close()
		} else if inst == "." {
			bf.Print()
		} else if inst == "," {
			bf.ReadByte()
		}
	}
}
