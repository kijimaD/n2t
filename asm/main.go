package main

import (
	"bufio"
	"io"
	"os"

	"github.com/kijimaD/n2t/asm/parser"
	"github.com/kijimaD/n2t/asm/symtable"
)

func main() {
}

type asm struct {
	romAddr  int
	symtable symtable.Symtable
}

func NewASM() asm {
	return asm{
		romAddr:  0,
		symtable: symtable.NewTable(),
	}
}

func (a *asm) run() {
	f, ferr := os.Open("prog.asm")
	if ferr != nil {
		panic(ferr)
	}
	buf := bufio.NewReaderSize(f, 1024)
	pg := parser.NewPG(buf)

	// シンボルテーブル追加
	for {
		_, err := pg.Advance()
		if err == io.EOF {
			break
		}
		switch pg.CommandType() {
		case parser.C_COMMAND:
			a.romAddr++
		case parser.A_COMMAND:
			a.romAddr++
		case parser.L_COMMAND:
			a.symtable.AddEntry(pg.Symbol(), a.romAddr)
		}
	}

	// FIXME: Readlineのカウンタをリセットしている
	f, ferr = os.Open("prog.asm")
	if ferr != nil {
		panic(ferr)
	}
	buf = bufio.NewReaderSize(f, 1024)
	pg.IN = buf

	// シンボル解決
	for {
		_, err := pg.Advance()
		if err == io.EOF {
			break
		}
	}
}
