package postusecase

import (
	"context"
	"fmt"
	"hackathon/internal/auth"
	"mime/multipart"
	"net/http"
)

// UploadFile はファイルのアップロードを行います。
func (pu *PostUsecase) UploadFile(ctx context.Context, r *http.Request) error {
	// Firebaseの初期化確認
	if auth.StorageBucket == nil {
		return fmt.Errorf("firebase is not initialized")
	}

	// リクエストからファイルを取得
	err := r.ParseMultipartForm(10 << 20) // 最大10MBのファイルを許可
	if err != nil {
		return fmt.Errorf("ファイルのアップロードに失敗しました: %v", err)
	}

	// "file" フィールドからファイルを取得
	file, handler, err := r.FormFile("file")
	if err != nil {
		return fmt.Errorf("ファイルが見つかりません: %v", err)
	}
	defer file.Close()

	// バリデーション: ファイルタイプとサイズの検証
	if err := validateFile(handler); err != nil {
		return err
	}

	// ファイルをDAO層にアップロード
	remoteFilename := fmt.Sprintf("uploads/%s", handler.Filename) // バケット内のパス
	contentType := handler.Header.Get("Content-Type")

	return pu.dao.UploadFile(ctx, file, remoteFilename, contentType)
}

// validateFile はアップロードされるファイルのバリデーションを行います。
func validateFile(handler *multipart.FileHeader) error {
	// 拡張子チェックなど、ファイルのバリデーションを実施
	if handler.Header.Get("Content-Type") != "image/jpeg" && handler.Header.Get("Content-Type") != "image/png" && handler.Header.Get("Content-Type") != "application/pdf" {
		return fmt.Errorf("サポートされていないファイル形式です")
	}

	// サイズチェック（例: 最大10MB）
	if handler.Size > 10<<20 {
		return fmt.Errorf("ファイルサイズが大きすぎます")
	}

	return nil
}
