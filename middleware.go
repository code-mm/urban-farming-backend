package main

import (
	"io/ioutil"
	"bytes"
	"net/http"
	"encoding/json"
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

func ContentTypeValidationJson(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// check content type in request header
	if r.Header.Get("Content-Type") != "application/json" {
		rw.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	// get body from request
	requestBody, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	// check if json in body is valid
	if !json.Valid([]byte(requestBody)) {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	// reset the request body
	r.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))

	next(rw, r)
}