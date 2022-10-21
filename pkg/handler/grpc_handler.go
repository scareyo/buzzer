package handler

import (
    "context"
    "fmt"
    "net"
    "sync"

    "google.golang.org/grpc"
    
    "github.com/scareyo/buzzer/pkg/event"

    pb "github.com/scareyo/buzzer/proto"
)

type GrpcHandler struct{
    pb.UnimplementedBuzzerServer
    ringing *bool
}

func (vh GrpcHandler) Start(listener net.Listener) {
    fmt.Println("Starting gRPC server")

    vh.ringing = new(bool)
    *vh.ringing = false

    event.CallUpdated.Register(vh)
    
    server := grpc.NewServer()
    pb.RegisterBuzzerServer(server, &vh)
    server.Serve(listener)
}

func (vh *GrpcHandler) ListenDoor(in *pb.ListenDoorRequest, stream pb.Buzzer_ListenDoorServer) error {
    fmt.Println("Received ListenDoorRequest: " + in.GetMessage())

    var wg sync.WaitGroup
    wg.Add(1)
    go func() {
        defer wg.Done()
        for {
            if *vh.ringing {
                fmt.Println("RING")
                reply := pb.RingDoorReply{Message: "ACTIVATE"}
                stream.Send(&reply)
                break
            }
        }
    }()

    wg.Wait()
    return nil
}

func (vh *GrpcHandler) OpenDoor(ctx context.Context, in *pb.OpenDoorRequest) (*pb.OpenDoorReply, error) {
    fmt.Println("Received OpenDoorRequest: " + in.GetMessage())
    return &pb.OpenDoorReply{Message: "Hello"}, nil
}

func (vh GrpcHandler) updateCall() {
}

func (vh GrpcHandler) Handle(payload event.CallUpdatedPayload) {
    switch payload.Status {
    case "ringing":
        fmt.Println("The phone is ringing")
        *vh.ringing = true
    case "completed":
        fmt.Println("The phone is no longer ringing")
        *vh.ringing = false
    }
}

