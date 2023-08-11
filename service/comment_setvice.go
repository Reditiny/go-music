package service

import (
	"go-music/dto"
	"go-music/model"
	"go-music/repository"
	"strconv"
)

type ICommentService interface {
	SongListComment(list int) []*model.Comment
	AddComment(comment dto.CommentDto)
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

func (CS CommentService) AddComment(comment dto.CommentDto) {
	// c := utils.CopyOne[dto.CommentDto, model.Comment](comment)
	// 前端传输的 songListId 为 string 类型 故在此做转换
	songListId, _ := strconv.Atoi(comment.SongListId)
	c := model.Comment{
		Content:    comment.Content,
		Type:       comment.Type,
		SongId:     comment.SongId,
		SongListId: songListId,
		UserId:     comment.UserId,
	}
	CS.commentRepository.AddComment(c)
}
