package benchmark

import (
	"fmt"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/adapters/storage/postgres_storage"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/adapters/storage/postgres_storage/migrations"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/adapters/storage/postgres_storage/test"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/entities"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/interfaces"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/services/storage"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"go.uber.org/multierr"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

func benchAddUser(repo storage.UsersRepository, n int) func(b *testing.B) {
	return func(b *testing.B) {
		b.N = n
		for i := 0; i < b.N; i++ {
			err := repo.AddUser(interfaces.UserDTO{
				Name:     uuid.NewString(),
				Password: uuid.NewString(),
				Role:     uuid.NewString(),
			})
			if err != nil {
				panic(err)
			}
		}
	}
}

func benchGetUser(repo storage.UsersRepository, n int) func(b *testing.B) {
	return func(b *testing.B) {
		b.N = n
		for i := 0; i < b.N; i++ {
			_, err := repo.GetUser(entities.AuthRequest{
				Username: "admin1",
				Password: "12345",
			})
			if err != nil {
				panic(err)
			}
		}
	}
}

func UsersSqlxRepositoryBenchmark(n int) (res []string, err error) {
	dbContainer, sqlxDB, err := test.SetupTestDatabase()
	if err != nil {
		return nil, err
	}
	defer func() {
		err = multierr.Combine(err, test.TeardownTestDatabase(dbContainer, sqlxDB))
	}()
	sqlxRepo := postgres_storage.NewPostgresStorage(sqlxDB)

	migration := migrations.NewMigration(sqlxDB)
	err = migration.UpTest()
	if err != nil {
		return nil, err
	}

	addUser := benchAddUser(sqlxRepo, n)
	getUser := benchGetUser(sqlxRepo, n)

	resultsAddUser := testing.Benchmark(addUser)
	res = append(res, fmt.Sprintf("sqlx.AddUser -- runs %5d times\tCPU: %5d ns/op\tRAM: %5d allocs/op %5d bytes/op\n",
		resultsAddUser.N, resultsAddUser.NsPerOp(), resultsAddUser.AllocsPerOp(), resultsAddUser.AllocedBytesPerOp(),
	))

	resultsGetUser := testing.Benchmark(getUser)
	res = append(res, fmt.Sprintf("sqlx.GetUser -- runs %5d times\tCPU: %5d ns/op\tRAM: %5d allocs/op %5d bytes/op\n",
		resultsGetUser.N, resultsGetUser.NsPerOp(), resultsGetUser.AllocsPerOp(), resultsGetUser.AllocedBytesPerOp(),
	))

	err = migration.DownTest()
	if err != nil {
		return nil, err
	}

	return res, err
}

func UsersGormRepositoryBenchmark(n int) (res []string, err error) {
	dbContainer, sqlxDB, err := test.SetupTestDatabase()
	if err != nil {
		return nil, err
	}
	defer func() {
		err = multierr.Combine(err, test.TeardownTestDatabase(dbContainer, sqlxDB))
	}()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: sqlxDB}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	gormRepo := &UsersGormRepository{gormDB: gormDB}

	migration := migrations.NewMigration(sqlxDB)
	err = migration.UpTest()
	if err != nil {
		return nil, err
	}

	addUser := benchAddUser(gormRepo, n)
	getUser := benchGetUser(gormRepo, n)

	resultsAddUser := testing.Benchmark(addUser)
	res = append(res, fmt.Sprintf("gorm.AddUser -- runs %5d times\t\tCPU: %5d ns/op\t\tRAM: %5d allocs/op %5d bytes/op\n",
		resultsAddUser.N, resultsAddUser.NsPerOp(), resultsAddUser.AllocsPerOp(), resultsAddUser.AllocedBytesPerOp(),
	))

	resultsGetUser := testing.Benchmark(getUser)
	res = append(res, fmt.Sprintf("gorm.GetUser -- runs %5d times\t\tCPU: %5d ns/op\t\tRAM: %5d allocs/op %5d bytes/op\n",
		resultsGetUser.N, resultsGetUser.NsPerOp(), resultsGetUser.AllocsPerOp(), resultsGetUser.AllocedBytesPerOp(),
	))

	err = migration.DownTest()
	if err != nil {
		return nil, err
	}

	return res, err
}
