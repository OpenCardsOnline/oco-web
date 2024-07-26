package services

import (
	"fmt"
	"testing"
)

func TestRandomStringGenerator(t *testing.T) {
	authService := &AuthService{}

	randomString1, err := authService.GenerateRandomString(16)
	if err != nil {
		fmt.Println(err)
	}
	randomString2, err := authService.GenerateRandomString(16)
	if err != nil {
		fmt.Println(err)
	}
	if string(randomString1) == string(randomString2) {
		t.Fatalf(`GenerateRandomString(16): Expected Random Values But Got Same Values Or Error [%s]`, randomString1)
	}
	if len(string(randomString1)) < 10 {
		t.Fatalf(`GenerateRandomString(16): Expected RandomString1 To Have Length > 10 But Got [%s]`, randomString1)
	}
}

func TestBCryptHashing(t *testing.T) {
	authService := &AuthService{}

	originalPassword := "$up3rS3cretP@ssw0rd!"
	hashedPassword, err := authService.HashPassword(originalPassword)
	if err != nil {
		fmt.Println(err)
	}
	if string(originalPassword) == string(hashedPassword) {
		t.Fatalf(`TestBCryptHashing(): Expected Hashed Password To Be Different Than Original - ORIGINAL [%s], HASHED: [%s]`, originalPassword, hashedPassword)
	}

	passwordCheckResult1 := authService.CheckPasswordHash(originalPassword, hashedPassword)
	if !passwordCheckResult1 {
		t.Fatalf(`TestBCryptHashing(): Expected Hashed Password To Be Verified - Value [%v]`, passwordCheckResult1)
	}

	passwordCheckResult2 := authService.CheckPasswordHash(originalPassword+"_wrong", hashedPassword)
	if passwordCheckResult2 {
		t.Fatalf(`TestBCryptHashing(): Expected Hashed Password To NOT Be Verified - Value [%v]`, passwordCheckResult2)
	}
}
