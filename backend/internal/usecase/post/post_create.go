package postusecase

import (
	"context"
	"database/sql"
	"fmt"
	"hackathon/db/sqlc/generated"
	"hackathon/domain"
	"time"
	"encoding/json"
	"github.com/google/uuid"
)

func (uc *PostUsecase) CreatePost(ctx context.Context, request domain.CreatePostRequest) (*domain.Post, error) {
	// 必須フィールドのバリデーション
	if request.UserID == "" || request.Content == "" {
		return nil, fmt.Errorf("ユーザーIDとコンテンツは必須です")
	}

	mediaUrlsJSON, err := json.Marshal(request.MediaURLs)
	if err != nil {
		return nil, fmt.Errorf("メディアURLのエンコードに失敗しました: %v", err)
	}

	var originalPostID sql.NullString
    if request.IsRepost {
        originalPostID = sql.NullString{String: *request.OriginalPostID, Valid: true}
    } else {
        // リポストでない場合は、original_post_id を NULL に設定
        originalPostID = sql.NullString{String: "", Valid: false}
    }


	var replyToID sql.NullString
	var rootPostID sql.NullString
	if request.IsReply {
		replyToID = sql.NullString{String: *request.ReplyToID, Valid: true}
		rootPostID = sql.NullString{String: *request.RootPostID, Valid: true}
	} else {
		// リプライでない場合は、reply_to_id を NULL に設定
		replyToID = sql.NullString{String: "", Valid: false}
		rootPostID = sql.NullString{String: "", Valid: false}
	}
	// ID生成
	postID := uuid.New().String()

	// DAO層を呼び出し
	arg := sqlc.CreatePostParams{
		ID:             postID,
		UserID:         sql.NullString{String: request.UserID, Valid: true},
		Content:        sql.NullString{String: request.Content, Valid: true},
		MediaUrls:      json.RawMessage(mediaUrlsJSON),
		Visibility:     sql.NullString{String: request.Visibility, Valid: true},
		Language:       sql.NullString{String: request.Language, Valid: true},
		Location:       sql.NullString{String: request.Location, Valid: true},
		Device:         sql.NullString{String: request.Device, Valid: true},
		OriginalPostID: originalPostID,
		ReplyToID:      replyToID,
		RootPostID:     rootPostID,
		IsRepost:       sql.NullBool{Bool: request.IsRepost, Valid: true},
		IsReply:        sql.NullBool{Bool: request.IsReply, Valid: true},
		CreatedAt:      sql.NullTime{Time: time.Now(), Valid: true},
	}

	err = uc.dao.CreatePost(ctx, arg)
	if err != nil {
		return nil, fmt.Errorf("投稿の作成に失敗しました: %v", err)
	}

	return &domain.Post{
		ID:         postID,
		UserID:     request.UserID,
		Content:    request.Content,
		MediaURLs:  request.MediaURLs,
		Visibility: request.Visibility,
		Language:   request.Language,
		Location:   request.Location,
		Device:     request.Device,
	}, nil
}