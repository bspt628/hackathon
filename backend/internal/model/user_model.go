package model

type User struct {
	ID       int	`json:"id"`
	Username string	`json:"username"`
	Email    string	`json:"email"`
	DisplayName string `json:"display_name"`
}

type UserCreate struct {
	Username string	`json:"username"`
	Email    string	`json:"email"`
	DisplayName string `json:"display_name"`
}
