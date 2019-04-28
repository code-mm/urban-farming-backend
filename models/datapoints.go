package models

import (
    "time"
)


type DataPointPh struct {
    ID                  int64
    Time                time.Time           `gorm:"type:timestamp with time zone;not null" validate:"regexp=^(\d{4})-(\d{2})-(\d{2})T(\d{2})\:(\d{2})\:(\d{2})\+(\d{2})\:(\d{2})$"`
    Value               float32             `gorm:"not null" validate:"min=0,max=10"`
    FarmID              int64               `gorm:"not null" json:"-"`
}

type DataPointOxygen struct {
    ID                  int64
    Time                time.Time           `gorm:"type:timestamp with time zone;not null" validate:"regexp=^(\d{4})-(\d{2})-(\d{2})T(\d{2})\:(\d{2})\:(\d{2})\+(\d{2})\:(\d{2})$"`
    Value               float32             `gorm:"not null" validate:"min=0,max=25"`
    FarmID              int64               `gorm:"not null" json:"-"`
}

type DataPointTemperature struct {
    ID                  int64
    Time                time.Time           `gorm:"type:timestamp with time zone;not null" validate:"regexp=^(\d{4})-(\d{2})-(\d{2})T(\d{2})\:(\d{2})\:(\d{2})\+(\d{2})\:(\d{2})$"`
    Value               float32             `gorm:"not null" validate:"min=-50,max=150"`
    FarmID              int64               `gorm:"not null" json:"-"`
}
