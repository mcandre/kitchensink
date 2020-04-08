// echoserver mirrors incoming HTTP request data as outgoing HTTP response data.

package main

import (
	"io"
	"net/http"
)

// main is CLI entrypoint.
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.Copy(w, r.Body)
	})
	http.ListenAndServe(":8080", nil)
}
