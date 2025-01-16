package server

import (
	"github.com/hansels/coda-payments-self-api/src/api"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"log"
	"net/http"
)

type Opts struct {
	ListenAddress string
	Delay         int64
}

type Handler struct {
	options     *Opts
	listenErrCh chan error
}

func New(o *Opts) *Handler {
	handler := &Handler{options: o}
	return handler
}

func (h *Handler) Run() {
	log.Printf("Listening on %s", h.options.ListenAddress)

	c := cors.New(cors.Options{
		AllowedHeaders: []string{"X-Requested-With", "Authorization", "Content-Type", "X-Authorization"},
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"FETCH", "GET", "POST", "DELETE", "PUT", "OPTIONS"},
	})

	router := httprouter.New()

	apiOpts := &api.Opts{Delay: h.options.Delay}
	api.New(apiOpts).Register(router)

	handler := c.Handler(router)
	h.listenErrCh <- http.ListenAndServe(h.options.ListenAddress, handler)
}

func (h *Handler) ListenError() <-chan error {
	return h.listenErrCh
}
