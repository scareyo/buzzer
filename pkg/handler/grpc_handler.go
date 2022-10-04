package handler

import (
    "context"
    "fmt"
    "net"

    "google.golang.org/grpc"

    pb "github.com/scareyo/buzzer/proto"
)

type GrpcHandler struct{
    pb.UnimplementedBuzzerServer
}

func (vh GrpcHandler) Start(listener net.Listener) {
    fmt.Println("Starting gRPC server")
    server := grpc.NewServer()
    pb.RegisterBuzzerServer(server, &GrpcHandler{})
    server.Serve(listener)
}

func (vh *GrpcHandler) ListenDoor(in *pb.ListenDoorRequest, stream pb.Buzzer_ListenDoorServer) error {
    fmt.Println("Received ListenDoorRequest: " + in.GetMessage())
    return nil
}

func (vh *GrpcHandler) OpenDoor(ctx context.Context, in *pb.OpenDoorRequest) (*pb.OpenDoorReply, error) {
    fmt.Println("Received OpenDoorRequest: " + in.GetMessage())
    return &pb.OpenDoorReply{Message: "Hello"}, nil
}

func (vh GrpcHandler) updateCall() {
}
