package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("rahasia")

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     uint   `json:"role"`
	jwt.StandardClaims
}

func GenerateJWT(email, username string, role uint) (tokenString string, err error) {
	expTime := time.Now().Add(5 * time.Minute)

	claims := &JWTClaim{
		Username: username,
		Email:    email,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err = token.SignedString(jwtKey)

	return
}

// validate jwt token
func ValidateToken(signedToken string) (email string, role uint, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)

	if err != nil {
		return
	}

	Claim, ok := token.Claims.(*JWTClaim)

	// jika claim gagal
	if !ok {
		err = errors.New("Cloud are not parse claims for token")
		return
	}

	// jika expired
	if Claim.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("Token expired")
		return
	}

	role = Claim.Role
	email = Claim.Email

	return
}
