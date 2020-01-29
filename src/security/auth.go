package security

import (
	"strconv"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/theodo/scalab/config"
	"github.com/theodo/scalab/src/entities"
	"github.com/theodo/scalab/src/repositories"
	"golang.org/x/crypto/bcrypt"
)

type userSigninProvider struct{}

func NewUserSigninProvider() userSigninProvider {
	return userSigninProvider{}
}

type userFromContextProvider struct{}

func NewUserFromContextProvider() userFromContextProvider {
	return userFromContextProvider{}
}

type userAccountCreator struct{}

func NewUserAccountCreator() userAccountCreator {
	return userAccountCreator{}
}

func (up *userSigninProvider) SignIn(cred Credentials, c echo.Context) (user entities.User, err error) {
	r := repositories.NewUserRepository()
	user, err = r.FindByEmail(cred.Username)
	if err != nil {
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cred.Password)); err != nil {
		return user, &PasswordMismatchError{}
	}

	sess, _ := session.Get(config.GetSession().CookieName, c)
	sess.Values["user_id"] = user.Id
	sess.Save(c.Request(), c.Response())

	return
}

func (up *userFromContextProvider) GetId(c echo.Context) (id int, err error) {
	sess, _ := session.Get(config.GetSession().CookieName, c)
	stringId, ok := sess.Values["user_id"].(string)
	if !ok {
		err = &InvalidAuthSessionError{}
		return
	}
	id, err = strconv.Atoi(stringId)
	if err != nil {
		err = &InvalidAuthSessionError{err}
		return
	}
	return
}

func (up *userFromContextProvider) Get(c echo.Context) (user entities.User, err error) {
	id, err := up.GetId(c)
	if err != nil {
		return
	}
	r := repositories.NewUserRepository()
	user, err = r.Find(id)
	return
}

func (uc *userAccountCreator) Create(cred ConfirmedCredentials) (user entities.User, err error) {
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
