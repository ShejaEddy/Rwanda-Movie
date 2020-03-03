package services

import (
	"github.com/jinzhu/gorm"
	"github.com/projects/rwanda-movie/models"
	"github.com/satori/go.uuid"
)

func DeletePost(db *gorm.DB, postId uuid.UUID) bool {
	var post models.Post
	post.ID = postId
	err := db.Delete(&post).Error

	if err != nil {
		panic(err)
		return false
	}

	return true

}
