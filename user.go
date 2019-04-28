package main

import (
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/gorilla/context"
    "./models"
)


func User(w http.ResponseWriter, r *http.Request) {
    var userResult models.User

    if db.Where("email = ?", context.Get(r, "userEmail").(string)).First(&userResult).Error != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    result, err := json.Marshal(userResult)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.Write(result)
}

func UserFarmList(w http.ResponseWriter, r *http.Request) {
    var farmResult []models.Farm
    if db.Select("farms.*").Joins("INNER JOIN user_farm_permissions ON farms.id = user_farm_permissions.farm_id").Joins("INNER JOIN users ON user_farm_permissions.user_id = users.id").Where("email = ?", context.Get(r, "userEmail").(string)).Find(&farmResult).Error != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    result, err := json.Marshal(farmResult)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(result)
 }

func UserFarmDatapointPh(w http.ResponseWriter, r *http.Request) {
    // check if farm exists
    params := mux.Vars(r)

    var farmResult models.Farm
    if db.Where("id = ?", params["farmId"]).First(&farmResult).Error != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    // check if user has read access to farm
    var farmPermission models.UserFarmPermission
    if db.Joins("JOIN users ON user_farm_permissions.user_id = users.id").Where("user_farm_permissions.farm_id=?", params["farmId"]).Where("user_farm_permissions.read=?", true).Where("users.email=?", context.Get(r, "userEmail").(string)).First(&farmPermission).Error != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    // fetch data from database
    var dataPoint []models.DataPointPh
    if db.Select("data_point_phs.*").Where("farm_id = ?", params["farmId"]).Find(&dataPoint).Error != nil {
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

func UserFarmDatapointOxygen(w http.ResponseWriter, r *http.Request) {
    // check if farm exists
    params := mux.Vars(r)

    var farmResult models.Farm
    if db.Where("id = ?", params["farmId"]).First(&farmResult).Error != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    // check if user has read access to farm
    var farmPermission models.UserFarmPermission
    if db.Joins("JOIN users ON user_farm_permissions.user_id = users.id").Where("user_farm_permissions.farm_id=?", params["farmId"]).Where("user_farm_permissions.read=?", true).Where("users.email=?", context.Get(r, "userEmail").(string)).First(&farmPermission).Error != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    // fetch data from database
    var dataPoint []models.DataPointOxygen
    if db.Select("data_point_oxygens.*").Where("farm_id = ?", params["farmId"]).Find(&dataPoint).Error != nil {
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

func UserFarmDatapointTemperature(w http.ResponseWriter, r *http.Request) {
    // check if farm exists
    params := mux.Vars(r)

    var farmResult models.Farm
    if db.Where("id = ?", params["farmId"]).First(&farmResult).Error != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    // check if user has read access to farm
    var farmPermission models.UserFarmPermission
    if db.Joins("JOIN users ON user_farm_permissions.user_id = users.id").Where("user_farm_permissions.farm_id=?", params["farmId"]).Where("user_farm_permissions.read=?", true).Where("users.email=?", context.Get(r, "userEmail").(string)).First(&farmPermission).Error != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    // fetch data from database
    var dataPoint []models.DataPointTemperature
    if db.Select("data_point_temperatures.*").Where("farm_id = ?", params["farmId"]).Find(&dataPoint).Error != nil {
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
