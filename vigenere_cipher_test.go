package extended_vigenere_cipher_test

import (
	evc "extended-vigenere-cipher"
	"extended-vigenere-cipher/table"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestVigenereCipher_Encrypt(t *testing.T) {
	plaintext := "AAAA"
	key := "BBBB"
	var vt table.VigenereTable

	charset := table.UpperCaseCharset()
	vt.Generate(charset)

	vc := evc.NewVigenereCipher(&vt)
	ciphertext, err := vc.Encrypt(plaintext, key)
	require.NoError(t, err)
	resPlaintext, err := vc.Decrypt(ciphertext, key)
	require.NoError(t, err)
	assert.Equal(t, plaintext, resPlaintext)
}
