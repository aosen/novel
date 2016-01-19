/*
Author: Aosen
Data: 2016-01-12
QQ: 316052486
Desc: 所有的模型定义都在这里
*/

package models

import "time"

//使用web api的app信息
type Application struct {
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
	Createtime   time.Time `orm:"type(date)"`
}

//小说排名信息
type Novelrank struct {
}
