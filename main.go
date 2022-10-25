package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		println("server ok...")
		switch r.URL.Path {
		case "/":
			switch r.Method {
			case "GET":
				w.Write([]byte("Hello World!"))
			case "POST":
				HandlePost(w, r)
			}
		default:
			http.NotFound(w, r)
		}
	})

	http.ListenAndServe(":3000", nil)

}

func HandlePost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("We have a problem reading the content! Please, try again."))
		w.WriteHeader(http.StatusBadGateway)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)

}
