package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/coffemanfp/beppin-server/database"
	dbu "github.com/coffemanfp/beppin-server/database/utils"
	errs "github.com/coffemanfp/beppin-server/errors"
	"github.com/coffemanfp/beppin-server/helpers"
	"github.com/coffemanfp/beppin-server/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// Login - Login user.
func Login(c echo.Context) (err error) {
	var m models.ResponseMessage
	var user models.User

	if err = c.Bind(&user); err != nil {
		return echo.ErrBadRequest
	}

	if !user.ValidateLogin() {
		m.Error = fmt.Sprintf("%v: %s", errs.ErrInvalidUserLogin)

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	db, err := database.Get()
	if err != nil {
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	dbUser, match, err := dbu.Login(db, user.Username, user.Password)
	if err != nil {
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	if !match {
		return echo.ErrUnauthorized
	}

	userI, err := helpers.ParseDBModelToModel(dbUser)
	if err != nil {
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	user = userI.(models.User)

	claim := models.Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token, err := claim.GenerateJWT()
	if err != nil {
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	m.Message = "Ok."
	m.Content = echo.Map{
		"token": token,
	}

	return c.JSON(http.StatusOK, m)
}
