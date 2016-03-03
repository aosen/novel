/*
构建索引
*/
package tasks

import (
	"log"
	"novel/models"
	"time"

	"github.com/aosen/search"
)

type Task struct {
	Execut func(docid uint64, arg string)
	Arg    string
	Docid  uint64
}

type ManagementCenter struct {
	Queue chan *Task
	//协程数
	Number int
	//任务队列最大容量
	Total int
}

//生成一个管理中心
func NewManagementCenter(number, total int) *ManagementCenter {
	return &ManagementCenter{
		Queue:  make(chan *Task, total),
		Number: number,
		Total:  total,
	}
}

//管理器开始运行
func (self *ManagementCenter) Start() {
	// 开启Number个goroutine
	for i := 0; i < self.Number; i++ {
		go func() {
			for {
				task, ok := <-self.Queue
				if !ok {
					break
				}

				task.Execut(task.Docid, task.Arg)
			}
		}()
	}
}

// 关门送客
func (self *ManagementCenter) Stop() {
	close(self.Queue)
}

// 添加任务
func (self *ManagementCenter) AddTask(task *Task) {
	self.Queue <- task
}

type IndexTask struct {
	searcher *search.Engine
}

func NewIndexTask(searcher *search.Engine) *IndexTask {
	return &IndexTask{
		searcher: searcher,
	}
}

func (self *IndexTask) Index() {
	doindex := func(docid uint64, arg string) {
		self.searcher.IndexDocument(docid, search.DocumentIndexData{
			Content: arg,
		})
	}
	//生成一个管理中心
	mc := NewManagementCenter(100, 100)
	go mc.Start()
	//每天更新一次
	for {
		novels, err := models.NewBaseModel().GetAllNovelForIndex()
		if err != nil {
			log.Println(err.Error())
		} else {
			for _, novel := range novels {
				mc.AddTask(&Task{
					Execut: doindex,
					Arg:    novel.Title + novel.Author,
					Docid:  uint64(novel.Id),
				})
			}
		}
		time.Sleep(86400)
	}
}
