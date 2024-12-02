package postusecase

import (
	"context"
	"database/sql"
	"fmt"
	"hackathon/db/sqlc/generated"
	"hackathon/domain"
	"encoding/json"
	"github.com/google/uuid"
)

func (uc *PostUsecase) CreatePost(ctx context.Context, request domain.CreatePostRequest) (*domain.Post, error) {
	// 必須フィールドのバリデーション
	if request.UserID == "" || request.Content == "" {
		return nil, fmt.Errorf("ユーザーIDとコンテンツは必須です")
	}

	// MediaURLs のエンコード
	mediaUrlsJSON, err := json.Marshal(request.MediaURLs)
	if err != nil {
		return nil, fmt.Errorf("メディアURLのエンコードに失敗しました: %v", err)
	}

	var originalPostID sql.NullString
	var replyToID sql.NullString
	var rootPostID sql.NullString

	// リポストの検証
	if request.IsRepost {
		if request.OriginalPostID == nil || *request.OriginalPostID == "" {
			return nil, fmt.Errorf("original_post_id が空です")
		}

		// original_post_id が存在するか確認
		exists, err := uc.dao.CheckPostExists(ctx, *request.OriginalPostID)
		if err != nil {
			return nil, fmt.Errorf("投稿の存在確認に失敗しました: %v", err)
		}
		if !exists {
			return nil, fmt.Errorf("指定された original_post_id は存在しません")
		}

		originalPostID = sql.NullString{String: *request.OriginalPostID, Valid: true}
	} else {
		originalPostID = sql.NullString{String: "", Valid: false}
	}

	// リプライの検証
	if request.IsReply {
		if request.ReplyToID == nil || *request.ReplyToID == "" {
			return nil, fmt.Errorf("reply_to_id が空です")
		}

		// reply_to_id が存在するか確認
		exists, err := uc.dao.CheckPostExists(ctx, *request.ReplyToID)
		if err != nil {
			return nil, fmt.Errorf("返信元投稿の存在確認に失敗しました: %v", err)
		}
		if !exists {
			return nil, fmt.Errorf("指定された reply_to_id は存在しません")
		}

		replyToID = sql.NullString{String: *request.ReplyToID, Valid: true}

		// root_post_id の検証
		if request.RootPostID == nil || *request.RootPostID == "" {
			return nil, fmt.Errorf("root_post_id が空です")
		}

		validRoot, err := uc.dao.CheckRootPostValidity(ctx, *request.RootPostID)
		if err != nil {
			return nil, fmt.Errorf("root_post_id の妥当性確認に失敗しました: %v", err)
		}
		if !validRoot {
			return nil, fmt.Errorf("指定された root_post_id は無効です（既に他の投稿に紐づけられています）")
		}

		rootPostID = sql.NullString{String: *request.RootPostID, Valid: true}
	} else {
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
		OriginalPostID: originalPostID,
		ReplyToID:      replyToID,
		RootPostID:     rootPostID,
		IsRepost:       sql.NullBool{Bool: request.IsRepost, Valid: true},
		IsReply:        sql.NullBool{Bool: request.IsReply, Valid: true},
	}

	err = uc.dao.CreatePost(ctx, arg)
	if err != nil {
		return nil, fmt.Errorf("投稿の作成に失敗しました: %v", err)
	}

	// 返信元投稿の reply_count をインクリメント
	if request.IsReply {
		err = uc.dao.IncrementReplyCount(ctx, *request.ReplyToID)
		if err != nil {
			return nil, fmt.Errorf("返信元投稿の返信数更新に失敗しました: %v", err)
		}
	}

	return &domain.Post{
		ID:         postID,
		UserID:     request.UserID,
		Content:    request.Content,
		MediaURLs:  request.MediaURLs,
		Visibility: request.Visibility,
	}, nil
}