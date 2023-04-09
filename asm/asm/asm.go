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

// 数字だけの場合は値(10進)、文字が入っている場合はシンボル
var SYM_REGEXP = `([0-9]+)|([0-9a-zA-Z_\.\$:]+)`

type asm struct {
	romAddr  int
	symtable symtable.Symtable
	in       string // FIXME: Readerにしたいけど、2回読み込みでseekをリセットする方法がわからない
	out      io.Writer
}

func NewASM(in string, out io.Writer) asm {
	return asm{
		romAddr:  0,
		symtable: symtable.NewTable(),
		in:       in,
		out:      out,
	}
}

func (a *asm) Run() {
	f, ferr := os.Open(a.in)
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
			// なんでromAddrが必要なんだろう?
			// ラベルは定義された行数に名前をつけたものだから。
			// 値が入ったシンボルとは意味が異なる
			a.symtable.AddEntry(pg.Symbol(), a.romAddr)
		}
	}

	// FIXME: Readlineのカウンタをリセットしている
	f, ferr = os.Open(a.in)
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
			fmt.Fprintf(a.out, "%s\n", bincode)
		}
	}
}
