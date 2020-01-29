package account

import "github.com/dgrijalva/jwt-go"

type AccessTokenClaims struct {
	jwt.StandardClaims
	UserID string
	Admin  bool
}

type RefreshTokenClaims struct {
	jwt.StandardClaims
	UserID string
	Sub    int
}
