package main

import (
    "log"
    "net/http"
)

func forwardURL(ResponseWriter http.ResponseWriter, Request *http.Request) {
    http.Redirect(ResponseWriter, Request, "https://evhs.schoolloop.com", 301)
}

func main() {
    http.HandleFunc("/go/", forwardURL)
    err := http.ListenAndServe(":8798", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}