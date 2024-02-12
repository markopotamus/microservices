package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/markopotamus/microservices-proto/golang/order"
	"github.com/markopotamus/microservices/order/config"
	"github.com/markopotamus/microservices/order/internal/ports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api  ports.APIPort
	port int
	order.UnimplementedOrderServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Run() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, err=%v", err)
	}

	grpcServer := grpc.NewServer()
	order.RegisterOrderServer(grpcServer, a)

	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port %d", a.port)
	}
}
