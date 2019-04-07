package main

import (
    "log"
    "net/http"
)

func main() {
    // create router for api endpoints
    router := BaseRouter()

    // start webserver loop
    log.Fatal(http.ListenAndServe(":8000", router))
}
