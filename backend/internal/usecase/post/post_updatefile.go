package postusecase

import (
	"context"
	"fmt"
	"io"
	"path/filepath"
	"time"
)

func (uc *PostUsecase) UploadFile(ctx context.Context, file io.Reader, originalFilename string) (string, error) {
    // Generate a unique file name
    extension := filepath.Ext(originalFilename)
    fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), extension)

    // Delegate to DAO
    if err := uc.dao.UploadToFirebase(ctx, file, fileName); err != nil {
        return "", fmt.Errorf("ファイルのアップロードに失敗しました: %v", err)
    }

    return fileName, nil
}
