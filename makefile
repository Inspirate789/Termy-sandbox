cli := $(go env GOBIN)/termy-cli
CLI_SRC := $(wildcard cmd/termy-cli/*.go)

.PHONY: run swagger cli unit integration report build-frontend run-frontend build-backend run-backend

run:
	docker compose --env-file env/deploy.env up -d --build

build-frontend:
	docker compose --env-file env/deploy.env build frontend

run-frontend:
	docker compose --env-file env/deploy.env up frontend -d

build-backend:
	docker compose --env-file env/deploy.env build postgres pgadmin backend backend2 backend3 backend-mirror nginx -d

run-backend:
	docker compose --env-file env/deploy.env up postgres pgadmin backend backend2 backend3 backend-mirror nginx -d

swagger:
	swag fmt
	swag init --parseDependency --parseInternal --parseDepth 1 -g cmd/backend/main.go -o swagger/

cli: $(CLI_SRC)
	go install ./cmd/termy-cli

unit:
	export ALLURE_OUTPUT_PATH=$(PWD)/test-reports && go test -short -v ./...

integration:
	export ALLURE_OUTPUT_PATH=$(PWD)/test-reports && go test -v ./internal/adapters/storage/postgres_storage/integration

e2e:
	export ALLURE_OUTPUT_PATH=$(PWD)/test-reports && go test -v ./internal/adapters/server

report:
	allure serve ./test-reports/allure-results

bench:
	go test -v ./internal/adapters/storage/postgres_storage/benchmark
	docker compose --env-file env/deploy.env up benchmark -d --build
	docker compose --env-file env/deploy.env up prometheus -d
	# sleep 5
	# curl -X POST http://localhost:2112/bench/sqlx?n=100000
	# curl -X POST http://localhost:2112/bench/gorm?n=100000
	# docker compose --env-file env/deploy.env down
	# http://localhost:9090/graph?g0.expr=go_memstats_alloc_bytes&g0.tab=0&g0.stacked=1&g0.show_exemplars=0&g0.range_input=10m&g1.expr=go_memstats_alloc_bytes_total&g1.tab=0&g1.stacked=1&g1.show_exemplars=0&g1.range_input=10m&g2.expr=go_memstats_heap_objects&g2.tab=0&g2.stacked=1&g2.show_exemplars=0&g2.range_input=10m&g3.expr=go_memstats_mallocs_total&g3.tab=0&g3.stacked=1&g3.show_exemplars=0&g3.range_input=10m&g4.expr=go_memstats_stack_inuse_bytes&g4.tab=0&g4.stacked=1&g4.show_exemplars=0&g4.range_input=10m&g5.expr=go_gc_duration_seconds_sum&g5.tab=0&g5.stacked=1&g5.show_exemplars=0&g5.range_input=10m&g6.expr=process_cpu_seconds_total&g6.tab=0&g6.stacked=1&g6.show_exemplars=0&g6.range_input=10m
