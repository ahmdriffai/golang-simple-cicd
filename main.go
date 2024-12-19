package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello dunia tipu-tipu sekali")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listening on : 7000")
	http.ListenAndServe(":7000", nil)
}
