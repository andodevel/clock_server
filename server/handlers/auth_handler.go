package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/andodevel/go-echo-template/db"
	"github.com/andodevel/go-echo-template/helpers"
	"github.com/andodevel/go-echo-template/models"
)

// TODO: 1 JWT token - 1 session
// TODO: Check if we could migrate this group to GraphQL
// TODO: Check if we can inject DBConnection to echo Context

// Register ...
func Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Give me your private pics, I give you membership :)"})
	}
}

// LoginView ...
func LoginView() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenCookie, _ := c.Cookie(helpers.JWTCookieKey)
		if tokenCookie != nil {
			jwtToken := tokenCookie.Value
			isValid := helpers.IsValidJWTToken(jwtToken)
			if isValid { // Valid JWTToken
				redirect := c.QueryParam("redirect")
				if !helpers.IsEmpty(redirect) {
					// Override latest AccessToken
					newAccessToken := helpers.GenerateAccessToken()
					return c.Redirect(http.StatusFound, redirect+"?accessToken="+newAccessToken) // FIXME: '?' might not correct if there are params in redirect already.
				}

				return c.Redirect(http.StatusFound, "/")
			}
		}
		var html, _ = helpers.ParseHTMLTemplateFile("login", "server/go-echo-template/templates//login.html", helpers.Map{})
		return c.HTML(http.StatusOK, html)
	}
}

// Login ...
func Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		// TODO: Validation lib?
		if helpers.IsBlank(username) || helpers.IsBlank(password) {
			return c.Redirect(http.StatusFound, "/auth/login")
			// return echo.ErrBadRequest
		}

		db := db.CurrentDBConn()
		var user models.User
		err := db.Where("username = ?", username).Take(&user).Error
		if err != nil {
			return c.Redirect(http.StatusFound, "/auth/login")
			// return err
		}

		// TODO: Fix this hardcode password
		if "password" != password {
			return c.Redirect(http.StatusFound, "/auth/login")
			// return echo.ErrForbidden
		}

		jwtToken, err := helpers.CreateJWTToken(&user)
		if err != nil {
			log.Println("Error Creating JWT token", err)
			return c.Redirect(http.StatusFound, "/auth/login")
			// return c.String(http.StatusInternalServerError, "Something went wrong")
		}

		// TODO: NEHHH!!! Token in cookie????
		c.SetCookie(&http.Cookie{
			Path:     "/",
			Name:     helpers.JWTCookieKey,
			Value:    jwtToken,
			HttpOnly: true,
			Expires:  time.Now().Add(10 * time.Minute),
		})
		// FIXME: if API call?
		// return c.JSON(http.StatusOK, echo.Map{
		// 	"message": "You were logged in! JWT was set in your cookie!",
		// 	"token":   token,
		// })
		redirect := c.QueryParam("redirect")
		if !helpers.IsEmpty(redirect) {
			// Override latest AccessToken
			newAccessToken := helpers.GenerateAccessToken()
			return c.Redirect(http.StatusFound, redirect+"?accessToken="+newAccessToken) // FIXME: '?' might not correct if there are params in redirect already.
		}
		return c.Redirect(http.StatusFound, "/")
	}

}

// Logout ...
func Logout() echo.HandlerFunc {
	// TODO: FIXME
	return func(c echo.Context) error {
		c.SetCookie(&http.Cookie{
			Path:     "/",
			Name:     helpers.JWTCookieKey,
			Value:    "",
			HttpOnly: true,
			Expires:  time.Unix(0, 0),
		})
		return c.Redirect(http.StatusFound, "/")
	}
}

// CurrentUser ...
func CurrentUser() echo.HandlerFunc {
	// TODO: fix issue of message: "invalid or expired jwt"
	return func(c echo.Context) error {
		tokenCookie, _ := c.Cookie(helpers.JWTCookieKey)
		var claims *helpers.JWTClaims
		if tokenCookie != nil {
			claims, _ = helpers.ParseJWTToken(tokenCookie.Value)
		}
		if claims == nil {
			return echo.ErrForbidden
		}
		return c.JSON(http.StatusOK, echo.Map{"message": echo.Map{
			"username": claims.Username,
		}})
	}
}
