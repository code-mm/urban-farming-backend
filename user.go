package main

import (
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/gorilla/context"
)


/*
 * user
 */
func User(w http.ResponseWriter, r *http.Request) {
    var user ModelUser
    if _, err := Db.QueryOne(&user, `SELECT * FROM "user" WHERE email = ?`, context.Get(r, "email").(string)); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    result, err := json.Marshal(user)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.Write(result)
}


/*
 * user device list
 */
 func UserDeviceList(w http.ResponseWriter, r *http.Request) {
    var device []ModelDevice
    if _, err := Db.Query(&device, `select device.* from device INNER JOIN user_device_access ON device.id = user_device_access.model_device_id INNER JOIN "user" ON user_device_access.model_user_id = "user".id WHERE "user".email = ?`, context.Get(r, "email").(string)); err != nil {
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
 * user device datapoint
 */
func UserDeviceDatapointPh(w http.ResponseWriter, r *http.Request) {
    // check if device exists
    params := mux.Vars(r)

    var device []ModelDevice
    exists, err := Db.Model(&device).Where("id = ?", params["deviceId"]).Exists()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    if exists == false {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    // check if user has read access to device
    var deviceAccess ModelUserDeviceAccess
    exists, err = Db.Model(&deviceAccess).Join("JOIN \"user\" ON model_user_device_access.model_user_id = \"user\".id").Where("model_user_device_access.model_device_id=?", params["deviceId"]).Where("model_user_device_access.read=?", true).Where("\"user\".email=?", context.Get(r, "email").(string)).Exists()
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    if exists == false {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    // fetch data from database
    var dataPoint []ModelDeviceDataPointPh
    if _, err := Db.Query(&dataPoint, `SELECT device_datapoint_ph.time, device_datapoint_ph.value FROM device_datapoint_ph WHERE model_device_id = ?`, params["deviceId"]); err != nil {
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

func UserDeviceDatapointOxygen(w http.ResponseWriter, r *http.Request) {
    // check if device exists
    params := mux.Vars(r)

    var device []ModelDevice
    exists, err := Db.Model(&device).Where("id = ?", params["deviceId"]).Exists()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    if exists == false {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    // check if user has read access to device
    var deviceAccess ModelUserDeviceAccess
    exists, err = Db.Model(&deviceAccess).Join("JOIN \"user\" ON model_user_device_access.model_user_id = \"user\".id").Where("model_user_device_access.model_device_id=?", params["deviceId"]).Where("model_user_device_access.read=?", true).Where("\"user\".email=?", context.Get(r, "email").(string)).Exists()
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    if exists == false {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    // fetch data from database
    var dataPoint []ModelDeviceDataPointPh
    if _, err := Db.Query(&dataPoint, `SELECT device_datapoint_oxygen.time, device_datapoint_oxygen.value FROM device_datapoint_oxygen WHERE model_device_id = ?`, params["deviceId"]); err != nil {
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

func UserDeviceDatapointTemperature(w http.ResponseWriter, r *http.Request) {
    // check if device exists
    params := mux.Vars(r)

    var device []ModelDevice
    exists, err := Db.Model(&device).Where("id = ?", params["deviceId"]).Exists()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    if exists == false {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    // check if user has read access to device
    var deviceAccess ModelUserDeviceAccess
    exists, err = Db.Model(&deviceAccess).Join("JOIN \"user\" ON model_user_device_access.model_user_id = \"user\".id").Where("model_user_device_access.model_device_id=?", params["deviceId"]).Where("model_user_device_access.read=?", true).Where("\"user\".email=?", context.Get(r, "email").(string)).Exists()
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    if exists == false {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    // fetch data from database
    var dataPoint []ModelDeviceDataPointPh
    if _, err := Db.Query(&dataPoint, `SELECT device_datapoint_temperature.time, device_datapoint_temperature.value FROM device_datapoint_temperature WHERE model_device_id = ?`, params["deviceId"]); err != nil {
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