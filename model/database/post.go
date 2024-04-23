package database

import (
	"gorm.io/gorm"
	"shopping-cart/infrastructure"
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
func (post *Post) Model() *gorm.DB { return infrastructure.Db.Model(post) }
func (post *Post) FindAll() error {
	return post.Model().Find(&post).Error
}
func (post *Post) Create() error {
	return post.Model().Create(post).Error
}

func (post *Post) Update(updateData Post) error {
	return post.Model().Updates(updateData).Error
}

type Posts []Post

func (posts *Posts) model() *gorm.DB { return infrastructure.Db.Model(posts) }
func (posts *Posts) FindAll() error {
	return posts.model().Find(&posts).Error
}
