// 自动生成模板MyChats
package chat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// MyChats 结构体
type MyChats struct {
      global.GVA_MODEL
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:知识库名称;"`
      ChatName  string `json:"chatName" form:"chatName" gorm:"column:chat_name;comment:知识库内容的描述和介绍;"`
      KBId  string `json:"kBId" form:"kBId" gorm:"column:kb_id;comment:知识库id;"`
      ChatType  *int `json:"chatType" form:"chatType" gorm:"column:chat_type;comment:对话类型;"`
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName MyChats 表名
func (MyChats) TableName() string {
  return "my_chats"
}

