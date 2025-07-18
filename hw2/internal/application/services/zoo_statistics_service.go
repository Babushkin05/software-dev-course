package services

import "github.com/Babushkin05/software-dev-course/hw2/internal/application/ports"

type ZooStatistics struct {
	TotalAnimals    int `json:"total_animals" example:"42"`
	FreeEnclosures  int `json:"free_enclosures" example:"5"`
	TotalEnclosures int `json:"total_enclosures" example:"15"`
}

type ZooStatisticsService struct {
	animals    ports.AnimalRepository
	enclosures ports.EnclosureRepository
}

func NewZooStatisticsService(a ports.AnimalRepository, e ports.EnclosureRepository) *ZooStatisticsService {
	return &ZooStatisticsService{a, e}
}

func (z *ZooStatisticsService) GetStatistics() (*ZooStatistics, error) {
	animalList, _ := z.animals.List()
	enclosureList, _ := z.enclosures.List()

	free := 0
	for _, e := range enclosureList {
		if len(e.AnimalIDs) < e.Capacity {
			free++
		}
	}

	return &ZooStatistics{
		TotalAnimals:    len(animalList),
		FreeEnclosures:  free,
		TotalEnclosures: len(enclosureList),
	}, nil
}
