package repository

import (
	"go-music/common"
	"go-music/model"
)

type ICommentRepository interface {
	SongListComment(list int) []*model.Comment
}

type CommentRepository struct {
}

func NewCommentRepository() ICommentRepository {
	return CommentRepository{}
}

func (c CommentRepository) SongListComment(list int) []*model.Comment {
	var comments []*model.Comment
	common.DB.Where("song_list_id = ?", list).Find(&comments)
	return comments
}
