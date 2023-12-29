package benchmark

import (
	"testing"
)

const benchN = 100_000

func TestUsers(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	resSqlx, err := UsersSqlxRepositoryBenchmark(benchN)
	if err != nil {
		panic(err)
	}
	t.Log(resSqlx)

	resGorm, err := UsersGormRepositoryBenchmark(benchN)
	if err != nil {
		panic(err)
	}
	t.Log(resGorm)
}
