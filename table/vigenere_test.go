package table_test

import (
	"extended-vigenere-cipher/table"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestVigenereTable_Set(t *testing.T) {

	charset := table.DefaultCharset()
	vt := table.NewVigenereTable(charset, charset, nil)
	require.NotEmpty(t, vt)

	for i := 0; i < len(vt.Cell); i++ {
		a := i // for iterate in tableType
		for j := 0; j < len(vt.Cell[i]); j++ {
			if a == len(vt.Cell[i]) {
				a = 0
			}
			require.Equal(t, charset[a], vt.Cell[i][j])
			a++
		}
	}
}

func TestVigenereTable_Generate(t *testing.T) {
	var vt table.VigenereTable

	c := vt.Generate(table.UpperCaseCharset(), table.UpperCaseCharset())
	assert.Equal(t, len(table.UpperCaseCharset()), len(c))
}

func TestVigenereTable_GetSubTextIndex(t *testing.T) {
	var vt table.VigenereTable

	charset := table.UpperCaseCharset()
	vt.TextCharset = charset

	i, err := vt.GetSubTextIndex("B")
	require.NoError(t, err)
	assert.Equal(t, 1, i)
}

func TestVigenereTable_GetSubKeyIndex(t *testing.T) {
	var vt table.VigenereTable

	charset := table.UpperCaseCharset()
	vt.KeyCharset = charset

	i, err := vt.GetSubKeyIndex("A")
	require.NoError(t, err)
	assert.Equal(t, 0, i)
}
func TestVigenereTable_GenerateYumnamTable(t *testing.T) {

	et, dt, err := table.GenerateYumnamTable(3, table.DefaultCharset(), true)
	require.NoError(t, err)
	et.Print()
	fmt.Println("====")
	dt.Print()

	et, dt, err = table.GenerateYumnamTable(3, table.DefaultCharset(), false)
	require.NoError(t, err)
}
