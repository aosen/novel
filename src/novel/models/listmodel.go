/*
Author: Aosen
Date: 2016-02-03
Desc:
获取小说列表
*/

package models

import "github.com/astaxie/beego/orm"

type NovelListModel struct {
	BaseModel
}

func NewNovelListModel() *NovelListModel {
	return &NovelListModel{}
}

func (self *NovelListModel) GetNovelList(firstid, secondid, page, limit int, np string) ([]map[string]interface{}, error) {
	var novels []*Novel
	o := orm.NewOrm()
	//获取小说列表
	if _, err := o.QueryTable("novel").Filter("firstid", firstid).Filter("secondid", secondid).Limit(limit, (page-1)*limit).All(&novels); err != nil {
		return nil, err
	}
	ret := []map[string]interface{}{}
	for _, novel := range novels {
		ret = append(ret, map[string]interface{}{
			"title":        novel.Title,
			"novelid":      novel.Id,
			"author":       novel.Author,
			"picture":      np + novel.Picture,
			"introduction": novel.Introduction,
		})
	}
	return ret, nil
}
