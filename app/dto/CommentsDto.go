package dto

type CommentsDto struct {
	Id      int64  `json:"id"`
	UserId  int64  `json:"user_id"`
	Content string `json:"content"`
	PostId  int64  `json:"post_id"`
}
