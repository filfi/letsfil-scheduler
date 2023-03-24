package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type pageInfo struct {
	Page     int    `json:"page" param:"page" query:"page" form:"page"  xml:"page"`
	PageSize int    `json:"page_size" param:"page_size" query:"page_size" form:"page_size"  xml:"page_size"`
	Sort     string `json:"sort" form:"sort" uri:"sort"`
}

func GeneratePaginationFromRequest(c *gin.Context) (pagination pageInfo) {
	if err := c.ShouldBind(&pagination); err != nil {
		fmt.Printf("参数绑定错误:%s\n", err)
	}

	// 校验参数
	if pagination.PageSize < 0 {
		pagination.PageSize = 10
	}
	if pagination.Page < 1 {
		pagination.Page = 1
	}

	if len(pagination.Sort) == 0 {
		pagination.Sort = "created_at desc"
	}

	return
}
