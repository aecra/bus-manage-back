package main

import (
	"github.com/gin-gonic/gin"
)

func postDriver(c *gin.Context) {
	var driver DriverInsert
	if c.ShouldBind(&driver) == nil {
		if res, err := CheckId("line", driver.LineId); !res || err != nil {
			c.JSON(404, gin.H{"error": "LineId is invalid"})
			return
		}

		result1, err := db.Exec(`INSERT INTO staff 
    (number,name,sex,native_place,id_number,phone,position,wages,office,entry_time) 
    VALUES(?,?,?,?,?,?,?,?,?,?)`,
			driver.Number, driver.Name, driver.Sex, driver.NativePlace, driver.IdNUmber,
			driver.Phone, driver.Positon, driver.Wages, driver.Office, driver.EntryTime)
		if err != nil {
			c.JSON(500, gin.H{"error": "Insert staff failed"})
			return
		}
		id, _ := result1.LastInsertId()
		db.Exec("INSERT INTO line_staff (line,staff) VALUES(?,?)", driver.LineId, id)

		c.JSON(200, gin.H{"error": ""})
		return
	}
	c.JSON(404, gin.H{"error": "invalid parameter"})
}

func postBus(c *gin.Context) {
	var bus BusInsert
	if c.ShouldBind(&bus) == nil {
		if res, err := CheckId("line", bus.LineId); !res || err != nil {
			c.JSON(404, gin.H{"error": "LineId is invalid"})
			return
		}
		if res, err := CheckId("staff", bus.BuyBy); !res || err != nil {
			c.JSON(500, gin.H{"error": "BuyBy is invalid"})
			return
		}

		_, err := db.Exec(`INSERT INTO bus 
    (model,people,color,plate,line,buy_time,buy_by) 
    VALUES(?,?,?,?,?,?,?)`,
			bus.Model, bus.People, bus.Color, bus.Plate, bus.LineId, bus.BuyTime, bus.BuyBy)
		if err != nil {
			c.JSON(500, gin.H{"error": "Insert staff failed"})
			return
		}

		c.JSON(200, gin.H{"error": ""})
		return
	}
	c.JSON(404, gin.H{"error": "invalid parameter"})
}

func postDriverViolation(c *gin.Context) {
	var dvi DriverViolationInsert
	if c.ShouldBind(&dvi) == nil {
		if res, err := CheckId("bus", dvi.Bus); !res || err != nil {
			c.JSON(404, gin.H{"error": "Bus is invalid"})
			return
		}
		if res, err := CheckId("station", dvi.Station); !res || err != nil {
			c.JSON(404, gin.H{"error": "Station is invalid"})
			return
		}
		if res, err := CheckId("staff", dvi.LiablePerson); !res || err != nil {
			c.JSON(404, gin.H{"error": "LiablePerson is invalid"})
			return
		}
		if res, err := CheckId("staff", dvi.InputBy); !res || err != nil {
			c.JSON(404, gin.H{"error": "InputBy is invalid"})
			return
		}

		_, err := db.Exec(`INSERT INTO violation 
    (liable_person,bus,station,violation_time,kind,input_by,input_time) 
    VALUES(?,?,?,?,?,?,?)`,
			dvi.LiablePerson, dvi.Bus, dvi.Station, dvi.ViolationTime, dvi.Violation, dvi.InputBy, dvi.InputTime)
		if err != nil {
			c.JSON(500, gin.H{"error": "Insert bus failed"})
			return
		}

		c.JSON(200, gin.H{"error": ""})
		return
	}
	c.JSON(404, gin.H{"error": "invalid parameter"})
}
