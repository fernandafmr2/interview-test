package service

import (
	"log"
	"interview-test/dto"
	"interview-test/entity"
	"interview-test/repository"

	"github.com/mashingan/smapping"
)

type AnimalService interface {
	Insert(b dto.AnimalCreateDTO) entity.Animal
	Update(b dto.AnimalUpdateDTO) entity.Animal
	Delete(b entity.Animal)
	All() []entity.Animal
	FindByID(animalID uint64) entity.Animal
}

type animalService struct {
	AnimalRepository repository.AnimalRepository
}

func NewAnimalService(animalRepo repository.AnimalRepository) AnimalService {
	return &animalService{
		AnimalRepository: animalRepo,
	}
}

func (service *animalService) Insert(b dto.AnimalCreateDTO) entity.Animal {
	animal := entity.Animal{}
	err := smapping.FillStruct(&animal, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.AnimalRepository.InsertAnimal(animal)
	return res
}

func (service *animalService) Update(b dto.AnimalUpdateDTO) entity.Animal {
	animal := entity.Animal{}
	err := smapping.FillStruct(&animal, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.AnimalRepository.UpdateAnimal(animal)
	return res
}

func (service *animalService) Delete(b entity.Animal) {
	service.AnimalRepository.DeleteAnimal(b)
}

func (service *animalService) All() []entity.Animal {
	return service.AnimalRepository.AllAnimal()
}

func (service *animalService) FindByID(animalID uint64) entity.Animal {
	return service.AnimalRepository.FindAnimalByID(animalID)
}