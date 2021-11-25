package publisher

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"

	"github.com/gosom/go-sse-events-demo/services"
)

type Publisher struct {
	di *services.Container
}

func New(di *services.Container) (*Publisher, error) {
	ans := Publisher{
		di: di,
	}
	return &ans, nil
}

func (o *Publisher) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			event := map[string]interface{}{
				"uuid": uuid.New().String(),
				"ts":   time.Now().UTC(),
			}
			b, err := json.Marshal(event)
			if err != nil {
				return err
			}
			if err := o.di.Rclient.Publish(ctx, o.di.Cfg.RedisChan, string(b)).Err(); err != nil {
				return err
			}
			time.Sleep(time.Millisecond * 100)
		}
	}
	return nil
}
