package tariff

import (
	"github.com/jinzhu/gorm"
)

type Storage struct {
	*gorm.DB
}

func (r *Storage) Get() ([]Tariffs, error) {
	var tariffs []Tariffs
	if err := r.Find(&tariffs, "active = ?", true).Error; err != nil {
		return nil, err
	}
	return tariffs, nil
}

func (r *Storage) GetById(tariffId uint) (*Tariffs, error) {
	var tariffs Tariffs
	if err := r.First(&tariffs, "id = ?", tariffId).Error; err != nil {
		return nil, err
	}
	return &tariffs, nil
}
