package controllers

import (
	repository_stock "chartmiru/repository/stock"
	repository_company "chartmiru/repository/company"
	"chartmiru/models"
	// "encoding/json"

	"github.com/astaxie/beego"
)

// Operations about Stocks
type StockController struct {
	beego.Controller
	response models.Response
}


// @Title Get
// @Description get all Stocks
// @Success 200 {object} models.Stock
// @router / [get]
func (s *StockController) Get() {
	companies, _ := repository_company.GetCompanies()
	var companyIds []uint64
	for _, company := range companies {
		companyIds = append(companyIds, company.Id)
	}
	stocks, _ := s.getStockList(companyIds)


	// 会社ID毎にデータをまとめる
	for _, stock := range stocks {
		for _, company := range companies {
			if stock.Company_id == company.Id {
				company.Stocks = append(company.Stocks, stock) 
			}
		}
	}

	// 会社の情報を付与する
	// for companyId, _ := range stocksPerCompany {
	// 	company, err := repository_company.GetCompany(companyId)
	// 	stocksPerCompany[companyId] = append(stocksPerCompany[companyId], company)
	// }

	// s.response.Data    = stocksPerCompany
	s.response.Code    = 200
	s.Data["json"] = companies
	s.ServeJSON()
}

func (s *StockController) getStockList(company_ids []uint64) ([]*models.Stock, error) {
	stocks, err := repository_stock.GetStockWithIds(company_ids)
	if err != nil {
		s.Abort("Db")
	}
	
	return stocks, err
}
