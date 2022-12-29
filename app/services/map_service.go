package services

import (
	"mindmap-go/app/models"
	"mindmap-go/app/repository"
)

type MapSvc struct {
	Repo repository.MapRepository
}

type MapService interface {
	CreateMap(mapForm *MapForm) (*models.Map, error)
	GetAllByUser(userID int) ([]*models.Map, error)
	GetMapByID(id int, userID int) (*models.Map, error)
	UpdateMap(mindMap *models.Map, req *models.MapUpdate) error
	DeleteMap(mindMap *models.Map) error
}

func NewMapService(repo repository.MapRepository) MapService {
	return &MapSvc{
		Repo: repo,
	}
}

func (m *MapSvc) CreateMap(mapForm *MapForm) (*models.Map, error) {
	mindMap := &models.Map{
		Name:      mapForm.Name,
		Desc:      mapForm.Description,
		CreatorID: mapForm.CreatorID,
		Cards: []models.Card{
			{
				Name:      "Mind Map",
				Text:      "Example long description for your new mind map.",
				CreatorID: mapForm.CreatorID,
			},
		},
	}
	err := m.Repo.CreateMap(mindMap)
	if err != nil {
		return nil, err
	}
	return mindMap, nil
}

func (m *MapSvc) GetAllByUser(userID int) ([]*models.Map, error) {
	return m.Repo.GetAllByUser(userID)
}

func (m *MapSvc) GetMapByID(id int, userID int) (*models.Map, error) {
	return m.Repo.GetMapByID(id, userID)
}

func (m *MapSvc) UpdateMap(mindMap *models.Map, req *models.MapUpdate) error {
	return m.Repo.UpdateMap(mindMap, req)
}

func (m *MapSvc) DeleteMap(mindMap *models.Map) error {
	return m.Repo.DeleteMap(mindMap)
}
