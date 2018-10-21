package main

import (
	"fmt"
	"net/http"
)

const port = 3000

func main() {
	fmt.Printf("\nRunning card generation tool...\nNavigate to http://localhost:%v\n", port)
	http.Handle("/", http.FileServer(http.Dir(".")))
	//http.Handle("/images/", http.StripPrefix("./images/", http.FileServer(http.Dir("./images"))))
	//http.Handle("/fonts/", http.StripPrefix("./fonts/", http.FileServer(http.Dir("./fonts"))))
	//http.Handle("/css/", http.StripPrefix("./css/", http.FileServer(http.Dir("./css"))))
	//http.Handle("/js/", http.StripPrefix("./js/", http.FileServer(http.Dir("./js"))))
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
