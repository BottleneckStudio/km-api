package user

import (
	cip "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/stretchr/testify/mock"
)

type MockedUserService struct {
	mock.Mock
}

func (m *MockedUserService) GetUser(in *cip.GetUserInput) (*cip.GetUserOutput, error) {
	args := m.Called(in)

	resp := args.Get(0)
	if resp == nil {
		return nil, args.Error(1)
	}

	return resp.(*cip.GetUserOutput), args.Error(1)
}
