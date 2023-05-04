package parser

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

type pg struct {
	CC    string // current command
	Bufin bufio.Reader
}

func NewPG(in io.Reader) *pg {
	bufin := bufio.NewReader(in)

	return &pg{
		CC:    "",
		Bufin: *bufin,
	}
}

type commandKind int

const (
	A_COMMAND commandKind = iota // アドレス命令 @xxx
	C_COMMAND                    // 計算命令 dest=comp;jump // dest もしくは jump はない可能性がある
	L_COMMAND                    // 疑似コマンド (xxx) シンボルxxxを定義する。アセンブリプログラムの他のところで用いることができる。定義される前の行でも使用できる
)

const C_COMMAND_REGEXP = `(?:(A?M?D?.*)=)?([^;]+)(?:;(.+))?`

func (pg *pg) Advance() (string, error) {
	var e error
	for {
		line, err := pg.Bufin.ReadString('\n')

		if err == io.EOF {
			e = err
			pg.CC = ""
			break
		}

		str := strings.TrimSuffix(string(line), "\n")
		idx := strings.Index(str, "//")
		if idx != -1 {
			str = str[:idx]
		}

		str = strings.TrimSpace(str)
		if str != "" {
			pg.CC = str
			break
		}
	}
	return pg.CC, e
}

func (pg *pg) CommandType() commandKind {
	if pg.CC[0] == '@' {
		return A_COMMAND
	} else if pg.CC[0] == '(' {
		return L_COMMAND
	} else {
		return C_COMMAND
	}
}

func (pg *pg) Symbol() string {
	var str string
	switch pg.CommandType() {
	case A_COMMAND:
		str = pg.CC[1:]
	case L_COMMAND:
		str = pg.CC[1 : len(pg.CC)-1]
	default:
		panic("can't symbolize!")
	}

	return str
}

// dest=comp;jump
func (pg *pg) Dest() string {
	var str string
	if pg.CommandType() == C_COMMAND {
		r := regexp.MustCompile(C_COMMAND_REGEXP)
		result := r.FindAllStringSubmatch(pg.CC, -1)
		str = result[0][1]
	} else {
		panic("this is not C command!")
	}
	return str
}

func (pg *pg) Comp() string {
	var str string
	if pg.CommandType() == C_COMMAND {
		r := regexp.MustCompile(C_COMMAND_REGEXP)
		result := r.FindAllStringSubmatch(pg.CC, -1)
		str = result[0][2]
	} else {
		panic("this is not C command!")
	}
	return str
}

func (pg *pg) Jump() string {
	var str string
	if pg.CommandType() == C_COMMAND {
		r := regexp.MustCompile(C_COMMAND_REGEXP)
		result := r.FindAllStringSubmatch(pg.CC, -1)
		str = result[0][3]
	} else {
		panic("this is not C command!")
	}
	return str
}
