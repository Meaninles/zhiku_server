package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/knowledgeBase"
	"github.com/flipped-aurora/gin-vue-admin/server/router/setting"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System        system.RouterGroup
	Example       example.RouterGroup
	Setting       setting.RouterGroup
	KnowledgeBase knowledgeBase.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
