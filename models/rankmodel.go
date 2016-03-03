/*
根据page limit 获取排行榜
*/

package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type NovelRankModel struct {
	BaseModel
}

func NewNovelRankModel() *NovelRankModel {
	return &NovelRankModel{}
}

func (self *NovelRankModel) GetRankList(page, limit int, picpath string) ([]map[string]interface{}, error) {
	var crs []*Clickrank
	var novels []*Novel
	o := orm.NewOrm()
	if _, err := o.QueryTable("clickrank").Filter("createtime", time.Now()).Limit(limit, page).All(&crs, "id", "Novelid"); err != nil {
		return nil, err
	}
	nids := []int{}
	for _, novel := range crs {
		nids = append(nids, novel.Novelid)
	}
	if _, err := o.QueryTable("novel").Filter("id__in", nids).All(&novels, "Id", "Title", "Novelpv", "Author", "Picture", "Firstid", "Secondid", "Introduction"); err != nil {
		return nil, err
	} else {
		ret := []map[string]interface{}{}
		for _, novel := range novels {
			ret = append(ret, map[string]interface{}{
				"title":        novel.Title,
				"novelid":      novel.Id,
				"novelpv":      novel.Novelpv,
				"author":       novel.Author,
				"picture":      picpath + novel.Picture,
				"first":        novel.Firstid,
				"second":       novel.Secondid,
				"introduction": novel.Introduction})
		}
		return ret, nil
	}
}
