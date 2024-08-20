package models

type Books struct {
	ID          int    `gorm:"primaryKey;column:id" json:"id"`
	Title       string `gorm:"column:title" json:"title"`
	Author      string `gorm:"column:author" json:"author"`
	Description string `gorm:"column:description" json:"description"`
	LaunchYear  int    `gorm:"column:launch_year" json:"launch_year"`
	ISBN        string `gorm:"column:isbn" json:"isbn"`
	CoverImage  string `gorm:"column:cover_image_url" json:"cover_image_url"`
}
