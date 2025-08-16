package bbs

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bbs"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type XkBBsService struct {
}

func (s *XkBBsService) CreatXkBbs(xkBbs *bbs.XkBbs) (err error) {
	err = global.GVA_DB.Create(xkBbs).Error
	return
}

func (s *XkBBsService) UpdateXkBbs(xkBbs *bbs.XkBbs) (err error) {
	err = global.GVA_DB.Updates(xkBbs).Error
	return
}

func (s *XkBBsService) DeleteXkBbs(xkBbs *bbs.XkBbs) (err error) {
	err = global.GVA_DB.Delete(xkBbs).Error
	return
}

// DeleteXkBbsById  根据id批量删除
func (s *XkBBsService) DeleteXkBbsById(id uint) (err error) {
	err = global.GVA_DB.Where("id = ?", id).Delete(&bbs.XkBbs{}).Error
	return
}

// GetXkBbs 根据id查询
func (s *XkBBsService) GetXkBbs(id uint) (xkBbs bbs.XkBbs, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&xkBbs).Error
	return
}

// LoadXkBbsPage 分页查询
func (s *XkBBsService) LoadXkBbsPage(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&bbs.XkBbs{})
	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Limit(limit).Offset(offset).Find(&list).Error
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
