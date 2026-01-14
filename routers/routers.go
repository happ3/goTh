package routers

import (
	"goTh/app/controllers"
	"goTh/app/middlewares"

	"github.com/gin-gonic/gin"
)

func RouterInit(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("/login", controllers.UserController{}.Login)
		userGroup.POST("/logout", controllers.UserController{}.Logout)
		userGroup.POST("/refreshToken", controllers.UserController{}.RefreshToken)
		userGroup.POST("/registerUser", controllers.UserController{}.RegisterUser)
		userGroup.POST("/authentication", middlewares.AuthMiddleware, controllers.UserController{}.Authentication)
		userGroup.GET("/userInfo", middlewares.AuthMiddleware, controllers.UserController{}.GetUserInfo)
	}

	commentsGroup := r.Group("/comments", middlewares.AuthMiddleware)
	{
		commentsGroup.POST("/addComments", controllers.CommentsController{}.AddComments)
		commentsGroup.POST("/findCommentByPostId", controllers.CommentsController{}.FindCommentByPostId)
	}

	postsGroup := r.Group("/posts", middlewares.AuthMiddleware)
	{
		postsGroup.POST("/pagePost", controllers.PostsController{}.PagePost)
		postsGroup.POST("/addPosts", controllers.PostsController{}.AddPosts)
		postsGroup.POST("/getPosts", controllers.PostsController{}.GetPosts)
		postsGroup.POST("/updatePost", controllers.PostsController{}.UpdatePost)
		postsGroup.POST("/delPost", controllers.PostsController{}.DelPost)
	}
}
