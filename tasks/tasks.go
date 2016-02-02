/*
Author: Aosen
Data: 2016-2-2
Desc: novel 任务列表
*/

package tasks

import (
	"log"
	"time"
)

func SysTask(settings map[string]string) {
	for {
		select {
		case <-time.After(1 * time.Hour):
			log.Println("lalalallalalalalla")
		}
	}
}
