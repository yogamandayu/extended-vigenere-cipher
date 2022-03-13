package extended_vigenere_cipher

import (
	"extended-vigenere-cipher/table"
	"extended-vigenere-cipher/util"
	"fmt"
	"strings"
)

/*
Extended Vigenere Cipher is a modified vignere cipher by combining :
1. Enhanced cipher by Gautam (2018) by adding ukey as shift value.
2. Random vigenere table by Yumnam (2012) by using different vigenere table for encrypt and decrypt.
3. Autokey by not repeating key usage, but use ciphertext or plaintext instead as a key.
4. Extending vigenere table character set.
*/

// TableType is a table used in encryption.
type TableType string
type AutokeyType string

const (
	RowUniqueType TableType = "row"
	ColUniqueType TableType = "col"

	PlaintextAutokey  AutokeyType = "plaintext"
	CiphertextAutokey AutokeyType = "ciphertext"
)

// Interface is an interface to be implemented by ExtendedVigenereCipher.
type Interface interface {
	Encrypt(plaintext, key string, ukey int) (string, error)
	Decrypt(ciphertext, key string, ukey int) (string, error)
}

// ExtendedVigenereCipher is a main struct of this encryption.
type ExtendedVigenereCipher struct {
	Table       *table.VigenereTable
	TableType   TableType
	AutokeyType AutokeyType
}

// NewExtendedVigenereCipher is a constructor.
func NewExtendedVigenereCipher(table *table.VigenereTable, tableType TableType, autokeyType AutokeyType) *ExtendedVigenereCipher {
	return &ExtendedVigenereCipher{
		Table:       table,
		TableType:   tableType,
		AutokeyType: autokeyType,
	}
}

var _ Interface = &ExtendedVigenereCipher{}

// Encrypt is to encrypt plaintext.
func (evc ExtendedVigenereCipher) Encrypt(plaintext, key string, ukey int) (string, error) {
	if evc.Table == nil {
		return "", fmt.Errorf("extended_vigenere_cipher.error.missing_table")
	}

	if string(evc.TableType) == "" || (evc.TableType != RowUniqueType && evc.TableType != ColUniqueType) {
		return "", fmt.Errorf("extended_vigenere_cipher.error.invalid_table_type")
	}

	var ciphertext strings.Builder

	tempKey := key
	indexKey := 0
	usableCharacter := len(evc.Table.TextCharset)

	for _, subText := range plaintext {

		var c string
		if evc.TableType == ColUniqueType {
			m, err := evc.Table.GetSubTextIndex(string(subText))
			if err != nil {
				return "", err
			}
			n, err := evc.Table.GetSubKeyIndex(string(tempKey[indexKey]))
			if err != nil {
				return "", err
			}

			// C = Et((M+Ukey) % Usable_Character,N)
			c = evc.Table.Cell[util.Mod(m+ukey, usableCharacter)][n]
		} else {
			n, err := evc.Table.GetSubTextIndex(string(subText))
			if err != nil {
				return "", err
			}
			m, err := evc.Table.GetSubKeyIndex(string(tempKey[indexKey]))
			if err != nil {
				return "", err
			}

			// C = Et(M,(N+UKey) % Usable_Character)
			c = evc.Table.Cell[m][util.Mod(n+ukey, usableCharacter)]
		}
		ciphertext.WriteString(c)

		indexKey++
		if evc.AutokeyType == CiphertextAutokey { //Autokey by ciphertext
			tempKey = tempKey[0:] + c
		} else if evc.AutokeyType == PlaintextAutokey {
			tempKey = tempKey[0:] + string(subText) //Autokey by plaintext
		} else {
			if indexKey == len(tempKey) { //RESET KEY
				indexKey = 0
			}
		}
	}

	return ciphertext.String(), nil
}

// Decrypt is to decrypt ciphertext.
func (evc ExtendedVigenereCipher) Decrypt(ciphertext, key string, ukey int) (string, error) {
	if evc.Table == nil {
		return "", fmt.Errorf("extended_vigenere_cipher.error.missing_table")
	}

	if string(evc.TableType) == "" || (evc.TableType != RowUniqueType && evc.TableType != ColUniqueType) {
		return "", fmt.Errorf("extended_vigenere_cipher.error.invalid_table_type")
	}

	var plaintext strings.Builder

	tempKey := key
	indexKey := 0
	usableCharacter := len(evc.Table.TextCharset)

	for _, subText := range ciphertext {
		n, err := evc.Table.GetSubTextIndex(string(subText))
		if err != nil {
			return "", err
		}

		m, err := evc.Table.GetSubKeyIndex(string(tempKey[indexKey]))
		if err != nil {
			return "", err
		}

		if evc.TableType == ColUniqueType {
			temp := n
			n = m
			m = temp
		}

		//P = ((Dt( N , M ) - UKey) % Usable_Character)
		a := evc.Table.Cell[m][n]
		b, err := evc.Table.GetSubTextIndex(a)
		if err != nil {
			return "", err
		}

		e := util.Mod(b-ukey, usableCharacter)

		sp := evc.Table.TextCharset[e]
		if err != nil {
			return "", err
		}
		plaintext.WriteString(sp)

		indexKey++
		if evc.AutokeyType == CiphertextAutokey { //Autokey by ciphertext
			tempKey = tempKey[0:] + string(subText) //VARYING KEY
		} else if evc.AutokeyType == PlaintextAutokey {
			tempKey = tempKey[0:] + sp //VARYING KEY
		} else {
			if indexKey == len(tempKey) { //RESET KEY
				indexKey = 0
			}
		}
	}

	return plaintext.String(), nil
}
