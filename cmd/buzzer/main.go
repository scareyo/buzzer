package main

import (
    "fmt"

    "github.com/scareyo/buzzer/pkg/config"
    "github.com/scareyo/buzzer/pkg/handler"
)

func main() {
    
    cfg, err := config.NewDefaultConfig()
    if err != nil {
        fmt.Println(err)
    }

    voice := new (handler.VoiceHandler)
    voice.Start(cfg.VoicePort)
}

