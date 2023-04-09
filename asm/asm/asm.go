package asm

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"

	"github.com/kijimaD/n2t/asm/code"
	"github.com/kijimaD/n2t/asm/parser"
	"github.com/kijimaD/n2t/asm/symtable"
)

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

// 数字だけの場合は値(10進)、文字が入っている場合はシンボル
var SYM_REGEXP = `([0-9]+)|([0-9a-zA-Z_\.\$:]+)`

func (a *asm) run() {
	f, ferr := os.Open("../prog.asm")
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
		case parser.A_COMMAND:
			a.romAddr++
		case parser.C_COMMAND:
			a.romAddr++
		case parser.L_COMMAND:
			a.symtable.AddEntry(pg.Symbol(), a.romAddr)
		}
	}

	// FIXME: Readlineのカウンタをリセットしている
	f, ferr = os.Open("../prog.asm")
	if ferr != nil {
		panic(ferr)
	}
	buf = bufio.NewReaderSize(f, 1024)
	pg.IN = buf
	r := regexp.MustCompile(SYM_REGEXP)

	// シンボル解決
	for {
		_, err := pg.Advance()
		if err == io.EOF {
			break
		}

		cmdType := pg.CommandType()
		var bincode string

		switch cmdType {
		case parser.A_COMMAND:
			result := r.FindAllStringSubmatch(pg.CC, -1)
			if result[0][1] != "" {
				// value
				value := result[0][1]
				i, _ := strconv.Atoi(value)
				bincode = fmt.Sprintf("%016b", i)
			} else if result[0][2] != "" {
				// symbol
				if a.symtable.Contains(pg.Symbol()) {
					address := a.symtable.GetAddress(pg.Symbol())
					bincode = fmt.Sprintf("%016b", address)
				} else {
					a.symtable.AddVariable(pg.Symbol())
					address := a.symtable.GetAddress(pg.Symbol())
					bincode = fmt.Sprintf("%016b", address)
				}
			}
		case parser.C_COMMAND:
			bincode = fmt.Sprintf("111" + code.Comp(pg.Comp()) + code.Dest(pg.Dest()) + code.Jump(pg.Jump()))
		}

		if cmdType != parser.L_COMMAND {
			// 書き込み
			fmt.Printf("v-> %#v %#v\n", bincode, pg.CC)
		}
	}
}