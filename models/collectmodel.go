/*
上传小说收藏量

2016－02-17

@aosen
*/

package models

import "github.com/astaxie/beego/orm"

type NovelCollectModel struct {
	BaseModel
}

func NewNovelCollectModel() *NovelCollectModel {
	return &NovelCollectModel{}
}

func (self *NovelCollectModel) PutCollect(novelid int) (map[string]interface{}, error) {
	o := orm.NewOrm()
	_, err := o.QueryTable("novel").Filter("id", novelid).Update(orm.Params{
		"novelcollect": orm.ColValue(orm.Col_Add, 1),
	})
	if err != nil {
		return nil, err
	}
	ret := map[string]interface{}{
		"novelid": novelid,
	}
	return ret, nil
}
