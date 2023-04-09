package main

import "testing"

func TestRun(t *testing.T) {
	// assert = assert.New(t)

	asm := NewASM()
	asm.run()
}
