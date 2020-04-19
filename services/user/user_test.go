package user

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	cip "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAccountDetails(t *testing.T) {
	t.Run("valid token", func(t *testing.T) {
		to := new(MockedUserService)
		client := New("id", "secret")
		// change provider to our mocked object
		client.Provider = to
		to.On(
			"GetUser",
			mock.MatchedBy(func(in *cip.GetUserInput) bool {
				return true
			}),
		).Return(&cip.GetUserOutput{
			Username: aws.String("test"),
			UserAttributes: []*cip.AttributeType{
				{
					Name:  aws.String("email"),
					Value: aws.String("test@testing.com"),
				},
			},
		}, nil)

		u := client.AccountDetails("test")

		assert.Equal(t, "test@testing.com", u.Email)
		to.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		to := new(MockedUserService)
		client := New("id", "secret")
		// change provider to our mocked object
		client.Provider = to
		to.On(
			"GetUser",
			mock.MatchedBy(func(in *cip.GetUserInput) bool {
				return true
			}),
		).Return(nil, errors.New(""))

		u := client.AccountDetails("test")

		assert.Nil(t, u)
		to.AssertExpectations(t)
	})
}
