package table

import (
	"fmt"
	"github.com/yogamandayu/extended-vigenere-cipher/util"
)

type VigenereTableInterface interface {
	Generate(textCharset, keyCharset []string) [][]string
	GetSubTextIndex(subText string) (int, error)
	GetSubKeyIndex(subKey string) (int, error)
	Print()
}

/*
This is how I imagine the concept of vigenere table.
I separated between key and text charset with the cell.
It will make the vigenere table have a flexibility of charset.
For example, you can randomize the key/text charset or the table too.

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

// NewVigenereTable is a constructor. if table is not provided, it will generate by key and text charset
func NewVigenereTable(textCharset, keyCharset []string, table [][]string) *VigenereTable {
	var v VigenereTable
	if len(table) == 0 {
		v.Cell = v.Generate(textCharset, keyCharset)
	}
	v.TextCharset = textCharset
	v.KeyCharset = keyCharset
	return &v
}

// Generate will generate vigenere table from a charset. Different position of charset will generate different table.
// Example : charset with A-Z & 0-9 | 0-9 & A-Z. will generate different table.
// Generated cell depends on TextCharset.
func (v *VigenereTable) Generate(textCharset, keyCharset []string) [][]string {
	cell := make([][]string, len(keyCharset))
	for i := 0; i < len(keyCharset); i++ {
		cell[i] = make([]string, len(textCharset))

		for j := 0; j < len(textCharset); j++ {
			x := j + (i % len(textCharset))

			if x > len(textCharset)-1 {
				x -= len(textCharset)
			}

			cell[i][j] = textCharset[x]
		}
	}
	return cell
}

// GetSubTextIndex is to get index location of specific text character.
func (v *VigenereTable) GetSubTextIndex(subText string) (int, error) {
	for i, s := range v.TextCharset {
		if s == subText {
			return i, nil
		}
	}
	return 0, fmt.Errorf("extended_vigenere_cipher.error.subtext_not_found")
}

// GetSubKeyIndex is to get index location of specific key character.
func (v *VigenereTable) GetSubKeyIndex(subKey string) (int, error) {
	for i, s := range v.KeyCharset {
		if s == subKey {
			return i, nil
		}
	}
	return 0, fmt.Errorf("extended_vigenere_cipher.error.subkey_not_found")
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

// GenerateYumnamTable will generate random vigenere table. it will return encrypt and decrypt table.
func GenerateYumnamTable(seed int, charset []string, colUnique bool) (encryptTable *VigenereTable, decryptTable *VigenereTable, err error) {

	encryptTable = NewVigenereTable(charset, charset, nil)
	decryptTable = NewVigenereTable(charset, charset, nil)
	size := len(encryptTable.Cell)

	arr := util.TwoDimensionArrayRandomInteger(seed, size)
	if colUnique {
		// Column Unique.
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				temp := encryptTable.Cell[j][i]
				encryptTable.Cell[j][i] = encryptTable.Cell[arr[i][j]][i]
				encryptTable.Cell[arr[i][j]][i] = temp
			}
		}

		//Formula Yumnan Cell : Dt (Et (M, N), N) = M
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				cell := encryptTable.Cell[i][j]

				et, err := encryptTable.GetSubKeyIndex(cell)
				if err != nil {
					return nil, nil, err
				}

				m := encryptTable.TextCharset[i]
				decryptTable.Cell[et][j] = m
			}
		}
	} else {
		// Row Unique.
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				temp := encryptTable.Cell[i][j]
				encryptTable.Cell[i][j] = encryptTable.Cell[i][arr[i][j]]
				encryptTable.Cell[i][(arr)[i][j]] = temp
			}
		}

		//Formula Yumnan Cell : Dt (M,Et(M,N)) = N
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				cell := encryptTable.Cell[i][j]

				et, err := encryptTable.GetSubKeyIndex(cell)
				if err != nil {
					return nil, nil, err
				}

				n := encryptTable.TextCharset[j]
				decryptTable.Cell[i][et] = n
			}
		}
	}
	return
}
