package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gosom/go-sse-events-demo/services"
)

type API struct {
	di  *services.Container
	srv *http.Server

	lock *sync.RWMutex

	incoming   chan chan string
	closing    chan chan string
	registered map[chan string]bool
}

func New(di *services.Container) (*API, error) {
	ans := API{
		di:         di,
		lock:       &sync.RWMutex{},
		incoming:   make(chan chan string),
		closing:    make(chan chan string),
		registered: make(map[chan string]bool),
	}
	ans.srv = &http.Server{
		Addr: di.Cfg.ServerAddr,
	}
	ans.srv.Handler = handler(&ans)
	return &ans, nil
}

func (o *API) Start(ctx context.Context) error {
	go o.listen(ctx)
	return o.srv.ListenAndServeTLS(o.di.Cfg.Crt, o.di.Cfg.Key)
}

func (o *API) listen(ctx context.Context) {
	subscriber := o.di.Rclient.Subscribe(ctx, o.di.Cfg.RedisChan)
	rchan := subscriber.Channel()
	for {
		select {
		case <-ctx.Done():
			return
		case c := <-o.incoming:
			log.Println("register")
			o.lock.Lock()
			o.registered[c] = true
			o.lock.Unlock()
		case c := <-o.closing:
			o.lock.Lock()
			delete(o.registered, c)
			o.lock.Unlock()
			log.Println("deregister")
		case msg := <-rchan:
			// dont' like that
			o.lock.RLock()
			for ch, _ := range o.registered {
				ch <- msg.Payload
			}
			o.lock.RUnlock()
		}
	}

}
func handler(srv *API) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "streaming not supported", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Connection", "keep-alive")

		messages := make(chan string)
		defer close(messages)
		srv.incoming <- messages

		defer func() {
			srv.closing <- messages
		}()

		for {
			select {
			case <-r.Context().Done():
				return
			case data := <-messages:
				fmt.Fprintf(w, "data: %s\n\n", data)
				flusher.Flush()
			}
		}
	}
}
