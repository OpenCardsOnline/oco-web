package services

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/opencardsonline/oco-web/config"
	"github.com/opencardsonline/oco-web/internal/database/entities"
	logger "github.com/opencardsonline/oco-web/internal/logging"
	"github.com/opencardsonline/oco-web/internal/models"
	"github.com/opencardsonline/oco-web/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	appConfig     *config.AppConfig
	db            *pgx.Conn
	_emailService *EmailService
	_userRepo     *repositories.UserRepository
}

func (_s *AuthService) New(appConfig *config.AppConfig, db *pgx.Conn) {
	_s.db = db
	_s.appConfig = appConfig
	_s._emailService = &EmailService{apiKey: appConfig.EmailAPIKey}
	_s._userRepo = &repositories.UserRepository{}
	_s._userRepo.New(db)
}

func (_s *AuthService) VerifyNewUser(token string) bool {
	userVerificationToken := _s._userRepo.GetUserVerificationTokenByToken(token)
	if userVerificationToken == nil {
		logger.Log.Error("Invalid Token", errors.New("the verification token could not be found"))
		return false
	}

	existingUser := _s._userRepo.GetUserById(int(userVerificationToken.UserId))
	if existingUser == nil {
		logger.Log.Error("This user does not exist!", errors.New("cannot find the specified user"))
		return false
	}

	return true
}

func (_s *AuthService) CreateNewUser(newUser models.NewUserRequest) *entities.UserEntity {
	existingUser := _s._userRepo.GetUserByEmail(newUser.Email)
	if existingUser != nil {
		logger.Log.Error("This user already exists!", errors.New("cannot create user because they already exist"))
		return nil
	}

	hashedPassword, err := _s.HashPassword(newUser.Password)
	if err != nil {
		logger.Log.Error("An error occurred when attempting to hash the password", err)
		return nil
	}

	createdUser := _s._userRepo.InsertNewUser(newUser.Email, newUser.Username, hashedPassword)

	verificationToken, err := _s.GenerateRandomString(32)
	if err != nil {
		logger.Log.Error("An error occurred when attempting to generate a verification token", err)
		return nil
	}

	finalToken := strconv.FormatInt(time.Now().Unix(), 10) + verificationToken

	htmlEmailContent := fmt.Sprintf(`
		<p>
			Welcome to OpenCardsOnline! To verify your account, please click the link: 
			<a href="%s/api/v1/auth/verify?token=%s">CLICK HERE</a>
		</p>
	`, _s.appConfig.APIBaseURL, finalToken)
	_s._emailService.SendEmail("OpenCardsOnline", "do-not-reply@vistatable.com", "OpenCardsOnline - Verify Your Account", createdUser.Username, createdUser.Email, htmlEmailContent)

	_s._userRepo.InsertNewUserVerificationToken(&createdUser.Id, &finalToken)

	return createdUser
}

func (_s *AuthService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (_s *AuthService) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (_s *AuthService) GenerateRandomString(n int) (string, error) {
	randomBytes := make([]byte, n)

	if _, err := io.ReadFull(rand.Reader, randomBytes); err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(randomBytes), nil
}
