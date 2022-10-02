package model

import (
    "github.com/scareyo/buzzer/pkg/event"
)

type Buzzer struct {
    active  bool
}

func (b Buzzer) Start() {
    event.CallUpdated.Register(b)
}

func (b *Buzzer) updateStatus(status string) {
    switch status {
    case "ringing":
        b.active = true
    case "completed":
        b.active = false
    }
}

func (b Buzzer) Handle(payload event.CallUpdatedPayload) {
    b.updateStatus(payload.Status)
}
