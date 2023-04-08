package code

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDest(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string
		expect string
	}{
		{
			input:  "null",
			expect: "000",
		},
		{
			input:  "M",
			expect: "001",
		},
		{
			input:  "AMD",
			expect: "111",
		},
	}

	for idx, tt := range tests {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			got := dest(tt.input)
			assert.Equal(tt.expect, got)
		})
	}
}

func TestComp(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string
		expect string
	}{
		{
			input:  "0",
			expect: "0101010",
		},
		{
			input:  "1",
			expect: "0111111",
		},
		{
			input:  "D|M",
			expect: "1010101",
		},
	}

	for idx, tt := range tests {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			got := comp(tt.input)
			assert.Equal(tt.expect, got)
		})
	}
}

func TestJump(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string
		expect string
	}{
		{
			input:  "null",
			expect: "000",
		},
		{
			input:  "JGT",
			expect: "001",
		},
		{
			input:  "JMP",
			expect: "111",
		},
	}

	for idx, tt := range tests {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			got := jump(tt.input)
			assert.Equal(tt.expect, got)
		})
	}
}
