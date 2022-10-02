package main

import (
    "fmt"

    "github.com/scareyo/buzzer/pkg/config"
    "github.com/scareyo/buzzer/pkg/handler"
    "github.com/scareyo/buzzer/pkg/model"
)

func main() {
    
    cfg, err := config.NewDefaultConfig()
    if err != nil {
        fmt.Println(err)
    }

    buzzer := new(model.Buzzer)
    buzzer.Start()

    voice := new(handler.VoiceHandler)
    voice.Start(cfg.VoicePort, cfg.Timeout)
}

