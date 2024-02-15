package entity

//go:generate mockgen -source=user_authentication.go -destination=mock/user_authentication.go -package=mock

type INewUserAuthentication interface {
	NewUserAuthentication(phoneNumber, password string) IUserAuthentication
}

type userAuthenticationFactory struct{}

func NewUserAuthenticationFactory() INewUserAuthentication {
	return &userAuthenticationFactory{}
}

func (u *userAuthenticationFactory) NewUserAuthentication(phoneNumber, password string) IUserAuthentication {
	return nil
}

type IUserAuthentication interface {
	PhoneNumber() string
	Password() string
	Validate() []error
}

type userAuthentication struct {
	phoneNumber string
	password    string
}

func (u *userAuthentication) PhoneNumber() string {
	return u.phoneNumber
}

func (u *userAuthentication) Password() string {
	return u.password
}

func (u *userAuthentication) Validate() []error {
	return nil
}
