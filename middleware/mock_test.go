package middleware

import (
	"github.com/BottleneckStudio/km-api/services/user"
	"github.com/stretchr/testify/mock"
)

type userServiceMock struct {
	mock.Mock
}

func (o *userServiceMock) AccountDetails(token string) *user.User {
	args := o.Called(token)

	resp := args.Get(0)
	if resp == nil {
		return nil
	}

	return resp.(*user.User)
}
