/*
上传小说点击数

2016-02-16

@aosen
*/

package models

import "github.com/astaxie/beego/orm"

type NovelPVModel struct {
	BaseModel
}

func NewNovelPVModel() *NovelPVModel {
	return &NovelPVModel{}
}

func (self *NovelPVModel) PutPV(novelid int) (map[string]interface{}, error) {
	o := orm.NewOrm()
	_, err := o.QueryTable("novel").Filter("id", novelid).Update(orm.Params{
		"novelpv": orm.ColValue(orm.Col_Add, 1),
	})
	if err != nil {
		return nil, err
	}
	ret := map[string]interface{}{
		"novelid": novelid,
	}
	return ret, nil
}
