package main

import (
	"anime/migrations/list"

	gm "github.com/ShkrutDenis/go-migrations"
	gmStore "github.com/ShkrutDenis/go-migrations/store"
)

func main() {
	gm.Run(getMigrationsList())
}

func getMigrationsList() []gmStore.Migratable {
	return []gmStore.Migratable{
		&list.CreateUserTable{},
		&list.CreatePostTable{},
		&list.UpdateUserTable{},
		&list.CreateAnimeTable{},
		&list.CreateEpisodeTable{},
	}
}
