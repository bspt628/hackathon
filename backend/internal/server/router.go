package server

import (
	"database/sql"
	"hackathon/internal/auth"
	"hackathon/internal/controller/user"
	"hackathon/internal/controller/follow"
	"hackathon/internal/controller/post"
	"github.com/gorilla/mux"
)

func NewRouter(dbConn *sql.DB) *mux.Router {
	router := mux.NewRouter()

	// コントローラーの初期化
	userController := usercontroller.NewUserController(dbConn)
	postController := postcontroller.NewPostController(dbConn)
	// likeController := controller.NewLikeController(dbConn)
	followController := followcontroller.NewFollowController(dbConn)
	// notificationController := controller.NewNotificationController(dbConn)
	// messageController := controller.NewMessageController(dbConn)

	// 認証ミドルウェアを適用するルートグループ
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(auth.FirebaseAuthMiddleware)

	// ユーザー関連
	router.HandleFunc("/users/signup", userController.CreateUser).Methods("POST")
	router.HandleFunc("/users/signin", userController.SignIn).Methods("POST")

	apiRouter.HandleFunc("/users/{id}", userController.GetUser).Methods("GET")
	apiRouter.HandleFunc("/users/{id}", userController.DeleteUser).Methods("DELETE")

	apiRouter.HandleFunc("/users/{id}/profile", userController.UpdateUserProfile).Methods("PUT")
	apiRouter.HandleFunc("/users/{id}/settings", userController.UpdateUserSettings).Methods("PUT")
	apiRouter.HandleFunc("/users/{id}/notification-settings", userController.UpdateUserNotifications).Methods("PUT")
	apiRouter.HandleFunc("/users/{id}/privacy", userController.UpdateUserPrivacy).Methods("PUT")
	apiRouter.HandleFunc("/users/{id}/ban-status", userController.UpdateUserBanStatus).Methods("PUT")
	apiRouter.HandleFunc("/users/{id}/username", userController.UpdateUserName).Methods("PUT")
	apiRouter.HandleFunc("/users/{id}/email", userController.UpdateUserEmail).Methods("PUT")
	router.HandleFunc("/api/users/email/{username}", userController.GetUserEmailByUsername).Methods("GET")
	apiRouter.HandleFunc("/users/password-reset/request", userController.HandlePasswordResetRequest).Methods("POST") 
	apiRouter.HandleFunc("/users/password-reset/reset", userController.ResetPassword).Methods("POST")                

	// 投稿関連
	apiRouter.HandleFunc("/posts", postController.CreatePost).Methods("POST")
	// router.HandleFunc("/api/posts/recent", postController.GetRecentPosts).Methods("GET")
	// router.HandleFunc("/api/posts/search", postController.SearchPostsByHashtag).Methods("GET")
	// router.HandleFunc("/api/timeline", postController.GetUserTimeline).Methods("GET")

	// // いいね機能
	// router.HandleFunc("/api/posts/{id}/likes", likeController.AddLike).Methods("POST")
	// router.HandleFunc("/api/posts/{id}/likes/count", likeController.UpdatePostLikesCount).Methods("PUT")

	// // フォロー機能
	apiRouter.HandleFunc("/users/{id}/follow", followController.AddFollow).Methods("POST")
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
