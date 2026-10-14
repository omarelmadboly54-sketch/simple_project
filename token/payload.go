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
	Email		string		`json:"email"`
	Role		string		`json:"role"`
	UserID		int64		`json:"user_id"`
	IssuedAt	time.Time	`json:"issued_at"`
	ExpiredAt	time.Time	`json:"expired_at"`
}

func NewPayload(userID int64,email string,role string,duration time.Duration)(*Payload,error){
	payload:= &Payload{
		Email:email ,
		Role: role,
		UserID:userID,
		IssuedAt: time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload,nil
}

func(p *Payload)Valid()error{
	if time.Now().After(p.ExpiredAt){
		return ErrExpiredToken
	}
	return nil
}
