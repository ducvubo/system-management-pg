GOOSE_DBSTRING ?= "avnadmin:AVNS_nejNZ1u-tCiAtEqOUdh@tcp(mysql-39588c9a-vminhduc8-88ed.a.aivencloud.com:13890)/SystemManagementPG"
GOOSE_MIGRATION_DIR ?= sql/schema
GOOSE_DRIVER ?= mysql

# Tên của ứng dụng của bạn
APP_NAME := server

# Chạy ứng dụng

docker_build:
	docker-compose up -d --build
	docker-compose ps
docker_stop:
	docker-compose down

dev:
	go run ./cmd/$(APP_NAME)

air:
	air

docker_up:
	docker compose up -d

up_by_one:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up-by-one
# create new a migration
create_migration:
	@goose -dir=$(GOOSE_MIGRATION_DIR) create $(name) sql
upse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up
downse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) down
resetse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) reset

sqlgen:
	sqlc generate

swag:
	rm -rf ./cmd/swag/docs/* && swag init -g ./cmd/server/main.go -o ./cmd/swag/docs

.PHONY: dev downse upse resetse docker_build docker_stop docker_up swag

.PHONY: air
