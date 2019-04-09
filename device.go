package main

import (
    "fmt"
    "io"
    "net/http"
    "github.com/gorilla/context"
)

func Device(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Content-Type", "application/json")
    io.WriteString(w, context.Get(r, "deviceIdentifier").(string))
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
