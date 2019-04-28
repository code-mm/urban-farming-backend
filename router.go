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
    authenticationSubrouter.HandleFunc("/gettoken/farm", AuthenticationGetTokenFarm).Methods("POST")
    authenticationSubrouter.HandleFunc("/gettoken/user", AuthenticationGetTokenUser).Methods("POST")
}

func FarmRouter(router *mux.Router) {
	deviceSubrouter := router.PathPrefix("/farm").Subrouter()
    deviceSubrouter.Handle("/", negroni.New(negroni.HandlerFunc(JwtTokenValidationFarm), negroni.WrapFunc(Farm))).Methods("GET")
    deviceSubrouter.Handle("/datapoint/ph", negroni.New(negroni.HandlerFunc(JwtTokenValidationFarm), negroni.WrapFunc(FarmDataPointPhList))).Methods("GET")
    deviceSubrouter.Handle("/datapoint/ph", negroni.New(negroni.HandlerFunc(JwtTokenValidationFarm), negroni.HandlerFunc(ContentTypeValidationJson), negroni.WrapFunc(FarmDataPointPhCreate))).Methods("POST")
    deviceSubrouter.Handle("/datapoint/oxygen", negroni.New(negroni.HandlerFunc(JwtTokenValidationFarm), negroni.WrapFunc(FarmDataPointOxygenList))).Methods("GET")
    deviceSubrouter.Handle("/datapoint/oxygen", negroni.New(negroni.HandlerFunc(JwtTokenValidationFarm), negroni.HandlerFunc(ContentTypeValidationJson), negroni.WrapFunc(FarmDataPointOxygenCreate))).Methods("POST")
    deviceSubrouter.Handle("/datapoint/temperature", negroni.New(negroni.HandlerFunc(JwtTokenValidationFarm), negroni.WrapFunc(FarmDataPointTemperatureList))).Methods("GET")
    deviceSubrouter.Handle("/datapoint/temperature", negroni.New(negroni.HandlerFunc(JwtTokenValidationFarm), negroni.HandlerFunc(ContentTypeValidationJson), negroni.WrapFunc(FarmDataPointTemperatureCreate))).Methods("POST")
}

func UserRouter(router *mux.Router) {
    userSubrouter := router.PathPrefix("/user").Subrouter()
    userSubrouter.Handle("/", negroni.New(negroni.HandlerFunc(JwtTokenValidationUser), negroni.WrapFunc(User))).Methods("GET")
    userSubrouter.Handle("/farm", negroni.New(negroni.HandlerFunc(JwtTokenValidationUser), negroni.WrapFunc(UserFarmList))).Methods("GET")
    userSubrouter.Handle("/farm/{farmId:[0-9]+}/datapoint/ph", negroni.New(negroni.HandlerFunc(JwtTokenValidationUser), negroni.WrapFunc(UserFarmDatapointPh))).Methods("GET")
    userSubrouter.Handle("/farm/{farmId:[0-9]+}/datapoint/oxygen", negroni.New(negroni.HandlerFunc(JwtTokenValidationUser), negroni.WrapFunc(UserFarmDatapointOxygen))).Methods("GET")
    userSubrouter.Handle("/farm/{farmId:[0-9]+}/datapoint/temperature", negroni.New(negroni.HandlerFunc(JwtTokenValidationUser), negroni.WrapFunc(UserFarmDatapointTemperature))).Methods("GET")
}
