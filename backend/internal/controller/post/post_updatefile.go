package postcontroller

import (
	"fmt"
	"net/http"
)

func (pc *PostController) UploadFile(w http.ResponseWriter, r *http.Request) {
    // Parse multipart form
    if err := r.ParseMultipartForm(10 << 20); err != nil {
        http.Error(w, fmt.Sprintf("ファイルのアップロードに失敗しました: %v", err), http.StatusBadRequest)
        return
    }

    file, handler, err := r.FormFile("file")
    if err != nil {
        http.Error(w, fmt.Sprintf("ファイルが見つかりません: %v", err), http.StatusBadRequest)
        return
    }
    defer file.Close()

    fileName, err := pc.postUsecase.UploadFile(r.Context(), file, handler.Filename)
    if err != nil {
        http.Error(w, fmt.Sprintf("ファイルのアップロードに失敗しました: %v", err), http.StatusInternalServerError)
        return
    }

    fileURL := fmt.Sprintf("https://storage.googleapis.com/%s/%s", "your-firebase-storage-bucket-name", fileName)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(fmt.Sprintf(`{"file_url":"%s"}`, fileURL)))
}
