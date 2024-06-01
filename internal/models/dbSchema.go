package models

import (
	"time"
)

type User struct {
	ID       uint   `gorm:"primarykey" json:"id"`
	Username string `gorm:"unique" json:"username"`
	Email    string `gorm:"unique" json:"email"`
	Password []byte `json:"-"`                          // ignored during response i.e private
	Posts    []Post `gorm:"foreignKey:UserID" json:"-"` // 1 to many (one user can have many posts)
}

type Post struct {
	PostID uint `gorm:"primarykey" json:"post_id"`
	UserID uint `gorm:"index;not null" json:"user_id"` // FK
	User   User `gorm:"foreignKey:UserID" json:"-"`

	PostTitle       string    `json:"post_title"`
	PostDescription string    `json:"post_description"`
	PostContent     string    `json:"post_content"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
	Likes           []Like    `gorm:"foreignKey:PostID" json:"-"` // 1 to many (one post can have many likes)
	Comments        []Comment `gorm:"foreignKey:PostID" json:"-"` // 1 to many (one post can have many comments)
	ImageURL        string    `json:"image_url"`
}

type Like struct {
	ID     uint `gorm:"primarykey" json:"id"`
	UserID uint `gorm:"index:idx_user_post;not null" json:"user_id"` // FK
	PostID uint `gorm:"index:idx_user_post;not null" json:"post_id"` // FK
}

type Comment struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	PostID    uint      `gorm:"index;not null" json:"post_id"` // FK
	UserID    uint      `gorm:"index;not null" json:"user_id"` // FK
	User      User      `gorm:"foreignKey:UserID"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
