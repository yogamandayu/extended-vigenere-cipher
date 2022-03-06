package table

import "fmt"

type VigenereTableInterface interface {
	Generate(charset []string)
	Set(table [][]string)
	Print()
}

type VigenereTable [][]string

var _ VigenereTableInterface = &VigenereTable{}

// Generate will generate vigenere table from a charset. Different position of charset will generate different table.
// Example : charset with A-Z & 0-9 | 0-9 & A-Z. will generate different table.
func (v *VigenereTable) Generate(charset []string) {

	size := len(charset)

	*v = make([][]string, size)

	for i := 0; i < size; i++ {
		(*v)[i] = make([]string, size)

		for j := 0; j < size; j++ {
			x := j + i

			if x > len(charset)-1 {
				x -= len(charset)
			}

			(*v)[i][j] = charset[x]
		}
	}
}

// Set is to set vigenere table with custom table.
// Example : row and column have different charset.
func (v *VigenereTable) Set(table [][]string) {
	*v = table
}

// Print will print vigenere table.
func (v *VigenereTable) Print() {
	for i := 0; i < len(*v); i++ {
		for j := 0; j < len((*v)[i]); j++ {
			fmt.Print((*v)[i][j])
		}
		fmt.Println()
	}
}
