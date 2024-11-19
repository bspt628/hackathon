package authController

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
)

// パスワードリセットリクエスト (トークン送信)
func (prc *PasswordResetController) HandlePasswordResetRequest(c *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required,email"`
	}

	// リクエスト解析
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なリクエスト"})
		return
	}

	// パスワードリセット処理
	err := prc.passwordResetUsecase.RequestPasswordReset(c.Request.Context(), req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("パスワードリセットに失敗しました: %v", err)})
		return
	}

	// 成功レスポンス
	c.JSON(http.StatusOK, gin.H{"message": "パスワードリセットメールを送信しました"})
}

// パスワード変更 (トークン検証とパスワード更新)
func (prc *PasswordResetController) ResetPassword(c *gin.Context) {
	var request struct {
		Token    string `json:"token" binding:"required"`
		Password string `json:"password" binding:"required,min=8"`
	}

	// リクエスト解析
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("リクエスト解析エラー: %v", err)})
		return
	}

	// パスワード更新処理
	err := prc.passwordResetUsecase.ResetPassword(c.Request.Context(), request.Token, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("パスワード更新失敗: %v", err)})
		return
	}

	// 成功レスポンス
	c.JSON(http.StatusOK, gin.H{"message": "パスワードが更新されました"})
}
