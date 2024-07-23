package repositories

import (
	"context"
	"fmt"

	"github.com/opencardsonline/oco-web/internal/database"
	"github.com/opencardsonline/oco-web/internal/database/entities"
)

func GetUserByEmail(email string) *entities.UserEntity {
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
	err := database.DB.QueryRow(context.Background(), sql, email).Scan(
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
		fmt.Println(err.Error())
		return nil
	}
	return &user
}

func GetUserById(id int) *entities.UserEntity {
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
	err := database.DB.QueryRow(context.Background(), sql, id).Scan(
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
		fmt.Print(err.Error())
		return nil
	}
	return &user
}

func InsertNewUser(email string, username string, passwordHash string) *entities.UserEntity {
	sql := `
		INSERT INTO public.users (
			email,
			username,
			password_hash
		) 
		VALUES ($1,$2,$3)
		RETURNING id
	`
	var lastInsertId int
	err := database.DB.QueryRow(context.Background(), sql, email, username, passwordHash).Scan(&lastInsertId)
	if err != nil {
		fmt.Print(err.Error())
		return nil
	}
	result := GetUserById(lastInsertId)
	if result == nil {
		return nil
	}
	return result
}

func InsertNewUserVerificationToken(userId *uint, hashedToken *string) {
	sql := `
		INSERT INTO public.user_verification_tokens (
			user_id,
			token_hash
		) 
		VALUES ($1,$2)
	`
	_, err := database.DB.Exec(context.Background(), sql, userId, hashedToken)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
}

func GetUserVerificationTokenByToken(token string) *entities.UserVerificationTokenEntity {
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
	err := database.DB.QueryRow(context.Background(), sql, token).Scan(
		&userVerificationTokenEntity.Id,
		&userVerificationTokenEntity.CreatedAt,
		&userVerificationTokenEntity.ModifiedAt,
		&userVerificationTokenEntity.IsArchived,
		&userVerificationTokenEntity.TokenHash,
		&userVerificationTokenEntity.UserId)
	if err != nil {
		fmt.Print(err.Error())
		return nil
	}
	return &userVerificationTokenEntity
}