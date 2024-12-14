package model

import "github.com/google/uuid"

type UserCreateRequest struct {
		Email        string `json:"email"`
		Password 	 string `json:"password"`
		Username     string `json:"username"`
		DisplayName  string `json:"display_name"`
}

type UserCreateDAORequest struct {
	ID          string
	Email       string
	Password    string
	Username    string
	DisplayName string
}

func NewUserCreateDAORequest(email, password, username, displayName string) UserCreateDAORequest {
	return UserCreateDAORequest{
		ID:		 uuid.New().String(),
		Email:       email,
		Password:    password,
		Username:    username,
		DisplayName: displayName,
	}
}