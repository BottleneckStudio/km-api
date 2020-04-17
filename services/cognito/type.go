package cognito

import cip "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"

// Provider ...
type Provider interface {
	SignUp(*cip.SignUpInput) (*cip.SignUpOutput, error)
	ConfirmSignUp(*cip.ConfirmSignUpInput) (*cip.ConfirmSignUpOutput, error)
	InitiateAuth(*cip.InitiateAuthInput) (*cip.InitiateAuthOutput, error)
	GetUser(*cip.GetUserInput) (*cip.GetUserOutput, error)
	AdminGetUser(*cip.AdminGetUserInput) (*cip.AdminGetUserOutput, error)
	UpdateUserAttributes(*cip.UpdateUserAttributesInput) (*cip.UpdateUserAttributesOutput, error)
	ChangePassword(*cip.ChangePasswordInput) (*cip.ChangePasswordOutput, error)
}
