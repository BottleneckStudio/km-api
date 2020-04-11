package post

// Post ...
type Post struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Username string `json:"username"`
	UserPic  string `json:"userpic"`
	Created  int64  `json:"created"`
	Updated  int64  `json:"updated"`
	Publish  int    `json:"publish"`
	Content  string `json:"content"`
}
