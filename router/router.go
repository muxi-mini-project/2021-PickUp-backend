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
	r.POST("/pickup/users", user.UserCreate)

	//r.POST("/pickup/users/login", user.UserLogin)
	r.DELETE("pickup/users", user.UserDelete)

	//用户主页及其信息修改
	r.GET("/pickup/users", user.ViewUser)
	r.GET("pickup/users/comment", user.ViewUsersComment)
	r.PUT("/pickup/users", user.UpdateUser)

	//用户修改密码
	r.PUT("/pickup/users/password", user.UpdatePassword)

	//用户评价
	r.PUT("/pickup/users/comment", user.UsersComment)
	//r.GET("/users/comment/{comment_id}", handler.ViewComment)

	//司机订单
	r.POST("/pickup/driver", driver.AddDriverRequirement)
	r.GET("/pickup/driver/:route_id", driver.ViewDRequirement)
	r.GET("/pickup/driver", driver.ViewDriverRequirement)
	r.DELETE("/pickup/driver", driver.DeleteDriverRequirement)
	r.PUT("/pickup/driver/confirm", driver.DriverConfirm)

	//乘客订单
	r.POST("/pickup/passenger", passenger.AddPassengerRequirement)
	r.GET("/pickup/passenger/:route_id", passenger.ViewPRequirement)
	r.GET("/pickup/passenger", passenger.ViewPassengerRequirement)
	r.DELETE("/pickup/passenger", passenger.DeletePassengerRequirement)

	//乘客确认
	r.PUT("/pickup/passenger/confirm", passenger.PassengerConfirm)

	//获取乘客与司机的匹配度
	r.GET("/pickup/user/recommend", user.Recommend)

	//常用路径
	r.POST("/pickup/route", route.AddNewRoute)
	r.GET("/pickup/route", route.ViewRoute)
	r.DELETE("/pickup/route/:route_id", route.DeleteRoute)

	//举报功能
	r.PUT("/pickup/users/report/:report_id", user.Report)

}
