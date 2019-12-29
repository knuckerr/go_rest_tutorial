package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/knuckerr/go_rest/api/models"
	"github.com/spf13/viper"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	ID       uint32 `json:"id"`
	jwt.StandardClaims
}

func Createtoken(u *models.User) (string, error) {
	t := time.Now().Add(40 * time.Minute)
	claims := &Claims{
		Username: u.Nickname,
		ID:       u.ID,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: t.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(viper.GetString("SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Vaildtoken(tokenstring string, claims *Claims) error {
	tkn, err := jwt.ParseWithClaims(tokenstring, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("SECRET_KEY")), nil
	})
	if !tkn.Valid {
		return errors.New("Invalid token")
	}
	if err != nil {
		return err
	}
	return nil

}

func Refreshtoken(tokenstring string) (string, error) {
	claims := &Claims{}
	err := Vaildtoken(tokenstring, claims)
	if err != nil {
		return "", err
	}
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		return "", errors.New("token must be older than 30 sec")
	}
	expirationTime := time.Now().Add(40 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token_new := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token_new.SignedString([]byte(viper.GetString("SECRET_KEY")))
	return tokenString, nil

}
