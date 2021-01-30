package router

import (
	"pickup/handler"
	"pickup/handler/driver"
	"pickup/handler/passenger"
	"pickup/handler/route"
	"pickup/handler/user"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {

	r.GET("/text", handler.Text)

	//此函数为一站式登录函数,若用户登录成功(首次),则将其学号和密码保存起来创造一个新用户
	r.POST("/users", user.UserCreate)

	//登录
	r.POST("/users/login", user.UserLogin)

	//用户主页及其信息修改
	r.GET("/users/:uid", user.ViewUser)
	r.PUT("/users/:uid/update", user.UpdateUser)

	//用户修改密码
	r.PUT("/users/:uid/password", user.UpdatePassword)

	//用户评价
	r.PUT("/users/:uid/comment", user.UsersComment)
	//r.GET("/users/comment/{comment_id}", handler.ViewComment)

	//司机订单
	r.POST("/driver", driver.AddDriverRequirement)
	r.GET("/driver/route", driver.ViewDRequirement)
	r.GET("/driver", driver.ViewDriverRequirement)
	r.DELETE("/driver", driver.DeleteDriverRequirement)
	r.PUT("/driver/confirm", driver.DriverConfirm)

	//乘客订单
	r.POST("/passenger", passenger.AddPassengerRequirement)
	r.GET("/passenger/route", passenger.ViewPRequirement)
	r.GET("/passenger", passenger.ViewPassengerRequirement)
	r.DELETE("/passenger", passenger.DeletePassengerRequirement)

	//乘客确认
	r.PUT("/passenger/confirm", passenger.PassengerConfirm)

	//获取乘客与司机的匹配度
	r.GET("/user/match", user.MatchDegree)

	//常用路径
	r.POST("/route", route.AddNewRoute)
	r.GET("/route", route.ViewRoute)
	r.DELETE("/route", route.DeleteRoute)

}
