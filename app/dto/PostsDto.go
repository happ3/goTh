package dto

type PostsDto struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId  int    `json:"user_id"`
	BaseDto
}
