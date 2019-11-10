package repositories

import (
	"context"
	"golang-kafka-v5-crud/cmd/producer/api/models"
)

type UserRepository interface {
	List(ctx context.Context) ([]*models.User, error)
}
