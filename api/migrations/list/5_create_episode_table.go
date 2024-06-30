package list

import (
	mysql "github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateEpisodeTable struct{}

func (m *CreateEpisodeTable) GetName() string {
	return "CreateEpisodeTable"
}

func (m *CreateEpisodeTable) Up(con *sqlx.DB) {
	table := mysql.NewTable("episodes", con)
	table.Column("id").Type("int").Autoincrement()
	table.PrimaryKey("id")
	table.Integer("anime_id")
	table.ForeignKey("anime_id").Reference("animes").SetKeyName("fk_anime_episodes").On("id")
	table.String("title", 500).Nullable()
	table.String("link", 500).Nullable()
	table.String("src", 500).Nullable()
	table.String("image", 500).Nullable()
	table.String("description", 1000).Nullable()
	table.Column("deleted_at").Type("timestamp").Nullable()
	table.WithTimestamps()

	table.MustExec()
}

func (m *CreateEpisodeTable) Down(con *sqlx.DB) {
	mysql.DropTable("episodes", con).MustExec()
}
