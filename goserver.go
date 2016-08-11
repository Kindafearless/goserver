package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
)

func main() {
    fmt.Println("Initializing Hello World Server...")

    // Retrieving url path and printing it after hello
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })

    // Serving at port 8080
    log.Fatal(http.ListenAndServe(":8080", nil))

}
