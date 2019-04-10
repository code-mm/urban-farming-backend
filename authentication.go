package main

import (
    "net/http"
    "time"
    "io"
    "golang.org/x/crypto/bcrypt"
    "github.com/SermoDigital/jose/jws"
    "github.com/SermoDigital/jose/crypto"
)

func AuthenticationGetTokenDevice(w http.ResponseWriter, r *http.Request) {
    // check if we got application/x-www-form-urlencoded content type
    if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
        w.WriteHeader(http.StatusUnsupportedMediaType)
        return
    }

    // extract device identifier and secret
    deviceIdentifier := r.FormValue("deviceIdentifier")
    deviceSecret := r.FormValue("deviceSecret")

    // check if identifier and secret are valid
//    device, err := Db.Model(new(ModelDevice)).Where("Identifier = ?", deviceIdentifier).Exists()
    device := new(ModelDevice)
    _, err := Db.QueryOne(device, `SELECT * FROM device WHERE identifier = ?`, deviceIdentifier)

    if err != nil {
        panic(err)
    }

    // compare hashed secret with plaintext secret and deny access if not equal
    if bcrypt.CompareHashAndPassword([]byte(device.Secret), []byte(deviceSecret)) != nil {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    // generate json web token
    claims := jws.Claims{}
    claims.SetSubject(deviceIdentifier)
    claims.SetIssuedAt(time.Now())
    claims.SetExpiration(time.Now().Add(time.Duration(3600) * time.Second))
    token, _ := jws.NewJWT(claims, crypto.SigningMethodHS256).Serialize([]byte(JwtSecret))

    // return the web token
    w.Header().Set("Content-Type", "application/jwt")
    io.WriteString(w, string(token[:]))
}