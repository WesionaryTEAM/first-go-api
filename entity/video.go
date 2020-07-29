package entity

type Person struct {
	Name  string `json:"name" binding:"required"`
	Age   int8   `json:"age" binding:"gte=1,lte=130"`
	Email string `json:"email" validate:"required,email"`
}

type Video struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	URL         string `json:"url" binding:"required,url"`
	Author      Person `json:"author" binding:"required"`
}
