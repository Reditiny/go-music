package dto

type CollectStatusDto struct {
	UserId int    `json:"userId"`
	Type   string `json:"type"`
	SongId int    `json:"songId"`
}
