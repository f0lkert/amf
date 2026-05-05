package processor

import (
	"github.com/f0lkert/amf/internal/sbi/consumer"
	"github.com/f0lkert/amf/pkg/app"
)

type ProcessorAmf interface {
	app.App

	Consumer() *consumer.Consumer
}

type Processor struct {
	ProcessorAmf
}

func NewProcessor(amf ProcessorAmf) (*Processor, error) {
	p := &Processor{
		ProcessorAmf: amf,
	}
	return p, nil
}
