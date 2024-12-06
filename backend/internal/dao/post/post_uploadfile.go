package postdao

import (
	"context"
	"fmt"
	"hackathon/internal/auth"
	"cloud.google.com/go/storage"
	"io"
)

// UploadFile はFirebase Storageにファイルをアップロードします。
func (dao *PostDAO) UploadFile(ctx context.Context, file io.Reader, remoteFilename, contentType string) error {
	// Firebase Storageバケットにファイルをアップロード
	writer := auth.StorageBucket.Object(remoteFilename).NewWriter(ctx)
	writer.ObjectAttrs.ContentType = contentType
	writer.ObjectAttrs.ACL = []storage.ACLRule{
		{
			Entity: storage.AllUsers,
			Role:   storage.RoleReader,
		},
	}

	// ファイルをアップロード
	if _, err := io.Copy(writer, file); err != nil {
		return fmt.Errorf("ファイルのアップロードに失敗しました: %v", err)
	}

	// アップロード完了後、Writerを閉じる
	if err := writer.Close(); err != nil {
		return fmt.Errorf("ファイルのアップロードに失敗しました: %v", err)
	}

	return nil
}
