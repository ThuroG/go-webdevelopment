
package main

import (
	"fmt"
	"net/http"
	"os"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello, World!</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Fprintln(os.Stdout, "Hello World")
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
