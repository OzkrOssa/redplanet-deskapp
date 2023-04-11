package repository

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type Consecutive struct {
	ID          int
	Consecutive int
	Create_At   time.Time
}
type DatabaseRepository interface {
	GetConsecutive() (int64, error)
	CreateConsecutive()
}

type databaseRepository struct {
	db  *sql.DB
	ctx context.Context
}

func NewDatabaseRepository(ctx context.Context, db *sql.DB) *databaseRepository {
	return &databaseRepository{db, ctx}
}

func (dr databaseRepository) GetConsecutive() (int, error) {
	var con Consecutive
	row, err := dr.db.QueryContext(dr.ctx, "SELECT * FROM mekanopayments ORDER BY id DESC LIMIT 1")
	if err != nil {
		log.Println(err)
	}
	defer row.Close()

	for row.Next() {

		err := row.Scan(&con.ID, &con.Consecutive, &con.Create_At)
		if err != nil {
			return 0, err
		}
	}
	return con.Consecutive, nil
}

func (dr databaseRepository) CreateConsecutive(consecutive int) error {
	_, err := dr.db.ExecContext(dr.ctx, "INSERT INTO mekanopayments (consecutive, create_at) VALUES (?,?)", consecutive, time.Now().Format("2006-01-02"))
	if err != nil {
		return err
	}
	return nil
}
