package spc

import "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"

func normalizePage(info request.PageInfo) request.PageInfo {
	if info.Page <= 0 {
		info.Page = 1
	}
	if info.PageSize <= 0 {
		info.PageSize = 10
	}
	if info.PageSize > 500 {
		info.PageSize = 500
	}
	return info
}
