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

func Timer(d time.Duration, task func()) {
	timer := time.NewTicker(d)
	for {
		select {
		case <-timer.C:
			task()
		}
	}
}

func SysTask(settings map[string]string) {
	defer func() {
		if x := recover(); x != nil {
			log.Printf("caught panic: %v", x)
		}
	}()
	//生成排名对象
	rankobj := NewRankTask()
	go Timer(1*time.Second, rankobj.PVRank)
	go Timer(1*time.Second, rankobj.CollectRank)
}
