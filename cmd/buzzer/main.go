package main

import (
    "fmt"
    "net"

    "github.com/scareyo/buzzer/pkg/config"
    "github.com/scareyo/buzzer/pkg/handler"
    "github.com/scareyo/buzzer/pkg/model"

    "github.com/soheilhy/cmux"
)

func main() {
    
    cfg, err := config.NewDefaultConfig()
    if err != nil {
        fmt.Println(err)
    }


    listener, err := net.Listen("tcp", ":" + cfg.Port)
    if err != nil {
        fmt.Println(err)
    }

    m := cmux.New(listener)
    grpcListener := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
    httpListener := m.Match(cmux.HTTP1Fast())

    buzzer := new(model.Buzzer)
    buzzer.Start()

    voice := new(handler.VoiceHandler)
    go voice.Start(httpListener, cfg.Timeout)

    grpc := new(handler.GrpcHandler)
    go grpc.Start(grpcListener)

    m.Serve()
}

