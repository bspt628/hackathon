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

type CreateRepostParams struct {
	ID                string
	UserID            string
	OriginalPostID    string
	IsQuoteRepost     bool
	AdditionalComment string
}

type DeleteRepostParams struct {
	UserID         string
	OriginalPostID string
}

type CheckRepostParams struct {
	UserID         string
	OriginalPostID string
}
