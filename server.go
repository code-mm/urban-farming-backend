package main

import (
    "os"
    "log"
    "net/http"
    "github.com/go-pg/pg"
)

// database handle
var Db *pg.DB

func main() {
    // read database related environment variables
    connectionOptions := GetDbConnectionOptions(os.Getenv("DbHostname"), os.Getenv("DbPort"), os.Getenv("DbUsername"), os.Getenv("DbPassword"), os.Getenv("DbDatabase"))

    // open database
    Db = OpenDb(connectionOptions)
    defer CloseDb(Db)

    // create tables if they do not exist
    InitDb(Db)

    // create router for api endpoints
    router := BaseRouter()
    DeviceRouter(router)

    // start webserver loop
    log.Fatal(http.ListenAndServe(":8000", router))
}
