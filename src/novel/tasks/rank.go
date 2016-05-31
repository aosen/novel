/*
Author: Aosen
Data: 2016-02-17
Desc: 所有与排名相关的任务
*/

package tasks

import (
	"log"
	"novel/models"
	"time"

	"github.com/astaxie/beego/orm"
)

type RankTask struct {
}

func NewRankTask() *RankTask {
	return &RankTask{}
}

func (self *RankTask) PVRank() {
	log.Println("start pv rank")
	nobj := models.NewBaseModel()
	if novels, err := nobj.GetAllNovel(); err != nil {
		log.Println(err.Error())
	} else {
		//根据pv排序
		novelspv := models.NovelsPv(novels)
		models.NovelPvSort(novelspv)
		npv := []*models.Novel(novelspv)
		now := time.Now()
		o := orm.NewOrm()
		//插入pv排行榜
		pvrank := []models.Clickrank{}
		for _, novel := range npv {
			pvrank = append(pvrank,
				models.Clickrank{
					Novelid:      novel.Id,
					Firstid:      novel.Firstid,
					Secondid:     novel.Secondid,
					Novelpv:      novel.Novelpv,
					Novelcollect: novel.Novelcollect,
					Createtime:   now,
				})
		}
		if _, err := o.InsertMulti(1, pvrank); err != nil {
			log.Println("insert pv rank:", err.Error())
		} else {
			log.Println("insert pv rank success")
		}
	}
}

func (self *RankTask) CollectRank() {
	log.Println("start collect rank")
	nobj := models.NewBaseModel()
	if novels, err := nobj.GetAllNovel(); err != nil {
		log.Println(err.Error())
	} else {
		//根据collect排序
		novelscollect := models.NovelsCollect(novels)
		models.NovelCollectSort(novelscollect)
		nco := []*models.Novel(novelscollect)
		now := time.Now()
		o := orm.NewOrm()
		//插入collect排行榜
		collectrank := []models.Collectrank{}
		for _, novel := range nco {
			collectrank = append(collectrank,
				models.Collectrank{
					Novelid:      novel.Id,
					Firstid:      novel.Firstid,
					Secondid:     novel.Secondid,
					Novelpv:      novel.Novelpv,
					Novelcollect: novel.Novelcollect,
					Createtime:   now,
				})
		}
		if _, err := o.InsertMulti(1, collectrank); err != nil {
			log.Println("insert collect rank:", err.Error())
		} else {
			log.Println("insert collect rank success")
		}
	}
}
