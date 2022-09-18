package auth

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
	"xm/domain"
)

var secret = []byte("BnFdv[DF9>2c`Oq!!(%^")

type claimskey int

var claimsKey claimskey

type AMech struct {
}

func AuthMechInit() *AMech {
	return &AMech{}
}
func (j *AMech) CreateToken(sub string, userInfo interface{}) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	expiration := time.Now().Add(time.Hour)
	token.Claims = &domain.JwtClaims{
		&jwt.RegisteredClaims{
			// Set the exp and sub claims.
			ExpiresAt: jwt.NewNumericDate(expiration),
			Subject:   sub,
		}, userInfo,
	}
	val, err := token.SignedString(secret)
	if err != nil {

		return "", err
	}
	return val, nil
}
func (j *AMech) GetClaimsFromToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
func (j *AMech) SetJWTClaimsContext(ctx context.Context, claims jwt.MapClaims) context.Context {
	return context.WithValue(ctx, claimsKey, claims)
}
