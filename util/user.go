package util

import "github.com/casdoor/casdoor-go-sdk/auth"

func GetUserId(claims *auth.Claims) (id string) {
	return claims.Name
}

func IsStudent(claims *auth.Claims) (b bool) {
	return claims.Tag != "老师"
}
func IsTeacher(claims *auth.Claims) (b bool) {
	return claims.Tag == "老师"
}