package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Topic_20190601_222641 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Topic_20190601_222641{}
	m.Created = "20190601_222641"

	migration.Register("Topic_20190601_222641", m)
}

// Run the migrations
func (m *Topic_20190601_222641) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.CreateTable("tb_topic", "InnoDB", "utf8")
	m.PriCol("id").SetAuto(true).SetDataType("int(11)").SetNullable(false).SetUnsigned(true)
	m.NewCol("name").SetDataType("varchar(50)").SetNullable(false)
	sql := m.GetSQL()
	m.SQL(sql)
}

// Reverse the migrations
func (m *Topic_20190601_222641) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("drop table tb_topic")
}
