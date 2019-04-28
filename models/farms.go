package models

import (
	"github.com/satori/go.uuid"
)


type Farm struct {
    ID                  int64
    Name                string          
    Identifier          uuid.UUID           `gorm:"type:uuid;not null;unique"`
    Secret              string              `gorm:"not null" json:"-"`
    Location            string
}
