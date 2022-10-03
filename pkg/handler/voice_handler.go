package handler

import (
    "fmt"
    "net"
    "net/http"
    
    "github.com/scareyo/buzzer/pkg/event"

    "github.com/twilio/twilio-go/twiml"
)

type VoiceHandler struct {
    timeout int
}
    
func (vh VoiceHandler) Start(listener net.Listener, timeout int) {
    fmt.Println("Starting HTTP server")

    vh.timeout = timeout

    mux := http.NewServeMux()

    mux.HandleFunc("/voice/receive", vh.receiveCall);
    mux.HandleFunc("/voice/update", vh.updateCall);

    server := &http.Server{Handler: mux}
    server.Serve(listener)
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
