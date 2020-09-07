package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/unknwon/com"
	"os"
	"path"
	"strconv"
	"time"
)

const (
	dbName        = "data/beeblog.db"
	sqlite3Driver = "sqlite3"
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index;auto_now_add"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index;auto_now"`
	TopicCount      int64
	TopicLastUserId int64
}

type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index;auto_now_add"`
	Updated         time.Time `orm:"index;auto_now"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index;auto_now"`
	ReplyCount      int64
	ReplyLastUserId int64
}

func RegisterDB() {
	// 检查数据库文件
	if !com.IsExist(dbName) {
		os.MkdirAll(path.Dir(dbName), os.ModePerm)
		os.Create(dbName)
	}

	// 注册模型
	orm.RegisterModel(new(Category), new(Topic))
	// 注册驱动("sqlite3"属于默认注册，此处代码可省略)
	orm.RegisterDriver(sqlite3Driver, orm.DRSqlite)
	// 注册默认数据库
	orm.RegisterDataBase("default", sqlite3Driver, dbName, 10)

}

func AddCategory(name string) error {
	o := orm.NewOrm()

	cate := &Category{Title: name}

	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return err
	}

	cate.Created = time.Now()
	cate.TopicTime = time.Now()
	_, err = o.Insert(cate)
	if err != nil {
		return err
	}

	return nil
}

func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	cate := &Category{Id: cid}
	_, err = o.Delete(cate)

	return err
}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()

	cates := make([]*Category, 0)

	qs := o.QueryTable("category")
	_, err := qs.All(&cates)

	return cates, err
}

func AddTopic(title, content string) error {
	o := orm.NewOrm()

	topic := &Topic{
		Title:   title,
		Content: content,
	}

	_, err := o.Insert(topic)

	return err
}

func GetAllTopics(isDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()

	topics := make([]*Topic, 0)

	qs := o.QueryTable("topic")

	var err error
	if isDesc {
		_, err = qs.OrderBy("-created").All(&topics)
	} else {
		_, err = qs.All(&topics)

	}

	return topics, err
}

func GetTopic(tid string) (*Topic, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	o := orm.NewOrm()

	topic := new(Topic)

	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return nil, err
	}

	topic.Views++
	_, err = o.Update(topic)

	return topic, err
}

func ModifyTopic(tid, title, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	topic := &Topic{Id: tidNum}

	err = o.Read(topic)
	if err == nil {
		topic.Title = title
		topic.Content = content
		_, err = o.Update(topic)
	}
	return err
}

func DeleteTopic(tid string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	topic := &Topic{Id: tidNum}
	_, err = o.Delete(topic)
	return err
}
