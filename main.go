package main

import (
	"fmt"
	"log"
	"net"

	connectDB "github.com/imagefinder/connect"
	"github.com/imagefinder/initializers"
	"github.com/imagefinder/models"
	pb "github.com/imagefinder/models"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedImageServiceServer
}

func init() {
	initializers.LoadEnv()
	connectDB.InitConnector()
}

func (s *server) PostImage(ctx context.Context, req *pb.PostImageRequest) (*pb.PostImageResponse, error) {
	return nil
}

func (s *server) GetImage(ctx context.Context, req *pb.GetImageRequest) (*pb.GetImageResponse, error) {
	var image models.Image
	id := req.ID

	formatedID := fmt.Sprintf("%v", id)

	connectDB.DB.Table("images").Where("ID = ?", formatedID).First(&image)
	return
}
func main() {
	// Start a gRPC server
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	server := &server{}
	pb.RegisterImageServiceServer(s, server)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	// Connect to the database

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

}
