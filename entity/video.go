package entity

type Video struct {
	Title       string `json:"title" binding:"required,min=2,max=10" validate:"title_contains_cool"`
	Description string `json:"description" binding:"required,max=20"`
	URL         string `json:"url" binding:"required,url"`
	Author      Person `json:"author" binding:"required"`
}
