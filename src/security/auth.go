package security

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/theodo/scalab/config"
	"github.com/theodo/scalab/src/entities"
	"github.com/theodo/scalab/src/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserFromCredentialProvider interface {
	Get(cred Credentials, c echo.Context) (entities.User, error)
}

type userFromCredentialProvider struct{}

func NewUserFromCredentialProvider() UserFromCredentialProvider {
	return userFromCredentialProvider{}
}

type UserFromContextProvider interface {
	Get(c echo.Context) (entities.User, error)
}

type userFromContextProvider struct{}

func NewUserFromContextProvider() UserFromContextProvider {
	return userFromContextProvider{}
}

type UserAccountCreator interface {
	Create(cred ConfirmedCredentials) (entities.User, error)
}

type userAccountCreator struct{}

func NewUserAccountCreator() UserAccountCreator {
	return userAccountCreator{}
}

func (up userFromCredentialProvider) Get(cred Credentials, c echo.Context) (user entities.User, err error) {
	r := repositories.NewUserRepository()
	user, err = r.FindByEmail(cred.Username)
	if err != nil {
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cred.Password)); err != nil {
		return user, &PasswordMismatchError{}
	}

	sess, _ := session.Get(config.GetSession().CookieName, c)
	sess.Values["email"] = user.Email
	sess.Save(c.Request(), c.Response())

	return
}

func (up userFromContextProvider) Get(c echo.Context) (user entities.User, err error) {
	sess, _ := session.Get(config.GetSession().CookieName, c)
	r := repositories.NewUserRepository()
	email, ok := sess.Values["email"].(string)
	if !ok {
		return
	}
	user, err = r.FindByEmail(email)
	return
}

func (uc userAccountCreator) Create(cred ConfirmedCredentials) (user entities.User, err error) {
	if cred.Password != cred.ConfirmedPassword {
		err = &ConfirmedPasswordMismatchError{}
		return
	}

	var password []byte
	password, err = bcrypt.GenerateFromPassword([]byte(cred.Password), config.GetSecurity().Cost)
	if err != nil {
		return
	}
	user = entities.User{Name: cred.Username, Email: cred.Username, Password: string(password)}

	r := repositories.NewUserRepository()
	user, err = r.Create(user)

	return
}
