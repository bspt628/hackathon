// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
	"database/sql"
)

type Querier interface {
	AddBlock(ctx context.Context, arg AddBlockParams) error
	AddFollow(ctx context.Context, arg AddFollowParams) error
	AddLike(ctx context.Context, arg AddLikeParams) error
	CreateNotification(ctx context.Context, arg CreateNotificationParams) error
	CreatePost(ctx context.Context, arg CreatePostParams) (sql.Result, error)
	CreateRepost(ctx context.Context, arg CreateRepostParams) error
	CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error)
	GetDMConversation(ctx context.Context, arg GetDMConversationParams) ([]Dm, error)
	GetRecentPosts(ctx context.Context, limit int32) ([]GetRecentPostsRow, error)
	GetUnreadNotifications(ctx context.Context, userid sql.NullString) ([]Notification, error)
	GetUserStats(ctx context.Context, id string) (GetUserStatsRow, error)
	GetUserTimeline(ctx context.Context, arg GetUserTimelineParams) ([]GetUserTimelineRow, error)
	SearchPostsByHashtag(ctx context.Context, arg SearchPostsByHashtagParams) ([]SearchPostsByHashtagRow, error)
	SendDM(ctx context.Context, arg SendDMParams) error
	UpdateFollowersCount(ctx context.Context, id string) error
	UpdatePostLikesCount(ctx context.Context, id string) error
	UpdateUserInfo(ctx context.Context, arg UpdateUserInfoParams) error
}

var _ Querier = (*Queries)(nil)
