package main

import (
    "fmt"
    "io"
    "net/http"
    "github.com/gorilla/context"
    "time"
    "github.com/SermoDigital/jose/jws"
    "github.com/SermoDigital/jose/crypto"
)

func Device(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    io.WriteString(w, context.Get(r, "deviceIdentifier").(string))
}

func DeviceAuthenticationGetToken(w http.ResponseWriter, r *http.Request) {
    // check if we got application/x-www-form-urlencoded content type
    if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
        w.WriteHeader(http.StatusUnsupportedMediaType)
        return
    }

    // extract device identifier and secret
    deviceIdentifier := r.FormValue("deviceIdentifier")
    deviceSecret := r.FormValue("deviceSecret")

    // check if identifier and secret are valid
    device := new(ModelDevice)
    exists, err := Db.Model(device).Where("Identifier = ?", deviceIdentifier).Where("Secret = ?", deviceSecret).Exists()

    if err != nil {
        panic(err)
    }

    if !exists {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    // generate json web token
    claims := jws.Claims{}
    claims.SetSubject(deviceIdentifier)
    claims.SetIssuedAt(time.Now())
    claims.SetExpiration(time.Now().Add(time.Duration(3600) * time.Second))
    token, _ := jws.NewJWT(claims, crypto.SigningMethodHS256).Serialize([]byte("your-256-bit-secret"))

    // return the web token
    w.Header().Set("Content-Type", "application/jwt")
    io.WriteString(w, string(token[:]))
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
