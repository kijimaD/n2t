package main

import (
	"fmt"
	"os"

	"github.com/kijimaD/n2t/asm/asm"
)

func main() {
	input := fmt.Sprintf("%s.asm", os.Args[1])
	output := fmt.Sprintf("%s.hack", os.Args[1])

	outfile, err := os.Create(output)
	if err != nil {
		panic(err)
	}

	a := asm.NewASM(input, outfile)
	a.Run()
}
