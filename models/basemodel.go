/*
Author: Aosen
Data: 2016-01-11
QQ: 316052486
Desc: models的基类，所有其他model都继承此类, 使用beego提供的orm
*/
package models

import (
	"log"
	"sort"
	"time"

	"github.com/aosen/goutils"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//服务器配置信息
type System struct {
	Id int
	K  string `orm:"size(50);unique"`
	V  string `orm:"size(100)"`
}

//用户信息表
type Userinfo struct {
	Id int
	//用户的唯一识别码
	Userauth   string    `orm:"size(30);unique"`
	Createtime time.Time `orm:"type(date)"`
}

//用户日至，只记录用户当天首次登陆
type Userlog struct {
	Id int
	//用户ID
	Userid int
	//登陆时间
	Logintime time.Time `orm:"type(date)"`
}

//一级分类表
type First struct {
	Id         int
	Firstname  string    `orm:"size(20);unique"`
	Updatetime time.Time `orm:"type(date)"`
	Createtime time.Time `orm:"type(date)"`
}

//二级分类表
type Second struct {
	Id         int
	Firstid    int
	Secondname string    `orm:"size(20);unique"`
	Updatetime time.Time `orm:"type(date)"`
	Createtime time.Time `orm:"type(date)"`
}

//小说简介信息
type Novel struct {
	Id           int
	Title        string    `orm:"size(200)"`
	Firstid      int       `orm:"index"`
	Secondid     int       `orm:"index"`
	Author       string    `orm:"size(50);index"`
	Introduction string    `orm:"type(text)"`
	Picture      string    `orm:"size(200)"`
	Novelsource  string    `orm:"size(200);unique"`
	Novelpv      int       `orm:"default(0)"`
	Novelcollect int       `orm:"default(0)"`
	Createtime   time.Time `orm:"type(date)"`
}

//根据novel结构体列表中的novelpv进行排序
type NovelsPv []*Novel

func (self NovelsPv) Len() int {
	return len(self)
}

func (self NovelsPv) Less(i, j int) bool {
	return self[i].Novelpv > self[j].Novelpv
}

func (self NovelsPv) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func NovelPvSort(nsp NovelsPv) {
	sort.Sort(nsp)
}

//根据novel结构体列表中的novelcollect进行排序
type NovelsCollect []*Novel

func (self NovelsCollect) Len() int {
	return len(self)
}

func (self NovelsCollect) Less(i, j int) bool {
	return self[i].Novelcollect > self[j].Novelcollect
}

func (self NovelsCollect) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func NovelCollectSort(nsc NovelsCollect) {
	sort.Sort(nsc)
}

//小说内容表
type Content struct {
	Id            int
	Novelid       int
	Title         string    `orm:"size(200);index"`
	Firstid       int       `orm:"index"`
	Secondid      int       `orm:"index"`
	Chapter       int       `orm:"index"`
	Subtitle      string    `orm:"size(200);index"`
	Text          string    `orm:"type(text)"`
	Contentsource string    `orm:"size(200);index"`
	Createtime    time.Time `orm:"type(date)"`
}

//小说收藏量排名信息
type Collectrank struct {
	Id int
	//小说ID
	Novelid int
	//小说一级分类ID
	Firstid int
	//小说二级分类ID
	Secondid int
	//小说pv
	Novelpv int
	//小说收藏量
	Novelcollect int
	Createtime   time.Time `orm:"type(date)"`
}

//小说点击量排名信息
type Clickrank struct {
	Id int
	//小说ID
	Novelid int
	//小说一级分类ID
	Firstid int
	//小说二级分类ID
	Secondid int
	//小说pv
	Novelpv int
	//小说收藏量
	Novelcollect int
	Createtime   time.Time `orm:"type(date)"`
}

//小说推荐列表 json格式的
type Recommendlist struct {
	Id            int
	Recommendlist string    `orm:"type(text)"`
	Updatetime    time.Time `orm:"type(date)"`
	Createtime    time.Time `orm:"type(date)"`
}

func init() {
	//读取配置文件信息
	settings := func(path string) map[string]string {
		return goutils.NewConfig().Load(path).GlobalContent()
	}("conf/app.conf")
	//获取settings中的信息
	dbinfo, ok := settings["DBINFO"]
	if !ok {
		log.Fatal("not found DBINFO in config file")
	}
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", dbinfo)
	orm.RegisterModel(
		new(First),
		new(Second),
		new(System),
		new(Userinfo),
		new(Userlog),
		new(Novel),
		new(Content),
		new(Collectrank),
		new(Clickrank),
		new(Recommendlist),
	)
}

type BaseModel struct {
}

func NewBaseModel() *BaseModel {
	return &BaseModel{}
}

func (self *BaseModel) GetAllNovel() ([]*Novel, error) {
	var novels []*Novel
	o := orm.NewOrm()
	//获取小说列表
	if _, err := o.QueryTable("novel").Limit(-1).All(&novels, "Id", "Firstid", "Secondid", "Novelpv", "Novelcollect"); err != nil {
		return nil, err
	}
	return novels, nil
}
