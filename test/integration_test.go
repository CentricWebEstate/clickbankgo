package test

import (
	"io/ioutil"
	"os"
	"testing"

	clickbank "github.com/centricwebestate/clickbankgo/pkg/v1"
)

var nullBytes = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}

func TestPayload(t *testing.T) {
	payloadBytes, err := os.Open("./data/encrypted_payload.json")
	if err != nil {
		t.Fatalf("failed loading payload: %v", err)
	}

	keyBytes, err := ioutil.ReadFile("./data/example_key.txt")
	if err != nil {
		t.Fatalf("failed loading key: %v", err)
	}

	key := clickbank.ObfuscateKey(keyBytes)
	decrypted_response, err := clickbank.DecryptRequestBody(payloadBytes, key)
	if err != nil {
		t.Fatalf("failed decrypting response: %v", err)
	}

	t.Logf("Response string: %s", decrypted_response)

	final_object, err := clickbank.DecodeResponse(decrypted_response)
	if err != nil {
		t.Fatalf("failed serialising response: %v", err)
	}

	t.Logf("Finished: %+v", final_object)
}
