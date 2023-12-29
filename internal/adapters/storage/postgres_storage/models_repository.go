package postgres_storage

import (
	"context"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/entities"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/pkg/sqlx_utils"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type ModelsPgRepository struct {
	conn *sqlx.DB
}

func (r *ModelsPgRepository) GetAllModels(layer string) ([]entities.Model, error) {
	args := map[string]any{
		"layer_name": layer,
	}
	models := make([]entities.Model, 0)
	err := sqlx_utils.NamedSelect(context.Background(), r.conn, &models, selectAllModels, args)
	if err != nil {
		return nil, err
	}
	return models, nil
}

func (r *ModelsPgRepository) SaveModels(layer string, models []string) ([]int, error) {
	args := map[string]any{
		"layer_name":   layer,
		"models_array": pq.Array(models),
	}
	id := make([]int, 0)
	err := sqlx_utils.NamedSelect(context.Background(), r.conn, &id, insertModels, args)
	if err != nil {
		return nil, err
	}
	return id, nil
}
