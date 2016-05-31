/*
小说下载
*/

package models

import (
	"errors"
	"novel/utils"

	"github.com/astaxie/beego/orm"
)

type NovelDownloadModel struct {
	BaseModel
}

func NewNovelDownloadModel() *NovelDownloadModel {
	return &NovelDownloadModel{}
}

//根据novelid生成json文本
func (self *NovelDownloadModel) GetNovelText(novelid int) (map[string]interface{}, error) {
	ret := map[string]interface{}{}
	//首先通过novelid获取小说的标题
	o := orm.NewOrm()
	novel := Novel{Id: novelid}
	//获取小说简介
	err := o.Read(&novel)
	if err == orm.ErrNoRows {
		return nil, errors.New("查询不到")
	} else if err == orm.ErrMissPK {
		return nil, errors.New("找不到主键")
	} else {
		ret["title"] = novel.Title
	}
	//获取章节列表
	var chapters []*Content
	//获取章节列表
	if _, err := o.QueryTable("content").Filter("novelid", novelid).All(&chapters, "Id", "Subtitle", "Chapter", "Text"); err != nil {
		return nil, err
	}
	chapterlist := utils.KVL{}
	for _, chapter := range chapters {
		chapterlist = chapterlist.Append(chapterlist, map[string]interface{}{
			"chapterid": chapter.Id,
			"subtitle":  chapter.Subtitle,
			"content":   chapter.Text,
			"chapter":   chapter.Chapter,
		})
	}
	utils.MapDicSortToMap(chapterlist)
	ret["chaptercontent"] = chapterlist
	ret["novelid"] = novelid
	return ret, nil
}
