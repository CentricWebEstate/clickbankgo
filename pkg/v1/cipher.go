package clickbank

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

type CryptStruct struct {
	key    []byte
	iv     []byte
	cipher cipher.Block
}

var ErrInvalidInput = errors.New("input is not a factor of the cipher's blocksize")
var ErrInvalidPadding = errors.New("decrypted output does not appear to be correctly PKCS5 padded")

/* Create a new CrypteStruct */
func NewCryptStruct(key []byte, iv []byte) (*CryptStruct, error) {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	return &CryptStruct{key, iv[:cipher.BlockSize()], cipher}, nil
}

/* Decrypt Something */
func (c *CryptStruct) Decrypt(input []byte) ([]byte, error) {
	ctx := cipher.NewCBCDecrypter(c.cipher, c.iv)

	// Ensure that the input meets the correct blocksize
	if len(input)%c.cipher.BlockSize() != 0 {
		return nil, ErrInvalidInput
	}

	text := input[c.cipher.BlockSize():]
	ctx.CryptBlocks(text, text)
	return input, nil
}

func (c *CryptStruct) DecryptPKCS5(input []byte) ([]byte, error) { 
	r, err := c.Decrypt(input)
	if err != nil {
		return nil, err
	}

	return PKCS5Unpad(r)
}

func PKCS5Unpad (input []byte) ([]byte, error) {
	// Expect the last byte of the string to have the length of padding
	length := len(input)
	padding_size := int(input[length-1])

	if len(input) < padding_size {
		return nil, ErrInvalidPadding
	}

	return input[:length-padding_size], nil
}