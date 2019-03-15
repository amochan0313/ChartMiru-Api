package repository

import (
	"github.com/astaxie/beego/orm"
	"chartmiru/models"

)

func init() {
	orm.RegisterModel(
        new(models.Stock),
	)
}

func GetStocks() ([]*models.Stock, error) {
	var stocks []*models.Stock
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").
	From("stock")
	sql := qb.String()
	o := orm.NewOrm()
	_, err := o.Raw(sql).QueryRows(&stocks)
	return stocks, err
}
