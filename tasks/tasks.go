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
		case <-time.After(1 * time.Minute):
			nobj := models.NewBaseModel()
			if novels, err := nobj.GetAllNovel(); err != nil {
				log.Println(err.Error())
			} else {
				log.Println(novels[0].Id)
			}
		}
	}
}
