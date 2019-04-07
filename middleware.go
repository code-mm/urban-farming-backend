package main

import (
	"net/http"
	"regexp"
	"strings"
	"github.com/gorilla/context"
	"github.com/SermoDigital/jose/jws"
    "github.com/SermoDigital/jose/crypto"
)

func DeviceJwtTokenValidation(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// check if authorization header in place and get it if so
	authorizationHeader := r.Header.Get("Authorization")

	// check format of authorization header
	authorizationHeaderRegex, _ := regexp.Compile("^Bearer [A-Za-z0-9.]+")

	if !authorizationHeaderRegex.MatchString(authorizationHeader) {
		rw.WriteHeader(http.StatusUnauthorized)
        return
	}

	// extract access token from authorization header
	accessToken := strings.TrimPrefix(authorizationHeader, "Bearer ")

	// check if jwt is valid
	jwt, err := jws.ParseJWT([]byte (accessToken))

	if err != nil {
		rw.WriteHeader(http.StatusUnauthorized)
        return
	}

	if err = jwt.Validate([]byte("your-256-bit-secret"), crypto.SigningMethodHS256); err != nil {
		rw.WriteHeader(http.StatusUnauthorized)
        return
	}

	// get the deviceIdentifier from json jwt and store it in the request context
	deviceIdentifier := jwt.Claims()["sub"]
	context.Set(r, "deviceIdentifier", deviceIdentifier)

    next(rw, r)
}