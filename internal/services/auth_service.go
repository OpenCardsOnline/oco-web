package services

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/opencardsonline/oco-web/config"
	"github.com/opencardsonline/oco-web/internal/database/entities"
	logger "github.com/opencardsonline/oco-web/internal/logging"
	"github.com/opencardsonline/oco-web/internal/models"
	"github.com/opencardsonline/oco-web/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

func VerifyNewUser(token string) bool {
	userVerificationToken := repositories.GetUserVerificationTokenByToken(token)
	if userVerificationToken == nil {
		logger.Log.Error("Invalid Token", errors.New("the verification token could not be found"))
		return false
	}

	existingUser := repositories.GetUserById(int(userVerificationToken.UserId))
	if existingUser == nil {
		logger.Log.Error("This user does not exist!", errors.New("cannot find the specified user"))
		return false
	}

	fmt.Println(existingUser)

	return true
}

func CreateNewUser(newUser models.NewUserRequest) *entities.UserEntity {
	existingUser := repositories.GetUserByEmail(newUser.Email)
	if existingUser != nil {
		logger.Log.Error("This user already exists!", errors.New("cannot create user because they already exist"))
		return nil
	}

	hashedPassword, err := HashPassword(newUser.Password)
	if err != nil {
		logger.Log.Error("An error occurred when attempting to hash the password", err)
		return nil
	}

	createdUser := repositories.InsertNewUser(newUser.Email, newUser.Username, hashedPassword)

	verificationToken, err := GenerateRandomString(32)
	if err != nil {
		logger.Log.Error("An error occurred when attempting to generate a verification token", err)
		return nil
	}

	finalToken := strconv.FormatInt(time.Now().Unix(), 10) + verificationToken

	fmt.Println("")
	fmt.Println(finalToken)

	htmlEmailContent := fmt.Sprintf(`
		<p>
			Welcome to OpenCardsOnline! To verify your account, please click the link: 
			<a href="%s/api/v1/auth/verify?token=%s">CLICK HERE</a>
		</p>
	`, config.AppConfiguration.APIBaseURL, finalToken)
	SendEmail("OpenCardsOnline", "do-not-reply@vistatable.com", "OpenCardsOnline - Verify Your Account", createdUser.Username, createdUser.Email, htmlEmailContent)

	repositories.InsertNewUserVerificationToken(&createdUser.Id, &finalToken)

	return createdUser
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateRandomString(n int) (string, error) {
	randomBytes := make([]byte, n)

	if _, err := io.ReadFull(rand.Reader, randomBytes); err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(randomBytes), nil
}
