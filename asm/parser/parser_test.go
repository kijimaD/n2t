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

	f, _ := os.Open("a.asm")
	buf := bufio.NewReaderSize(f, 1024)

	pg := NewPG(buf)

	v, _ := pg.advance()
	assert.Equal("@i", v)
	v, _ = pg.advance()
	assert.Equal("M=1", v)
	v, _ = pg.advance()
	assert.Equal("@sum", v)
	v, _ = pg.advance()
	assert.Equal("M=0", v)
	v, _ = pg.advance()
	assert.Equal("M=0", v)
	v, _ = pg.advance()
	assert.Equal("M=0", v)

	v, err := pg.advance()
	assert.Equal("", v)
	assert.Equal(io.EOF, err)
}

func TestCommandType(t *testing.T) {
	assert := assert.New(t)

	buf := bytes.NewBufferString("dummy")
	pg := NewPG(buf)

	pg.cc = "@sum"
	v := pg.commandType()
	assert.Equal(A_COMMAND, v)

	pg.cc = "dest=comp;"
	v = pg.commandType()
	assert.Equal(C_COMMAND, v)

	pg.cc = "(sym)"
	v = pg.commandType()
	assert.Equal(L_COMMAND, v)
}

func TestSymbol(t *testing.T) {
	assert := assert.New(t)

	buf := bytes.NewBufferString("dummy")
	pg := NewPG(buf)

	pg.cc = "@sum"
	v := pg.symbol()
	assert.Equal("sum", v)

	pg.cc = "(sym)"
	v = pg.symbol()
	assert.Equal("sym", v)
}

func TestDest(t *testing.T) {
	assert := assert.New(t)

	buf := bytes.NewBufferString("dummy")
	pg := NewPG(buf)

	pg.cc = "dest=comp;jump"
	v := pg.dest()
	assert.Equal("dest", v)

	pg.cc = "comp;jump"
	v = pg.dest()
	assert.Equal("", v)
}

func TestComp(t *testing.T) {
	assert := assert.New(t)

	buf := bytes.NewBufferString("dummy")
	pg := NewPG(buf)

	pg.cc = "dest=comp;jump"
	v := pg.comp()
	assert.Equal("comp", v)
}

func TestJump(t *testing.T) {
	assert := assert.New(t)

	buf := bytes.NewBufferString("dummy")
	pg := NewPG(buf)

	pg.cc = "dest=comp;jump"
	v := pg.jump()
	assert.Equal("jump", v)

	pg.cc = "dest=comp"
	v = pg.jump()
	assert.Equal("", v)
}
