package table_test

import (
	"extended-vigenere-cipher/table"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestVigenereTable_Set(t *testing.T) {
	var vt table.VigenereTable

	charset := table.DefaultCharset()

	vt.Generate(charset)
	require.NotEmpty(t, vt)

	for i := 0; i < len(vt); i++ {
		a := i // for iterate in tableType
		for j := 0; j < len(vt[i]); j++ {
			if a == len(vt[i]) {
				a = 0
			}
			require.Equal(t, charset[a], vt[i][j])
			a++
		}
	}
	vt.Print()
}
