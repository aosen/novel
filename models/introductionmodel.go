/*
获取小说简介
2016-02-04
@aosen
*/

package models

import (
	"errors"

	"github.com/astaxie/beego/orm"
)

type NovelIntroductionModel struct {
	BaseModel
}

func NewNovelIntroductionModel() *NovelIntroductionModel {
	return &NovelIntroductionModel{}
}

func (self *NovelIntroductionModel) GetChapterNum(novelid int) (int64, error) {
	o := orm.NewOrm()
	return o.QueryTable("content").Filter("novelid", novelid).Count()
}

func (self *NovelIntroductionModel) GetNovelIntroduction(novelid int, np string) (map[string]interface{}, error) {
	o := orm.NewOrm()
	novel := Novel{Id: novelid}
	//获取小说简介
	err := o.Read(&novel)
	if err == orm.ErrNoRows {
		return nil, errors.New("查询不到")
	} else if err == orm.ErrMissPK {
		return nil, errors.New("找不到主键")
	} else {
		if cnt, e := self.GetChapterNum(novelid); e != nil {
			return nil, e
		} else {
			return map[string]interface{}{
				"title":        novel.Title,
				"novelid":      novel.Id,
				"author":       novel.Author,
				"picture":      np + novel.Picture,
				"introduction": novel.Introduction,
				"chapternum":   cnt,
			}, nil
		}
	}
}
