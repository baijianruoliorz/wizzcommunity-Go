package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"wizzcommunity/logic"
)

/*
*  @author liqiqiorz
*  @data 2020/10/22 10:03
 */
//跟社区相关

//CommunityHandler 可以查询到所有社区
func CommunityHandler(c *gin.Context) {
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

//CommunityDetailHandler 社区分类详情
func CommunityDetailHandler(c *gin.Context) {
	//	1.获取社区id
	//	这里是获取url参数
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	//	根据社区id获取社区详情
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
