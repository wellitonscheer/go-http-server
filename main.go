package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func increment(num *int) {
	*num += 1
}

func setMessage(msg *string, newMsg string) {
	*msg = newMsg
}

func main() {
	a := 0
	message := "message"

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/increment", func(w http.ResponseWriter, r *http.Request) { increment(&a); fmt.Fprintf(w, "number a: %v\n", a) })
	http.HandleFunc("/headers", headers)

	http.HandleFunc("GET /message", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "%v\n", message) })
	http.HandleFunc("POST /message/{msg}", func(w http.ResponseWriter, r *http.Request) { setMessage(&message, r.PathValue("msg")) })

	http.ListenAndServe(":6969", nil)
}
