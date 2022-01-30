package vo
type CreatePost struct {
	CategoryId uint `json:"category_id" bindding:"required"`
	Title string `json:"title" bindding:"required,max=10"`
	HeadImg string `json:"head_img"`
	Content string `json:"content" bindding:"required"`
}
