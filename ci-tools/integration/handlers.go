package integration

import (
	"encoding/json"
	"fmt"
)

type Handler interface {
	GetEventPtr() interface{}
	Invoke() error
}

type HandlerMap struct {
	registered map[string]Handler
}

func NewHandlerMap() *HandlerMap {
	return &HandlerMap{registered: map[string]Handler{}}
}

func (h *HandlerMap) Register(event string, handler Handler) {
	h.registered[event] = handler
}

func (h *HandlerMap) Lookup(event string) (handler Handler, err error) {
	handler, ok := h.registered[event]
	if !ok {
		err = fmt.Errorf("no handler for event '%s'", event)
	}
	return
}

func ParseEvent(rawPayload []byte, target Handler) (err error) {
	out := target.GetEventPtr()
	err = json.Unmarshal(rawPayload, out)
	if err != nil {
		err = fmt.Errorf("failed to parse event: %w", err)
	}
	return
}
