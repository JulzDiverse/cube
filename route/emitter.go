package route

import (
	"encoding/json"
	"fmt"

	nats "github.com/nats-io/go-nats"
)

const publisherSubject = "router.register"

type Publisher interface {
	Publish(subj string, data []byte) error
}

type NATSPublisher struct {
	NatsClient *nats.Conn
}

func (p *NATSPublisher) Publish(subj string, data []byte) error {
	return p.NatsClient.Publish(subj, data)
}

type RouteEmitter struct {
	Publisher Publisher
	Scheduler TaskScheduler
	Work      <-chan []RegistryMessage
}

func (r *RouteEmitter) Start() {
	r.Scheduler.Schedule(func() error {
		select {
		case batch := <-r.Work:
			go r.emit(batch)
		}

		return nil
	})
}

func (r *RouteEmitter) emit(batch []RegistryMessage) {
	for _, route := range batch {
		routeJson, err := json.Marshal(route)
		if err != nil {
			fmt.Println("Faild to marshal route message:", err.Error())
			continue
		}

		if err = r.Publisher.Publish(publisherSubject, routeJson); err != nil {
			fmt.Println("failed to publish route:", err.Error())
		}
	}
}
