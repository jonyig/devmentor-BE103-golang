package database

import (
	"devmentor-BE103-golang/infrastructure"
	"gorm.io/gorm"
	"time"
)

// Post model
type Post struct {
	Id        int       `json:"id",gorm:"primarykey"`
	Title     string    `json:"title",gorm:"column:title"`
	Content   string    `json:"content",gorm:"column:content"`
	CreatedAt time.Time `json:"created_at",gorm:"column:created_at"`
}

func (Post) TableName() string {
	return "posts"
}

func (post *Post) model() *gorm.DB { return infrastructure.Db.Model(post) }

func (post *Post) FindAll() error {
	return post.model().Find(&post).Error
}

func (post *Post) Create() error {
	return post.model().Create(post).Error
}

type Posts []Post

func (posts *Posts) model() *gorm.DB { return infrastructure.Db.Model(posts) }

func (posts *Posts) FindAll() error {
	return posts.model().Find(&posts).Error
}
