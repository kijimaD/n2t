package parser

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdvance(t *testing.T) {
	assert := assert.New(t)

	f, _ := os.Open("../test.asm")
	buf := bufio.NewReaderSize(f, 1024)

	pg := NewPG(buf)

	v, _ := pg.Advance()
	assert.Equal("@i", v)
	v, _ = pg.Advance()
	assert.Equal("M=1", v)
	v, _ = pg.Advance()
	assert.Equal("@sum", v)
	v, _ = pg.Advance()
	assert.Equal("M=0", v)
	v, _ = pg.Advance()
	assert.Equal("M=0", v)
	v, _ = pg.Advance()
	assert.Equal("M=0", v)

	v, err := pg.Advance()
	assert.Equal("", v)
	assert.Equal(io.EOF, err)
}

func TestCommandType(t *testing.T) {
	assert := assert.New(t)

	buf := bytes.NewBufferString("dummy")
	pg := NewPG(buf)

	pg.CC = "@sum"
	v := pg.CommandType()
	assert.Equal(A_COMMAND, v)

	pg.CC = "dest=comp;"
	v = pg.CommandType()
	assert.Equal(C_COMMAND, v)

	pg.CC = "(sym)"
	v = pg.CommandType()
	assert.Equal(L_COMMAND, v)
}

func TestSymbol(t *testing.T) {
	assert := assert.New(t)

	buf := bytes.NewBufferString("dummy")
	pg := NewPG(buf)

	pg.CC = "@sum"
	v := pg.Symbol()
	assert.Equal("sum", v)

	pg.CC = "(sym)"
	v = pg.Symbol()
	assert.Equal("sym", v)
}

func TestDest(t *testing.T) {
	assert := assert.New(t)

	buf := bytes.NewBufferString("dummy")
	pg := NewPG(buf)

	pg.CC = "dest=comp;jump"
	v := pg.Dest()
	assert.Equal("dest", v)

	pg.CC = "comp;jump"
	v = pg.Dest()
	assert.Equal("", v)
}

func TestComp(t *testing.T) {
	assert := assert.New(t)

	buf := bytes.NewBufferString("dummy")
	pg := NewPG(buf)

	pg.CC = "dest=comp;jump"
	v := pg.Comp()
	assert.Equal("comp", v)
}

func TestJump(t *testing.T) {
	assert := assert.New(t)

	buf := bytes.NewBufferString("dummy")
	pg := NewPG(buf)

	pg.CC = "dest=comp;jump"
	v := pg.Jump()
	assert.Equal("jump", v)

	pg.CC = "dest=comp"
	v = pg.Jump()
	assert.Equal("", v)
}
