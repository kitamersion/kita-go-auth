package events

import "sync"

type EventBus struct {
	handlers map[EventName][]EventHandler
	mu       sync.RWMutex
}

var EventBusGo *EventBus

func InitalizeEventBus() {
	EventBusGo = &EventBus{
		handlers: make(map[EventName][]EventHandler),
	}
}

// Subscribe registers an EventHandler for a specific event name.
func (bus *EventBus) Subscribe(eventName EventName, handler EventHandler) {
	bus.mu.Lock()
	defer bus.mu.Unlock()
	bus.handlers[eventName] = append(bus.handlers[eventName], handler)
}

// Publish broadcasts an event to all registered handlers for its name.
func (bus *EventBus) Publish(event Event) {
	bus.mu.RLock()
	defer bus.mu.RUnlock()

	if handlers, found := bus.handlers[event.Name()]; found {
		for _, handler := range handlers {
			go handler.Handle(event) // Execute handlers asynchronously
		}
	}
}
