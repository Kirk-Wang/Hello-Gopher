package dao

import (
	"log"

	"github.com/Kirk-Wang/Hello-Gopher/lottery/wheel/models"
	"github.com/go-xorm/xorm"
)

type BlackipDao struct {
	engine *xorm.Engine
}

func NewBlackipDao(engine *xorm.Engine) *BlackipDao {
	return &BlackipDao{
		engine: engine,
	}
}

func (d *BlackipDao) Get(id int) *models.LtBlackip {
	data := &models.LtBlackip{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *BlackipDao) GetAll() []models.LtBlackip {
	datalist := make([]models.LtBlackip, 0)
	err := d.engine.
		Desc("id").
		Find(&datalist)
	if err != nil {
		log.Println("backip_dao.GetAll error=", err)
		return datalist
	}
	return datalist
}

func (d *BlackipDao) CountAll() int64 {
	num, err := d.engine.Count(&models.LtBlackip{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

// func (d *BlackipDao) Delete(id int) error {
// 	// 软删除
// 	data := &models.LtBlackip{
// 		Id:        id,
// 		SysStatus: 1,
// 	}
// 	_, err := d.engine.Id(data.Id).Update(data)
// 	return err
// }

func (d *BlackipDao) Update(data *models.LtBlackip, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *BlackipDao) Create(data *models.LtBlackip) error {
	_, err := d.engine.Insert(data)
	return err
}

func (d *BlackipDao) GetById(ip string) *models.LtBlackip {
	datalist := make([]models.LtBlackip, 0)
	err := d.engine.Where("ip=?", ip).
		Desc("id").
		Limit(1).
		Find(&datalist)

	if err != nil || len(datalist) < 1 {
		return nil
	} else {
		return &datalist[0]
	}
}
