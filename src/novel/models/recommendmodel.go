/*
 获取推荐列表
 @aosen
*/

package models

import "github.com/astaxie/beego/orm"

type NovelRecommendModel struct {
	BaseModel
}

func NewNovelRecommendModel() *NovelRecommendModel {
	return &NovelRecommendModel{}
}

func (self *NovelRecommendModel) GetList(picpath string) ([]map[string]interface{}, error) {
	var reclist []*Recommend
	o := orm.NewOrm()
	//获取推荐列表
	if _, err := o.QueryTable("recommend").All(&reclist, "Id", "Tagid", "Novelid", "Top"); err != nil {
		return nil, err
	}
	ret := []map[string]interface{}{}
	for k := range RecommendMap {
		el := map[string]interface{}{}
		novellist := []map[string]interface{}{}
		for _, n := range reclist {
			if n.Tagid == k {
				el["tag"] = RecommendMap[n.Tagid]
				el["tagid"] = n.Tagid
				novel, err := self.GetNovel(n.Id)
				if err != nil {
					continue
				} else {
					novellist = append(novellist, map[string]interface{}{
						"title":      novel["title"],
						"secondname": self.GetSecondName(novel["second"].(int)),
						"novelid":    novel["novelid"],
						"intro":      novel["introduction"],
						"picture":    picpath + novel["picture"].(string),
					})
				}
			}
		}
		el["novellist"] = novellist[0:5]
		ret = append(ret, el)
	}
	return ret, nil
}

func (self *NovelRecommendModel) GetMore(tagid int, picpath string) (map[string]interface{}, error) {
	var reclist []*Recommend
	o := orm.NewOrm()
	//获取推荐列表
	if _, err := o.QueryTable("recommend").Filter("tagid", tagid).All(&reclist, "Id", "Tagid", "Novelid", "Top"); err != nil {
		return nil, err
	}
	el := map[string]interface{}{}
	novellist := []map[string]interface{}{}
	ret := map[string]interface{}{}
	for _, n := range reclist {
		el["tag"] = RecommendMap[n.Tagid]
		el["tagid"] = n.Tagid
		novel, err := self.GetNovel(n.Id)
		if err != nil {
			continue
		} else {
			novellist = append(novellist, map[string]interface{}{
				"title":      novel["title"],
				"secondname": self.GetSecondName(novel["second"].(int)),
				"novelid":    novel["novelid"],
				"intro":      novel["introduction"],
				"picture":    picpath + novel["picture"].(string),
				"author":     novel["author"],
			})
		}
	}
	el["novellist"] = novellist
	ret = el
	return ret, nil
}
