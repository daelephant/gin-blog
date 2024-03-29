/**
 * @Author: yin
 * @Description:pagination
 * @Software: GoLand
 * @Version: 1.0.0
 * @Time : 2019-09-01 19:40
 */
package util

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"

	"gin-blog/pkg/setting"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}

	return result
}
