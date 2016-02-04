/*
获取小说章节
2016-02-04
@aosen
*/

package models

import (
	"novel/utils"

	"github.com/astaxie/beego/orm"
)

type NovelChapterModel struct {
	BaseModel
}

func NewNovelChapterModel() *NovelChapterModel {
	return &NovelChapterModel{}
}

func (self *NovelChapterModel) GetChapterList(novelid int) ([]map[string]interface{}, error) {
	var chapters []*Content
	o := orm.NewOrm()
	//获取章节列表
	if _, err := o.QueryTable("content").Filter("novelid", novelid).All(&chapters, "Id", "Novelid", "Title", "Subtitle", "Chapter"); err != nil {
		return nil, err
	}
	ret := utils.KVL{}
	for _, chapter := range chapters {
		ret = ret.Append(ret, map[string]interface{}{
			"title":     chapter.Title,
			"novelid":   chapter.Novelid,
			"subtitle":  chapter.Subtitle,
			"chapterid": chapter.Id,
			"chapter":   chapter.Chapter,
		})
	}
	utils.MapDicSortToMap(ret)
	return ret, nil
}
