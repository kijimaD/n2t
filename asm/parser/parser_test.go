package parser

import (
	"bufio"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasMoreCommands(t *testing.T) {
}

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

}

func TestSymbol(t *testing.T) {

}

func TestDest(t *testing.T) {

}

func TestComp(t *testing.T) {

}

func TestJump(t *testing.T) {

}
