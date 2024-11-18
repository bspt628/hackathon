package server

import (
	"database/sql"
	"hackathon/internal/controller/user"
	"hackathon/internal/controller/auth"

	"hackathon/internal/auth"
	"github.com/gorilla/mux"
)

func NewRouter(dbConn *sql.DB) *mux.Router {
	router := mux.NewRouter()

	// コントローラーの初期化
	userController := controller.NewUserController(dbConn)
	// postController := controller.NewPostController(dbConn)
	// likeController := controller.NewLikeController(dbConn)
	// followController := controller.NewFollowController(dbConn)
	// notificationController := controller.NewNotificationController(dbConn)
	// messageController := controller.NewMessageController(dbConn)

	// 認証ミドルウェアを適用するルートグループ
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(auth.AuthMiddleware)

	// トークン生成エンドポイントを追加
	authController := authController.NewAuthController(dbConn)
	router.HandleFunc("/auth/signin", authController.SignIn).Methods("POST")

	// ユーザー関連
	router.HandleFunc("/api/users", userController.CreateUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", userController.GetUser).Methods("GET")
	apiRouter.HandleFunc("/users/{id}", userController.DeleteUser).Methods("DELETE")
	// router.HandleFunc("/api/users/{id}", userController.UpdateUserInfo).Methods("PUT")
	// router.HandleFunc("/api/users/{id}/stats", userController.GetUserStats).Methods("GET")

	// // 投稿関連
	// router.HandleFunc("/api/posts", postController.CreatePost).Methods("POST")
	// router.HandleFunc("/api/posts/recent", postController.GetRecentPosts).Methods("GET")
	// router.HandleFunc("/api/posts/search", postController.SearchPostsByHashtag).Methods("GET")

	// // タイムライン
	// router.HandleFunc("/api/timeline", postController.GetUserTimeline).Methods("GET")

	// // いいね機能
	// router.HandleFunc("/api/posts/{id}/likes", likeController.AddLike).Methods("POST")
	// router.HandleFunc("/api/posts/{id}/likes/count", likeController.UpdatePostLikesCount).Methods("PUT")

	// // フォロー機能
	// router.HandleFunc("/api/users/{id}/follow", followController.AddFollow).Methods("POST")
	// router.HandleFunc("/api/users/{id}/followers/count", followController.UpdateFollowersCount).Methods("PUT")

	// // リポスト機能
	// router.HandleFunc("/api/posts/{id}/repost", postController.CreateRepost).Methods("POST")

	// // ブロック機能
	// router.HandleFunc("/api/users/{id}/block", userController.AddBlock).Methods("POST")

	// // ダイレクトメッセージ
	// router.HandleFunc("/api/messages", messageController.SendDM).Methods("POST")
	// router.HandleFunc("/api/messages/{userId}", messageController.GetDMConversation).Methods("GET")

	// // 通知機能
	// router.HandleFunc("/api/notifications", notificationController.CreateNotification).Methods("POST")
	// router.HandleFunc("/api/notifications/unread", notificationController.GetUnreadNotifications).Methods("GET")

	return router
}