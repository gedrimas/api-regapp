package helpers

import (
	jwt "github.com/dgrijalva/jwt-go"
	"fmt"
	"time"
	//"api-regapp/helpers"
)

type JwtSignedDetails struct {
	Email     string
	Username  string
	User_id   string
	User_type string
	Company_id string
	jwt.StandardClaims
}

var SECRET_KEY = EnvFileVal("SECRET_KEY")


func ValidateToken(signedToken string) (claims *JwtSignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtSignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*JwtSignedDetails)
	if !ok {
		msg = fmt.Sprintf("This token is incorrect. Sorry!")
		msg = err.Error()
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("Ooops looks like your token has expired")
		msg = err.Error()
		return
	}
	return claims, msg
}