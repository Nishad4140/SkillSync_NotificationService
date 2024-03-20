package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/Nishad4140/SkillSync_NotificationService/db"
	"github.com/Nishad4140/SkillSync_NotificationService/initializer"
	"github.com/Nishad4140/SkillSync_ProtoFiles/pb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf(err.Error())
	}
	addr := os.Getenv("DB_KEY")
	db, err := db.InitDB(addr)
	if err != nil {
		log.Fatalf("error connecting to database")
	}
	lis, err := net.Listen("tcp", ":4007")
	if err != nil {
		log.Fatalf("error while listening on the port 4007")
	}
	fmt.Println("notification service running on the port 4007")
	services := initializer.Initializer(db)
	server := grpc.NewServer()
	pb.RegisterNotificationServiceServer(server, services)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to listening on the port 4007")
	}
}
