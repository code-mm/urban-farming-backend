package main

import (
    "net/http"
    "encoding/json"
    "github.com/gorilla/context"
    "gopkg.in/validator.v2"
)


/*
 * device
 */
func Device(w http.ResponseWriter, r *http.Request) {
    var device ModelDevice
    if _, err := Db.QueryOne(&device, `SELECT * FROM device WHERE identifier = ?`, context.Get(r, "deviceIdentifier").(string)); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    result, err := json.Marshal(device)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.Write(result)
}


/*
 *  device datapoint ph
 */
func DeviceDataPointPhList(w http.ResponseWriter, r *http.Request) {
    var dataPoint []ModelDeviceDataPointPh

    // query database for data points
    if _, err := Db.Query(&dataPoint, `SELECT device_datapoint_ph.time, device_datapoint_ph.value FROM device_datapoint_ph INNER JOIN device ON device_datapoint_ph.model_device_id = device.id WHERE device.identifier = ?`, context.Get(r, "deviceIdentifier").(string)); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    // marshal data to json
    result, err := json.Marshal(dataPoint)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
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
func DeviceDataPointOxygenList(w http.ResponseWriter, r *http.Request) {
    var dataPoint []ModelDeviceDataPointOxygen

    // query database for data points
    if _, err := Db.Query(&dataPoint, `SELECT device_datapoint_oxygen.time, device_datapoint_oxygen.value FROM device_datapoint_oxygen INNER JOIN device ON device_datapoint_oxygen.model_device_id = device.id WHERE device.identifier = ?`, context.Get(r, "deviceIdentifier").(string)); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    // marshal data to json
    result, err := json.Marshal(dataPoint)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(result)
}

func DeviceDataPointOxygenCreate(w http.ResponseWriter, r *http.Request) {
    var dataPoint ModelDeviceDataPointOxygen

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
 * device datapoint temperature
 */
func DeviceDataPointTemperatureList(w http.ResponseWriter, r *http.Request) {
    var dataPoint []ModelDeviceDataPointTemperature

    // query database for data points
    if _, err := Db.Query(&dataPoint, `SELECT device_datapoint_temperature.time, device_datapoint_temperature.value FROM device_datapoint_temperature INNER JOIN device ON device_datapoint_temperature.model_device_id = device.id WHERE device.identifier = ?`, context.Get(r, "deviceIdentifier").(string)); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    // marshal data to json
    result, err := json.Marshal(dataPoint)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(result)
}

func DeviceDataPointTemperatureCreate(w http.ResponseWriter, r *http.Request) {
    var dataPoint ModelDeviceDataPointTemperature

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
