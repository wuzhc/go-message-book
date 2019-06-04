package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Comment_20190601_222646 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Comment_20190601_222646{}
	m.Created = "20190601_222646"

	migration.Register("Comment_20190601_222646", m)
}

// Run the migrations
func (m *Comment_20190601_222646) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.CreateTable("tb_comment", "InnoDB", "utf8")
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("int(11)").SetUnsigned(true)
	m.NewCol("topic_id").SetDataType("int(11)").SetNullable(false).SetUnsigned(true)
	m.NewCol("user_name").SetDataType("varchar(50)").SetNullable(false)
	m.NewCol("create_time").SetDataType("datetime")
	m.NewCol("text").SetDataType("text")
	sql := m.GetSQL()
	m.SQL(sql)
}

// Reverse the migrations
func (m *Comment_20190601_222646) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("drop table tb_comment")
}
