package main

import (
	"fmt"
	"net/http"
)

const port = 3000

func main() {
	fmt.Printf("\nRunning card generation tool...\nNavigate to http://localhost:%v\n", port)
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
