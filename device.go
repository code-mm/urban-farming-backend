package main

import (
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/context"
    "gopkg.in/validator.v2"
)


/*
 * device
 */
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


/*
 *  device datapoint ph
 */
func DeviceDataPointPhList(w http.ResponseWriter, r *http.Request) {
    deviceDataPointPh := new([]ModelDeviceDataPointPh)
    _, err := Db.QueryOne(deviceDataPointPh, `SELECT device_datapoint_ph.time, device_datapoint_ph.value FROM device_datapoint_ph INNER JOIN device ON device_datapoint_ph.model_device_id = device.id WHERE device.identifier = ?`, context.Get(r, "deviceIdentifier").(string))

    result, err := json.Marshal(deviceDataPointPh)
    if err != nil {
        log.Panic(err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(result)
}

func DeviceDataPointPhCreate(w http.ResponseWriter, r *http.Request) {
    var dataPoint ModelDeviceDataPointPh

    // decode json from request
    if err := json.NewDecoder(r.Body).Decode(&dataPoint); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    // validate submitted data
    if err := validator.Validate(dataPoint); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    // add device id to the model
    if _, err := Db.QueryOne(&(dataPoint.ModelDeviceId), `SELECT device.id FROM device WHERE identifier = ?`, context.Get(r, "deviceIdentifier").(string)); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    // write model to database
    if err := Db.Insert(&dataPoint); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    // return http status created
    w.WriteHeader(http.StatusCreated)
}


/*
 * device datapoint oxygen
 */
func DeviceDataPointOxygen(w http.ResponseWriter, r *http.Request) {
    deviceDataPointOxygen := new([]ModelDeviceDataPointOxygen)
    _, err := Db.QueryOne(deviceDataPointOxygen, `SELECT device_datapoint_oxygen.time, device_datapoint_oxygen.value FROM device_datapoint_oxygen INNER JOIN device ON device_datapoint_oxygen.model_device_id = device.id WHERE device.identifier = ?`, context.Get(r, "deviceIdentifier").(string))

    result, err := json.Marshal(deviceDataPointOxygen)
    if err != nil {
        log.Panic(err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(result)
}


/*
 * device datapoint temperature
 */
func DeviceDataPointTemperature(w http.ResponseWriter, r *http.Request) {
    DeviceDataPointTemperature := new([]ModelDeviceDataPointTemperature)
    _, err := Db.QueryOne(DeviceDataPointTemperature, `SELECT device_datapoint_temperature.time, device_datapoint_temperature.value FROM device_datapoint_temperature INNER JOIN device ON device_datapoint_temperature.model_device_id = device.id WHERE device.identifier = ?`, context.Get(r, "deviceIdentifier").(string))

    result, err := json.Marshal(DeviceDataPointTemperature)
    if err != nil {
        log.Panic(err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(result)
}
