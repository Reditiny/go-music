package service

import (
	"go-music/model"
	"go-music/repository"
)

type ICommentService interface {
	SongListComment(list int) []*model.Comment
}

type CommentService struct {
	commentRepository repository.ICommentRepository
}

func NewCommentService() ICommentService {
	return CommentService{commentRepository: repository.NewCommentRepository()}
}

func (CS CommentService) SongListComment(list int) []*model.Comment {
	return CS.commentRepository.SongListComment(list)
}
