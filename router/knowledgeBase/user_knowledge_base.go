package knowledgeBase

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserKnowledgeBaseRouter struct {
}

// InitUserKnowledgeBaseRouter 初始化 UserKnowledgeBase 路由信息
func (s *UserKnowledgeBaseRouter) InitUserKnowledgeBaseRouter(Router *gin.RouterGroup) {
	userKnowledgeBaseRouter := Router.Group("userKnowledgeBase").Use(middleware.OperationRecord())
	userKnowledgeBaseRouterWithoutRecord := Router.Group("userKnowledgeBase")
	var userKnowledgeBaseApi = v1.ApiGroupApp.KnowledgeBaseApiGroup.UserKnowledgeBaseApi
	{
		userKnowledgeBaseRouter.POST("createUserKnowledgeBase", userKnowledgeBaseApi.CreateUserKnowledgeBase)   // 新建UserKnowledgeBase
		userKnowledgeBaseRouter.DELETE("deleteUserKnowledgeBase", userKnowledgeBaseApi.DeleteUserKnowledgeBase) // 删除UserKnowledgeBase
		userKnowledgeBaseRouter.DELETE("deleteUserKnowledgeBaseByIds", userKnowledgeBaseApi.DeleteUserKnowledgeBaseByIds) // 批量删除UserKnowledgeBase
		userKnowledgeBaseRouter.PUT("updateUserKnowledgeBase", userKnowledgeBaseApi.UpdateUserKnowledgeBase)    // 更新UserKnowledgeBase
	}
	{
		userKnowledgeBaseRouterWithoutRecord.GET("findUserKnowledgeBase", userKnowledgeBaseApi.FindUserKnowledgeBase)        // 根据ID获取UserKnowledgeBase
		userKnowledgeBaseRouterWithoutRecord.GET("getUserKnowledgeBaseList", userKnowledgeBaseApi.GetUserKnowledgeBaseList)  // 获取UserKnowledgeBase列表
	}
}
