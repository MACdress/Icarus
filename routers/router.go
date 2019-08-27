package routers

import "github.com/gin-gonic/gin"

func main()  {
	router:=gin.Default()
	user :=router.Group("/user")
	{
		user.POST("/login")
	}

	_ = router.Run(":8080")
}
