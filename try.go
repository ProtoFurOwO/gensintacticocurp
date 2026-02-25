package main

import (
	"fmt"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	fmt.Println("Servidor iniciado en localhost:8000...")
	http.ListenAndServe(":8000", nil)
}
