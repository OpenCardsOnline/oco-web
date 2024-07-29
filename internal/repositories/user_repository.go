package repositories

import (
	"github.com/opencardsonline/oco-web/internal/database"
	"github.com/opencardsonline/oco-web/internal/database/entities"
	logger "github.com/opencardsonline/oco-web/logging"
)

type UserRepository struct {
	db *database.AppDBConn
}

func (_r *UserRepository) New(db *database.AppDBConn) {
	_r.db = db
}

func (_r *UserRepository) GetUserByEmail(email string) *entities.UserEntity {
	sql := `
		SELECT
			id,
			created_at,
			modified_at,
			is_archived,
			email,
			username,
			is_verified,
			password_hash,
			last_login,
			failed_login_attempts,
			is_banned,
			ban_reason,
			can_use_api_keys
		FROM public.users
		WHERE email = $1
	`
	var user entities.UserEntity
	err := _r.db.QueryRow(sql, email).Scan(
		&user.Id,
		&user.CreatedAt,
		&user.ModifiedAt,
		&user.IsArchived,
		&user.Email,
		&user.Username,
		&user.IsVerified,
		&user.PasswordHash,
		&user.LastLogin,
		&user.FailedLoginAttempts,
		&user.IsBanned,
		&user.BanReason,
		&user.CanUseAPIKeys)
	if err != nil {
		logger.Log.Error("an error occurred when attempting to get the user by email", "UserRepository.GetUserByEmail", err)
		return nil
	}
	return &user
}

func (_r *UserRepository) GetUserById(id int) *entities.UserEntity {
	sql := `
		SELECT
			id,
			created_at,
			modified_at,
			is_archived,
			email,
			username,
			is_verified,
			password_hash,
			last_login,
			failed_login_attempts,
			is_banned,
			ban_reason,
			can_use_api_keys
		FROM public.users
		WHERE id = $1
	`
	var user entities.UserEntity
	err := _r.db.QueryRow(sql, id).Scan(
		&user.Id,
		&user.CreatedAt,
		&user.ModifiedAt,
		&user.IsArchived,
		&user.Email,
		&user.Username,
		&user.IsVerified,
		&user.PasswordHash,
		&user.LastLogin,
		&user.FailedLoginAttempts,
		&user.IsBanned,
		&user.BanReason,
		&user.CanUseAPIKeys)
	if err != nil {
		logger.Log.Error("an error occurred when attempting to get the user by id", "UserRepository.GetUserById", err)
		return nil
	}
	return &user
}

func (_r *UserRepository) InsertNewUser(email string, username string, passwordHash string) *entities.UserEntity {
	sql := `
		INSERT INTO public.users (
			email,
			username,
			password_hash
		) 
		VALUES ($1,$2,$3)
		RETURNING id
	`

	lastInsertId, err := _r.db.InsertRowAndReturnLastID(sql, email, username, passwordHash)
	if err != nil {
		logger.Log.Error("an error occurred when attempting to insert new user", "UserRepository.InsertNewUser", err)
		return nil
	}
	result := _r.GetUserById(*lastInsertId)
	if result == nil {
		return nil
	}
	return result
}

func (_r *UserRepository) InsertNewUserVerificationToken(userId *uint, hashedToken *string) {
	sql := `
		INSERT INTO public.user_verification_tokens (
			user_id,
			token_hash
		) 
		VALUES ($1,$2)
	`
	err := _r.db.ExecuteQueryWithNoReturn(sql, userId, hashedToken)
	if err != nil {
		logger.Log.Error("an error occurred when attempting to insert new user verification token", "UserRepository.InsertNewUserVerificationToken", err)
		return
	}
}

func (_r *UserRepository) GetUserVerificationTokenByToken(token string) *entities.UserVerificationTokenEntity {
	sql := `
		SELECT
			id,
			created_at,
			modified_at,
			is_archived,
			token_hash,
			user_id
		FROM public.user_verification_tokens
		WHERE token_hash = $1
	`
	var userVerificationTokenEntity entities.UserVerificationTokenEntity
	err := _r.db.QueryRow(sql, token).Scan(
		&userVerificationTokenEntity.Id,
		&userVerificationTokenEntity.CreatedAt,
		&userVerificationTokenEntity.ModifiedAt,
		&userVerificationTokenEntity.IsArchived,
		&userVerificationTokenEntity.TokenHash,
		&userVerificationTokenEntity.UserId)
	if err != nil {
		logger.Log.Error("an error occurred when attempting to get user verification token", "UserRepository.GetUserVerificationTokenByToken", err)
		return nil
	}
	return &userVerificationTokenEntity
}
