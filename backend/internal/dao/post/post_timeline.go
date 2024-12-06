package postdao

import (
	"context"
	"database/sql"
	"hackathon/db/sqlc/generated"
	"hackathon/internal/model"
	"hackathon/internal/utils"
)


func (dao *PostDAO) GetAllPosts(ctx context.Context, limit int32) ([]model.PostAll, error) {
    rows, err := dao.queries.GetAllPosts(ctx, limit)
	if err != nil {
		return nil, err
	}

	var posts []model.PostAll
	for _, row := range rows {
		posts = append(posts, convertGetAllPostsRowToPostAll(row))
	}

	return posts, nil
}

func (dao *PostDAO) GetFollowingUsersPosts(ctx context.Context, userID string, limit int32)([]model.PostAll, error) {
    rows, err := dao.queries.GetFollowingUsersPosts(ctx, sqlc.GetFollowingUsersPostsParams{
        FollowerID: sql.NullString{String: userID, Valid: true},
        Limit:      limit,
    })
	if err != nil {
		return nil, err
	}

	var posts []model.PostAll
	for _, row := range rows {
		posts = append(posts, convertGetFollowedUsersPostsRowToPostAll(row))
	}

	return posts, nil
}

func (dao *PostDAO) GetPostByID(ctx context.Context, postID string) (model.PostAll, error) {
	row, err := dao.queries.GetPost(ctx, postID)
	if err != nil {
		return model.PostAll{}, err
	}

	return convertGetPostByIDRowToPostAll(row), nil
}

func convertGetFollowedUsersPostsRowToPostAll(row sqlc.GetFollowingUsersPostsRow) model.PostAll {
	return model.PostAll{
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


func convertGetAllPostsRowToPostAll(row sqlc.GetAllPostsRow) model.PostAll {
	return model.PostAll{
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

func convertGetPostByIDRowToPostAll(row sqlc.GetPostRow) model.PostAll {
	return model.PostAll{
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