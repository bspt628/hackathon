package server

import (
	"database/sql"
	"hackathon/internal/auth"
	followcontroller "hackathon/internal/controller/follow"
	likecontroller "hackathon/internal/controller/like"
	postcontroller "hackathon/internal/controller/post"
	usercontroller "hackathon/internal/controller/user"
	repostcontroller "hackathon/internal/controller/repost"
	notificationcontroller "hackathon/internal/controller/notification"
	"hackathon/internal/gemini"

	"github.com/gorilla/mux"
)

func NewRouter(dbConn *sql.DB) *mux.Router {
	router := mux.NewRouter()
	// router.Use(auth.CORS)

	// コントローラーの初期化
	userController := usercontroller.NewUserController(dbConn)
	postController := postcontroller.NewPostController(dbConn)
	likeController := likecontroller.NewLikeController(dbConn)
	followController := followcontroller.NewFollowController(dbConn)
	repostController := repostcontroller.NewRepostController(dbConn)
	notificationController := notificationcontroller.NewNotificationController(dbConn)
	// messageController := controller.NewMessageController(dbConn)

	// 認証ミドルウェアを適用するルートグループ
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(auth.FirebaseAuthMiddleware)

	// ユーザー関連
	router.HandleFunc("/users/signup", userController.CreateUser).Methods("POST")
	router.HandleFunc("/users/signin", userController.SignIn).Methods("POST")

	apiRouter.HandleFunc("/users/user/{id}", userController.GetUser).Methods("GET")
	apiRouter.HandleFunc("/users/delete", userController.DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/users/email/{username}", userController.GetUserEmailByUsername).Methods("GET")
	apiRouter.HandleFunc("/users/firebase", userController.GetUserIDByFirebaseUID).Methods("GET")
	
	apiRouter.HandleFunc("/users/password-reset/request", userController.PasswordResetRequest).Methods("POST") 
	apiRouter.HandleFunc("/users/password-reset/reset", userController.ResetPassword).Methods("POST")  

	apiRouter.HandleFunc("/generate-content", gemini.GenerateContentHandler).Methods("POST")


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
	apiRouter.HandleFunc("/users/username", userController.UpdateUserUsername).Methods("PUT")
	apiRouter.HandleFunc("/users/email", userController.UpdateUserEmail).Methods("PUT")
	

	// 投稿関連
	apiRouter.HandleFunc("/posts", postController.CreatePost).Methods("POST")
	// apiRouter.HandleFunc("/posts/{id}", postController.GetPost).Methods("GET")
	apiRouter.HandleFunc("/posts/{id}", postController.DeletePost).Methods("DELETE")
	apiRouter.HandleFunc("/posts/{id}/restore", postController.RestorePost).Methods("PUT")
	apiRouter.HandleFunc("/posts/timeline/all", postController.GetAllPosts).Methods("GET")
	apiRouter.HandleFunc("/posts/timeline/follow", postController.GetFollowingUsersPosts).Methods("GET")
	apiRouter.HandleFunc("/posts/timeline/one/{id}", postController.GetPostByID).Methods("GET")
	apiRouter.HandleFunc("/posts/upload", postController.UploadFile).Methods("POST")
	// router.HandleFunc("/api/posts/recent", postController.GetRecentPosts).Methods("GET")
	// router.HandleFunc("/api/posts/search", postController.SearchPostsByHashtag).Methods("GET")
	// router.HandleFunc("/api/timeline", postController.GetUserTimeline).Methods("GET")

	// // いいね機能
	apiRouter.HandleFunc("/likes/{id}", likeController.CreateLike).Methods("POST")
	apiRouter.HandleFunc("/likes/{id}", likeController.DeleteLike).Methods("DELETE")
	apiRouter.HandleFunc("/likes/{id}/status", likeController.GetLikeStatus).Methods("GET")
	apiRouter.HandleFunc("/likes/{id}/count", likeController.GetPostLikesCount).Methods("GET")

	// router.HandleFunc("/api/likes/{id}/count", likeController.UpdatePostLikesCount).Methods("PUT")

	// // フォロー機能
	apiRouter.HandleFunc("/follow/{id}", followController.AddFollow).Methods("POST")
	apiRouter.HandleFunc("/follow/{id}", followController.RemoveFollow).Methods("DELETE")
	apiRouter.HandleFunc("/follow/{id}/status", followController.GetFollowStatus).Methods("GET")
	apiRouter.HandleFunc("/follow/followers/count", followController.GetFollowersCount).Methods("GET")
	apiRouter.HandleFunc("/follow/followings/count", followController.GetFollowingsCount).Methods("GET")

	apiRouter.HandleFunc("/follow/followers/all", followController.GetFollowers).Methods("GET")
	apiRouter.HandleFunc("/follow/followings/all", followController.GetFollowings).Methods("GET")
	apiRouter.HandleFunc("/follow/ff", followController.GetFollowersAndFollowings).Methods("GET")

	// // リポスト機能
	apiRouter.HandleFunc("/repost", repostController.CreateRepost).Methods("POST")
	apiRouter.HandleFunc("/repost", repostController.DeleteRepost).Methods("DELETE")
	apiRouter.HandleFunc("/repost/{id}/status", repostController.GetRepostStatus).Methods("GET")

	// // ブロック機能
	// router.HandleFunc("/api/users/{id}/block", userController.AddBlock).Methods("POST")

	// // ダイレクトメッセージ
	// router.HandleFunc("/api/messages", messageController.SendDM).Methods("POST")
	// router.HandleFunc("/api/messages/{userId}", messageController.GetDMConversation).Methods("GET")

	// // 通知機能
	apiRouter.HandleFunc("/notifications", notificationController.CreateNotifications).Methods("POST")
	apiRouter.HandleFunc("/notifications/count/unread", notificationController.CountUnreadNotifications).Methods("GET")
	apiRouter.HandleFunc("/notifications/count/all", notificationController.CountAllNotifications).Methods("GET")
	apiRouter.HandleFunc("/notifications/{id}/read", notificationController.MarkNotificationsAsRead).Methods("PUT")

	return router
}
