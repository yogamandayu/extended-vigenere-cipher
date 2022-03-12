package extended_vigenere_cipher

import (
	"extended-vigenere-cipher/table"
	"fmt"
	"strings"
)

type VigenereCipherInterface interface {
	Encrypt(plaintext, key string) (string, error)
	Decrypt(ciphertext, key string) (string, error)
}

type VigenereCipher struct {
	VigenereTable *table.VigenereTable
}

func NewVigenereCipher(table *table.VigenereTable) *VigenereCipher {
	return &VigenereCipher{
		VigenereTable: table,
	}
}

var _ VigenereCipherInterface = &VigenereCipher{}

// Encrypt method do the traditional way.
func (v VigenereCipher) Encrypt(plaintext, key string) (string, error) {
	if v.VigenereTable == nil {
		return "", fmt.Errorf("vigenere_cipher.error.missing_table")
	}

	var ciphertext strings.Builder
	var ik int // iterating through key purpose.
	for _, r := range plaintext {
		s := string(r)
		k := string(key[ik])

		indexS, indexK := -1, -1
		for j := 0; j < len((v.VigenereTable.Cell)[0]); j++ {
			if (v.VigenereTable.Cell)[0][j] == s {
				indexS = j
				break
			}
		}
		if indexS == -1 {
			return "", fmt.Errorf("vigenere_cipher.error.plaintext_char_not_found")
		}

		for j := 0; j < len(v.VigenereTable.Cell); j++ {
			if (v.VigenereTable.Cell)[0][j] == k {
				indexK = j
				break
			}
		}
		if indexK == -1 {
			return "", fmt.Errorf("vigenere_cipher.error.key_char_not_found")
		}

		ciphertext.WriteString((v.VigenereTable.Cell)[indexK][indexS])
		ik++
		if ik >= len(key) {
			ik = 0
		}
	}

	return ciphertext.String(), nil
}

// Decrypt method do the traditional way.
func (v VigenereCipher) Decrypt(ciphertext, key string) (string, error) {
	if v.VigenereTable == nil {
		return "", fmt.Errorf("vigenere_cipher.error.missing_table")
	}

	var plaintext strings.Builder
	var ik int // iterating through key purpose.
	for _, r := range ciphertext {
		s := string(r)
		k := string(key[ik])

		indexS, indexK := -1, -1

		for j := 0; j < len(v.VigenereTable.Cell); j++ {
			if (v.VigenereTable.Cell)[0][j] == k {
				indexK = j
				break
			}
		}
		if indexK == -1 {
			return "", fmt.Errorf("vigenere_cipher.error.key_char_not_found")
		}

		for j := 0; j < len((v.VigenereTable.Cell)[indexK]); j++ {
			if (v.VigenereTable.Cell)[indexK][j] == s {
				indexS = j
				break
			}
		}
		if indexS == -1 {
			return "", fmt.Errorf("vigenere_cipher.error.ciphertext_char_not_found")
		}

		plaintext.WriteString((v.VigenereTable.Cell)[0][indexS])
		ik++
		if ik >= len(key) {
			ik = 0
		}
	}

	return plaintext.String(), nil
}
