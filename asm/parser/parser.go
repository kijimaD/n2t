package parser

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

type pg struct {
	cc string // current command
	in io.Reader
}

func NewPG(in io.Reader) *pg {
	return &pg{
		cc: "",
		in: in,
	}
}

type commandKind int

const (
	A_COMMAND commandKind = iota // @xxx
	C_COMMAND                    // dest=comp;jump // dest もしくは jump はない可能性がある
	L_COMMAND                    // (xxx) 疑似コマンド。中身はシンボル
)

const C_COMMAND_REGEXP = `(?:(A?M?D?.*)=)?([^;]+)(?:;(.+))?`

func (pg *pg) Advance() (string, error) {
	var e error
	for {
		bu := bufio.NewReaderSize(pg.in, 1024)
		line, _, err := bu.ReadLine()
		if err == io.EOF {
			e = err
			pg.cc = ""
			break
		}

		str := strings.TrimSuffix(string(line), "\n")
		idx := strings.Index(str, "//")
		if idx != -1 {
			str = str[:idx]
		}

		str = strings.TrimSpace(str)
		if str != "" {
			pg.cc = str
			break
		}
	}
	return pg.cc, e
}

func (pg *pg) commandType() commandKind {
	if pg.cc[0] == '@' {
		return A_COMMAND
	} else if pg.cc[0] == '(' {
		return L_COMMAND
	} else {
		return C_COMMAND
	}
}

func (pg *pg) symbol() string {
	var str string
	switch pg.commandType() {
	case A_COMMAND:
		str = pg.cc[1:]
	case L_COMMAND:
		str = pg.cc[1 : len(pg.cc)-1]
	default:
		panic("can't symbolize!")
	}

	return str
}

// dest=comp;jump
func (pg *pg) dest() string {
	var str string
	if pg.commandType() == C_COMMAND {
		r := regexp.MustCompile(C_COMMAND_REGEXP)
		result := r.FindAllStringSubmatch(pg.cc, -1)
		str = result[0][1]
	} else {
		panic("this is not C command!")
	}
	return str
}

func (pg *pg) comp() string {
	var str string
	if pg.commandType() == C_COMMAND {
		r := regexp.MustCompile(C_COMMAND_REGEXP)
		result := r.FindAllStringSubmatch(pg.cc, -1)
		str = result[0][2]
	} else {
		panic("this is not C command!")
	}
	return str
}

func (pg *pg) jump() string {
	var str string
	if pg.commandType() == C_COMMAND {
		r := regexp.MustCompile(C_COMMAND_REGEXP)
		result := r.FindAllStringSubmatch(pg.cc, -1)
		str = result[0][3]
	} else {
		panic("this is not C command!")
	}
	return str
}
