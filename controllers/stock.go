package controllers

import (
	"chartmiru/repository/stock"
	// "encoding/json"

	"github.com/astaxie/beego"
)

// Operations about Stocks
type StockController struct {
	beego.Controller
}


// @Title Get
// @Description get all Stocks
// @Success 200 {object} models.Stock
// @router / [get]
func (s *StockController) Get() {
	
	stocks, _ := s.getStockList()
	s.response.Data    = stocks
	s.response.Code    = 200
	s.Data["json"] = s.response
	s.ServeJSON()
}

func (s *StockController) getStockList() (models.Stock, error) {
	stocks, err := repository.GetStocks()

	if err != nil {
		c.Abort("Db")
	}
	
	return stocks, err
}
