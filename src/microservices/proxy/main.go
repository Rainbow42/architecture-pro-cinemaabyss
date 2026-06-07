package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type config struct {
	port                  string
	monolithURL           string
	moviesServiceURL      string
	eventsServiceURL      string
	gradualMigration      bool
	moviesMigrationPercent int
}

func main() {
	cfg := loadConfig()

	monolith, err := url.Parse(cfg.monolithURL)
	if err != nil {
		log.Fatal(err)
	}
	movies, err := url.Parse(cfg.moviesServiceURL)
	if err != nil {
		log.Fatal(err)
	}
	events, err := url.Parse(cfg.eventsServiceURL)
	if err != nil {
		log.Fatal(err)
	}

	monolithProxy := httputil.NewSingleHostReverseProxy(monolith)
	moviesProxy := httputil.NewSingleHostReverseProxy(movies)
	eventsProxy := httputil.NewSingleHostReverseProxy(events)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, "Strangler Fig Proxy is healthy")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if strings.HasPrefix(path, "/api/events") {
			eventsProxy.ServeHTTP(w, r)
			return
		}

		if strings.HasPrefix(path, "/api/movies") {
			if shouldRouteToMovies(cfg) {
				moviesProxy.ServeHTTP(w, r)
				return
			}
			monolithProxy.ServeHTTP(w, r)
			return
		}

		if strings.HasPrefix(path, "/api/") {
			monolithProxy.ServeHTTP(w, r)
			return
		}

		http.NotFound(w, r)
	})

	log.Printf("Proxy started on port %s (migration %d%%)", cfg.port, cfg.moviesMigrationPercent)
	log.Fatal(http.ListenAndServe(":"+cfg.port, nil))
}

func loadConfig() config {
	gradual := os.Getenv("GRADUAL_MIGRATION") == "true"
	percent, _ := strconv.Atoi(os.Getenv("MOVIES_MIGRATION_PERCENT"))
	if percent < 0 {
		percent = 0
	}
	if percent > 100 {
		percent = 100
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	return config{
		port:                  port,
		monolithURL:           getEnv("MONOLITH_URL", "http://localhost:8080"),
		moviesServiceURL:      getEnv("MOVIES_SERVICE_URL", "http://localhost:8081"),
		eventsServiceURL:      getEnv("EVENTS_SERVICE_URL", "http://localhost:8082"),
		gradualMigration:      gradual,
		moviesMigrationPercent: percent,
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func shouldRouteToMovies(cfg config) bool {
	if !cfg.gradualMigration {
		return true
	}
	if cfg.moviesMigrationPercent >= 100 {
		return true
	}
	if cfg.moviesMigrationPercent <= 0 {
		return false
	}
	return rand.IntN(100) < cfg.moviesMigrationPercent
}
