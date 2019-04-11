package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/context"
)

func Device(w http.ResponseWriter, r *http.Request) {
    device := new(ModelDevice)
    _, err := Db.QueryOne(device, `SELECT * FROM device WHERE identifier = ?`, context.Get(r, "deviceIdentifier").(string))

    result, err := json.Marshal(device)
    if err != nil {
        log.Panic(err)
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.Write(result)
}

func DeviceDataPoint(w http.ResponseWriter, r *http.Request) {

    fmt.Fprintf(w, "DeviceDataPoint")
}

func DeviceDataPointPh(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "DeviceDataPointPh")
}

func DeviceDataPointOxygen(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "DeviceDataPointOxygen")
}

func DeviceDataPointTemperature(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "DeviceDataPointTemperature")
}
