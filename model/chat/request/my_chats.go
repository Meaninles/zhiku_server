package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/chat"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type MyChatsSearch struct{
    chat.MyChats
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
