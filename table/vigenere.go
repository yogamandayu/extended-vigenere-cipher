package table

import "fmt"

type VigenereTableInterface interface {
	Generate() error
	Set(table [][]string)
	Print()
}

/*
This is how I imagine the concept of vigenere table.
I separated between key and text charset with the cell.
It will make the vigenere table have a flexibility of charset.
For example, you can randomize the key and text charset or the table too.

Key			Text			Cell
0|A|		0|A|			ABCDEF
1|B|		1|B|			BCDEFA
2|C|		2|C|			CDEFAB
3|D|		3|D|			DEFABC
4|E|		4|E|			EFABCD
5|F|		5|F|			FABCDE

A,A = Cell[0][0] = A
D,C = Cell[3][2] = F

*/

// VigenereTable is main entity of vigenere table.
// KeyCharset represent each row.
// TextCharset represent each column.
type VigenereTable struct {
	TextCharset []string
	KeyCharset  []string
	Cell        [][]string
}

var _ VigenereTableInterface = &VigenereTable{}

func NewVigenereTable(textCharset, keyCharset []string) *VigenereTable {
	return &VigenereTable{
		TextCharset: textCharset,
		KeyCharset:  keyCharset,
	}
}

// Generate will generate vigenere table from a charset. Different position of charset will generate different table.
// Example : charset with A-Z & 0-9 | 0-9 & A-Z. will generate different table.
// Generated cell depends on TextCharset.
func (v *VigenereTable) Generate() error {
	if len(v.TextCharset) == 0 {
		return fmt.Errorf("vigenere_cipher.error.missing_text_charset")
	}
	if len(v.KeyCharset) == 0 {
		return fmt.Errorf("vigenere_cipher.error.missing_key_charset")
	}
	if len(v.KeyCharset) > len(v.TextCharset) {
		return fmt.Errorf("vigenere_cipher.error.invalid_charset_length")
	}

	v.Cell = make([][]string, len(v.KeyCharset))
	for i := 0; i < len(v.KeyCharset); i++ {
		v.Cell[i] = make([]string, len(v.TextCharset))

		for j := 0; j < len(v.TextCharset); j++ {
			x := j + i

			if x > len(v.TextCharset)-1 {
				x -= len(v.TextCharset)
			}

			v.Cell[i][j] = v.TextCharset[x]
		}
	}
	return nil
}

// Set is to set vigenere table with custom table.
// Example : row and column have different charset.
func (v *VigenereTable) Set(table [][]string) {
	v.Cell = table
}

// Print will print vigenere table.
func (v *VigenereTable) Print() {
	for i := 0; i < len(v.Cell); i++ {
		for j := 0; j < len((v.Cell)[i]); j++ {
			fmt.Print((v.Cell)[i][j])
		}
		fmt.Println()
	}
}
