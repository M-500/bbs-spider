package middlewares

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-12 11:50

// PaginationMiddleware
//
//	@Description: 分页中间件，因为不想每次都写解析分页的数据
//	@return gin.HandlerFunc
func PaginationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		pageStr := c.DefaultQuery("page", "1")
		pageSizeStr := c.DefaultQuery("pageSize", "10")

		page, err := strconv.ParseInt(pageStr, 10, 64)
		if err != nil || page <= 0 {
			page = 1
		}

		pageSize, err := strconv.ParseInt(pageSizeStr, 10, 64)
		if err != nil || pageSize <= 0 {
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		limit := pageSize

		c.Set("offset", offset)
		c.Set("limit", limit)

		c.Next()
	}
}
