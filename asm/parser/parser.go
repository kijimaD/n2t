package parser

import (
	"bufio"
	"fmt"
	"io"
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

func (pg *pg) advance() (string, error) {
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
			fmt.Println(str)

			break
		}
	}
	return pg.cc, e
}
