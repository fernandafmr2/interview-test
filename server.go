package main

import (
	"interview-test/config"
	"interview-test/controller"
	"interview-test/repository"
	"interview-test/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetUpDatabaseConnection()
	animeRepository repository.AnimalRepository = repository.NewAnimalRepository(db)
	animeService    service.AnimalService       = service.NewAnimalService(animeRepository)
	animalController controller.AnimalController = controller.NewAnimalController(animeService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()


	bookRoutes := r.Group("api/")
	{
		bookRoutes.GET("/", animalController.All)
		bookRoutes.POST("/", animalController.Insert)
		bookRoutes.GET("/:id", animalController.FindByID)
		bookRoutes.PUT("/", animalController.Update)
		bookRoutes.DELETE("/:id", animalController.Delete)
	}

	r.Run(":9000")
}
