package postgres_storage

import (
	"context"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/entities"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/pkg/sqlx_utils"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type ModelElementsPgRepository struct {
	conn *sqlx.DB
}

func (r *ModelElementsPgRepository) GetAllModelElements(layer string) ([]entities.ModelElement, error) {
	args := map[string]any{
		"layer_name": layer,
	}
	var elements []entities.ModelElement
	err := sqlx_utils.NamedSelect(context.Background(), r.conn, &elements, selectAllModelElements, args)
	if err != nil {
		return nil, err
	}
	return elements, nil
}

func (r *ModelElementsPgRepository) SaveModelElements(layer string, modelElements []string) ([]int, error) {
	args := map[string]any{
		"layer_name":     layer,
		"elements_array": pq.Array(modelElements),
	}
	id := make([]int, 0)
	err := sqlx_utils.NamedSelect(context.Background(), r.conn, &id, insertModelElements, args)
	if err != nil {
		return nil, err
	}
	return id, nil
}
