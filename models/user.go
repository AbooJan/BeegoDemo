package models

type User struct {
	ID     int64
	Name   string 	`json:"name"`
	Age    int  	`json:"age"`
	Gender int   	`json:"gender"`
	Height float64 	`json:"height"`
}
