package entity

type Person struct {
	ID    uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name  string `json:"name" binding:"required" gorm:"type:varchar(32)"`
	Age   int8   `json:"age" binding:"gte=1,lte=130" `
	Email string `json:"email" validate:"required,email" gorm:"type:varchar(255)"`
}

type Video struct {
	ID          uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Title       string `json:"title" binding:"required" gorm:"type:varchar(100)"`
	Description string `json:"description" binding:"required" gorm:"type:varchar(200)"`
	URL         string `json:"url" binding:"required,url" gorm:"type:varchar(200);UNIQUE"`
	Author      Person `json:"author" binding:"required" gorm:"foreignkey:PersonID"`
	PersonID    uint64 `json:"-"`
}
