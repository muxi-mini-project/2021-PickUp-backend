package router

import (
	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()

	//注册
	r.POST("/users", AddNewUsers)

	//登录
	r.POST("/users/login", LoginUser)

	//用户主页及其信息修改
	r.GET("/users/{user_id}")
	r.PUT("/users/{user_id}", UpdateUser)

	//用户修改密码
	r.PUT("/users/password", UpdatePassword)

	//用户评价
	r.PUT("/users/comment/{comment_id}", UsersComment)
	r.GET("/users/comment/{comment_id}")

	//司机订单
	r.POST("/driver", AddDriverRequirement)
	r.GET("/driver")
	r.DELETE("/driver", DeleteDriverRequirement)

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
