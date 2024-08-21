package models

import (
	"github.com/dgrijalva/jwt-go"
)

type Books struct {
	IdBook      int    `gorm:"primaryKey;column:id_book" json:"id_book"`
	IdUser	    int    `gorm:"column:id_user" json:"id_user"`
	Title       string `gorm:"column:title" json:"title"`
	Author      string `gorm:"column:author" json:"author"`
	Description string `gorm:"column:description" json:"description"`
	LaunchYear  int    `gorm:"column:launch_year" json:"launch_year"`
	ISBN        string `gorm:"column:isbn" json:"isbn"`
	CoverImage  string `gorm:"column:cover_image_url" json:"cover_image_url"`
}

type GetJoinBooks struct {
	IdBook      int    `gorm:"primaryKey;column:id_book" json:"id_book"`
	IdUser	    int    `gorm:"column:id_user" json:"id_user"`
	Name		string `gorm:"column:name" json:"name,omitempty"`
	Title       string `gorm:"column:title" json:"title"`
	Author      string `gorm:"column:author" json:"author"`
	Description string `gorm:"column:description" json:"description"`
	LaunchYear  int    `gorm:"column:launch_year" json:"launch_year"`
	ISBN        string `gorm:"column:isbn" json:"isbn"`
	CoverImage  string `gorm:"column:cover_image_url" json:"cover_image_url"`
}

type Users struct {
	IdUser   uint   `gorm:"primaryKey;column:id_user" json:"id_user"`
	IdRole   int    `gorm:"column:id_role" json:"id_role"`
	Name     string `gorm:"column:name" json:"name"`
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	Email    string `gorm:"column:email" json:"email"`
}

type Roles struct {
	IdRole   int    `gorm:"primaryKey;column:id_role" json:"id_role"`
	RoleName string `gorm:"column:role_name" json:"role_name"`
}

type JWTClaims struct {
	jwt.StandardClaims
	IdUser uint `json:"id_user"`
	IdRole int  `json:"id_role"`
}