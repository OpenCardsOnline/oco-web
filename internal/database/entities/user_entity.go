package entities

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type UserEntity struct {
	Id                  uint               `json:"id"`
	CreatedAt           pgtype.Timestamptz `json:"created_at"`
	ModifiedAt          pgtype.Timestamptz `json:"modified_at"`
	IsArchived          bool               `json:"is_archived"`
	Email               string             `json:"email"`
	Username            string             `json:"username"`
	IsVerified          bool               `json:"is_verified"`
	PasswordHash        string             `json:"-"`
	LastLogin           pgtype.Timestamptz `json:"last_login"`
	FailedLoginAttempts uint               `json:"-"`
	IsBanned            bool               `json:"is_banned"`
	BanReason           pgtype.Text        `json:"ban_reason"`
	CanUseAPIKeys       bool               `json:"can_use_api_keys"`
}

type UserVerificationTokenEntity struct {
	Id         uint               `json:"id"`
	CreatedAt  pgtype.Timestamptz `json:"created_at"`
	ModifiedAt pgtype.Timestamptz `json:"modified_at"`
	IsArchived bool               `json:"is_archived"`
	TokenHash  string             `json:"token_hash"`
	UserId     uint               `json:"user_id"`
}
