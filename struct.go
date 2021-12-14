package main

import (
	"database/sql"
	"time"
)

type DriverInput struct {
	Fleet int32 `form:"fleet"`
	Line  int32 `form:"line"`
}

type Driver struct {
	Id            sql.NullInt32   `db:"id" json:"id"`
	FleetId       sql.NullInt32   `db:"fleet_id" json:"fleetId"`
	Fleet         sql.NullString  `db:"fleet" json:"fleet"`
	LineId        sql.NullInt32   `db:"line_id" json:"lineId"`
	Line          sql.NullString  `db:"line" json:"line"`
	Number        sql.NullString  `db:"number" json:"number"`
	Name          sql.NullString  `db:"name" json:"name"`
	Sex           sql.NullString  `db:"sex" json:"sex"`
	NativePlace   sql.NullString  `db:"native_place" json:"nativePlace"`
	IdNUmber      sql.NullString  `db:"id_number" json:"idNumber"`
	Phone         sql.NullString  `db:"phone" json:"phone"`
	Positon       sql.NullString  `db:"position" json:"position"`
	Wages         sql.NullFloat64 `db:"wages" json:"wages"`
	Office        sql.NullString  `db:"office" json:"office"`
	EntryTime     sql.NullTime    `db:"entry_time" json:"entryTime"`
	DepartureTime sql.NullTime    `db:"departure_time" json:"departureTime"`
}

type DriverInsert struct {
	Number      string    `db:"number" form:"number"`
	Name        string    `db:"name" form:"name"`
	Sex         string    `db:"sex" form:"sex"`
	NativePlace string    `db:"native_place" form:"nativePlace"`
	IdNUmber    string    `db:"id_number" form:"idNumber"`
	Phone       string    `db:"phone" form:"phone"`
	Positon     string    `db:"position" form:"position"`
	Wages       float64   `db:"wages" form:"wages"`
	Office      string    `db:"office" form:"office"`
	EntryTime   time.Time `db:"entry_time" form:"entryTime"`
	LineId      int32     `db:"line_id" form:"lineId"`
}

type DriverViolationInput struct {
	Driver int32     `form:"driver"`
	Number string    `form:"number"`
	Start  time.Time `form:"start"`
	End    time.Time `form:"end"`
}

type DriverViolation struct {
	Id            sql.NullInt32  `db:"id" json:"id"`
	Number        sql.NullString `db:"number" json:"number"`
	Name          sql.NullString `db:"name" json:"name"`
	Fleet         sql.NullString `db:"fleet" json:"fleet"`
	Line          sql.NullString `db:"line" json:"line"`
	Station       sql.NullString `db:"station" json:"station"`
	BusPlate      sql.NullString `db:"bus_plate" json:"busPlate"`
	Violation     sql.NullString `db:"violation" json:"violation"`
	ViolationTime sql.NullTime   `db:"violation_time" json:"violationTime"`
	InputBy       sql.NullString `db:"input_by" json:"inputBy"`
	InputTime     sql.NullTime   `db:"input_time" json:"inputTime"`
}

type DriverViolationInsert struct {
	LiablePerson  int32     `form:"liablePerson"`
	Bus           int32     `form:"bus"`
	Station       int32     `form:"station"`
	Violation     int32     `form:"violation"`
	ViolationTime time.Time `form:"violationTime"`
	InputBy       int32     `form:"inputBy"`
	InputTime     time.Time `orm:"inputTime"`
}

type FleetViolationInput struct {
	Fleet int32     `form:"fleet"`
	Start time.Time `form:"start"`
	End   time.Time `form:"end"`
}

type FleetViolation struct {
	Id        sql.NullInt32  `db:"id" json:"id"`
	Number    sql.NullInt32  `db:"number" json:"number"`
	Violation sql.NullString `db:"violation" json:"violation"`
}

type BusInsert struct {
	Model   string    `form:"model"`
	People  int32     `form:"people"`
	Color   string    `form:"color"`
	Plate   string    `form:"plate"`
	LineId  int32     `form:"lineId"`
	BuyTime time.Time `form:"buyTime"`
	BuyBy   int32     `form:"buyBy"`
}

type AFleet struct {
	Id   sql.NullInt32  `db:"id" json:"id"`
	Name sql.NullString `db:"name" json:"name"`
}

type ALineInput struct {
	Fleet int32 `form:"fleet"`
}

type ALine struct {
	Id   sql.NullInt32  `db:"id" json:"id"`
	Name sql.NullString `db:"name" json:"name"`
}

type AStationInput struct {
	Fleet int32 `form:"fleet"`
	Line  int32 `form:"line"`
}

type AStation struct {
	Id   sql.NullInt32  `db:"id" json:"id"`
	Name sql.NullString `db:"name" json:"name"`
}

type AViolationKind struct {
	Id   sql.NullInt32  `db:"id" json:"id"`
	Name sql.NullString `db:"name" json:"name"`
}

type ABusInput struct {
	Fleet int32 `form:"fleet"`
	Line  int32 `form:"line"`
}

type ABus struct {
	Id    sql.NullInt32  `db:"id" json:"id"`
	Plate sql.NullString `db:"plate" json:"plate"`
}

type AStaff struct {
	Id   sql.NullInt32  `db:"id" json:"id"`
	Name sql.NullString `db:"name" json:"name"`
}
