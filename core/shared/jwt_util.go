package shared

import (
	"errors"
	"time"

	"multiverse/core/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetUserEmail(c *gin.Context) (string, error) {
	email, ok := c.Get("user_email")
	if !ok {
		return "", errors.New("claim not found")
	}
	return email.(string), nil
}

func ValidateAndGetClaims(tokenString string) (jwt.MapClaims, error) {
	secret := []byte(config.Configs.Secrets.AuthServerJwtSecret)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signature")
		}
		return secret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}

}

// CreateToken creates a new token for a specific username and duration
func CreateJwtToken(email string, duration time.Duration) (string, error) {
	type Claims struct {
		Email string `json:"email"`
		jwt.StandardClaims
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
	})
	return jwtToken.SignedString([]byte(config.Configs.Secrets.AuthServerJwtSecret))
}
