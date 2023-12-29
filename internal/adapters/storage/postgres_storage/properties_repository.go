package postgres_storage

import (
	"context"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/entities"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/interfaces"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/pkg/sqlx_utils"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type PropertiesPgRepository struct {
	conn *sqlx.DB
}

func (r *PropertiesPgRepository) GetAllProperties() ([]entities.Property, error) {
	properties := make([]entities.Property, 0)
	err := sqlx_utils.NamedSelect(context.Background(), r.conn, &properties, selectAllProperties, make(map[string]any))
	if err != nil {
		return nil, err
	}
	return properties, nil
}

func (r *PropertiesPgRepository) GetPropertiesByUnit(layer string, unit interfaces.SearchUnitDTO) ([]entities.Property, error) {
	args := map[string]any{
		"layer_name": layer,
		"lang":       unit.Lang,
		"unit_text":  unit.Text,
	}
	properties := make([]entities.Property, 0)
	err := sqlx_utils.NamedSelect(context.Background(), r.conn, &properties, selectPropertiesByUnit, args)
	if err != nil {
		return nil, err
	}
	return properties, nil
}

func (r *PropertiesPgRepository) SaveProperties(properties []string) ([]int, error) {
	args := map[string]any{
		"properties_array": pq.Array(properties),
	}
	id := make([]int, 0)
	err := sqlx_utils.NamedSelect(context.Background(), r.conn, &id, insertProperties, args)
	if err != nil {
		return nil, err
	}
	return id, nil
}
