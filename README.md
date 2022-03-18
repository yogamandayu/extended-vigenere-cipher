# Extended Vigenere Cipher

Extended vigenere cipher is a modification of vigenere cipher with method :

1. Random vigenere table by Yumnam (2012) in "Generalization of Vigenere Cipher" is using separate random vigenere table
   for encrypt and decrypt.
3. Autokey by Chris C. (2015) in "Cryptology Notes" is to extend the key if the message is larger than the key by using
   the plaintext or the ciphertext.
4. Gautam et. al. (2018) in "Enhanced Cipher Technique Using Vigenere and Modified Caesar Cipher" add an additional key
   to shift when choosing ciphertext in the vigenere table.
5. Extending vigenere table character set usage.

## Vigenere Table

This is how I imagine the concept of vigenere table in this project. I separated between key and text charset with the
cell. It will make the vigenere table have a flexibility of charset. For example, you can randomize the key/text charset
or the table too.

```
Key	Text		Cell
0|A|	0|A|		ABCDEF
1|B|	1|B|		BCDEFA
2|C|	2|C|		CDEFAB
3|D|	3|D|		DEFABC
4|E|	4|E|		EFABCD
5|F|	5|F|		FABCDE

A,A = Cell[0][0] = A
D,C = Cell[3][2] = F
```

## Usage

Prepare the requirement :

1. Plaintext
2. Key
3. UKey (integer)
4. Random vigenere table from Yumnam's journal. You can use this code for generating Yumnam's table.

```
et, dt, err = table.GenerateYumnamTable(3, table.DefaultCharset(), true)

Arguments :
1. Seed value for generate number
2. List of charset.
3. Boolean if table use row or column unique (read journal for detail) 
```

If you already have the separate encrypt and decrypt table, you need to register it first.

```
et := table.NewVigenereTable(textCharset, keyCharset, encryptCell)
dt := table.NewVigenereTable(textCharset, keyCharset, decryptCell)
```

After the requirement is fulfilled, you just pass the requirements for encrypt or decrypt.

```
e = NewExtendedVigenereCipher(et, ColUniqueType, CiphertextAutokey)
ciphertext, err := e.Encrypt(plaintext, key, ukey)

e = NewExtendedVigenereCipher(dt, ColUniqueType, CiphertextAutokey)
plaintext, err := e.Decrypt(ciphertext, key, ukey)
```
## Journal
[Google Drive](https://drive.google.com/file/d/14jG781UfoLB3hu-f9alzHmkhv5v0cLed/view?usp=sharing)

## License

[MIT](https://choosealicense.com/licenses/mit/)
