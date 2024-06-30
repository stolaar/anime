package list

import (
	mysql "github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateAnimeTable struct{}

func (m *CreateAnimeTable) GetName() string {
	return "CreateAnimeTable"
}

func (m *CreateAnimeTable) Up(con *sqlx.DB) {
	table := mysql.NewTable("animes", con)
	table.Column("id").Type("int").Autoincrement()
	table.PrimaryKey("id")
	table.String("title", 500).Nullable()
	table.String("link", 500).Nullable()
	table.String("src", 500).Nullable()
	table.String("image", 500).Nullable()
	table.String("description", 1000).Nullable()
	table.Column("deleted_at").Type("timestamp").Nullable()
	table.WithTimestamps()

	table.MustExec()
}

func (m *CreateAnimeTable) Down(con *sqlx.DB) {
	mysql.DropTable("animes", con).MustExec()
}
