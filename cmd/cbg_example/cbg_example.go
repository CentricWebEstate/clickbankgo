package main

import (
	"flag"
	"fmt"
	"net/http"

	clickbank "github.com/centricwebestate/clickbankgo/pkg/v1"
)

var default_cb_key []byte = []byte("ABCD1234EFGH5678")

func handlerFromKey(key []byte) func(http.ResponseWriter, *http.Request) {
	key = clickbank.ObfuscateKey(key)
	return func(w http.ResponseWriter, r *http.Request) {
		decrypted_response, err := clickbank.DecryptRequestBody(r.Body, key)
		if err != nil {
			fmt.Fprintf(w, "%#v", err)
		}

		final_object, err := clickbank.DecodeResponse(decrypted_response)
		if err != nil {
			fmt.Fprintf(w, "%#v", err)
		}

		fmt.Fprintf(w, "%#v", final_object)
	}
}

func main() {
	key := flag.String("key", string(default_cb_key), "The clickbank key to decode the data")
	flag.Parse()
	http.HandleFunc("/", handlerFromKey([]byte(*key)))
	http.ListenAndServe(":8080", nil)
}
