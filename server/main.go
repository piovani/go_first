package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", HelloHandle)
	http.ListenAndServe(":8887", nil)

}

func HelloHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello All")
}