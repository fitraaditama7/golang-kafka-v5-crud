package repo

import (
	"context"
	"database/sql"
	"golang-kafka-v5-crud/cmd/producer/api/helper/utils"
	"golang-kafka-v5-crud/cmd/producer/api/models"
	"golang-kafka-v5-crud/cmd/producer/api/repositories"
)

type userRepository struct {
	db *sql.DB
}

func NewPSQLRepository(db *sql.DB) repositories.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) List(ctx context.Context) ([]*models.User, error) {
	var err error
	var user []*models.User
	var done = make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		query := `SELECT id, email, username, password, created_at FROM user`

		rows, err := r.db.QueryContext(ctx, query)
		if err != nil {
			ch <- false
			return
		}

		for rows.Next() {
			row := new(models.User)
			err = rows.Scan(
				&row.ID,
				&row.Email,
				&row.Username,
				&row.Password,
				&row.CreatedAt,
			)

			if err != nil {
				ch <- false
				return
			}
			user = append(user, row)
		}
	}(done)

	if utils.OK(done) {
		return user, nil
	}
	return nil, err
}
