package auth

import (
	"os"
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/golang-jwt/jwt/v5"
)

var (
	secret = os.Getenv("JWT_SECRET_KEY")
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetMiddlewareConfig() jwtware.Config {
	config := jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(secret)},
		ContextKey: "jwt",
	}

	return config
}

func Generate(login *Login) (string, error) {
	claims := jwt.MapClaims{
		"email": login.Email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
		"admin": true,
	}

	t := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	token, err := t.SignedString(secret)

	if err != nil {
		return "", err
	}

	return token, nil
}
