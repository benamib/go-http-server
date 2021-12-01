package main

import (
	"fmt"
	"net/http"
)

const PORT = 8080

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello Workd\n")
	})
	fmt.Printf("Serving on port :%d\n", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}
