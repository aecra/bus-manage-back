package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func main() {
	var err error
	db, err = sqlx.Open("mysql", "user:password@tcp(ipv4:port)/database?parseTime=True")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer db.Close()

	router := gin.Default()
	router.Use(Cors())

	// 查询某个车队下的司机基本信息
	router.GET("/Driver", getDriver)
	// 查询某名司机在某个时间段的违章详细信息
	router.GET("/DriverViolation", getDriverViolation)
	// 查询某个车队在某个时间段的违章统计信息
	router.GET("/FleetViolation", getFleetViolation)
	// 录入司机的基本信息
	router.POST("/Driver", postDriver)
	// 录入汽车基本信息
	router.POST("/Bus", postBus)
	// 录入司机的违章信息
	router.POST("/DriverViolation", postDriverViolation)

	// 辅助查询
	router.GET("/AFleet", getAFleet)
	router.GET("/ALine", getALine)
	router.GET("/AStation", getAStation)
	router.GET("/AViolationKind", getAViolationKind)
	router.GET("/ABus", getABus)
	router.GET("/AStaff", getAStaff)
	// 启动服务
	router.Run(":7814")
}
