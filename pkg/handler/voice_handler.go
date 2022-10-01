package handler

import (
    "fmt"
    "net/http"

    "github.com/twilio/twilio-go/twiml"
)

type VoiceHandler struct{}
    
func (vh VoiceHandler) Start(port string) {
    fmt.Println("Starting buzzer server on port " + port)
    http.HandleFunc("/voice", vh.handleVoiceCall);

    err := http.ListenAndServe(":" + port, nil)
    if err != nil {
        fmt.Println("error starting http server")
    }
}

func (vh VoiceHandler) handleVoiceCall(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Voice call received")

    say := &twiml.VoiceSay{
        Message: "Hey Sam",
    }

    pause := &twiml.VoicePause{
        Length: "10",
    }

    verbList := []twiml.Element{pause, say}
    twimlResult, err := twiml.Voice(verbList)
    if err != nil {
        fmt.Println(err)
    }

    w.Header().Set("Content-Type", "application/xml")
    w.Write([]byte(twimlResult))
}
