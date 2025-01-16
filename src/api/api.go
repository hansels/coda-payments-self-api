package api

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
	"time"
)

type Opts struct {
	Delay int64
}

type API struct {
	Opts *Opts
}

func New(opts *Opts) *API {
	return &API{Opts: opts}
}

func (a *API) Register(router *httprouter.Router) {
	// Health Check Endpoint
	router.GET("/ping", a.Ping)

	// Mirror Endpoint
	router.POST("/self", a.Self)
}

func (a *API) Ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("[GET] /ping Called!")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))

	// Delay the response
	log.Println("Delaying the response for", a.Opts.Delay, "seconds")
	time.Sleep(time.Duration(a.Opts.Delay) * time.Second)

	return
}

func (a *API) Self(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("[POST] /self Called!")

	if _, err := io.Copy(w, r.Body); err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// Delay the response
	log.Println("Delaying the response for", a.Opts.Delay, "seconds")
	time.Sleep(time.Duration(a.Opts.Delay) * time.Second)

	return
}
