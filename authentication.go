package main

import (
    "net/http"
    "time"
    "io"
    "golang.org/x/crypto/bcrypt"
    "github.com/SermoDigital/jose/jws"
    "github.com/SermoDigital/jose/crypto"
    "./models"
)


func AuthenticationGetTokenFarm(w http.ResponseWriter, r *http.Request) {
    // check if we got application/x-www-form-urlencoded content type
    if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
        w.WriteHeader(http.StatusUnsupportedMediaType)
        return
    }

    // get farm data from database
    var farmResult models.Farm
    if db.Where("identifier = ?", r.FormValue("identifier")).First(&farmResult).Error != nil {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    // compare hashed secret with plaintext secret and deny access if not equal
    if bcrypt.CompareHashAndPassword([]byte(farmResult.Secret), []byte(r.FormValue("secret"))) != nil {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    // generate json web token
    claims := jws.Claims{}
    claims.SetSubject(r.FormValue("identifier"))
    claims.SetIssuedAt(time.Now())
    claims.SetExpiration(time.Now().Add(time.Duration(jwtSettings.ValidityFarm) * time.Second))
    token, _ := jws.NewJWT(claims, crypto.SigningMethodHS256).Serialize([]byte(jwtSettings.Secret))

    // return the web token
    w.Header().Set("Content-Type", "application/jwt")
    io.WriteString(w, string(token[:]))
}

func AuthenticationGetTokenUser(w http.ResponseWriter, r *http.Request) {
    // check if we got application/x-www-form-urlencoded content type
    if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
        w.WriteHeader(http.StatusUnsupportedMediaType)
        return
    }

    // get user data from database
    var userResult models.User
    if db.Where("email = ?", r.FormValue("email")).First(&userResult).Error != nil {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    // compare hashed password with plaintext password and deny access if not equal
    if bcrypt.CompareHashAndPassword([]byte(userResult.Password), []byte(r.FormValue("password"))) != nil {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    // generate json web token
    claims := jws.Claims{}
    claims.SetSubject(r.FormValue("email"))
    claims.SetIssuedAt(time.Now())
    claims.SetExpiration(time.Now().Add(time.Duration(jwtSettings.ValidityUser) * time.Second))
    token, _ := jws.NewJWT(claims, crypto.SigningMethodHS256).Serialize([]byte(jwtSettings.Secret))

    // return the web token
    w.Header().Set("Content-Type", "application/jwt")
    io.WriteString(w, string(token[:]))
}
