package token

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const minSecretKeySize = 32

type JWTMaker struct {
	secretKey string
}

func NewJwtMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid secret key size:must be at least%d characters",minSecretKeySize)
	}
	return &JWTMaker{secretKey: secretKey},nil
}

func (maker *JWTMaker)CreateToken(email string,role string,duration time.Duration)(string,error){
	payload,err:=NewPayload(email,role,duration)
	if err!=nil{
		return "",err
	}
	jwtToken:=jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"email":payload.Email,
		"role":payload.Role,
		"exp":payload.ExpiredAt.Unix(),
		"iat":payload.IssuedAt.Unix(),
	})
	return jwtToken.SignedString([]byte(maker.secretKey))

}

func (maker *JWTMaker) VerifyToken(tokenString string) (*Payload, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, ErrInvalidToken
        }
        return []byte(maker.secretKey), nil
    })

    if err != nil {
        return nil, ErrInvalidToken
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        return nil, ErrInvalidToken
    }

    email, _ := claims["email"].(string)
    role, _ := claims["role"].(string)
    
    expFloat, ok := claims["exp"].(float64)
    if !ok {
        return nil, ErrInvalidToken
    }
    expiredAt := time.Unix(int64(expFloat), 0)

    if time.Now().After(expiredAt) {
        return nil, ErrExpiredToken
    }

    return &Payload{
        Email:     email,
        Role:      role,
        ExpiredAt: expiredAt,
    }, nil
}