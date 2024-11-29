package jwtunused
import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("your-secret-key") // 必ず環境変数などで管理すること

// クレーム（JWTのペイロード）
type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

// トークンを生成する関数
func GenerateToken(userID string, expiryDuration time.Duration) (string, error) {
	expirationTime := time.Now().Add(expiryDuration)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// トークンを生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// トークンを検証する関数
func ValidateToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("トークンが期限切れです")
		}
		return nil, errors.New("トークンが無効です")
	}

	if !token.Valid {
		return nil, errors.New("トークンが無効です")
	}

	return claims, nil
}
