package models

import (
	// "strconv"
	"time"
)

func init() {
	
}

type Stock struct {
	Id       uint64 
	Company_id uint
	Open uint
	Close uint
	Volume uint
	High uint
	Low uint
	Data_date time.Time `orm:"type(date)"`
	Created_at time.Time `orm:"auto_now_add;type(datetime)"`
}
