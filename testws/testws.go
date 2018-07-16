package main

import "fmt"
import "net/http"
import "google.golang.org/appengine"

//http handler below
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Best Blockie Ever!") // Response to request
  }

// a fuction which takes input and creates a block

func main() {
	//starts web server 
	http.HandleFunc("/testws", indexHandler) // set endpoint handler
		appengine.Main() // starts the server to receive requests
}