package main

import (
    "github.com/go-pg/pg"
    "github.com/go-pg/pg/orm"
)

func GetDbConnectionOptions(hostname string, port string, username string, password string, database string) *pg.Options {
    var dbOptions pg.Options
    dbOptions.Network = "tcp"

    // check hostname and port
    dbOptions.Addr = hostname + ":" + port

    // check username
    dbOptions.User = username

    // check password
    dbOptions.Password = password

    // check database
    dbOptions.Database = database

    return &dbOptions
}

func OpenDb(connectionOptions *pg.Options) *pg.DB {
    Db = pg.Connect(connectionOptions)
    return Db
}

func CloseDb(Db *pg.DB) {
    Db.Close()
}


func InitDb(Db *pg.DB) error {
    models := []interface{} {
        (*ModelUser)(nil),
        (*ModelDevice)(nil),
        (*ModelDeviceSetting)(nil),
        (*ModelDeviceDataPointPh)(nil),
        (*ModelDeviceDataPointOxygen)(nil),
        (*ModelDeviceDataPointTemperature)(nil),
        (*ModelUserDeviceAccess)(nil),
    }

    for _, model := range models {
        err := Db.CreateTable(model, &orm.CreateTableOptions{
            Temp: false,
            IfNotExists: true,
            FKConstraints: true,
        })
        if err != nil {
            panic(err)
        }
    }
    return nil
}
