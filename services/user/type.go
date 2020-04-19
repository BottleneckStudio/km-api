package user

import cip "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"

// Provider ...
type Provider interface {
	GetUser(*cip.GetUserInput) (*cip.GetUserOutput, error)
}

// Client main struct
type Client struct {
	ClientID     string
	ClientSecret string
	Provider     Provider
}

// User ...
type User struct {
	ID       string
	Bio      string
	Name     string
	Email    string
	Location string
	Website  string
	Picture  string
	Username string
}
