package symtable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEntryOperation(t *testing.T) {
	assert := assert.New(t)

	table := NewTable()
	table.AddVariable("TEST1")
	v := table.GetAddress("TEST1")
	assert.Equal(16, v)

	table.AddVariable("TEST2")
	v = table.GetAddress("TEST2")
	assert.Equal(17, v)
}

func TestContains(t *testing.T) {
	assert := assert.New(t)

	table := NewTable()

	ok := table.Contains("TEST1")
	assert.Equal(false, ok)

	table.AddVariable("TEST1")
	ok = table.Contains("TEST1")
	assert.Equal(true, ok)
}
