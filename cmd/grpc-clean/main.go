package main

import (
	"fmt"
	"log"
	"net"

	dbConfig "github.com/zerokkcoder/grpc-clean/internal/db"
	"github.com/zerokkcoder/grpc-clean/internal/models"
	interfaces "github.com/zerokkcoder/grpc-clean/pkg/v1"
	handler "github.com/zerokkcoder/grpc-clean/pkg/v1/handler/grpc"
	repo "github.com/zerokkcoder/grpc-clean/pkg/v1/repository"
	"github.com/zerokkcoder/grpc-clean/pkg/v1/usecase"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func main() {
	// connect to db
	db := dbConfig.DbConn()
	migrations(db)

	// add a listener address
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("ERROR STARTING THE SERVER: %v", err)
	}

	// start the grpc server
	grpcServer := grpc.NewServer()

	userUseCase := initUserServer(db)
	handler.NewServer(grpcServer, userUseCase)

	// start the server
	log.Fatal(grpcServer.Serve(lis))
}

func initUserServer(db *gorm.DB) interfaces.UseCaseInterface {
	userRepo := repo.New(db)
	return usecase.New(userRepo)
}

func migrations(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Migrated")
	}
}
