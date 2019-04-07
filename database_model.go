package main

import (
    "time"
    "github.com/satori/go.uuid"
)

/*
 * user
 */
type ModelUser struct {
    tableName           struct{}        `sql:"user"`
    Id                  int             `sql:",pk"`
    Firstname           string          `sql:",notnull"`
    Lastname            string          `sql:",notnull"`
    Email               string          `sql:",notnull"`
}

/*
 * workspace
 */
type ModelWorkspace struct {
    tableName           struct{}        `sql:"workspace"`
    Id                  int             `sql:",pk"`
    Name                string          `sql:",notnull"`
    ModelUserId         int             `sql:"on_delete:RESTRICT, on_update:CASCADE"`
    ModelUser           *ModelUser
}

/*
 * device
 */
type ModelDevice struct {
    tableName           struct{}        `sql:"device"`
    Id                  int64           `sql:",pk"`
    Name                string
    Identifier          uuid.UUID       `sql:",type:uuid,unique,notnull"`
    Secret              string          `sql:",notnull"`
    ModelWorkspaceId    int             `sql:"on_delete:RESTRICT, on_update:CASCADE`
    ModelWorkspace      *ModelWorkspace
}

type ModelDeviceDataPointPh struct {
    tableName           struct{}        `sql:"device_datapoint_ph"`
    Id                  int64           `sql:",pk"`
    Time                time.Time       `sql:",notnull"`
    Value               float32         `sql:",notnull"`
    ModelDeviceId       int             `sql:"on_delete:RESTRICT, on_update:CASCADE`
    ModelDevice         *ModelDevice
}

type ModelDeviceDataPointOxygen struct {
    tableName           struct{}        `sql:"device_datapoint_oxygen"`
    Id                  int64           `sql:",pk"`
    Time                time.Time       `sql:",notnull"`
    Value               float32         `sql:",notnull"`
    ModelDeviceId       int             `sql:"on_delete:RESTRICT, on_update:CASCADE`
    ModelDevice         *ModelDevice
}

type ModelDeviceDataPointTemperature struct {
    tableName           struct{}        `sql:"device_datapoint_temperature"`
    Id                  int64           `sql:",pk"`
    Time                time.Time       `sql:",notnull"`
    Value               float32         `sql:",notnull"`
    ModelDeviceId       int             `sql:"on_delete:RESTRICT, on_update:CASCADE`
    ModelDevice         *ModelDevice
}

type ModelDeviceSetting struct {
    tableName           struct{}        `sql:"device_setting"`
    Id                  int64           `sql:",pk"`
    Key                 string          `sql:",notnull"`
    Value               string          `sql:",notnull"`
    ModelDeviceId       int             `sql:"on_delete:RESTRICT, on_update:CASCADE`
    ModelDevice         *ModelDevice
}
