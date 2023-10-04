package observer

import (
	"fmt"
	"time"

	"github.com/JunNishimura/design-pattern-with-go/observer/publisher"
)

type Logger struct{}

func (l *Logger) Update(data any) { fmt.Printf("logger: %v\n", data) }

type Mailer struct{}

func (m *Mailer) Update(data any) { fmt.Printf("mailer: %v\n", data) }

var p *publisher.Publisher

func main() {
	p = publisher.New()
	logger := &Logger{}
	mailer := &Mailer{}

	p.Subscribe(logger)
	p.Subscribe(mailer)
	defer func() {
		p.UnSubscribe(logger)
		p.UnSubscribe(mailer)
	}()

	eventTrigger()
}

func eventTrigger() {
	for i := 0; i < 10; i++ {
		p.Notify(i)
		time.Sleep(5 * time.Second)
	}
}
