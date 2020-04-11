package handler

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/BottleneckStudio/km-api/services/post"
)

// GetPost ...
func GetPost(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer

	id := r.URL.Query().Get(":id")
	post := post.Post{
		ID:       id,
		Author:   "Tibur",
		Username: "thetiburshow",
		Created:  0,
		Updated:  0,
		Content:  "<h1>Yessir!</h1>",
	}
	body, err := json.Marshal(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	json.HTMLEscape(&buf, body)
	w.WriteHeader(http.StatusOK)
	w.Write(buf.Bytes())
}

// GetPosts ...
func GetPosts(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer

	p := post.Post{
		ID:       "1234",
		Author:   "Tibur",
		Username: "thetiburshow",
		Created:  0,
		Updated:  0,
		Content:  "<h1>Yessir!</h1>",
	}
	posts := []post.Post{}
	posts = append(posts, p)
	body, err := json.Marshal(posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	json.HTMLEscape(&buf, body)
	w.WriteHeader(http.StatusOK)
	w.Write(buf.Bytes())
}
