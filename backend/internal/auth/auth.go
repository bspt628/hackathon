package auth

import (
	"context"
	"fmt"
	"firebase.google.com/go"
	"firebase.google.com/go/auth"
	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
	"log"
)

var FirebaseAuthClient *auth.Client
var StorageClient *storage.Client
var StorageBucket *storage.BucketHandle
var BucketName = "term6-hiroto-uchida.firebasestorage.app" // バケット名を変数で保持

// Firebase Admin SDKの初期化
func InitFirebase() error {
	log.Println("Firebase initializing")

	// Firebaseの設定
	config := &firebase.Config{
		StorageBucket: "gs://" + BucketName, // バケット名
	}
	opt := option.WithCredentialsFile("FirebaseAdminSDK.json") // サービスアカウントのJSONキー

	// Firebaseアプリを初期化
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		return fmt.Errorf("error initializing app: %v", err)
	}

	// Firebase認証クライアントを取得
	authClient, err := app.Auth(context.Background())
	if err != nil {
		return fmt.Errorf("error getting Auth client: %v", err)
	}

	// Google Cloud Storageクライアントを初期化
	StorageClient, err = storage.NewClient(context.Background(), opt)
	if err != nil {
		return fmt.Errorf("error initializing storage client: %v", err)
	}

	// デフォルトのバケットを取得
	StorageBucket = StorageClient.Bucket(BucketName)

	// 認証クライアントをグローバル変数として保持
	FirebaseAuthClient = authClient
	log.Println("Firebase initialized")
	return nil
}

// StorageBucketのバケット名を返すメソッドを追加
func GetBucketName() string {
	return BucketName
}


// Firebaseにユーザーを作成
func CreateFirebaseUser(email, password, username, displayName string) (string, error) {
	params := (&auth.UserToCreate{}).
		Email(email).
		Password(password).
		DisplayName(displayName).
		Disabled(false)

	userRecord, err := FirebaseAuthClient.CreateUser(context.Background(), params)
	if err != nil {
		return "", fmt.Errorf("error creating user: %v", err)
	}

	return userRecord.UID, nil
	// userRecordからuidを取り出す
}

func DeleteFirebaseUser(uid string) error {
	err := FirebaseAuthClient.DeleteUser(context.Background(), uid)
	if err != nil {
		return fmt.Errorf("failed to delete Firebase user: %v", err)
	}
	return nil
}

// Firebase認証を使ってユーザー情報を取得
func GetUserInfo(uid string) (*auth.UserRecord, error) {
	userRecord, err := FirebaseAuthClient.GetUser(context.Background(), uid)
	if err != nil {
		return nil, fmt.Errorf("error getting user: %v", err)
	}
	return userRecord, nil
}

// FirebaseのIDトークンを検証し、UIDを取得する
func VerifyIDToken(idToken string) (*auth.Token, error) {
	token, err := FirebaseAuthClient.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return nil, fmt.Errorf("error verifying ID token: %v", err)
	}
	return token, nil
}
