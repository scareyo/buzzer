package event

var CallUpdated callUpdated

type CallUpdatedPayload struct {
    Status string
}

type callUpdated struct {
    handlers []interface{ Handle(CallUpdatedPayload) }
}

func (c *callUpdated) Register(handler interface{ Handle(CallUpdatedPayload) }) {
    c.handlers = append(c.handlers, handler)
}

func (c callUpdated) Trigger(payload CallUpdatedPayload) {
    for _, handler := range c.handlers {
        go handler.Handle(payload)
    }
}
