package knowledgeBase

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/knowledgeBase"
	knowledgeBaseReq "github.com/flipped-aurora/gin-vue-admin/server/model/knowledgeBase/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserKnowledgeBaseApi struct {
}

var userKnowledgeBaseService = service.ServiceGroupApp.KnowledgeBaseServiceGroup.UserKnowledgeBaseService

// CreateUserKnowledgeBase 创建UserKnowledgeBase
// @Tags UserKnowledgeBase
// @Summary 创建UserKnowledgeBase
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body knowledgeBase.UserKnowledgeBase true "创建UserKnowledgeBase"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userKnowledgeBase/createUserKnowledgeBase [post]
func (userKnowledgeBaseApi *UserKnowledgeBaseApi) CreateUserKnowledgeBase(c *gin.Context) {
	var userKnowledgeBase knowledgeBase.UserKnowledgeBase
	err := c.ShouldBindJSON(&userKnowledgeBase)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	userKnowledgeBase.UserId = userId
	userKnowledgeBase.CreatedBy = userId
	verify := utils.Rules{
		"UserId":                   {utils.NotEmpty()},
		"KnowledgeBaseName":        {utils.NotEmpty()},
		"KnowledgeBaseDescription": {utils.NotEmpty()},
	}
	if err := utils.Verify(userKnowledgeBase, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := userKnowledgeBaseService.CreateUserKnowledgeBase(userKnowledgeBase); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteUserKnowledgeBase 删除UserKnowledgeBase
// @Tags UserKnowledgeBase
// @Summary 删除UserKnowledgeBase
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body knowledgeBase.UserKnowledgeBase true "删除UserKnowledgeBase"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userKnowledgeBase/deleteUserKnowledgeBase [delete]
func (userKnowledgeBaseApi *UserKnowledgeBaseApi) DeleteUserKnowledgeBase(c *gin.Context) {
	var userKnowledgeBase knowledgeBase.UserKnowledgeBase
	err := c.ShouldBindJSON(&userKnowledgeBase)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userKnowledgeBase.DeletedBy = utils.GetUserID(c)
	if err := userKnowledgeBaseService.DeleteUserKnowledgeBase(userKnowledgeBase); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteUserKnowledgeBaseByIds 批量删除UserKnowledgeBase
// @Tags UserKnowledgeBase
// @Summary 批量删除UserKnowledgeBase
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserKnowledgeBase"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /userKnowledgeBase/deleteUserKnowledgeBaseByIds [delete]
func (userKnowledgeBaseApi *UserKnowledgeBaseApi) DeleteUserKnowledgeBaseByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := userKnowledgeBaseService.DeleteUserKnowledgeBaseByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateUserKnowledgeBase 更新UserKnowledgeBase
// @Tags UserKnowledgeBase
// @Summary 更新UserKnowledgeBase
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body knowledgeBase.UserKnowledgeBase true "更新UserKnowledgeBase"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userKnowledgeBase/updateUserKnowledgeBase [put]
func (userKnowledgeBaseApi *UserKnowledgeBaseApi) UpdateUserKnowledgeBase(c *gin.Context) {
	var userKnowledgeBase knowledgeBase.UserKnowledgeBase
	err := c.ShouldBindJSON(&userKnowledgeBase)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userKnowledgeBase.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"KnowledgeBaseName":        {utils.NotEmpty()},
		"KnowledgeBaseDescription": {utils.NotEmpty()},
	}
	if err := utils.Verify(userKnowledgeBase, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := userKnowledgeBaseService.UpdateUserKnowledgeBase(userKnowledgeBase); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindUserKnowledgeBase 用id查询UserKnowledgeBase
// @Tags UserKnowledgeBase
// @Summary 用id查询UserKnowledgeBase
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query knowledgeBase.UserKnowledgeBase true "用id查询UserKnowledgeBase"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userKnowledgeBase/findUserKnowledgeBase [get]
func (userKnowledgeBaseApi *UserKnowledgeBaseApi) FindUserKnowledgeBase(c *gin.Context) {
	var userKnowledgeBase knowledgeBase.UserKnowledgeBase
	err := c.ShouldBindQuery(&userKnowledgeBase)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reuserKnowledgeBase, err := userKnowledgeBaseService.GetUserKnowledgeBase(userKnowledgeBase.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reuserKnowledgeBase": reuserKnowledgeBase}, c)
	}
}

// GetUserKnowledgeBaseList 分页获取UserKnowledgeBase列表
// @Tags UserKnowledgeBase
// @Summary 分页获取UserKnowledgeBase列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query knowledgeBaseReq.UserKnowledgeBaseSearch true "分页获取UserKnowledgeBase列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userKnowledgeBase/getUserKnowledgeBaseList [get]
func (userKnowledgeBaseApi *UserKnowledgeBaseApi) GetUserKnowledgeBaseList(c *gin.Context) {
	var pageInfo knowledgeBaseReq.UserKnowledgeBaseSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := userKnowledgeBaseService.GetUserKnowledgeBaseInfoList(pageInfo); err != nil {
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
