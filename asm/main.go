package main

import (
	"fmt"
	"os"

	"github.com/kijimaD/n2t/asm/asm"
)

func main() {
	inname := fmt.Sprintf("%s.asm", os.Args[1])
	outname := fmt.Sprintf("%s.hack", os.Args[1])

	fp, err := os.Open(inname)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	outfile, err := os.Create(outname)
	if err != nil {
		panic(err)
	}

	a := asm.NewASM(fp, outfile)
	a.Run()
}
