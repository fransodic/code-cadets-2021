package engine

import (
	"context"
	"log"
)

// Engine is the main component, responsible for consuming received bets and calculated bets,
// processing them and publishing the resulting bets.
type Engine struct {
	consumer  Consumer
	handler   Handler
	publisher Publisher
}

// New creates and returns a new engine.
func New(consumer Consumer, handler Handler, publisher Publisher) *Engine {
	return &Engine{
		consumer:  consumer,
		handler:   handler,
		publisher: publisher,
	}
}

// Start will run the engine.
func (e *Engine) Start(ctx context.Context) {
	err := e.processBets(ctx)
	if err != nil {
		log.Println("Engine failed to process bets received:", err)
		return
	}

	err = e.processEventUpdates(ctx)
	if err != nil {
		log.Println("Engine failed to process bets calculated:", err)
		return
	}

	<-ctx.Done()
}

func (e *Engine) processBets(ctx context.Context) error {
	consumedBetsReceived, err := e.consumer.ConsumeBets(ctx)
	if err != nil {
		return err
	}

	resultingBets := e.handler.HandleBets(ctx, consumedBetsReceived)
	e.publisher.PublishBetsCalculated(ctx, resultingBets)

	return nil
}

func (e *Engine) processEventUpdates(ctx context.Context) error {
	consumedBetsCalculated, err := e.consumer.ConsumeEventUpdates(ctx)
	if err != nil {
		return err
	}

	resultingBets := e.handler.HandleEventUpdates(ctx, consumedBetsCalculated)
	e.publisher.PublishBetsCalculated(ctx, resultingBets)

	return nil
}
