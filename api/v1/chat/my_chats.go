package chat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/chat"
	chatReq "github.com/flipped-aurora/gin-vue-admin/server/model/chat/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MyChatsApi struct {
}

var myChatsService = service.ServiceGroupApp.ChatServiceGroup.MyChatsService

// CreateMyChats 创建MyChats
// @Tags MyChats
// @Summary 创建MyChats
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body chat.MyChats true "创建MyChats"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /myChats/createMyChats [post]
func (myChatsApi *MyChatsApi) CreateMyChats(c *gin.Context) {
	var myChats chat.MyChats
	err := c.ShouldBindJSON(&myChats)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	myChats.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"UserId":   {utils.NotEmpty()},
		"ChatName": {utils.NotEmpty()},
		"KBId":     {utils.NotEmpty()},
		"ChatType": {utils.NotEmpty()},
	}
	if err := utils.Verify(myChats, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := myChatsService.CreateMyChats(myChats); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMyChats 删除MyChats
// @Tags MyChats
// @Summary 删除MyChats
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body chat.MyChats true "删除MyChats"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /myChats/deleteMyChats [delete]
func (myChatsApi *MyChatsApi) DeleteMyChats(c *gin.Context) {
	var myChats chat.MyChats
	err := c.ShouldBindJSON(&myChats)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	myChats.DeletedBy = utils.GetUserID(c)
	if err := myChatsService.DeleteMyChats(myChats); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMyChatsByIds 批量删除MyChats
// @Tags MyChats
// @Summary 批量删除MyChats
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MyChats"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /myChats/deleteMyChatsByIds [delete]
func (myChatsApi *MyChatsApi) DeleteMyChatsByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := myChatsService.DeleteMyChatsByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMyChats 更新MyChats
// @Tags MyChats
// @Summary 更新MyChats
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body chat.MyChats true "更新MyChats"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /myChats/updateMyChats [put]
func (myChatsApi *MyChatsApi) UpdateMyChats(c *gin.Context) {
	var myChats chat.MyChats
	err := c.ShouldBindJSON(&myChats)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	myChats.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"UserId":   {utils.NotEmpty()},
		"ChatName": {utils.NotEmpty()},
		"KBId":     {utils.NotEmpty()},
		"ChatType": {utils.NotEmpty()},
	}
	if err := utils.Verify(myChats, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := myChatsService.UpdateMyChats(myChats); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMyChats 用id查询MyChats
// @Tags MyChats
// @Summary 用id查询MyChats
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query chat.MyChats true "用id查询MyChats"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /myChats/findMyChats [get]
func (myChatsApi *MyChatsApi) FindMyChats(c *gin.Context) {
	var myChats chat.MyChats
	err := c.ShouldBindQuery(&myChats)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if remyChats, err := myChatsService.GetMyChats(myChats.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remyChats": remyChats}, c)
	}
}

// GetMyChatsList 分页获取MyChats列表
// @Tags MyChats
// @Summary 分页获取MyChats列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query chatReq.MyChatsSearch true "分页获取MyChats列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /myChats/getMyChatsList [get]
func (myChatsApi *MyChatsApi) GetMyChatsList(c *gin.Context) {
	var pageInfo chatReq.MyChatsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := myChatsService.GetMyChatsInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (myChatsApi *MyChatsApi) GetKBIds(c *gin.Context) {
	userId := utils.GetUserID(c)
	if kBIds, err := myChatsService.GetKBIds(userId); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(kBIds, c)
	}
}
