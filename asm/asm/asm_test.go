package asm

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun1(t *testing.T) {
	assert := assert.New(t)
	inbuf := strings.NewReader(`@i
	        M=1 // i=1
	        @sum
	        M=0
	(LOOP)
	        @i
	        D=M
	        @100
	        D=D-A
	        @END
	        D;JGT
	        @i
	        D=M
	        @sum
	        M=D+M
	        @i
	        M=M+1
	        @LOOP
	        0;JMP
	(END)
	        @END
	        0;JMP
	`)
	outbuf := &bytes.Buffer{}

	asm := NewASM(inbuf, outbuf)
	asm.Run()

	expect := `0000000000010000
1110111111001000
0000000000010001
1110101010001000
0000000000010000
1111110000010000
0000000001100100
1110010011010000
0000000000010010
1110001100000001
0000000000010000
1111110000010000
0000000000010001
1111000010001000
0000000000010000
1111110111001000
0000000000000100
1110101010000111
0000000000010010
1110101010000111
`
	assert.Equal(expect, outbuf.String())
}
