package controllers

import (
    "github.com/astaxie/beego"
    "chartmiru/models"
)

type ErrorController struct {
    beego.Controller
    response models.Response
}

func (e *ErrorController) ErrorCompanyNotFound() {
    e.response.Message = "Company Not Found"
    e.response.Data    = nil
    e.response.Code    = 404002
    e.Data["json"]     = e.response
    e.ServeJSON()
}

func (e *ErrorController) ErrorDb() {
    e.response.Message = "database error"
    e.response.Data    = nil
    e.response.Code    = 500001
    e.Data["json"]     = e.response
    e.ServeJSON()
}

func (e *ErrorController) ErrorStockNotFound() {
    e.response.Message = "Stock Not Found"
    e.response.Data    = nil
    e.response.Code    = 404001
    e.Data["json"]     = e.response
    e.ServeJSON()
}
