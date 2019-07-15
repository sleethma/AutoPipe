package main

import "net/http"

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello World!"))
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
