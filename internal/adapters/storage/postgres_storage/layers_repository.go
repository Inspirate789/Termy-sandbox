package postgres_storage

import (
	"context"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/pkg/sqlx_utils"
	"github.com/jmoiron/sqlx"
	"slices"
)

type LayersPgRepository struct {
	conn *sqlx.DB
}

func (r *LayersPgRepository) LayerExist(layer string) (bool, error) {
	layers := make([]string, 0)
	err := sqlx_utils.NamedSelect(context.Background(), r.conn, &layers, selectLayerNames, make(map[string]any))
	if err != nil {
		return false, err
	}
	return slices.Contains(layers, layer), nil
}

func (r *LayersPgRepository) GetAllLayers() ([]string, error) {
	layers := make([]string, 0)
	err := sqlx_utils.NamedSelect(context.Background(), r.conn, &layers, selectLayerNames, make(map[string]any))
	if err != nil {
		return nil, err
	}
	return layers, nil
}

func (r *LayersPgRepository) SaveLayer(name string) error {
	args := map[string]any{
		"layer_name": name,
	}
	_, err := sqlx_utils.NamedExec(context.Background(), r.conn, createLayer, args)
	return err
}
