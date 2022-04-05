package auth

import (
	"fmt"

	"github.com/nodev918/go-grpc-api-gateway/pkg/auth/pb"
	"github.com/nodev918/go-grpc-api-gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	// using withInsecure() because no SSL running
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())
	
	if err != nil{
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAuthServiceClient(cc)
}