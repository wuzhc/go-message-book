package controllers

import (
	"fmt"
	"my-web/models"
	"time"
)

// "github.com/astaxie/beego"

// HomeController operations for Home
type HomeController struct {
	baseController
}

// URLMapping ...
func (c *HomeController) URLMapping() {
	c.Mapping("Index", c.Index)
}

// 留言板首页
func (c *HomeController) Index() {
	c.Data["name"] = "添加留言"

	// 主题列表
	topics, err := models.GetAllTopic(nil, nil, nil, nil, 0, 10)
	if err != nil {
		c.CustomAbort(500, err.Error())
	}

	c.Data["topics"] = topics
	c.Display()
}

// 留言板列表
func (c *HomeController) List() {
	c.Data["name"] = "留言列表"

	// 留言列表
	comments, err := models.GetAllComment(nil, nil, []string{"id"}, []string{"desc"}, 0, -1)
	if err != nil {
		c.CustomAbort(500, err.Error())
	}

	resp := make([]map[string]interface{}, 0)
	for _, v := range comments {
		temp := make(map[string]interface{})
		temp["id"] = v.(models.Comment).Id
		temp["userName"] = v.(models.Comment).UserName
		temp["createTime"] = v.(models.Comment).CreateTime
		temp["text"] = v.(models.Comment).Text
		tp, _ := models.GetTopicById(v.(models.Comment).TopicId)
		temp["topicName"] = tp.Name
		resp = append(resp, temp)
	}

	c.Data["comments"] = resp
	c.Display()
}

// 保存测试
func (c *HomeController) Save() {
	text := c.GetString("text")
	if len(text) == 0 {
		c.RspParamError("text参数错误")
	}

	topicID, _ := c.GetInt64("topic_id")
	if topicID == 0 {
		c.RspParamError("topic_id参数错误")
	}

	userName := c.GetString("user_name")
	if len(userName) == 0 {
		c.RspParamError("user_name参数错误")
	}

	comment := models.Comment{
		Text:       text,
		TopicId:    topicID,
		UserName:   userName,
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
	}

	_, err := models.AddComment(&comment)
	if err != nil {
		c.RspFail(err.Error())
	} else {
		c.RspSuccess("保存成功")
	}
}

// 删除回复
func (c *HomeController) Delete() {
	id, err := c.GetInt64("id")
	if err != nil {
		c.RspParamError(err.Error())
	}

	if id == 0 {
		c.RspParamError("id不能为空")
	}

	err = models.DeleteComment(id)
	if err == nil {
		c.RspSuccess("删除成功")
	} else {
		c.RspFail("删除失败,原因为:" + err.Error())
	}
}

// 获取
func (c *HomeController) Get() {
	params := c.Ctx.Input.Params()
	fmt.Println(params)
	fmt.Println("--------------------------", c.Ctx.Input.Param(":id"))
	// fmt.Println("--------------------------", c.Ctx.Input.Param(":user"))

	id, _ := c.GetInt64("id")
	topic, _ := models.GetTopicById(id)
	data := make(map[string]interface{})
	data["id"] = topic.Id
	data["name"] = topic.Name
	c.RspData(data)
}

func (c *HomeController) GetAll() {
	data, _ := models.GetAllTopic(nil, nil, nil, nil, 0, 10)
	for _, v := range data {
		c.Ctx.WriteString(v.(models.Topic).Name)
		c.Ctx.WriteString("<br>")
	}
}
