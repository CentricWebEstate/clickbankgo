package main

import (
	"fmt"
	"net/http"

	clickbank "github.com/centricwebestate/clickbankgo/pkg/v1"
)

var cb_key []byte = []byte("SOMEKEY")
var key []byte = clickbank.ObfuscateKey(cb_key)

func handler(w http.ResponseWriter, r *http.Request) {
	decrypted_response, _ := clickbank.DecryptRequestBody(r.Body, key)
	final_object, _ := clickbank.DecodeResponse(decrypted_response)
	fmt.Fprintf(w, "%#v", final_object)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
