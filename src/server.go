package main

import (
	"log"
	"os"

	"orchestrator/controllers"

	"github.com/Kamva/mgm/v2"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {

	connectionString := os.ExpandEnv("mongodb://$MONGODB_USER:$MONGODB_PASS@MONGODB_HOST:$MONGODB_PORT")
	log.Printf("mongo connection string: %s", connectionString)

	err := mgm.SetDefaultConfig(nil, "jobs", options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	app := fiber.New()

	app.Get("/api/jobs", controllers.GetAllJobs)
	app.Get("/api/jobs/:id", controllers.GetJobByID)
	app.Post("/api/jobs", controllers.CreateJob)
	app.Patch("/api/jobs/:id", controllers.ToggleJobStatus)
	app.Delete("/api/jobs/:id", controllers.DeleteJob)

	app.Listen(3000)
}