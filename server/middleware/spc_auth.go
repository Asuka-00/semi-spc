package middleware

import (
	"errors"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
)

// SpcCollectAuth SPC数据采集认证中间件（支持JWT或API Token）
func SpcCollectAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 先尝试JWT验证（从x-token header）
		token := utils.GetToken(c)
		if token != "" {
			j := utils.NewJWT()
			claims, err := j.ParseToken(token)
			if err == nil {
				// JWT有效，检查是否在黑名单
				if isBlacklist(token) {
					response.FailWithDetailed(gin.H{"reload": true}, "您的Token已被注销，请重新登录", c)
					c.Abort()
					return
				}
				// JWT验证通过，设置claims
				c.Set("claims", claims)
				c.Next()
				return
			}
			// JWT验证失败，继续尝试API Token
		}

		// 2. 尝试API Token验证（从X-API-Token或x-token header）
		apiToken := c.GetHeader("X-API-Token")
		if apiToken == "" {
			apiToken = c.GetHeader("x-token")
		}
		if apiToken == "" {
			response.FailWithMessage("未提供认证凭据，请在header中提供x-token（JWT）或X-API-Token（API Token）", c)
			c.Abort()
			return
		}

		// 3. 在sys_api_token表中查找该token
		var tokenRecord system.SysApiToken
		err := global.GVA_DB.Where("token = ? AND status = ?", apiToken, true).
			Preload("User").
			First(&tokenRecord).Error
		if err != nil {
			response.FailWithMessage("API Token无效或已失效", c)
			c.Abort()
			return
		}

		// 4. 检查是否过期
		if tokenRecord.ExpiresAt.Before(time.Now()) {
			response.FailWithMessage("API Token已过期", c)
			c.Abort()
			return
		}

		// 5. 解析API Token中的JWT claims（API Token本质是长期JWT）
		j := utils.NewJWT()
		claims, err := j.ParseToken(apiToken)
		if err != nil {
			response.FailWithMessage("API Token解析失败: "+err.Error(), c)
			c.Abort()
			return
		}

		// 6. 设置claims到context
		c.Set("claims", claims)
		c.Next()
	}
}

// VerifyApiToken 验证API Token是否有效（辅助函数）
func VerifyApiToken(token string) error {
	var tokenRecord system.SysApiToken
	err := global.GVA_DB.Where("token = ? AND status = ?", token, true).First(&tokenRecord).Error
	if err != nil {
		return errors.New("token无效")
	}
	if tokenRecord.ExpiresAt.Before(time.Now()) {
		return errors.New("token已过期")
	}
	return nil
}
