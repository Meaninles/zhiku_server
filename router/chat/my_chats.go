package chat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MyChatsRouter struct {
}

// InitMyChatsRouter 初始化 MyChats 路由信息
func (s *MyChatsRouter) InitMyChatsRouter(Router *gin.RouterGroup) {
	myChatsRouter := Router.Group("myChats").Use(middleware.OperationRecord())
	myChatsRouterWithoutRecord := Router.Group("myChats")
	var myChatsApi = v1.ApiGroupApp.ChatApiGroup.MyChatsApi
	{
		myChatsRouter.POST("createMyChats", myChatsApi.CreateMyChats)             // 新建MyChats
		myChatsRouter.DELETE("deleteMyChats", myChatsApi.DeleteMyChats)           // 删除MyChats
		myChatsRouter.DELETE("deleteMyChatsByIds", myChatsApi.DeleteMyChatsByIds) // 批量删除MyChats
		myChatsRouter.PUT("updateMyChats", myChatsApi.UpdateMyChats)              // 更新MyChats
	}
	{
		myChatsRouterWithoutRecord.GET("findMyChats", myChatsApi.FindMyChats)       // 根据ID获取MyChats
		myChatsRouterWithoutRecord.GET("getMyChatsList", myChatsApi.GetMyChatsList) // 获取MyChats列表
		myChatsRouterWithoutRecord.GET("getKBIds", myChatsApi.GetKBIds)
	}
}
