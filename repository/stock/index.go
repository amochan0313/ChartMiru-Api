package repository

import (
	"github.com/astaxie/beego/orm"
	"chartmiru/models"
	"strconv"
	"strings"
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


func GetStockWithIds(ids []uint64) ([]*models.Stock, error) {
	var castedIds []string
    for _, id := range ids {
        castedIds = append(castedIds, strconv.FormatUint(id, 10))
    }
	formatedIds := strings.Join(castedIds, ",")
	qb, _ := orm.NewQueryBuilder("mysql")
	var stocks []*models.Stock	
	qb.Select("*").
	From("stock").
	Where("company_id").
	In(formatedIds)

	sql := qb.String()
	o := orm.NewOrm()

	_, err := o.Raw(sql).QueryRows(&stocks)
	return stocks, err	
}
