package knowledgeBase

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/knowledgeBase"
	knowledgeBaseReq "github.com/flipped-aurora/gin-vue-admin/server/model/knowledgeBase/request"
	pb "github.com/flipped-aurora/gin-vue-admin/server/proto"
	"gorm.io/gorm"
)

type UserKnowledgeBaseService struct {
}

// CreateUserKnowledgeBase 创建UserKnowledgeBase记录
// Author [piexlmax](https://github.com/piexlmax)
func (userKnowledgeBaseService *UserKnowledgeBaseService) CreateUserKnowledgeBase(userKnowledgeBase knowledgeBase.UserKnowledgeBase) (err error) {
	res, err := global.GVA_KBINDEX_RPC_CLIENT.Create(context.Background(), &pb.KBIndexRequest{
		Name:        "1",
		Url:         "1",
		Tag:         "1",
		Key:         "1",
		Description: "1",
		Indexed:     0,
		KbId:        "1",
	})
	if err != nil {
		fmt.Println("怎么了？？？？》》》》》》》》")
		return err
	}
	fmt.Println("o不ok>>>>>")
	fmt.Println(res.Ok)

	err = global.GVA_DB.Create(&userKnowledgeBase).Error
	return err
}

// DeleteUserKnowledgeBase 删除UserKnowledgeBase记录
// Author [piexlmax](https://github.com/piexlmax)
func (userKnowledgeBaseService *UserKnowledgeBaseService) DeleteUserKnowledgeBase(userKnowledgeBase knowledgeBase.UserKnowledgeBase) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&knowledgeBase.UserKnowledgeBase{}).Where("id = ?", userKnowledgeBase.ID).Update("deleted_by", userKnowledgeBase.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&userKnowledgeBase).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteUserKnowledgeBaseByIds 批量删除UserKnowledgeBase记录
// Author [piexlmax](https://github.com/piexlmax)
func (userKnowledgeBaseService *UserKnowledgeBaseService) DeleteUserKnowledgeBaseByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&knowledgeBase.UserKnowledgeBase{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&knowledgeBase.UserKnowledgeBase{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateUserKnowledgeBase 更新UserKnowledgeBase记录
// Author [piexlmax](https://github.com/piexlmax)
func (userKnowledgeBaseService *UserKnowledgeBaseService) UpdateUserKnowledgeBase(userKnowledgeBase knowledgeBase.UserKnowledgeBase) (err error) {
	err = global.GVA_DB.Save(&userKnowledgeBase).Error
	return err
}

// GetUserKnowledgeBase 根据id获取UserKnowledgeBase记录
// Author [piexlmax](https://github.com/piexlmax)
func (userKnowledgeBaseService *UserKnowledgeBaseService) GetUserKnowledgeBase(id uint) (userKnowledgeBase knowledgeBase.UserKnowledgeBase, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&userKnowledgeBase).Error
	return
}

// GetUserKnowledgeBaseInfoList 分页获取UserKnowledgeBase记录
// Author [piexlmax](https://github.com/piexlmax)
func (userKnowledgeBaseService *UserKnowledgeBaseService) GetUserKnowledgeBaseInfoList(info knowledgeBaseReq.UserKnowledgeBaseSearch) (list []knowledgeBase.UserKnowledgeBase, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&knowledgeBase.UserKnowledgeBase{})
	var userKnowledgeBases []knowledgeBase.UserKnowledgeBase
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&userKnowledgeBases).Error
	return userKnowledgeBases, total, err
}
