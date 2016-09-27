package main

import (
    "fmt"
    "log"
    "strings"
    "net/http"
)

func RoutingActions(w http.ResponseWriter, r *http.Request){
    r.ParseForm()  // parse arguments, you have to call this by yourself

    // Print information to browser and stdout
    fmt.Println("URL Path: ", r.URL.Path)
    fmt.Fprintln(w, "Url Path: ", r.URL.Path)

    fmt.Println("URL Scheme: ", r.URL.Scheme)
    fmt.Fprintln(w, "URL Scheme: ", r.URL.Scheme)

    fmt.Println("Absolute URL: ", r.URL.IsAbs())
    fmt.Fprintln(w, "Absolute URL: ", r.URL.IsAbs())

    // - Form Info
    fmt.Println("Form Information: ", r.Form)
    fmt.Fprintln(w, "Form Information: ", r.Form)
    // - key values
    for k, v := range r.Form {
        fmt.Fprintf(w, "Key: %v", k)
        fmt.Fprintf(w, ", Value: %v \n", strings.Join(v, ""))

        fmt.Printf("Key: %v", k)
        fmt.Printf(", Value: %v \n", strings.Join(v, ""))
    }
}

func main() {
    fmt.Println("Initializing Server...")

    // Retrieving url path and printing it after hello
    http.HandleFunc("/", RoutingActions)

    // Serving at port 8080
    go func() {
      fmt.Println("Listening on 8080")
      log.Fatal(http.ListenAndServe(":8080", nil))
    }()

    // Serving at port 8443
    fmt.Println("Listening on 8443")
    log.Fatal(http.ListenAndServeTLS(":8443", "server.pem", "server.key", nil))
}
