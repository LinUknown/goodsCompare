package main

import (
	"github.com/gin-gonic/gin"
	"imooc.com/controller"
	"imooc.com/middleware"
	"net/http"
)

func main()  {
	// Engin
	//router := gin.Default()
	router := gin.New()
	// 加载html文件，即template包下所有文件
	router.LoadHTMLGlob("template/*")
	router.StaticFS("/static",http.Dir("./static"))
	router.StaticFS("/template",http.Dir("./template"))
	// 曲线救国，把template也放到静态
	// 原因在于 ajax 回调直接重定向了，想到好办法后再改

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,"login.html",gin.H{
			"title":"Main website",
		})
	})

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"login.html",gin.H{
			"title":"Main website",
		})
	})

	// 路由组
	user := router.Group("/user")
	{   // 请求参数在请求路径上
		//user.GET("/get/:id/:username",controller.QueryById)
		user.GET("/query",controller.QueryParam)
		//user.POST("/insert",controller.InsertNewUser)
		user.GET("/form",controller.RenderForm)// 跳转html页面
		user.POST("/login",controller.Login)
		user.POST("/register",controller.Register)
	}

	goods:= router.Group("/goods")
	goods.Use(middleware.Authorize())
	{
		goods.GET("/get_search", func(c *gin.Context) {
			c.HTML(http.StatusOK,"search.html",gin.H{
				"title":"欢迎使用省心比价网",
			})
		})
		goods.GET("/search",controller.Search)
		goods.POST("/price_history",controller.PriceHistory)
	}

	// 指定地址和端口号
	router.Run(":9091")
}