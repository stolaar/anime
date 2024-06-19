package repository

import (
	"log"

	"github.com/stolaar/anime/database"
	"github.com/stolaar/anime/model"
)

var (
	animes   = []*model.Anime{}
	episodes = []*model.Episode{}
)

func FindAnimeByQuery(q string) []*model.Anime {
	result := []*model.Anime{}
	db := database.NewConnection()

	defer db.Close()

	rows, err := db.Query(`SELECT * FROM anime where title like '%?%'`, &q)
	if err != nil {
		return result
	}

	defer rows.Close()

	for rows.Next() {
		model := &model.Anime{}
		err = rows.Scan(model.Id, model.Title, model.Link, model.Image)
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, model)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func SaveScrappedAnimes(scrapped []*model.Anime) {
	animes = append(animes, scrapped...)
}

func SaveScrappedEpisodes(scrapped []*model.Episode) {
	episodes = append(episodes, scrapped...)
}

func FindAnimeEpisodesById(id string) []*model.Episode {
	result := []*model.Episode{}

	for _, ep := range episodes {
		if ep.AnimeId == id {
			result = append(result, ep)
		}
	}

	return result
}

func FindAnimeById(id string) *model.Anime {
	result := &model.Anime{}

	for _, an := range animes {
		if an.Id == id {
			result = an
			break
		}
	}

	return result
}

func UpdateEpisodeSrc(link string, src string) {
	for _, ep := range episodes {
		if ep.Link == link {
			ep.Src = src
		}
	}
}
