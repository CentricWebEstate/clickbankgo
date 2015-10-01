package clickbank

import "github.com/spacemonkeygo/openssl"

type CryptStruct struct {
	key    []byte
	iv     []byte
	cipher *openssl.Cipher
}

/* Create a new CrypteStruct */
func NewCryptStruct(key []byte, iv []byte) (*CryptStruct, error) {
	cipher, err := openssl.GetCipherByName("aes-256-cbc")
	if err != nil {
		return nil, err
	}

	return &CryptStruct{key, iv, cipher}, nil
}

/* Decrypt Something */
func (c *CryptStruct) Decrypt(input []byte) ([]byte, error) {
	ctx, err := openssl.NewDecryptionCipherCtx(c.cipher, nil, c.key, c.iv)
	if err != nil {
		return nil, err
	}

	cipherbytes, err := ctx.DecryptUpdate(input)
	if err != nil {
		return nil, err
	}

	finalbytes, err := ctx.DecryptFinal()
	if err != nil {
		return nil, err
	}

	cipherbytes = append(cipherbytes, finalbytes...)
	return cipherbytes, nil
}
