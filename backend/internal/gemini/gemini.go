package gemini

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"cloud.google.com/go/vertexai/genai"
)

const (
	location  = "asia-northeast1"
	modelName = "gemini-1.5-flash-002"
	projectID = "term6-hiroto-uchida" // プロジェクトID
)

// GenerateContentHandler handles HTTP requests and generates content from a prompt.
func GenerateContentHandler(w http.ResponseWriter, r *http.Request) {
	// POSTメソッドのみ許可
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// リクエストボディからJSONデータを読み取る
	var request struct {
		Prompt string `json:"prompt"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// プロンプトが空かチェック
	if request.Prompt == "" {
		http.Error(w, "Prompt is required", http.StatusBadRequest)
		return
	}

	// Geminiにプロンプトを送信
	response, err := generateContentFromText(projectID, request.Prompt)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error generating content: %v", err), http.StatusInternalServerError)
		return
	}

	// レスポンスをJSONとして返す
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)
}

func generateContentFromText(projectID, promptText string) ([]byte, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, projectID, location)
	if err != nil {
		return nil, fmt.Errorf("error creating client: %w", err)
	}

	// Geminiにプロンプトを送信
	gemini := client.GenerativeModel(modelName)
	prompt := genai.Text(promptText)
	resp, err := gemini.GenerateContent(ctx, prompt)
	if err != nil {
		return nil, fmt.Errorf("error generating content: %w", err)
	}

	// 結果をJSONにシリアライズ
	return json.MarshalIndent(resp, "", "  ")
}
