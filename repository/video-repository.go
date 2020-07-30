package repository

import (
	"go-jwt/entity"

	"github.com/jinzhu/gorm"
)

type VideoRepository interface {
	Save(video entity.Video)
	Update(video entity.Video)
	Delete(video entity.Video)
	FindAll() []entity.Video
}

type database struct {
	connection *gorm.DB
}

func NewVideoRepositroy() VideoRepository {
	return nil
}
