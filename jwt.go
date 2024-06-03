package helpers

import (
    "time"
    "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key")

type JWTClaim struct {
    Email string `json:"email"`
    jwt.StandardClaims
}

func GenerateJWT(email string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &JWTClaim{
        Email: email,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    return tokenString, err
}

func ValidateToken(signedToken string) (*JWTClaim, error) {
    token, err := jwt.ParseWithClaims(
        signedToken,
        &JWTClaim{},
        func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        },
    )

    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(*JWTClaim)
    if !ok {
        return nil, err
    }

    return claims, nil
}