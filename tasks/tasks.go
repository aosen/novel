/*
Author: Aosen
Data: 2016-2-2
Desc: novel 任务列表
*/

package tasks

import (
	"log"
	"novel/models"
	"time"

	"github.com/astaxie/beego/orm"
)

func SysTask(settings map[string]string) {
	defer func() {
		if x := recover(); x != nil {
			log.Printf("caught panic: %v", x)
		}
	}()
	for {
		select {
		//小说排行计算
		case <-time.After(10 * time.Minute):
			nobj := models.NewBaseModel()
			if novels, err := nobj.GetAllNovel(); err != nil {
				log.Println(err.Error())
			} else {
				//根据pv排序
				novelspv := models.NovelsPv(novels)
				models.NovelPvSort(novelspv)
				//根据collect排序
				novelscollect := models.NovelsCollect(novels)
				models.NovelCollectSort(novelscollect)
				npv := []*models.Novel(novelspv)
				nco := []*models.Novel(novelscollect)
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
				log.Println("00000000000000000000000000")
			}
		}
	}
}
