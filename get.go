package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func getDriver(c *gin.Context) {
	var input DriverInput
	if c.ShouldBind(&input) == nil {
		var drivers []Driver
		switch {
		case input.Line != 0:
			err := db.Select(&drivers, "select * from driver where fleet_id=?;", input.Line)
			if err != nil {
				c.JSON(500, gin.H{"error": "Database operation failed"})
				return
			}
			c.JSON(200, drivers)
			return
		case input.Fleet != 0:
			err := db.Select(&drivers, "select * from driver where line_id=?;", input.Fleet)
			if err != nil {
				c.JSON(500, gin.H{"error": "Database operation failed"})
				return
			}
			c.JSON(200, drivers)
			return
		}
	}
	c.JSON(404, gin.H{"error": "invalid parameter"})
}

func getDriverViolation(c *gin.Context) {
	var input DriverViolationInput
	if c.ShouldBind(&input) == nil {
		var driverViolation []DriverViolation
		switch {
		case input.Number != "":
			sqlStatement := `select * from driver_violation 
			where number=? and violation_time between ? and ?;`
			err := db.Select(&driverViolation, sqlStatement, input.Number,
				input.Start.Format("2006-01-02 ::"), input.End.Format("2006-01-02 ::"))
			if err != nil {
				c.JSON(500, gin.H{"error": "Database operation failed"})
				return
			}
			c.JSON(200, driverViolation)
			return
		case input.Driver != 0:
			sqlStatement := `select * from driver_violation 
			where id=? and violation_time between ? and ?;`
			err := db.Select(&driverViolation, sqlStatement, input.Driver,
				input.Start.Format("2006-01-02 ::"), input.End.Format("2006-01-02 ::"))
			if err != nil {
				c.JSON(500, gin.H{"error": "Database operation failed"})
				return
			}
			c.JSON(200, driverViolation)
			return
		}
	}
	c.JSON(404, gin.H{"error": "invalid parameter"})
}

func getFleetViolation(c *gin.Context) {
	var input FleetViolationInput
	if c.ShouldBind(&input) == nil {
		var fleetViolation []FleetViolation
		err := db.Select(&fleetViolation, `SELECT l.fleet AS id,
    count(l.fleet) AS number,
    vk.name AS violation
		FROM violation AS v
    	JOIN staff AS s ON s.id = v.liable_person
    	JOIN line_staff AS ls ON ls.staff = s.id
    	JOIN line AS l ON l.id = ls.line
    	JOIN violation_kind AS vk ON vk.id = v.kind
			where l.fleet=? and v.violation_time between ? and ?
		GROUP BY v.kind;`, input.Fleet,
			input.Start.Format("2006-01-02 ::"), input.End.Format("2006-01-02 ::"))
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(500, gin.H{"error": "Database operation failed"})
			return
		}
		c.JSON(200, fleetViolation)
		return
	}
	c.JSON(404, gin.H{"error": "invalid parameter"})
}

func getAFleet(c *gin.Context) {
	var aFleet []AFleet
	err := db.Select(&aFleet, "select id,name from fleet;")
	if err != nil {
		c.JSON(500, gin.H{"error": "Database operation failed"})
		return
	}
	c.JSON(200, aFleet)
}

func getALine(c *gin.Context) {
	var aLineInput ALineInput
	if c.ShouldBindQuery(&aLineInput) == nil {
		var aLine []ALine
		var err error
		switch {
		case aLineInput.Fleet != 0:
			err = db.Select(&aLine, "select id,name from line where fleet=?;", aLineInput.Fleet)
			if err != nil {
				c.JSON(500, gin.H{"error": "Database operation failed"})
				return
			}
			c.JSON(200, aLine)
			return
		default:
			err = db.Select(&aLine, "select id,name from line;")
			if err != nil {
				c.JSON(500, gin.H{"error": "Database operation failed"})
				return
			}
			c.JSON(200, aLine)
			return
		}
	}
	c.JSON(404, gin.H{"error": "invalid parameter"})
}

func getAStation(c *gin.Context) {
	var aStationInput AStationInput
	if c.ShouldBindQuery(&aStationInput) == nil {
		var aStation []AStation
		var err error
		switch {
		case aStationInput.Line != 0:
			err = db.Select(&aStation, `select s.id as id,s.name as name 
			from station as s 
			join line_station as ls on ls.station=s.id 
			where ls.line=?`, aStationInput.Line)
			if err != nil {
				c.JSON(500, gin.H{"error": "Database operation failed"})
				return
			}
			c.JSON(200, aStation)
			return
		case aStationInput.Fleet != 0:
			err = db.Select(&aStation, `select s.id as id,s.name as name 
			from station as s 
			join line_station as ls on ls.station=s.id 
			join line as l on l.id=ls.line 
			where l.fleet=?`, aStationInput.Fleet)
			if err != nil {
				c.JSON(500, gin.H{"error": "Database operation failed"})
				return
			}
			c.JSON(200, aStation)
			return
		default:
			err = db.Select(&aStation, "select id,name from station;")
			if err != nil {
				c.JSON(500, gin.H{"error": "Database operation failed"})
				return
			}
			c.JSON(200, aStation)
			return
		}
	}
	c.JSON(404, gin.H{"error": "invalid parameter"})
}

func getAViolationKind(c *gin.Context) {
	var aViolationKind []AViolationKind
	err := db.Select(&aViolationKind, "select id,name from violation_kind;")
	if err != nil {
		c.JSON(500, gin.H{"error": "Database operation failed"})
		return
	}
	c.JSON(200, aViolationKind)
}

func getABus(c *gin.Context) {
	var aBusInput ABusInput
	if c.ShouldBindQuery(&aBusInput) == nil {
		var aBus []ABus
		var err error
		switch {
		case aBusInput.Line != 0:
			err = db.Select(&aBus, `select id,plate from bus where line=?`,
				aBusInput.Line)
			if err != nil {
				c.JSON(500, gin.H{"error": "Database operation failed"})
				return
			}
			c.JSON(200, aBus)
			return
		case aBusInput.Fleet != 0:
			err = db.Select(&aBus, `select b.id as id,b.plate as plate 
			from bus as b 
			join line as l on l.id=b.line 
			where l.fleet=?`, aBusInput.Fleet)
			if err != nil {
				fmt.Println(err)
				c.JSON(500, gin.H{"error": "Database operation failed"})
				return
			}
			c.JSON(200, aBus)
			return
		default:
			err = db.Select(&aBus, "select id,plate from bus;")
			if err != nil {
				c.JSON(500, gin.H{"error": "Database operation failed"})
				return
			}
			c.JSON(200, aBus)
			return
		}
	}
	c.JSON(404, gin.H{"error": "invalid parameter"})
}

func getAStaff(c *gin.Context) {
	var aVStaff []AStaff
	err := db.Select(&aVStaff, "select id,concat(name, '(',number, ')') as name from staff;")
	if err != nil {
		c.JSON(500, gin.H{"error": "Database operation failed"})
		return
	}
	c.JSON(200, aVStaff)
}
