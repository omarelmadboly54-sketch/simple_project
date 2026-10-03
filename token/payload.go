package token

import (
	"errors"
	"time"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken=errors.New("token is expired")
)

type Payload struct{
	Email  		string		`jsopn:"email"`
	Role   		string		`json:"role"`
	IssuedAt 	time.Time	`json:"issued_at"`
	ExpiredAt	time.Time	`json:"expired_at"`
}

func NewPayload(email string,role string,duration time.Duration)(*Payload,error){
	payload:=&Payload{
		Email: email,
		Role: role,
		IssuedAt: time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload,nil
}

func (payload *Payload) Valid() error{
	if time.Now().After(payload.ExpiredAt){
		return ErrExpiredToken
	}
	return nil
}
