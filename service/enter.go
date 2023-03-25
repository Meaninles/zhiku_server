package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/chat"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/knowledgeBase"
	"github.com/flipped-aurora/gin-vue-admin/server/service/setting"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup        system.ServiceGroup
	ExampleServiceGroup       example.ServiceGroup
	SettingServiceGroup       setting.ServiceGroup
	KnowledgeBaseServiceGroup knowledgeBase.ServiceGroup
	ChatServiceGroup          chat.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
