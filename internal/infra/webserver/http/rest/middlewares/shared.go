package middlewares

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/golang-jwt/jwt"
)

func extractPermissionLevel(tokenString string) (int, error) {
	var pm int
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return -1, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		pm, err = strconv.Atoi(fmt.Sprint(claims["permission_level"]))
		if err != nil {
			return -1, errors.New("invalid token")
		}
	}

	return pm, nil
}
