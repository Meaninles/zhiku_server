// 自动生成模板UserKnowledgeBase
package knowledgeBase

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// UserKnowledgeBase 结构体
type UserKnowledgeBase struct {
	global.GVA_MODEL
	UserId                   uint   `gorm:"index;column:user_id;comment:用户id"`
	KnowledgeBaseName        string `json:"knowledgeBaseName" form:"knowledgeBaseName" gorm:"index;column:knowledge_base_name;comment:知识库名称;"`
	KnowledgeBaseDescription string `json:"knowledgeBaseDescription" form:"knowledgeBaseDescription" gorm:"column:knowledge_base_description;comment:知识库内容的描述和介绍;"`
	CreatedBy                uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy                uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy                uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName UserKnowledgeBase 表名
func (UserKnowledgeBase) TableName() string {
	return "user_knowledge_base"
}
