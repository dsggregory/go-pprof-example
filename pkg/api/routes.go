package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/http/pprof"
	"time"
)

func setupPprofRoutes(router *mux.Router) {
	// These are straight out of init() from /net/http/pprof but we must add to our non-default router.
	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	router.HandleFunc("/debug/pprof/trace", pprof.Trace)

	// For gorilla mux, we need to explicitly route these which are handled by pprof.Index()
	router.HandleFunc("/debug/pprof/allocs", pprof.Index)
	router.HandleFunc("/debug/pprof/block", pprof.Index)
	router.HandleFunc("/debug/pprof/goroutine", pprof.Index)
	router.HandleFunc("/debug/pprof/heap", pprof.Index)
	router.HandleFunc("/debug/pprof/mutex", pprof.Index)
	router.HandleFunc("/debug/pprof/threadcreate", pprof.Index)
	router.HandleFunc("/debug/pprof/trace", pprof.Index)
}

func setup(debug bool) (*mux.Router) {
	router := mux.NewRouter()

	if debug {
		// Standard pprof endpoints are always enabled. These are additional and possibly expensive.
		log.Print("Debug endpoints enabled")
		router.HandleFunc("/debug/pprof/trace", pprof.Trace)
		router.PathPrefix("/debug").HandlerFunc(pprof.Index)
	}

	setupPprofRoutes(router)

	return router
}

// NewServer constructs and returns an http.Server that's ready
// to serve connections.
func NewServer(addr string, debug bool) (*http.Server) {
	router := setup(debug)
	return &http.Server{
		Handler: router,
		Addr:    addr,

		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: time.Minute * 2,
		ReadTimeout:  time.Minute * 2,
	}
}


