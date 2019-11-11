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
	var errs = make(chan error)
	var user []*models.User
	var done = make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		query := `SELECT id, email, username, password, created_at FROM users`

		rows, err := r.db.QueryContext(ctx, query)

		if err != nil {
			errs <- err
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
				errs <- err
				ch <- false
				return
			}
			user = append(user, row)
		}
		ch <- true
	}(done)

	select {
	case ok := <-done:
		if ok {
			return user, nil
		}
	case err = <-errs:
		return nil, err
	}
	return nil, err
}

func (r *userRepository) Detail(ctx context.Context, id int) (*models.User, error) {
	var user models.User
	var err error
	var done = make(chan bool)

	go func(ch chan<- bool) {
		var query = `SELECT id, email, username, password, created_at FROM users WHERE id = $1`
		err = r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.CreatedAt)
		if err != nil {
			ch <- false
			return
		}
		if err == sql.ErrNoRows {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if utils.OK(done) {
		return &user, nil
	}
	return nil, err
}
