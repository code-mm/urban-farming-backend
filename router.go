package main

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func BaseRouter() *mux.Router {
	router := mux.NewRouter()
	return router
}

func AuthenticationRouter(router *mux.Router) {
    authenticationSubrouter := router.PathPrefix("/authentication").Subrouter()
    authenticationSubrouter.HandleFunc("/gettoken/device", AuthenticationGetTokenDevice).Methods("POST")
}

func DeviceRouter(router *mux.Router) {
	deviceSubrouter := router.PathPrefix("/device").Subrouter()
    deviceSubrouter.Handle("/", negroni.New(negroni.HandlerFunc(JwtTokenValidationDevice), negroni.WrapFunc(Device))).Methods("GET")
    deviceSubrouter.Handle("/datapoint/ph", negroni.New(negroni.HandlerFunc(JwtTokenValidationDevice), negroni.WrapFunc(DeviceDataPointPhList))).Methods("GET")
    deviceSubrouter.Handle("/datapoint/ph", negroni.New(negroni.HandlerFunc(JwtTokenValidationDevice), negroni.HandlerFunc(ContentTypeValidationJson), negroni.WrapFunc(DeviceDataPointPhCreate))).Methods("POST")
    deviceSubrouter.Handle("/datapoint/oxygen", negroni.New(negroni.HandlerFunc(JwtTokenValidationDevice), negroni.WrapFunc(DeviceDataPointOxygenList))).Methods("GET")
    deviceSubrouter.Handle("/datapoint/oxygen", negroni.New(negroni.HandlerFunc(JwtTokenValidationDevice), negroni.HandlerFunc(ContentTypeValidationJson), negroni.WrapFunc(DeviceDataPointOxygenCreate))).Methods("POST")
    deviceSubrouter.Handle("/datapoint/temperature", negroni.New(negroni.HandlerFunc(JwtTokenValidationDevice), negroni.WrapFunc(DeviceDataPointTemperatureList))).Methods("GET")
    deviceSubrouter.Handle("/datapoint/temperature", negroni.New(negroni.HandlerFunc(JwtTokenValidationDevice), negroni.HandlerFunc(ContentTypeValidationJson), negroni.WrapFunc(DeviceDataPointTemperatureCreate))).Methods("POST")
}