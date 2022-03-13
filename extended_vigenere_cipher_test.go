package extended_vigenere_cipher_test

import (
	evc "extended-vigenere-cipher"
	"extended-vigenere-cipher/table"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExtendedVigenereCipher_Encrypt(t *testing.T) {

	plaintext := "Hello World!"
	key := "s3cr3T"
	ukey := 3

	et, dt, err := table.GenerateYumnamTable(3, table.DefaultCharset(), false)
	require.NoError(t, err)

	e := evc.NewExtendedVigenereCipher(et, evc.RowUniqueType, evc.CiphertextAutokey)
	ct1, err := e.Encrypt(plaintext, key, ukey)
	require.NoError(t, err)

	e = evc.NewExtendedVigenereCipher(dt, evc.RowUniqueType, evc.CiphertextAutokey)
	pt1, err := e.Decrypt(ct1, key, ukey)
	require.NoError(t, err)
	assert.Equal(t, plaintext, pt1)

	et, dt, err = table.GenerateYumnamTable(3, table.DefaultCharset(), true)
	require.NoError(t, err)

	e = evc.NewExtendedVigenereCipher(et, evc.ColUniqueType, evc.CiphertextAutokey)
	ct2, err := e.Encrypt(plaintext, key, ukey)
	require.NoError(t, err)

	e = evc.NewExtendedVigenereCipher(dt, evc.ColUniqueType, evc.CiphertextAutokey)
	pt2, err := e.Decrypt(ct2, key, ukey)
	require.NoError(t, err)
	assert.Equal(t, plaintext, pt2)

	assert.Equal(t, ct1, ct2)
	assert.Equal(t, pt1, pt2)
}
