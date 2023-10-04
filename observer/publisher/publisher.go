package publisher

type Subscriber interface {
	Update(data any)
}

type Publisher struct {
	subscribers []Subscriber
}

func New() *Publisher {
	return &Publisher{}
}

func (o *Publisher) Subscribe(s Subscriber) {
	o.subscribers = append(o.subscribers, s)
}

func (o *Publisher) UnSubscribe(s Subscriber) {
	newsubscribers := make([]Subscriber, 0, len(o.subscribers))
	for _, subscriber := range o.subscribers {
		if subscriber != s {
			newsubscribers = append(newsubscribers, subscriber)
		}
	}
	o.subscribers = newsubscribers
}

func (o *Publisher) Notify(data any) {
	for _, observer := range o.subscribers {
		observer.Update(data)
	}
}
