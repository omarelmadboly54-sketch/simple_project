package token

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const minsecretkeysize=32

type JWTMaker struct{
	secretKey string
}

func NewJwtMaker(secretKey string)(Maker,error){
	if len(secretKey)<minsecretkeysize{
		return nil,fmt.Errorf("Invalid SecretKey size must be at least:%d",minsecretkeysize)
	}
	return &JWTMaker{secretKey: secretKey},nil
}

func (maker *JWTMaker)CreateToken(userID int64,email string, role string, duration time.Duration)(string,error){
	Payload,err:=NewPayload(userID,email,role,duration)
	if err!=nil{
		return "",err
	}
	jwtToken:=jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"email":Payload.Email,
		"role":Payload.Role,
        "user_id":Payload.UserID,
		"exp":Payload.ExpiredAt.Unix(),
		"iat":Payload.IssuedAt.Unix(),
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

    userIDFloat, ok := claims["user_id"].(float64)
    if !ok {
        return nil, ErrInvalidToken
    }
    userID := int64(userIDFloat)

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
        UserID: userID,
        Email:     email,
        Role:      role,
        ExpiredAt: expiredAt,
    }, nil
}
