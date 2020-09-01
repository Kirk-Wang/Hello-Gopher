package dao

import (
	"log"

	"github.com/Kirk-Wang/Hello-Gopher/lottery/wheel/models"
	"github.com/go-xorm/xorm"
)

type CodeDao struct {
	engine *xorm.Engine
}

func NewCodeDao(engine *xorm.Engine) *CodeDao {
	return &CodeDao{
		engine: engine,
	}
}

func (d *CodeDao) Get(id int) *models.LtCode {
	data := &models.LtCode{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *CodeDao) GetAll() []models.LtCode {
	datalist := make([]models.LtCode, 0)
	err := d.engine.
		Desc("id").
		Find(&datalist)
	if err != nil {
		log.Println("code_dao.GetAll error=", err)
		return datalist
	}
	return datalist
}

func (d *CodeDao) CountAll() int64 {
	num, err := d.engine.Count(&models.LtCode{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *CodeDao) Delete(id int) error {
	// 软删除
	data := &models.LtCode{
		Id:        id,
		SysStatus: 1,
	}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

func (d *CodeDao) Update(data *models.LtCode, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *CodeDao) Create(data *models.LtCode) error {
	_, err := d.engine.Insert(data)
	return err
}
