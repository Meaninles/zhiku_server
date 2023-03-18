package setting

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ApiKeySettingRouter struct {
}

// InitApiKeySettingRouter 初始化 ApiKeySetting 路由信息
func (s *ApiKeySettingRouter) InitApiKeySettingRouter(Router *gin.RouterGroup) {
	apiKeySettingRouter := Router.Group("apiKeySetting").Use(middleware.OperationRecord())
	apiKeySettingRouterWithoutRecord := Router.Group("apiKeySetting")
	var apiKeySettingApi = v1.ApiGroupApp.SettingApiGroup.ApiKeySettingApi
	{
		apiKeySettingRouter.POST("createApiKeySetting", apiKeySettingApi.CreateApiKeySetting)             // 新建ApiKeySetting
		apiKeySettingRouter.DELETE("deleteApiKeySetting", apiKeySettingApi.DeleteApiKeySetting)           // 删除ApiKeySetting
		apiKeySettingRouter.DELETE("deleteApiKeySettingByIds", apiKeySettingApi.DeleteApiKeySettingByIds) // 批量删除ApiKeySetting
		apiKeySettingRouter.PUT("updateApiKeySetting", apiKeySettingApi.UpdateApiKeySetting)              // 更新ApiKeySetting
		apiKeySettingRouter.POST("useApiKeySetting", apiKeySettingApi.UseApiKeySetting)
	}
	{
		apiKeySettingRouterWithoutRecord.GET("findApiKeySetting", apiKeySettingApi.FindApiKeySetting)       // 根据ID获取ApiKeySetting
		apiKeySettingRouterWithoutRecord.GET("getApiKeySettingList", apiKeySettingApi.GetApiKeySettingList) // 获取ApiKeySetting列表
	}
}
