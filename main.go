package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/stolaar/anime/service"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/episode", func(w http.ResponseWriter, r *http.Request) {
		link, err := url.QueryUnescape(r.URL.Query().Get("link"))
		if err != nil {
			log.Fatal(err)
		}
		src := service.GetEpisodeSrc(link)

		src, _ = url.QueryUnescape(src)
		w.Write([]byte(src))
	})

	r.Get("/anime", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")

		animes, err := json.Marshal(service.SearchAnimes(query))
		if err != nil {
			log.Fatal(err)
		}

		w.Write([]byte(animes))
	})

	r.Get("/anime-episodes/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		res, err := json.Marshal(service.GetAnimeEpisodes(id))
		if err != nil {
			log.Fatal(err)
		}

		w.Write([]byte(res))
	})

	http.ListenAndServe(":3333", r)
}
