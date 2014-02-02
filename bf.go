package main

import (
	"os"
	"fmt"
	"io/ioutil"
)

type Brainfuck struct {
	Instructions []byte
	Pointer, At int
	Memory, Loops []int
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

//// BF instructions as functions!
// +
func (bf *Brainfuck) Plus() {
	bf.Memory[bf.At]++
	return
}

// -
func (bf *Brainfuck) Minus() {
	bf.Memory[bf.At]--
	return
}

// >
func (bf *Brainfuck) Next() {
	bf.At++
	return
}

// <
func (bf *Brainfuck) Last() {
	bf.At--
	return
}

// [
func (bf *Brainfuck) Open() {
	fmt.Print("!")
	return
}

// ]
func (bf *Brainfuck) Close() {
	fmt.Print("?")
	return
}

// .
func (bf *Brainfuck) Print() {
	fmt.Printf("%c", bf.Memory[bf.At])
	return
}


func main() {
	var bf Brainfuck
	if text, err := ioutil.ReadAll(os.Stdin); err == nil {
		bf = Brainfuck{
			// Strip the newline Stdin gives us.
			Instructions: text[:len(text)-1],
			Pointer: 0,
			At: 0,
			Memory: make([]int, 30),
		}
	} else {
		panic(err) // programming
	}
	for next := bf.NextInstruction(); next != 0; next = bf.NextInstruction() {
		inst := string(next)
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
		}
	}
}
