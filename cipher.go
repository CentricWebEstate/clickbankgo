package clickbank

import (
	"crypto/aes"
	"crypto/cipher"
)

type CryptStruct struct {
	key    []byte
	iv     []byte
	cipher cipher.Block
}

/* Create a new CrypteStruct */
func NewCryptStruct(key []byte, iv []byte) (*CryptStruct, error) {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	return &CryptStruct{key, iv, cipher}, nil
}

/* Decrypt Something */
func (c *CryptStruct) Decrypt(input []byte) ([]byte, error) {
	ctx := cipher.NewCBCDecrypter(c.cipher, c.iv)
	text := input[c.cipher.BlockSize():]
	ctx.CryptBlocks(text, text)
	return input, nil
}
