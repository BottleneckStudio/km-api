package user

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cip "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

// New creates a new instance of Client
func New(id, secret string) *Client {
	return &Client{id, secret, provider()}
}

// AccountDetails ...
func (c *Client) AccountDetails(token string) *User {
	input := &cip.GetUserInput{}
	input.SetAccessToken(token)

	out, err := c.Provider.GetUser(input)
	if err != nil {
		log.Println("GetUser error", err.Error())
		return nil
	}

	return newUserFromAttributes(out.UserAttributes)
}

// provider returns a new cognito identity service
func provider() Provider {
	region := os.Getenv("AWS_REGION")
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))

	return cip.New(sess)
}

func newUserFromAttributes(attrs []*cip.AttributeType) *User {
	user := &User{}

	// attr mappings
	am := map[string]string{}
	for _, a := range attrs {
		am[*a.Name] = *a.Value
	}

	user.ID = am["sub"]
	user.Bio = am["profile"]
	user.Name = am["name"]
	user.Email = am["email"]
	user.Location = am["locale"]
	user.Website = am["website"]
	user.Picture = am["picture"]
	user.Username = am["nickname"]

	return user
}
