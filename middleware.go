package main

import (
	"net/http"
	"github.com/gorilla/context"
	"github.com/SermoDigital/jose/jws"
    "github.com/SermoDigital/jose/crypto"
)

func JwtTokenValidationDevice(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// Parse jwt from request header authorization field
	token, err := jws.ParseJWTFromRequest(r)

	if err != nil {
		rw.WriteHeader(http.StatusUnauthorized)
        return
	}

	if err = token.Validate([]byte(JwtSecret), crypto.SigningMethodHS256); err != nil {
		rw.WriteHeader(http.StatusUnauthorized)
        return
	}

	// get the deviceIdentifier from json jwt and store it in the request context
	deviceIdentifier := token.Claims()["sub"]
	context.Set(r, "deviceIdentifier", deviceIdentifier)

    next(rw, r)
}