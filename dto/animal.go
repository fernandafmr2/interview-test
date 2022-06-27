package dto

// used by client when create a new animal
type AnimalCreateDTO struct {
	Name       string `json:"name" form:"name" binding:"required"`
	Class string `json:"class" form:"class" binding:"required"`
	Legs      uint64 `json:"legs" form:"legs" binding:"required"`
}

// used by client when updating a animal
type AnimalUpdateDTO struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Name       string `json:"name" form:"name" binding:"required"`
	Class string `json:"class" form:"class" binding:"required"`
	Legs      uint64 `json:"legs" form:"legs" binding:"required"`
}
