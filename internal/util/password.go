package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (hashedpassword string, err error) {
	byte, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err!=nil{
		return "",fmt.Errorf("failed to hash password:%w",err)
	}
	return string(byte),nil
}

func VerfiyPassword (hashedpassword,password string)bool{
	err:=bcrypt.CompareHashAndPassword([]byte(hashedpassword),[]byte(password))
	return err == nil
}


