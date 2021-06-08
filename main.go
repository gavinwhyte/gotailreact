package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

func main() {

    r := http.NewServeMux()

    r.HandleFunc("/route1", index)
    r.HandleFunc("/route2", index)
    buildHandler := http.FileServer(http.Dir("build"))
    r.Handle("/", buildHandler)

    srv := &http.Server{
        Handler:      r,
        Addr:         "127.0.0.1:8080",
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    fmt.Println("Server started on PORT 8080")
    log.Fatal(srv.ListenAndServe())

}

func index(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "build/index.html")
}
