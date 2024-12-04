package postdao

import (
	"context"
	"database/sql"
	"hackathon/db/sqlc/generated"
	"hackathon/domain"
	"hackathon/internal/utils"
)


func (dao *PostDAO) GetAllPosts(ctx context.Context, limit int32) ([]domain.PostAll, error) {
    rows, err := dao.queries.GetAllPosts(ctx, limit)
	if err != nil {
		return nil, err
	}

	var posts []domain.PostAll
	for _, row := range rows {
		posts = append(posts, convertGetAllPostsRowToPostAll(row))
	}

	return posts, nil
}

func (dao *PostDAO) GetFollowingUsersPosts(ctx context.Context, userID string, limit int32)([]domain.PostAll, error) {
    rows, err := dao.queries.GetFollowingUsersPosts(ctx, sqlc.GetFollowingUsersPostsParams{
        FollowerID: sql.NullString{String: userID, Valid: true},
        Limit:      limit,
    })
	if err != nil {
		return nil, err
	}

	var posts []domain.PostAll
	for _, row := range rows {
		posts = append(posts, convertGetFollowedUsersPostsRowToPostAll(row))
	}

	return posts, nil
}

func convertGetFollowedUsersPostsRowToPostAll(row sqlc.GetFollowingUsersPostsRow) domain.PostAll {
	return domain.PostAll{
		ID:             row.ID,
		UserID:         utils.ConvertNullString(row.UserID),
		Content:        utils.ConvertNullString(row.Content),
		CreatedAt:      utils.ConvertNullTime(row.CreatedAt),
		UpdatedAt:      utils.ConvertNullTime(row.UpdatedAt),
		IsRepost:       utils.ConvertNullBool(row.IsRepost),
		OriginalPostID: utils.ConvertNullString(row.OriginalPostID),
		ReplyToID:      utils.ConvertNullString(row.ReplyToID),
		RootPostID:     utils.ConvertNullString(row.RootPostID),
		IsReply:        utils.ConvertNullBool(row.IsReply),
		MediaUrls:      row.MediaUrls, // JSONデータなのでそのまま渡す
		LikesCount:     utils.ConvertNullInt32(row.LikesCount),
		RepostsCount:   utils.ConvertNullInt32(row.RepostsCount),
		RepliesCount:   utils.ConvertNullInt32(row.RepliesCount),
		ViewsCount:     utils.ConvertNullInt32(row.ViewsCount),
		Visibility:     utils.ConvertNullString(row.Visibility),
		IsPinned:       utils.ConvertNullBool(row.IsPinned),
		IsDeleted:      utils.ConvertNullBool(row.IsDeleted),
		Username:       row.Username, // 非NULLフィールド
		DisplayName:    utils.ConvertNullString(row.DisplayName),
	}
}


func convertGetAllPostsRowToPostAll(row sqlc.GetAllPostsRow) domain.PostAll {
	return domain.PostAll{
		ID:             row.ID,
		UserID:         utils.ConvertNullString(row.UserID),
		Content:        utils.ConvertNullString(row.Content),
		CreatedAt:      utils.ConvertNullTime(row.CreatedAt),
		UpdatedAt:      utils.ConvertNullTime(row.UpdatedAt),
		IsRepost:       utils.ConvertNullBool(row.IsRepost),
		OriginalPostID: utils.ConvertNullString(row.OriginalPostID),
		ReplyToID:      utils.ConvertNullString(row.ReplyToID),
		RootPostID:     utils.ConvertNullString(row.RootPostID),
		IsReply:        utils.ConvertNullBool(row.IsReply),
		MediaUrls:      row.MediaUrls, // JSONデータなのでそのまま渡す
		LikesCount:     utils.ConvertNullInt32(row.LikesCount),
		RepostsCount:   utils.ConvertNullInt32(row.RepostsCount),
		RepliesCount:   utils.ConvertNullInt32(row.RepliesCount),
		ViewsCount:     utils.ConvertNullInt32(row.ViewsCount),
		Visibility:     utils.ConvertNullString(row.Visibility),
		IsPinned:       utils.ConvertNullBool(row.IsPinned),
		IsDeleted:      utils.ConvertNullBool(row.IsDeleted),
		Username:       row.Username, // 非NULLフィールド
		DisplayName:    utils.ConvertNullString(row.DisplayName),
	}
}
