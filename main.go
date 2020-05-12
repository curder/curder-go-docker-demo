package main

import (
    "log"
    "net/http"
)

func main() {
    log.Println("start server...")

    http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
        writer.WriteHeader(http.StatusOK)
        _, _ = writer.Write([]byte("hello docker!"))
    })

    log.Fatal(http.ListenAndServe(":80", nil))
}
