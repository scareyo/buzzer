package handler

import (
    "fmt"
    "net/http"
    
    "github.com/scareyo/buzzer/pkg/event"

    "github.com/twilio/twilio-go/twiml"
)

type VoiceHandler struct {
    timeout int
}
    
func (vh VoiceHandler) Start(port string, timeout int) {
    fmt.Println("Starting buzzer server on port " + port)

    vh.timeout = timeout

    http.HandleFunc("/voice/receive", vh.receiveCall);
    http.HandleFunc("/voice/update", vh.updateCall);

    err := http.ListenAndServe(":" + port, nil)
    if err != nil {
        fmt.Println("error starting http server")
    }
}

func (vh VoiceHandler) receiveCall(w http.ResponseWriter, r *http.Request) {
    timeout := fmt.Sprint(vh.timeout)

    fmt.Println("Voice call received. Waiting up to " + timeout + " seconds")

    // Let the call ring for the configured timeout duration (seconds)
    pause := &twiml.VoicePause{
        Length: timeout,
    }

    verbList := []twiml.Element{pause}
    twimlResult, err := twiml.Voice(verbList)
    if err != nil {
        fmt.Println(err)
    }

    // Reply to Twilio with the TwiML
    w.Header().Set("Content-Type", "application/xml")
    w.Write([]byte(twimlResult))

    // Raise CallUpdated event
    r.ParseForm()
    status := r.PostForm["CallStatus"][0]
    event.CallUpdated.Trigger(event.CallUpdatedPayload {
        Status: status,
    })
}

func (vh VoiceHandler) updateCall(w http.ResponseWriter, r *http.Request) {
    // Raise CallUpdated event
    r.ParseForm()
    status := r.PostForm["CallStatus"][0]
    event.CallUpdated.Trigger(event.CallUpdatedPayload {
        Status: status,
    })
}
