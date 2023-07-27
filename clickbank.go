package clickbank

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
)

var ErrCouldNotDecode = errors.New("could not Base64 decode string")

func DecryptRequestBody(body io.Reader, key []byte) ([]byte, error) {
	var json_response EncryptedNotification
	var encrypted_response []byte
	json_decoder := json.NewDecoder(body)
	err := json_decoder.Decode(&json_response)
	if err != nil {
		return nil, err
	}

	// Extract the IV needed for cipher
	iv, err := base64.StdEncoding.DecodeString(json_response.IV)
	if err != nil {
		return nil, ErrCouldNotDecode
	}

	encrypted_response, err = base64.StdEncoding.DecodeString(json_response.Notification)
	if err != nil {
		return nil, ErrCouldNotDecode
	}

	crypt, err := NewCryptStruct(key, iv)
	if err != nil {
		return nil, err
	}

	decrypted_response, err := crypt.DecryptPKCS5(encrypted_response)
	if err != nil {
		return nil, err
	}

	return decrypted_response, nil
}

func DecodeResponse(encoded_response []byte) (*ClickbankNotification, error) {
	var final_object *ClickbankNotification
	err := json.Unmarshal(encoded_response, final_object)
	if err != nil {
		return nil, err
	}

	return final_object, nil
}

func ObfuscateKey(key []byte) []byte {
	sha := sha1.Sum(key)
	hexSha := make([]byte, hex.EncodedLen(len(sha)))
	hex.Encode(hexSha, sha[:])
	return hexSha[:32]
}
