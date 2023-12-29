# Benchmark SQLX vs GORM

## Pure shell
### SQLX
```shell
sqlx.AddUser -- runs 100000 times   CPU:  416092 ns/op    RAM:    38 allocs/op  2160 bytes/op
sqlx.GetUser -- runs 100000 times   CPU:  353662 ns/op    RAM:    52 allocs/op  2592 bytes/op
```
### GORM
```shell
gorm.AddUser -- runs 100000 times   CPU:  596098 ns/op    RAM:    98 allocs/op  6631 bytes/op
gorm.GetUser -- runs 100000 times   CPU:  402725 ns/op    RAM:    82 allocs/op  4908 bytes/op
```

## Docker with Prometheus metrics
### SQLX
```shell
sqlx.AddUser -- runs 100000 times   CPU: 1932968 ns/op    RAM:    39 allocs/op  3437 bytes/op
sqlx.GetUser -- runs 100000 times   CPU: 1918009 ns/op    RAM:    53 allocs/op  3883 bytes/op
```
### GORM
```shell
gorm.AddUser -- runs 100000 times   CPU: 3703048 ns/op    RAM:   101 allocs/op  9355 bytes/op
gorm.GetUser -- runs 100000 times   CPU: 1990463 ns/op    RAM:    83 allocs/op  6454 bytes/op
```
