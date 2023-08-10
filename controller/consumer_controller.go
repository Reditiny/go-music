package controller

import (
	"github.com/gin-gonic/gin"
	"go-music/dto"
	"go-music/model"
	"go-music/response"
	"go-music/service"
	"strconv"
)

type IConsumerController interface {
	Login(c *gin.Context)
	Add(c *gin.Context)
	GetById(c *gin.Context)
}

type ConsumerController struct {
	consumerService service.IConsumerService
}

func NewConsumerController() IConsumerController {
	return ConsumerController{consumerService: service.NewConsumerService()}
}

func (CC ConsumerController) Login(c *gin.Context) {
	var user dto.LoginDto
	if err := c.ShouldBind(&user); err != nil {
		response.Error(c, nil, err.Error())
		return
	}
	consumer := CC.consumerService.Login(user)
	if len(consumer) != 0 {
		response.Success(c, consumer, "登陆成功")
	} else {
		response.Fail(c, nil, "用户名或密码不存在")
	}
}

func (CC ConsumerController) Add(c *gin.Context) {
	var consumer model.Consumer
	if err := c.ShouldBind(&consumer); err != nil {
		response.Error(c, nil, err.Error())
	}
	success := CC.consumerService.Add(consumer)
	if success {
		response.Success(c, nil, "注册成功")
	} else {
		response.Fail(c, nil, "用户名已存在")
	}
}

func (CC ConsumerController) GetById(c *gin.Context) {
	idStr := c.Query("id")
	if idStr == "" {
		response.Fail(c, nil, "参数缺失")
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Fail(c, nil, "参数错误")
		return
	}
	data := CC.consumerService.GetById(id)
	response.Success(c, data, "请求成功")
}
