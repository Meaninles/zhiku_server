package setting

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/setting"
	settingReq "github.com/flipped-aurora/gin-vue-admin/server/model/setting/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

type ApiKeySettingApi struct {
}

var apiKeySettingService = service.ServiceGroupApp.SettingServiceGroup.ApiKeySettingService

// CreateApiKeySetting 创建ApiKeySetting
// @Tags ApiKeySetting
// @Summary 创建ApiKeySetting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body setting.ApiKeySetting true "创建ApiKeySetting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apiKeySetting/createApiKeySetting [post]
func (apiKeySettingApi *ApiKeySettingApi) CreateApiKeySetting(c *gin.Context) {
	var apiKeySetting setting.ApiKeySetting
	err := c.ShouldBindJSON(&apiKeySetting)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	apiKeySetting.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"ServiceName": {utils.NotEmpty()},
		"ApiKey":      {utils.NotEmpty()},
	}
	if err := utils.Verify(apiKeySetting, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := apiKeySettingService.CreateApiKeySetting(apiKeySetting); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		fmt.Println("xixi   ", os.Getenv("OPENAI_API_KEY"))
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteApiKeySetting 删除ApiKeySetting
// @Tags ApiKeySetting
// @Summary 删除ApiKeySetting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body setting.ApiKeySetting true "删除ApiKeySetting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /apiKeySetting/deleteApiKeySetting [delete]
func (apiKeySettingApi *ApiKeySettingApi) DeleteApiKeySetting(c *gin.Context) {
	var apiKeySetting setting.ApiKeySetting
	err := c.ShouldBindJSON(&apiKeySetting)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	apiKeySetting.DeletedBy = utils.GetUserID(c)
	if err := apiKeySettingService.DeleteApiKeySetting(apiKeySetting); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteApiKeySettingByIds 批量删除ApiKeySetting
// @Tags ApiKeySetting
// @Summary 批量删除ApiKeySetting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ApiKeySetting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /apiKeySetting/deleteApiKeySettingByIds [delete]
func (apiKeySettingApi *ApiKeySettingApi) DeleteApiKeySettingByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := apiKeySettingService.DeleteApiKeySettingByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateApiKeySetting 更新ApiKeySetting
// @Tags ApiKeySetting
// @Summary 更新ApiKeySetting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body setting.ApiKeySetting true "更新ApiKeySetting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /apiKeySetting/updateApiKeySetting [put]
func (apiKeySettingApi *ApiKeySettingApi) UpdateApiKeySetting(c *gin.Context) {
	var apiKeySetting setting.ApiKeySetting
	err := c.ShouldBindJSON(&apiKeySetting)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	apiKeySetting.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"ServiceName": {utils.NotEmpty()},
		"ApiKey":      {utils.NotEmpty()},
	}
	if err := utils.Verify(apiKeySetting, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := apiKeySettingService.UpdateApiKeySetting(apiKeySetting); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindApiKeySetting 用id查询ApiKeySetting
// @Tags ApiKeySetting
// @Summary 用id查询ApiKeySetting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query setting.ApiKeySetting true "用id查询ApiKeySetting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /apiKeySetting/findApiKeySetting [get]
func (apiKeySettingApi *ApiKeySettingApi) FindApiKeySetting(c *gin.Context) {
	var apiKeySetting setting.ApiKeySetting
	err := c.ShouldBindQuery(&apiKeySetting)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reapiKeySetting, err := apiKeySettingService.GetApiKeySetting(apiKeySetting.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reapiKeySetting": reapiKeySetting}, c)
	}
}

// GetApiKeySettingList 分页获取ApiKeySetting列表
// @Tags ApiKeySetting
// @Summary 分页获取ApiKeySetting列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query settingReq.ApiKeySettingSearch true "分页获取ApiKeySetting列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apiKeySetting/getApiKeySettingList [get]
func (apiKeySettingApi *ApiKeySettingApi) GetApiKeySettingList(c *gin.Context) {
	var pageInfo settingReq.ApiKeySettingSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := apiKeySettingService.GetApiKeySettingInfoList(pageInfo); err != nil {
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

func (apiKeySettingApi *ApiKeySettingApi) UseApiKeySetting(c *gin.Context) {
	var apiKeySetting setting.ApiKeySetting
	err := c.ShouldBindJSON(&apiKeySetting)
	if err != nil {
		response.FailWithMessage("错误的请求", c)
		return
	}
	err = os.Setenv("OPENAI_API_KEY", apiKeySetting.ApiKey)
	if err != nil || apiKeySetting.ApiKey == "" {
		response.FailWithMessage("错误的请求，API key为空", c)
		return
	}

	// TODO delete
	id := utils.GetUserID(c)
	fmt.Println("xixixi: ", id)

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		apiKeySettingInUsing, _ := apiKeySettingService.GetApiKeySettingInUsing()
		if err := tx.Model(&setting.ApiKeySetting{}).Where("id = ?", apiKeySettingInUsing.ID).Update("using", 0).Error; err != nil {
			return err
		}
		return nil
	})

	if err := apiKeySettingService.UpdateApiKeySetting(apiKeySetting); err != nil {
		global.GVA_LOG.Error("更新启用api key失败!", zap.Error(err))
		response.FailWithMessage("更新启用api key失败", c)
	} else {
		response.OkWithMessage("更新启用api key成功", c)
	}
}
