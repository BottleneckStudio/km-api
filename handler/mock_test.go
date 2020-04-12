package handler

import (
	"github.com/BottleneckStudio/km-api/services/post"
	"github.com/stretchr/testify/mock"
)

type postMock struct {
	mock.Mock
}

func (p *postMock) GetCount() int64 {
	args := p.Called()
	resp := args.Get(0)
	return resp.(int64)
}

func (p *postMock) CreatePost(vals map[string]interface{}) *post.Post {
	args := p.Called(vals)
	resp := args.Get(0)

	if resp == nil {
		return nil
	}

	return resp.(*post.Post)
}

func (p *postMock) GetPosts() []*post.Post {
	args := p.Called()
	resp := args.Get(0)

	if resp == nil {
		return nil
	}

	return resp.([]*post.Post)
}

func (p *postMock) GetUserPosts(username string) []*post.Post {
	args := p.Called(username)
	resp := args.Get(0)

	if resp == nil {
		return nil
	}

	return resp.([]*post.Post)
}

func (p *postMock) UpdatePost(id, cover, content string, created int64) error {
	args := p.Called(id, cover, content, created)
	_ = args.Get(0)
	return args.Error(0)
}

func (p *postMock) GetPost(id string) *post.Post {
	args := p.Called()
	resp := args.Get(0)

	if resp == nil {
		return nil
	}

	return resp.(*post.Post)
}
