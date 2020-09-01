package dao

import (
	"github.com/Kirk-Wang/Hello-Gopher/iris-1/models"
	"github.com/go-xorm/xorm"
)

type SuperstarDao struct {
	engine *xorm.Engine // 基于XORM抽象出 CRUD
}

func NewSuperstarDao(engine *xorm.Engine) *SuperstarDao {
	return &SuperstarDao{
		engine: engine,
	}
}

func (d *SuperstarDao) Get(id int) *models.StarInfo {
	data := &models.StarInfo{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		// 等于0，表明也是没有读出来，所以这里有两种选择
		// return nil
		data.Id = 0
		return data
	}
}

func (d *SuperstarDao) GetAll() []models.StarInfo {
	datalist := make([]models.StarInfo, 0)
	// 对 id 做一个降序的排序
	err := d.engine.Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *SuperstarDao) Search(country string) []models.StarInfo {
	datalist := make([]models.StarInfo, 0)
	err := d.engine.Where("country=?", country).
		Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *SuperstarDao) Delete(id int) error {
	// 不做物理删除，做个逻辑删除（状态的更新）--> GetAll 要处理一下
	data := &models.StarInfo{Id: id, SysStatus: 1}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

func (d *SuperstarDao) Update(data *models.StarInfo, columns []string) error {
	// MustCols -> 空字符串，0，false 也会强制被更新
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *SuperstarDao) Create(data *models.StarInfo) error {
	_, err := d.engine.Insert(data)
	return err
}
