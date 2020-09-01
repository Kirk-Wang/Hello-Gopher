package services

import (
	"github.com/Kirk-Wang/Hello-Gopher/lottery/wheel/dao"
	"github.com/Kirk-Wang/Hello-Gopher/lottery/wheel/models"
)

type GiftService interface {
	GetAll() []models.LtGift
	CountAll() int64
	Get(id int) *models.LtGift
	Delete(id int) error
	Update(data *models.LtGift, columns []string) error
	Create(data *models.LtGift) error
}

type giftService struct {
	dao *dao.GiftDao
}

func NewGiftService() GiftService {
	return &giftService{
		dao: dao.NewGiftDao(nil),
	}
}

func (s *giftService) GetAll() []models.LtGift {
	return s.dao.GetAll()
}
func (s *giftService) CountAll() int64 {
	return s.dao.CountAll()
}
func (s *giftService) Get(id int) *models.LtGift {
	return s.dao.Get(id)
}
func (s *giftService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *giftService) Update(data *models.LtGift, columns []string) error {
	return s.dao.Update(data, columns)
}
func (s *giftService) Create(data *models.LtGift) error {
	return s.dao.Create(data)
}
