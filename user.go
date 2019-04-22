package main

import (
    "net/http"
    "encoding/json"
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