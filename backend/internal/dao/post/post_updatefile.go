package postdao

import (
	"context"
	"fmt"
	"io"
	"cloud.google.com/go/storage"
)

func (dao *PostDAO) UploadToFirebase(ctx context.Context, file io.Reader, fileName string) error {
    // Initialize Firebase Storage client
    client, err := storage.NewClient(ctx)
    if err != nil {
        return fmt.Errorf("firebase Storageの初期化に失敗しました: %v", err)
    }
    defer client.Close()

    bucketName := "gs://term6-hiroto-uchida.firebasestorage.app"
    bucket := client.Bucket(bucketName)
    object := bucket.Object(fileName)

    writer := object.NewWriter(ctx)
    if _, err := io.Copy(writer, file); err != nil {
        return fmt.Errorf("ファイルのアップロードに失敗しました: %v", err)
    }
    if err := writer.Close(); err != nil {
        return fmt.Errorf("アップロード中にエラーが発生しました: %v", err)
    }

    // Make file publicly accessible (optional)
    if err := object.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
        return fmt.Errorf("ファイルの公開設定に失敗しました: %v", err)
    }

    return nil
}
