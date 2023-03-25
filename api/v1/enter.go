package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/chat"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/knowledgeBase"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/setting"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup        system.ApiGroup
	ExampleApiGroup       example.ApiGroup
	SettingApiGroup       setting.ApiGroup
	KnowledgeBaseApiGroup knowledgeBase.ApiGroup
	ChatApiGroup          chat.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
