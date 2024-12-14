package model

import (
	"database/sql"
	"hackathon/db/sqlc/generated"
)

// ToCreateUserParams は UserCreateDAORequest を sqlc.CreateUserParams に変換する
func ToCreateUserParams(request UserCreateDAORequest, firebaseUID string) sqlc.CreateUserParams {
	return sqlc.CreateUserParams{
		ID:           request.ID,
		FirebaseUid:  firebaseUID,
		Email:        request.Email,
		PasswordHash: request.Password,
		Username:     request.Username,
		DisplayName:  sql.NullString{String: request.DisplayName, Valid: request.DisplayName != ""},
	}
}