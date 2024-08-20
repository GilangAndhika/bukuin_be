package models

type Book struct {
	ID          int    `gorm:"primaryKey;autoIncrement;column:id" json:"id"` 
	Title       string `gorm:"column:title" json:"title"`
	Author      string `gorm:"column:author" json:"author"`
	Description string `gorm:"column:description" json:"description"`
	LaunchYear  int    `gorm:"column:launch_year" json:"launch_year"`
	ISBN        string `gorm:"column:isbn;unique" json:"isbn"`
}
