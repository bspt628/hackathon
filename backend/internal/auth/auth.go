package auth

import (
	"context"
	"fmt"

	"firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var firebaseAuthClient *auth.Client
var authClient *auth.Client

// Firebase Admin SDKの初期化
func InitFirebase() error {
	opt := option.WithCredentialsFile("/Users/uchidahiroto/hackathon/backend/Term6_Hiroto_Uchida_Firebase_Admin_SDK.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return fmt.Errorf("error initializing app: %v", err)
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		return fmt.Errorf("error getting Auth client: %v", err)
	}

	firebaseAuthClient = client
	return nil
}

// Firebaseにユーザーを作成
func CreateFirebaseUser(email, password, username, displayName string) (*auth.UserRecord, error) {
	params := (&auth.UserToCreate{}).
		Email(email).
		Password(password).
		DisplayName(displayName).
		Disabled(false)

	userRecord, err := firebaseAuthClient.CreateUser(context.Background(), params)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %v", err)
	}

	return userRecord, nil
}
// Firebase認証を使ってユーザー情報を取得する例
func GetUserInfo(uid string) (*auth.UserRecord, error) {
    userRecord, err := authClient.GetUser(context.Background(), uid)
    if err != nil {
        return nil, fmt.Errorf("error getting user: %v", err)
    }
    return userRecord, nil
}
