package router

import (
	"github.com/gin-gonic/gin"
)

func router() {
	r := gin.Default()

	//注册
	r.POST("/users", addNewUsers)

	//登录
	r.POST("/users/login", loginUser)

	//用户主页及其信息修改
	r.GET("/users/{user_id}")
	r.PUT("/users/{user_id}", updateUser)

	//用户修改密码
	r.PUT("/users/password", updatePassword)

	//用户评价
	r.PUT("/users/comment/{comment_id}", usersComment)
	r.GET("/users/comment/{comment_id}")

	//司机订单
	r.POST("/driver", addDriverRequirement)
	r.GET("/driver")
	r.DELETE("/driver", deleteDriverRequirement)

	//司机确认(注意乘客先确认,司机后确认)
	r.POST("​/driver​/confirm​/{confirm_id}", driverConfirm)

	//乘客订单
	r.POST("/passenger", addPassengerRequirement)
	r.GET("/passenger")
	r.DELETE("/passenger", deletePassengerRequirement)

	//乘客确认
	r.POST("/passenger/confirm/{confirm_id}", passengerConfirm)

	//常用路径
	r.POST("/route/{route_id}", addNewRoute)
	r.GET("/route/{route_id}")
	r.DELETE("/route/{route_id}", deleteRoute)

	r.Run()
}
