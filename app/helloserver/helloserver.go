// hello world web server
// localhost:8000/hello?name=<name> will return Hello, <name>!
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", hello)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "try hello")
}

func hello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if len(name) != 0 {
		fmt.Fprintf(w, "Hello, %s!\n", name)
	} else {
		fmt.Fprintf(w, "Hello!\n")
	}
}