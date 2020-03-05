module github.com/andodevel/clock_server

go 1.13

require (
	// Runtime dependencies
	github.com/99designs/gqlgen v0.9.3
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gorilla/sessions v1.2.0
	github.com/jinzhu/gorm v1.9.10
	github.com/joho/godotenv v1.3.0
	github.com/labstack/echo-contrib v0.6.0
	github.com/labstack/echo/v4 v4.1.10
	github.com/skip2/go-qrcode v0.0.0-20191027152451-9434209cb086
	github.com/vektah/gqlparser v1.1.2
	// Dev dependencies. TODO: Remove from built package
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/tools v0.0.0-20200304193943-95d2e580d8eb // indirect
)
