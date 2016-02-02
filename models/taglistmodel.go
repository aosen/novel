/*
Author: Aosen
Data: 2016-01-11
QQ: 316052486
Desc: 获取taglist
*/

package models

import "github.com/astaxie/beego/orm"

type TagListModel struct {
	BaseModel
}

func NewTagListModel() *TagListModel {
	return &TagListModel{}
}

func (self *TagListModel) GetTagList() ([]map[string]interface{}, error) {
	var firsts []*First
	var seconds []*Second
	o := orm.NewOrm()
	//获取一级分类列表
	if _, err := o.QueryTable("first").All(&firsts); err != nil {
		return nil, err
	}
	//获取二级分类列表
	if _, err := o.QueryTable("second").All(&seconds); err != nil {
		return nil, err
	}
	//生成一个一级分类字典
	firstdict := make(map[int]string)
	for _, first := range firsts {
		firstdict[first.Id] = first.Firstname
	}

	ret := []map[string]interface{}{}
	for _, second := range seconds {
		ret = append(ret, map[string]interface{}{
			"firstid":    second.Firstid,
			"firstname":  firstdict[second.Firstid],
			"secondid":   second.Id,
			"secondname": second.Secondname,
		})
	}

	return ret, nil
}
