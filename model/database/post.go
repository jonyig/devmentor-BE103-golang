package database

import (
	"devmentor-BE103-golang/infrastructure"
	"fmt"
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

func (post *Post) FindOne() error {
	return post.model().First(post).Error
}

func (post *Post) Create() error {
	return post.model().Create(post).Error
}

type Posts []Post

func (posts *Posts) model() *gorm.DB { return infrastructure.Db.Model(posts) }

func (posts *Posts) FindAll() error {
	return posts.model().FindInBatches(posts, 100, func(tx *gorm.DB, batch int) error {
		fmt.Printf("第 %d 批，查詢到 %d 筆資料：\n", batch, len(*posts))
		for _, user := range *posts {
			fmt.Printf("ID: %d, Name: %s, Age: %d\n", user.Title, user.Content, user.CreatedAt)
		}

		return nil
	}).Error
}
