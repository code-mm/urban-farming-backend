package models


type User struct {
    ID                  int64
    Firstname           string              `gorm:"not null"`
    Lastname            string              `gorm:"not null"`
    Email               string              `gorm:"not null;unique"`
    Password            string              `gorm:"not_null" json:"-"`
}
