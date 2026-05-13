package database

import (
	"database/sql"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID        int64     `bun:",pk,autoincrement"`
	Name      string    `bun:",notnull"`
	Email     string    `bun:",unique,notnull"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

// Using pgdriver (recommended)
var sqldb = sql.OpenDB(pgdriver.NewConnector(
	pgdriver.WithDSN(""),
))
var DB_CLIENT = bun.NewDB(sqldb, pgdialect.New())

// func main() {
// 	_, err := DB_CLIENT.NewCreateTable().
// 		Model((*User)(nil)).
// 		IfNotExists().Exec(context.Background())

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	user := &User{Name: "John Doe", Email: "john@example.com"}

// 	_, err = DB_CLIENT.NewInsert().Model(user).Exec(context.Background())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }
