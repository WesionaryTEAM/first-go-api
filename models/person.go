package models

import (
	"cloud-upload/config"
	"time"

	"gorm.io/gorm"
)

type Person struct {
	ID        uint64    `json:"id" gorm:"primary_key;size:28;not null;unique;auto_increment"`
	Name      string    `json:"name"`
	Photo     string    `json:"photo"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

func (p *Person) TableName() string {
	return "persons"
}

// func (p *Person) Save(db *gorm.DB) (*Person, error) {
// 	var err error
// 	err = db.Debug().Create(&p).Error
// 	if err != nil {
// 		return &Person{}, err
// 	}
// 	return p, nil
// }

func AddPerson(p *Person) (err error) {
	if err = config.DB.Create(p).Error; err != nil {
		return err
	}
	return nil
}

func (p *Person) Get(db *gorm.DB) (*[]Person, error) {
	var err error
	var persons []Person
	err = db.Debug().Find(&persons).Error
	if err != nil {
		return &[]Person{}, err
	}
	return &persons, nil
}

func (p *Person) Retrieve(db *gorm.DB, id string) (*Person, error) {
	err := db.Debug().Model(&Person{}).Where("id = ?", id).Take(&p).Error
	if err != nil {
		return &Person{}, err
	}
	// if gorm.IsRecordNotFoundError(err) {
	// 	return &Person{}, err
	// }
	return p, nil
}

func (p *Person) Delete(db *gorm.DB, id string) (bool, error) {
	db = db.Debug().Model(&Person{}).Where("id = ?", id).Take(&Person{}).Delete(&Person{})

	if db.Error != nil {
		return false, db.Error
	}
	return true, nil
}
