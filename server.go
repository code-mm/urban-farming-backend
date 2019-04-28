package main

import (
    "log"
    "net/http"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)


// database settings
var dbSettings *DatabaseSettings

// jwt settings
var jwtSettings *JwtSettings

// database handle
var db *gorm.DB

func main() {
    // set database settings
    dbSettings = DatabaseSettingsNew()

    //set jwt settings
    jwtSettings = JwtSettingsNew()

    // initialize database
    DbOpen(&db, dbSettings)
    db.LogMode(true)
    defer DbClose(db)

    DbInit(db)

    // create router for api endpoints
    router := BaseRouter()
    AuthenticationRouter(router)
    FarmRouter(router)
    UserRouter(router)

    // start webserver loop
    log.Fatal(http.ListenAndServe(":8000", router))
}
