package models

import (
	"time"

	"gorm.io/gorm"
)

type Person struct {
	ID        uint64    `json:"id" gorm:"primary_key;size:28;not null;unique"`
	Name      string    `json:"name"`
	Photo     string    `json:"photo"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

func (p *Person) TableName() string {
	return "persons"
}

func (p *Person) Save(db *gorm.DB) (*Person, error) {
	var err error
	err = db.Debug().Create(&p).Error
	if err != nil {
		return &Person{}, err
	}
	return p, nil
}
