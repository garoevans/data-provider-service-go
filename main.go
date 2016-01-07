package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s({\"data\":\"Test Resopnse\"})", r.URL.Query().Get("callback"))
}

// Form comment
type Form struct {
	Name, Email string
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var f Form
	err := dec.Decode(&f)
	if err == nil {
		fmt.Printf("%s: %s\n", f.Name, f.Email)
	} else {
		fmt.Println("Something went wrong")
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	mux.HandleFunc("/submit/", submitHandler)

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	handler := cors.Default().Handler(mux)
	http.ListenAndServe("localhost:8081", handler)
}
