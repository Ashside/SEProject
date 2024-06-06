package main

import (
	"github.com/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}
func main() {

	r := gin.Default()
	r.Use(corsMiddleware())

	apiGroup := r.Group("/api")

	// 处理/api/usr路由组
	usrGroup := apiGroup.Group("/usr")
	// 处理/api/usr/loginCheck的OPTIONS请求
	// 该请求用于检查用户登录状态
	// 输入：form表单，包含username和password
	// 输出：json格式，包含success、reason和type字段
	//usrGroup.OPTIONS("/loginCheck", api.LoginCheckPost)
	usrGroup.POST("/loginCheck", api.LoginCheckPost)
	// 处理/api/usr/registerCheck的OPTIONS请求
	// 该请求用于检查用户注册状态
	// 输入：form表单，包含username, password, type字段
	// 输出：json格式，包含success、reason字段
	usrGroup.POST("/registerCheck", api.RegisterCheckPost)

	// 处理/api/questionBank路由组
	questionBankGroup := apiGroup.Group("/questionBank")
	// 处理/api/questionBank/addQuestion路由组
	addQuestionGroup := questionBankGroup.Group("/addQuestion")
	// 处理/api/questionBank/addQuestion/simpleAnswer的OPTIONS请求
	// 该请求用于添加题目
	// 输入：form表单，包含question, answer, difficulty, subject, username字段
	// 输出：json格式，包含success、reason字段
	addQuestionGroup.POST("/simpleAnswer", api.AddSimpleAnswerPost)

	_ = r.Run(":8081")
}
