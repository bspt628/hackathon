package controller

import (
    "net/http"
    "github.com/gin-gonic/gin" // Ginフレームワークを使用する場合の例です。
    "hackathon/internal/usecase"
)

type UserController struct {
    userUsecase *usecase.UserUsecase // Usecaseへの依存性注入
}

func NewUserController(userUsecase *usecase.UserUsecase) *UserController {
    return &UserController{userUsecase: userUsecase}
}

func (u *UserController) GetUser(c *gin.Context) {
    id := c.Param("id")
    user, err := u.userUsecase.GetUserByID(c.Request.Context(), id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.JSON(http.StatusOK, user)
}