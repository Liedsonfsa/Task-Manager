package main

import (
	"log"

	"github.com/Liedsonfsa/Task-Manager/internal/api"
	"github.com/Liedsonfsa/Task-Manager/internal/config"
	"github.com/Liedsonfsa/Task-Manager/internal/repository"
	"github.com/Liedsonfsa/Task-Manager/internal/service"
	"github.com/Liedsonfsa/Task-Manager/pkg/database"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.Connect(database.DBConfig(cfg.DBConfig))
	if err != nil {
		log.Fatal("Failed to connect to database: %v", err)
	}
	defer db.Close()

	taskRepo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := api.NewTaskHandler(taskService)

	router := api.NewRouter(taskHandler)
	log.Printf("Starting server on %s", cfg.ServerAddress)
	if err = router.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}