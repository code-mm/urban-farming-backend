package models


type UserFarmPermission struct {
    ID                  int64
    Read                bool                `gorm:"not null"`
    Write               bool                `gorm:"not null"`
    UserID              int64               `gorm:"not null"`
    FarmID              int64               `gorm:"not null"`
}
