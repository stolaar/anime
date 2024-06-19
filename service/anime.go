package service

import (
	"strconv"

	"github.com/stolaar/anime/model"
	"github.com/stolaar/anime/repository"
)

func SearchAnimes(q string) []*model.Anime {
	result := repository.FindAnimeByQuery(q)

	if len(result) > 0 {
		return result
	}

	result, err := ScrapeAnimesByQuery(q)
	if err != nil {
		return result
	}
	repository.SaveScrappedAnimes(result)

	return result
}

func GetAnimeEpisodes(id string) []*model.Episode {
	result := repository.FindAnimeEpisodesById(id)

	if len(result) > 0 {
		return result
	}

	anime := repository.FindAnimeById(id)
	lastEpisodeNumber, err := GetNumberOfEpisodes(anime.Link)
	if err != nil {
		return result
	}
	num, err := strconv.Atoi(lastEpisodeNumber)
	if err != nil {
		return result
	}

	for i := 1; i <= num; i++ {
		si := strconv.Itoa(i)
		result = append(result, &model.Episode{
			Title:   si,
			AnimeId: anime.Id,
			Link:    baseUrl + anime.Link + "/" + "ep-" + si,
		})
	}

	repository.SaveScrappedEpisodes(result)

	return result
}

func GetEpisodeSrc(link string) string {
	scrapped, err := ScrapeEpisode(link)
	if err != nil {
		return ""
	}
	repository.UpdateEpisodeSrc(link, scrapped.Src)

	return scrapped.Src
}
