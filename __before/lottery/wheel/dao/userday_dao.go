package dao

import (
	"log"

	"github.com/Kirk-Wang/Hello-Gopher/lottery/wheel/models"
	"github.com/go-xorm/xorm"
)

type UserDayDao struct {
	engine *xorm.Engine
}

func NewUserDayDao(engine *xorm.Engine) *UserDayDao {
	return &UserDayDao{
		engine: engine,
	}
}

func (d *UserDayDao) Get(id int) *models.LtUserday {
	data := &models.LtUserday{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *UserDayDao) GetAll() []models.LtUserday {
	datalist := make([]models.LtUserday, 0)
	err := d.engine.
		Desc("id").
		Find(&datalist)
	if err != nil {
		log.Println("user_dao.GetAll error=", err)
		return datalist
	}
	return datalist
}

func (d *UserDayDao) CountAll() int64 {
	num, err := d.engine.Count(&models.LtUserday{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

// func (d *UserDayDao) Delete(id int) error {
// 	// 软删除
// 	data := &models.LtUserday{
// 		Id:        id,
// 		SysStatus: 1,
// 	}
// 	_, err := d.engine.Id(data.Id).Update(data)
// 	return err
// }

func (d *UserDayDao) Update(data *models.LtUserday, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *UserDayDao) Create(data *models.LtUserday) error {
	_, err := d.engine.Insert(data)
	return err
}
