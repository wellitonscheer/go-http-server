package main

import (
	"encoding/json"
	"fmt"
	"github.com/wellitonscheer/go-http-server/whats"
	"io"
	"net/http"
)

type sendBody struct {
	To string `json:"to"`
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	a := 0
	message := "message"

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/increment", func(w http.ResponseWriter, r *http.Request) { increment(&a); fmt.Fprintf(w, "number a: %v\n", a) })

	http.HandleFunc("GET /message", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "%v\n", message) })
	http.HandleFunc("POST /message/{msg}", func(w http.ResponseWriter, r *http.Request) { setMessage(&message, r.PathValue("msg")) })

	http.HandleFunc("POST /send/{msg}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Message to send: ", r.PathValue("msg"))

		b, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error to read body")
			return
		}

		var bodyStruct sendBody
		errB := json.Unmarshal(b, &bodyStruct)
		if errB != nil {
			fmt.Println("Error to Unmarshal")
			return
		}

		errS := whats.SendMessage(bodyStruct.To)
		if errS != nil {
			fmt.Printf("Failed to send message: %v\n", errS)
		} else {
			fmt.Println("Message sent successfully")
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err.Error())
	}
}

func increment(num *int) {
	*num += 1
}

func setMessage(msg *string, newMsg string) {
	*msg = newMsg
}
