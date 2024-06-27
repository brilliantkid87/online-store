package middleware

import (
	"net/http"
	"strings"
	jwtToken "synapsis/pkg/jwt"

	"log"

	"github.com/labstack/echo/v4"
)

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")

		if token == "" {
			log.Println("unauthorized", http.StatusBadRequest)
			return c.JSON(http.StatusUnauthorized, "unauthorized")
		}

		token = strings.Split(token, " ")[1]
		claims, err := jwtToken.DecodeToken(token)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, "unathorized")
		}

		c.Set("userLogin", claims)
		return next(c)
	}
}
