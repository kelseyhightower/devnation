package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var html = `
<html>
<h1>Hello Dev Nation!</h1>
<p>%s</p>
</html>
`

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, html, hostname)
	})
	http.ListenAndServe(":8080", nil)
}
