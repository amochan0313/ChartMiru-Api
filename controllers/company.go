package controllers

import (
	"chartmiru/repository/company"
	// "github.com/astaxie/beego/logs"
	// "encoding/json"d
	// "strconv"

	_ "fmt"
	"chartmiru/models"
	"github.com/astaxie/beego"
)

// Operations about company
type CompanyController struct {
	beego.Controller
	response models.Response
}

// @Title Get
// @Description find company by companyid
// @Param id path int true "the companyid you want to get"
// @Success 200 {object} models.Company
// @router /company/:id [get]
func (c *CompanyController) Get() {
	// 単体取得
	if id, _ := c.GetInt(":id"); id != 0 {
		company, _ := c.getCompanyItem(id)
    	c.response.Data    = company
  	  	c.response.Code    = 200
		c.Data["json"] = c.response
		c.ServeJSON()
		return
	}

	// 複数取得
	companies, _ := repository.GetCompanies()
	c.response.Data    = companies
	c.response.Code    = 200
	c.Data["json"] = c.response
	c.ServeJSON()
}

func (c *CompanyController) getCompanyItem(companyId int) (models.Company, error) {
	company, err := repository.GetCompany(companyId)
	
	emptyCompany := models.Company{}
	if ; company == emptyCompany {
		c.Abort("CompanyNotFound")
	}

	if err != nil {
		c.Abort("Db")
	}
	
	return company, err
}

func (c *CompanyController) getCompanyList() ([]*models.Company, error) {
	companies, err := repository.GetCompanies()
	if err != nil {
		c.Abort("Db")
	}
	
	return companies, err
}
