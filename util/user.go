package util

import "github.com/casdoor/casdoor-go-sdk/auth"

func GetUserId(claims *auth.Claims) (id string) {
	return claims.Name
}
