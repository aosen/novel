/*
Author: Aosen
Data: 2016-2-2
Desc: novel 任务列表
golang下有大神提供的cron库，
go下的cron库使用的是 github.com/robfig/cron，
最终使用的是 github.com/jakecoffman/cron，
后者也是前者的改进版，主要增加了个RemoveJob的函数来移除特定的任务。
http://ju.outofmemory.cn/entry/65356
*/
//使用示例:
//    c := cron.New()
//    spec := "*/5 * * * * ?"
//    c.AddFunc(spec, func() {
//        i++
//        log.Println("cron running:", i)
//    })
//    c.Start()
package tasks

import (
	"log"

	"github.com/jakecoffman/cron"
)

//执行定时器任务
func SysTask(settings map[string]string) {
	defer func() {
		if x := recover(); x != nil {
			log.Printf("caught panic: %v", x)
		}
	}()
	c := cron.New()
	//生成排名对象
	rankobj := NewRankTask()
	c.AddFunc("0 0 * * * ?", rankobj.PVRank, "pvrank")
	c.AddFunc("0 0 * * * ?", rankobj.CollectRank, "collectrank")
	c.Start()
}
