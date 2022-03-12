package table_test

import (
	"extended-vigenere-cipher/table"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestVigenereTable_Set(t *testing.T) {

	charset := table.DefaultCharset()
	vt := table.NewVigenereTable(charset, charset)

	err := vt.Generate()
	require.NoError(t, err)
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
	vt.Print()
}
