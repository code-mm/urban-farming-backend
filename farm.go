package main

import (
    "net/http"
    "encoding/json"
    "github.com/gorilla/context"
    "gopkg.in/validator.v2"
    "./models"
)


func Farm(w http.ResponseWriter, r *http.Request) {
    var farmResult models.Farm
    if db.Where("identifier = ?", context.Get(r, "farmIdentifier").(string)).First(&farmResult).Error != nil {
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

func FarmDataPointPhList(w http.ResponseWriter, r *http.Request) {
    var dataPoint []models.DataPointPh

    // query database for data points
    if db.Select("data_point_phs.id, data_point_phs.time, data_point_phs.value").Joins("INNER JOIN farms ON data_point_phs.farm_id = farms.id").Where("identifier = ?", context.Get(r, "farmIdentifier").(string)).Find(&dataPoint).Error != nil {
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

func FarmDataPointPhCreate(w http.ResponseWriter, r *http.Request) {
    var dataPoint models.DataPointPh

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

    // add farm id to the model
    var farm models.Farm
    if db.Select("ID").Where("identifier = ?", context.Get(r, "farmIdentifier").(string)).First(&farm).Error != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    dataPoint.FarmID = farm.ID

    // write model to database
    if db.Create(&dataPoint).Error != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    // return http status created
    w.WriteHeader(http.StatusCreated)
}

func FarmDataPointOxygenList(w http.ResponseWriter, r *http.Request) {
    var dataPoint []models.DataPointOxygen

    // query database for data points
    if db.Select("data_point_oxygens.id, data_point_oxygens.time, data_point_oxygens.value").Joins("INNER JOIN farms ON data_point_oxygens.farm_id = farms.id").Where("identifier = ?", context.Get(r, "farmIdentifier").(string)).Find(&dataPoint).Error != nil {
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

func FarmDataPointOxygenCreate(w http.ResponseWriter, r *http.Request) {
    var dataPoint models.DataPointOxygen

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

    // add farm id to the model
    var farm models.Farm
    if db.Select("ID").Where("identifier = ?", context.Get(r, "farmIdentifier").(string)).First(&farm).Error != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    dataPoint.FarmID = farm.ID

    // write model to database
    if db.Create(&dataPoint).Error != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    // return http status created
    w.WriteHeader(http.StatusCreated)
}

func FarmDataPointTemperatureList(w http.ResponseWriter, r *http.Request) {
    var dataPoint []models.DataPointTemperature

    // query database for data points
    if db.Select("data_point_temperatures.id, data_point_temperatures.time, data_point_temperatures.value").Joins("INNER JOIN farms ON data_point_temperatures.farm_id = farms.id").Where("identifier = ?", context.Get(r, "farmIdentifier").(string)).Find(&dataPoint).Error != nil {
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

func FarmDataPointTemperatureCreate(w http.ResponseWriter, r *http.Request) {
    var dataPoint models.DataPointTemperature

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

    // add farm id to the model
    var farm models.Farm
    if db.Select("ID").Where("identifier = ?", context.Get(r, "farmIdentifier").(string)).First(&farm).Error != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    dataPoint.FarmID = farm.ID

    // write model to database
    if db.Create(&dataPoint).Error != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    // return http status created
    w.WriteHeader(http.StatusCreated)
}
