package models

import (
	// "strconv"
	"time"
)
type Company struct {
	Id       uint64 
	Name     string `orm:"size(100)"`
	Market   string `orm:"size(50)"`
	Initialized bool
	Created_at time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
}
