package repository

import (
	"github.com/astaxie/beego/orm"
	"chartmiru/models"
)

func init() {
	orm.RegisterModel(
        new(models.Company),
	)
}

func GetCompanies() ([]*models.Company, error) {
	var companies []*models.Company	
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").
	From("company")
	sql := qb.String()
	o := orm.NewOrm()
	_, err := o.Raw(sql).QueryRows(&companies)
	return companies, err
}

func GetCompany(id int) (models.Company, error) {
	var company models.Company
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").
	From("company").
	Where("id = ?")
	sql := qb.String()
	o := orm.NewOrm()
	err := o.Raw(sql, id).QueryRow(&company)
	return company, err	
}

func GetCompanyWithIds(ids []*int) ([]*models.Company, error){
	qb, _ := orm.NewQueryBuilder("mysql")
	var companies []*models.Company	
	qb.Select("*").
	From("company").
	Where("id").
	In("?")
	sql := qb.String()
	o := orm.NewOrm()
	_, err := o.Raw(sql).QueryRows(&companies)
	return companies, err	
}
