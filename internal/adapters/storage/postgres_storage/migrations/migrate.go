package migrations

import (
	"context"
	"fmt"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/pkg/sqlx_utils"
	"github.com/jmoiron/sqlx"
)

type Migration struct {
	db *sqlx.DB
}

func NewMigration(db *sqlx.DB) *Migration {
	return &Migration{db: db}
}

func (m *Migration) createCommonTables() error {
	_, err := sqlx_utils.NamedExec(context.Background(), m.db, createCommonTablesQuery, make(map[string]any))
	return err
}

func (m *Migration) addUsers() error {
	_, err := sqlx_utils.NamedExec(context.Background(), m.db, addUsersQuery, make(map[string]any))
	return err
}

func (m *Migration) fillCommonTables() error {
	_, err := sqlx_utils.NamedExec(context.Background(), m.db, fillCommonTablesQuery, make(map[string]any))
	return err
}

func (m *Migration) initStoredProcedures() error {
	_, err := sqlx_utils.NamedExec(context.Background(), m.db, initStoredProceduresQuery, make(map[string]any))
	return err
}

func (m *Migration) createLayer(layer string) error {
	query := fmt.Sprintf(createLayerQuery, layer, layer, layer, layer, layer, layer, layer, layer, layer, layer,
		layer, layer, layer, layer, layer, layer, layer, layer, layer, layer, layer, layer, layer, layer, layer)
	_, err := sqlx_utils.NamedExec(context.Background(), m.db, query, make(map[string]any))
	return err
}

func (m *Migration) fillLayer(layer string) error {
	query := fmt.Sprintf(fillLayerQuery,
		layer, layer, layer, layer, layer, layer, layer, layer, layer, layer, layer, layer,
	)
	_, err := sqlx_utils.NamedExec(context.Background(), m.db, query, make(map[string]any))
	return err
}

func (m *Migration) dropLayer(layer string) error {
	query := fmt.Sprintf(dropLayerQuery, layer, layer, layer)
	_, err := sqlx_utils.NamedExec(context.Background(), m.db, query, make(map[string]any))
	return err
}

func (m *Migration) dropCommonTables() error {
	_, err := sqlx_utils.NamedExec(context.Background(), m.db, dropCommonTablesQuery, make(map[string]any))
	return err
}

const (
	testLayer = "test_layer"
)

func (m *Migration) UpProd() error {
	if err := m.createCommonTables(); err != nil {
		return err
	}
	if err := m.addUsers(); err != nil {
		return err
	}
	return m.initStoredProcedures()
}

func (m *Migration) UpTest() error {
	if err := m.createCommonTables(); err != nil {
		return err
	}
	if err := m.fillCommonTables(); err != nil {
		return err
	}
	if err := m.initStoredProcedures(); err != nil {
		return err
	}
	if err := m.createLayer(testLayer); err != nil {
		return err
	}
	return m.fillLayer(testLayer)
}

func (m *Migration) DownTest() error {
	if err := m.dropLayer(testLayer); err != nil {
		return err
	}
	return m.dropCommonTables()
}
