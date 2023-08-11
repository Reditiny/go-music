package controller

import (
	"github.com/gin-gonic/gin"
	"go-music/constant"
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
	Update(c *gin.Context)
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
		response.Success(c, consumer, constant.LOGIN_SUCCESS)
	} else {
		response.Fail(c, nil, constant.LOGIN_FAIL)
	}
}

func (CC ConsumerController) Add(c *gin.Context) {
	var consumer model.Consumer
	if err := c.ShouldBind(&consumer); err != nil {
		response.Error(c, nil, err.Error())
	}
	success := CC.consumerService.Add(consumer)
	if success {
		response.Success(c, nil, constant.REGISTER_SUCCESS)
	} else {
		response.Fail(c, nil, constant.USER_EXIST)
	}
}

func (CC ConsumerController) GetById(c *gin.Context) {
	idStr := c.Query("id")
	if idStr == "" {
		response.Fail(c, nil, constant.PARAM_FAIL_GET)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Fail(c, nil, constant.PARAM_FAIL_PARSE)
		return
	}
	data := CC.consumerService.GetById(id)
	response.Success(c, data, constant.REQUSEST_SUCCESS)
}

func (CC ConsumerController) Update(c *gin.Context) {
	var user model.Consumer
	err := c.ShouldBind(&user)
	if err != nil {
		response.Fail(c, nil, constant.PARAM_FAIL_PARSE)
		return
	}
	// TODO 更新用户信息
}
