package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func slash(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "ha\n")
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

func main() {
	a := 1

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/increment", func(w http.ResponseWriter, r *http.Request) { increment(&a); fmt.Fprintf(w, "number a: %v\n", a) })
	http.HandleFunc("/", slash)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":6969", nil)
}
