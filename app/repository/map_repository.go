package repository

import (
	"errors"
	"gorm.io/gorm"
	"mindmap-go/app/models"
	"mindmap-go/internal/database"
	"mindmap-go/utils"
)

type MapRepository interface {
	CreateMap(mindMap *models.Map) error
	GetAllByUser(userID int) ([]*models.Map, error)
	GetMapByID(id int, userID int) (*models.Map, error)
	UpdateMap(mindMap *models.Map, req *models.MapUpdate) error
	DeleteMap(mindMap *models.Map) error
}

func NewMapRepository(database *database.Database) MapRepository {
	return &MapRepo{
		DB: database,
	}
}

type MapRepo struct {
	DB *database.Database
}

func (m *MapRepo) CreateMap(mindMap *models.Map) error {
	return m.DB.Connection.Create(&mindMap).Error
}

func (m *MapRepo) GetAllByUser(userID int) ([]*models.Map, error) {
	var res []*models.Map
	err := m.DB.Connection.Where("creator_id = ?", userID).Find(&res).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*models.Map{}, nil
		}
		return nil, err
	}
	return res, nil
}

func (m *MapRepo) GetMapByID(id int, userID int) (*models.Map, error) {
	var res *models.Map
	err := m.DB.Connection.Where("id = ? AND creator_id = ?", id, userID).First(&res).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &utils.NonExistentEntryError{Message: "No map found with such id."}
		}
		return nil, err
	}
	return res, nil
}

func (m *MapRepo) UpdateMap(mindMap *models.Map, req *models.MapUpdate) error {
	return m.DB.Connection.Save(&mindMap).Error
}

func (m *MapRepo) DeleteMap(mindMap *models.Map) error {
	return m.DB.Connection.Delete(&mindMap).Error
}
