/*
根据chapterid获取小说内容, 返回
title
subtitle
novelid
content
chapterid
prev
next

2016-2-15

@aosen
*/

package models

import (
	"log"

	"github.com/astaxie/beego/orm"
)

type NovelContentModel struct {
	BaseModel
}

func NewNovelContentModel() *NovelContentModel {
	return &NovelContentModel{}
}

func (self *NovelContentModel) GetContent(chapterid int) (map[string]interface{}, error) {
	var content Content
	o := orm.NewOrm()
	//获取小说内容
	if err := o.QueryTable("content").Filter("id", chapterid).One(&content, "Id", "Novelid", "Title", "Subtitle", "Text"); err != nil {
		log.Println("00000000000000000000")
		return nil, err
	}
	//根据chapterid获取上一章节的chapterid和下一章节的chapterid
	var (
		pre  Content
		next Content
	)
	//获取上一章id
	if err := o.QueryTable("content").Filter("novelid", content.Novelid).Filter("id__lt", chapterid).OrderBy("-id").Limit(1).One(&pre, "Id"); err != nil {
		pre.Id = 0
	}
	//获取下一章Id
	if err := o.QueryTable("content").Filter("novelid", content.Novelid).Filter("id__gt", chapterid).OrderBy("id").Limit(1).One(&next, "Id"); err != nil {
		next.Id = 0
	}

	//返回结果
	ret := map[string]interface{}{
		"title":     content.Title,
		"subtitle":  content.Subtitle,
		"novelid":   content.Novelid,
		"content":   content.Text,
		"chapterid": chapterid,
		"prev":      pre.Id,
		"next":      next.Id,
	}
	return ret, nil
}
