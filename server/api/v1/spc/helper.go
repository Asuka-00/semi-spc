package spc

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func parsePage(c *gin.Context) request.PageInfo {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	if pageInfo.Page <= 0 {
		pageInfo.Page = 1
	}
	if pageInfo.PageSize <= 0 {
		pageInfo.PageSize = 10
	}
	return pageInfo
}

func parseUintQuery(c *gin.Context, key string) uint {
	v, _ := strconv.ParseUint(c.Query(key), 10, 64)
	return uint(v)
}

func bindJSONID(c *gin.Context) (uint, error) {
	var req request.GetById
	if err := c.ShouldBindJSON(&req); err != nil {
		return 0, err
	}
	if req.ID == 0 {
		id := parseUintQuery(c, "ID")
		if id == 0 {
			id = parseUintQuery(c, "id")
		}
		return id, nil
	}
	return req.Uint(), nil
}

func bindQueryID(c *gin.Context) (uint, error) {
	var req request.GetById
	_ = c.ShouldBindQuery(&req)
	if req.ID == 0 {
		id := parseUintQuery(c, "ID")
		if id == 0 {
			id = parseUintQuery(c, "id")
		}
		return id, nil
	}
	return req.Uint(), nil
}

func failCreate(c *gin.Context, err error) {
	global.GVA_LOG.Error("创建失败!", zap.Error(err))
	response.FailWithMessage("创建失败: "+err.Error(), c)
}

func failUpdate(c *gin.Context, err error) {
	global.GVA_LOG.Error("更新失败!", zap.Error(err))
	response.FailWithMessage("更新失败: "+err.Error(), c)
}

func failDelete(c *gin.Context, err error) {
	global.GVA_LOG.Error("删除失败!", zap.Error(err))
	response.FailWithMessage("删除失败: "+err.Error(), c)
}

func failGet(c *gin.Context, err error) {
	global.GVA_LOG.Error("获取失败!", zap.Error(err))
	response.FailWithMessage("获取失败: "+err.Error(), c)
}

func okPage(c *gin.Context, list interface{}, total int64, page request.PageInfo) {
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     page.Page,
		PageSize: page.PageSize,
	}, "获取成功", c)
}
