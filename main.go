package main

import (
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello world."))
}

func main() {
    http.HandleFunc("/", hello)
    http.ListenAndServe(":8888", nil)
    fmt.Println("listen 127.0.0.1:8888")
}
