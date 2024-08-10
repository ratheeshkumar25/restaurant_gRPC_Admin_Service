package di

import (
	"log"

	"github.com/ratheeshkumar/restaurant_admin_serviceV1/config"
	"github.com/ratheeshkumar/restaurant_admin_serviceV1/internal/db"
	"github.com/ratheeshkumar/restaurant_admin_serviceV1/internal/handlers"
	"github.com/ratheeshkumar/restaurant_admin_serviceV1/internal/menu"
	"github.com/ratheeshkumar/restaurant_admin_serviceV1/internal/repositories"
	"github.com/ratheeshkumar/restaurant_admin_serviceV1/internal/server"
	"github.com/ratheeshkumar/restaurant_admin_serviceV1/internal/services"
)

func Init() {
	config.LoadConfig()

	db := db.ConnectDB()
	menuClient, err := menu.ClientDial()
	if err != nil {
		log.Fatalf("something went wrong when dialing menu client %v", err)
	}

	adminRepo := repositories.NewAdminRepo(db)
	adminSvc := services.NewAdminService(adminRepo, menuClient)
	adminHandler := handlers.NewAdminHandler(adminSvc)
	server.NewGrpcServer(adminHandler)
}
