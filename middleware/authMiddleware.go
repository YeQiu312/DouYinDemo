package middleware

import (
	"douyin/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//该apk发送的token不在请求头中
		//authHeader := c.GetHeader("Authorization")

		Token := c.Query("token") // 先尝试从 URL 查询参数中获取
		if Token == "" {
			// 如果 URL 查询参数中不存在，再从请求体中获取
			Token = c.PostForm("token")
		}

		if Token == "" {
			// 如果还是获取不到 token，说明请求中没有提供 token 参数
			c.JSON(http.StatusBadRequest, util.Response{
				StatusCode: -1,
				StatusMsg:  "请求失败，缺少 token 参数",
			})
			return
		}

		// 校验 token 是否正确
		claims, err := util.ParseToken(Token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, util.UserInfoResponse{
				StatusCode: -1,
				StatusMsg:  "token错误",
			})
		}

		// 将解析后的userID保存到context中，方便后续调用
		c.Set("userID", claims.UserID)

		c.Next()
	}
}
