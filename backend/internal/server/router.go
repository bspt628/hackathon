package server

import (
	"database/sql"
	"hackathon/internal/auth"
	"hackathon/internal/controller/user"
	"hackathon/internal/controller/follow"
	"hackathon/internal/controller/post"
	"hackathon/internal/controller/like"
	"github.com/gorilla/mux"
)

func NewRouter(dbConn *sql.DB) *mux.Router {
	router := mux.NewRouter()

	// コントローラーの初期化
	userController := usercontroller.NewUserController(dbConn)
	postController := postcontroller.NewPostController(dbConn)
	likeController := likecontroller.NewLikeController(dbConn)
	followController := followcontroller.NewFollowController(dbConn)
	// notificationController := controller.NewNotificationController(dbConn)
	// messageController := controller.NewMessageController(dbConn)

	// 認証ミドルウェアを適用するルートグループ
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(auth.FirebaseAuthMiddleware)

	// ユーザー関連
	router.HandleFunc("/users/signup", userController.CreateUser).Methods("POST")
	router.HandleFunc("/users/signin", userController.SignIn).Methods("POST")

	apiRouter.HandleFunc("/users/get/{id}", userController.GetUser).Methods("GET")
	apiRouter.HandleFunc("/users/delete", userController.DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/users/email/{username}", userController.GetUserEmailByUsername).Methods("GET")
	
	apiRouter.HandleFunc("/users/password-reset/request", userController.PasswordResetRequest).Methods("POST") 
	apiRouter.HandleFunc("/users/password-reset/reset", userController.ResetPassword).Methods("POST")  

	// apiRouter.HandleFunc("/users/profile/{id}", userController.GetUserProfile).Methods("GET")
	// apiRouter.HandleFunc("/users/settings/{id}", userController.GetUserSettings).Methods("GET")
	// apiRouter.HandleFunc("/users/notification-settings/{id}", userController.GetUserNotifications).Methods("GET")
	// apiRouter.HandleFunc("/users/privacy/{id}", userController.GetUserPrivacy).Methods("GET")
	// apiRouter.HandleFunc("/users/ban-status/{id}", userController.GetUserBanStatus).Methods("GET")
	// apiRouter.HandleFunc("/users/username/{id}", userController.GetUserName).Methods("GET")
	// apiRouter.HandleFunc("/users/email/{id}", userController.GetUserEmail).Methods("GET")

	apiRouter.HandleFunc("/users/profile", userController.UpdateUserProfile).Methods("PUT")
	apiRouter.HandleFunc("/users/settings", userController.UpdateUserSettings).Methods("PUT")
	apiRouter.HandleFunc("/users/notification-settings", userController.UpdateUserNotifications).Methods("PUT")
	apiRouter.HandleFunc("/users/privacy", userController.UpdateUserPrivacy).Methods("PUT")
	apiRouter.HandleFunc("/users/ban-status", userController.UpdateUserBanStatus).Methods("PUT")
	apiRouter.HandleFunc("/users/username", userController.UpdateUserName).Methods("PUT")
	apiRouter.HandleFunc("/users/email", userController.UpdateUserEmail).Methods("PUT")

	// 投稿関連
	apiRouter.HandleFunc("/posts", postController.CreatePost).Methods("POST")
	// apiRouter.HandleFunc("/posts/{id}", postController.GetPost).Methods("GET")
	apiRouter.HandleFunc("/posts/{id}", postController.DeletePost).Methods("DELETE")
	apiRouter.HandleFunc("/posts/{id}/restore", postController.RestorePost).Methods("PUT")
	// router.HandleFunc("/api/posts/recent", postController.GetRecentPosts).Methods("GET")
	// router.HandleFunc("/api/posts/search", postController.SearchPostsByHashtag).Methods("GET")
	// router.HandleFunc("/api/timeline", postController.GetUserTimeline).Methods("GET")

	// // いいね機能
	apiRouter.HandleFunc("/likes", likeController.CreateLike).Methods("POST")
	apiRouter.HandleFunc("/likes", likeController.DeleteLike).Methods("DELETE")
	// router.HandleFunc("/api/likes/{id}/count", likeController.UpdatePostLikesCount).Methods("PUT")

	// // フォロー機能
	apiRouter.HandleFunc("/follow/{id}", followController.AddFollow).Methods("POST")
	apiRouter.HandleFunc("/follow/{id}", followController.RemoveFollow).Methods("DELETE")
	apiRouter.HandleFunc("/follow/status/{id}", followController.GetFollowStatus).Methods("GET")
	apiRouter.HandleFunc("/follow/followers/count", followController.UpdateAndGetFollowersCount).Methods("GET")
	apiRouter.HandleFunc("/follow/followings/count", followController.UpdateAndGetFollowingsCount).Methods("GET")

	apiRouter.HandleFunc("/follow/followers/all", followController.GetFollowers).Methods("GET")
	apiRouter.HandleFunc("/follow/followings/all", followController.GetFollowings).Methods("GET")
	apiRouter.HandleFunc("/follow/ff", followController.GetFollowersAndFollowings).Methods("GET")

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
