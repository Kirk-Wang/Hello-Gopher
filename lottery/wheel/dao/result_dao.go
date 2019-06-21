package dao

import (
	"log"

	"github.com/Kirk-Wang/Hello-Gopher/lottery/wheel/models"
	"github.com/go-xorm/xorm"
)

type ResultDao struct {
	engine *xorm.Engine
}

func NewResultDao(engine *xorm.Engine) *ResultDao {
	return &ResultDao{
		engine: engine,
	}
}

func (d *ResultDao) Get(id int) *models.LtResult {
	data := &models.LtResult{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *ResultDao) GetAll() []models.LtResult {
	datalist := make([]models.LtResult, 0)
	err := d.engine.
		Desc("id").
		Find(&datalist)
	if err != nil {
		log.Println("result_dao.GetAll error=", err)
		return datalist
	}
	return datalist
}

func (d *ResultDao) CountAll() int64 {
	num, err := d.engine.Count(&models.LtResult{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *ResultDao) Delete(id int) error {
	// 软删除
	data := &models.LtResult{
		Id:        id,
		SysStatus: 1,
	}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

func (d *ResultDao) Update(data *models.LtResult, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *ResultDao) Create(data *models.LtResult) error {
	_, err := d.engine.Insert(data)
	return err
}
