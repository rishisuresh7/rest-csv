package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"rest-csv/models"
)

type Auth interface {
	Login(username, password string) (*models.User, error)
}

type auth struct {
	username string
	hash     string
	secret   string
}

func NewAuth(u, h, a string) Auth {
	return &auth{
		username: u,
		hash:     h,
		secret:   a,
	}
}

func (a *auth) Login(username, password string) (*models.User, error) {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	sum := hex.EncodeToString(hasher.Sum(nil))
	if username != a.username || sum != a.hash {
		return nil, fmt.Errorf("Login: invalid credentials")
	}

	currTime := time.Now()
	claims := jwt.MapClaims{
		"sub": username,
		"iat": currTime.Unix(),
		"exp": currTime.Add(12 * time.Hour),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(a.secret))
	if err != nil {
		return nil, fmt.Errorf("Login: unable to sign JWT")
	}

	return &models.User{
		Username:       username,
		Authentication: fmt.Sprintf("Bearer %s", signedToken),
	}, nil
}
