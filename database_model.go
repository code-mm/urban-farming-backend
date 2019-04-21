package main

import (
    "time"
    "github.com/satori/go.uuid"
)

/*
 * user
 */
type ModelUser struct {
    tableName           struct{}            `sql:"user"`
    Id                  int                 `sql:",pk" json:"-"`
    Firstname           string              `sql:",notnull"`
    Lastname            string              `sql:",notnull"`
    Email               string              `sql:",notnull"`
    Password            string              `sql:",notnull" json:"-"`
}

/*
 * device access
 */
type ModelUserDeviceAccess struct {
    tableName           struct{}            `sql:"user_device_access"`
    Id                  int                 `sql:",pk" json:"-"`
    Read                bool                `sql:",notnull, default:false"`
    Write               bool                `sql:",notnull, default:false"`
    ModelUserId         int                 `sql:"on_delete:CASCADE, on_update:CASCADE`
    ModelUser           *ModelUser          `json:"-"`
    ModelDeviceId       int64               `sql:"on_delete:CASCADE, on_update:CASCADE`
    ModelDevice         *ModelDevice        `json:"-"`
}

/*
 * device
 */
type ModelDevice struct {
    tableName           struct{}            `sql:"device"`
    Id                  int64               `sql:",pk" json:"-"`
    Name                string          
    Identifier          uuid.UUID           `sql:",type:uuid,unique,notnull"`
    Secret              string              `sql:",notnull" json:"-"`
    
}

type ModelDeviceDataPointPh struct {
    tableName           struct{}            `sql:"device_datapoint_ph"`
    Id                  int64               `sql:",pk" json:"-"`
    Time                time.Time           `sql:",notnull" validate:"regexp=^(\d{4})-(\d{2})-(\d{2})T(\d{2})\:(\d{2})\:(\d{2})\+(\d{2})\:(\d{2})$"`
    Value               float32             `sql:",notnull" validate:"min=0,max=10"`
    ModelDeviceId       int64               `sql:"on_delete:CASCADE, on_update:CASCADE" json:"-"`
    ModelDevice         *ModelDevice        `json:"-"`
}

type ModelDeviceDataPointOxygen struct {
    tableName           struct{}            `sql:"device_datapoint_oxygen"`
    Id                  int64               `sql:",pk" json:"-"`
    Time                time.Time           `sql:",notnull" validate:"regexp=^(\d{4})-(\d{2})-(\d{2})T(\d{2})\:(\d{2})\:(\d{2})\+(\d{2})\:(\d{2})$"`
    Value               float32             `sql:",notnull" validate:"min=0,max=25"`
    ModelDeviceId       int64               `sql:"on_delete:CASCADE, on_update:CASCADE" json:"-"`
    ModelDevice         *ModelDevice        `json:"-"`
}

type ModelDeviceDataPointTemperature struct {
    tableName           struct{}            `sql:"device_datapoint_temperature"`
    Id                  int64               `sql:",pk" json:"-"`
    Time                time.Time           `sql:",notnull" validate:"regexp=^(\d{4})-(\d{2})-(\d{2})T(\d{2})\:(\d{2})\:(\d{2})\+(\d{2})\:(\d{2})$"`
    Value               float32             `sql:",notnull" validate:"min=-50,max=150"`
    ModelDeviceId       int64               `sql:"on_delete:CASCADE, on_update:CASCADE" json:"-"`
    ModelDevice         *ModelDevice        `json:"-"`
}

type ModelDeviceSetting struct {
    tableName           struct{}            `sql:"device_setting"`
    Id                  int64               `sql:",pk"`
    Key                 string              `sql:",notnull"`
    Value               string              `sql:",notnull"`
    ModelDeviceId       int64               `sql:"on_delete:CASCADE, on_update:CASCADE`
    ModelDevice         *ModelDevice        `json:"-"`
}
