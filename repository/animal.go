package repository

import (
	"interview-test/entity"

	"gorm.io/gorm"
)

type AnimalRepository interface {
	InsertAnimal(b entity.Animal) entity.Animal
	UpdateAnimal(b entity.Animal) entity.Animal
	DeleteAnimal(b entity.Animal)
	AllAnimal() []entity.Animal
	FindAnimalByID(AnimalID uint64) entity.Animal
}

type AnimalConnection struct {
	connection *gorm.DB
}

func NewAnimalRepository(dbConn *gorm.DB) AnimalRepository {
	return &AnimalConnection{
		connection: dbConn,
	}
}

func (db *AnimalConnection) InsertAnimal(b entity.Animal) entity.Animal {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *AnimalConnection) UpdateAnimal(b entity.Animal) entity.Animal {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *AnimalConnection) DeleteAnimal(b entity.Animal) {
	db.connection.Delete(&b)
}

func (db *AnimalConnection) FindAnimalByID(AnimalID uint64) entity.Animal {
	var Animal entity.Animal
	db.connection.Preload("User").Find(&Animal, AnimalID)
	return Animal
}

func (db *AnimalConnection) AllAnimal() []entity.Animal {
	var Animals []entity.Animal
	db.connection.Preload("User").Find(&Animals)
	return Animals
}
