package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"hash"
)

type filter struct {
	bitfield [100]bool
}

var (
	hasher = sha1.New()
)

func main() {
	fmt.Println("vim-go")
	f := filter{}
	f.set("HELLO")
	f.set("hello1")
	f.set("world")

	if f.get("world") {
		fmt.Println("world is true")
	}
	//	fmt.Printf("%v\n", f.bitfield)
}

func (f *filter) set(s string) {
	f.bitfield[f.hashPosition(s)] = true
}

func (f *filter) get(s string) bool {
	return f.bitfield[f.hashPosition(s)]
}

func (f *filter) hashPosition(s string) int {
	hs := createHash(hasher, s)
	if hs < 0 {
		hs = -hs
	}
	return hs % len(f.bitfield)
}

func createHash(h hash.Hash, input string) int {
	bits := h.Sum([]byte(input))
	println(bits)
	buffer := bytes.NewBuffer(bits)
	result, _ := binary.ReadVarint(buffer)
	println("Input String: ", input, "Result: ", int(result))
	return int(result) // cast the int64
}
