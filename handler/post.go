package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/BottleneckStudio/km-api/services/post"
	"github.com/go-chi/chi"
)

// GetPost ...
func GetPost(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer

	id := chi.URLParam(r, "id")

	ps := r.Context().Value("postService").(interface {
		GetPost(id string) *post.Post
	})

	post := ps.GetPost(id)
	if post == nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	body, err := json.Marshal(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	json.HTMLEscape(&buf, body)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, buf.String())
}

// GetPosts ...
func GetPosts(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer

	ps := r.Context().Value("postService").(interface {
		GetPosts() []*post.Post
	})

	posts := ps.GetPosts()
	body, err := json.Marshal(posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	json.HTMLEscape(&buf, body)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, buf.String())
}

// CreatePost ...
func CreatePost(w http.ResponseWriter, r *http.Request) {
	var payload post.Post
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Bad Input", http.StatusBadRequest)
		return
	}

	attr := map[string]interface{}{}
	attr["publish"] = payload.Publish
	attr["content"] = payload.Content
	attr["title"] = payload.Title
	attr["cover"] = payload.Username
	attr["author"] = payload.Author
	attr["userPic"] = payload.UserPic
	attr["username"] = payload.Username

	ps := r.Context().Value("postService").(interface {
		CreatePost(params map[string]interface{}) *post.Post
	})

	p := ps.CreatePost(attr)
	if p == nil {
		http.Error(w, "Failed to Create Post", http.StatusBadRequest)
		return
	}

	response, _ := json.Marshal(p)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(response))
}
