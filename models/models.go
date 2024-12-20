package models

type Post struct {
	Id        int    `json:"id"`
	Body      string `json:"body"`
	UserId    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

type Comment struct {
	Id     int    `json:"id"`
	UserId int    `json:"user_id"`
	PostId int    `json:"post_id"`
	Body   string `json:"body"`
}

type CommentLike struct {
	Id     int    `json:"id"`
	UserId int    `json:"user_id"`
	PostId int    `json:"post_id"`
	Body   string `json:"body"`
}
type PostLike struct {
	Id     int    `json:"id"`
	UserId int    `json:"user_id"`
	PostId int    `json:"post_id"`
	Body   string `json:"body"`
}
