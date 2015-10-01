package clickbank

import "encoding/base64"
import "encoding/json"
import "io"
import "crypto/sha1"

func DecryptRequestBody(body io.ReadCloser, key []byte) ([]byte, error) {
	var json_response EncryptedNotification
	var encrypted_response []byte
	json_decoder := json.NewDecoder(body)
	err := json_decoder.Decode(&json_response)
	if err != nil {
		return nil, err
	}

	// Extract the IV needed for cipher
	iv, err := base64.StdEncoding.DecodeString(json_response.IV)
	encrypted_response, err = base64.StdEncoding.DecodeString(json_response.Notification)
	crypt, _ := NewCryptStruct(key, iv)
	decrypted_response, err := crypt.Decrypt(encrypted_response)
	if err != nil {
		return nil, err
	}

	return decrypted_response, nil
}

func DecodeResponse(encoded_response []byte) (ClickbankNotification, error) {
	var final_object ClickbankNotification
	err := json.Unmarshal(encoded_response, &final_object)
	if err != nil {
		return ClickbankNotification{}, err
	}

	return final_object, nil
}

func ObfuscateKey(key []byte) []byte {
	hash := sha1.New()
	return hash.Sum(key)[:32]
}
