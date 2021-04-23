package routers

import (
	"ginDemo02/pkg/setting"
	v1 "ginDemo02/routers/api/v1"
	"github.com/gin-gonic/gin"
)

// 初始化路由
func InitRouter() *gin.Engine {
	engine := gin.New()

	// 使用中间件
	// 等同于 gin.Default()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	// 设置运行模式
	gin.SetMode(setting.RunMode)

	apiV1 := engine.Group("/v1")
	{
		// 获取标签列表
		apiV1.GET("/tags", v1.GetTags)
		// 新建标签
		apiV1.POST("/tags", v1.AddTag)
		// 编辑标签
		apiV1.PUT("/tags/:id", v1.EditTag)
		// 删除标签
		apiV1.DELETE("/tags/:id", v1.DeleteTag)


		// 获取文章列表
		apiV1.GET("/articles", v1.GetArticles)
		// 获取单篇文章
		apiV1.GET("/articles/:id", v1.GetArticle)
		// 新建文章
		apiV1.POST("/articles", v1.AddArticle)
		// 编辑文章
		apiV1.PUT("/articles/:id", v1.EditArticle)
		// 删除文章
		apiV1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return engine
}
