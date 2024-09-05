IMAGE_NAME ?= tel-product
IMAGE_TAG ?= 0.0.0

DB_DRIVER ?= postgres
DB_STRING ?= postgresql://admin:admin@localhost:5432/product

postgres:
	docker run -itd \
	-p 5432:5432 \
	-e POSTGRES_USER=admin \
	-e POSTGRES_PASSWORD=admin \
	-e POSTGRES_DB=product \
	--name postgres \
	postgres:16-alpine

migrate-status:
	go run cmd/migration/*.go -dir=migrations $(DB_DRIVER) ${DB_STRING} status

migrate-up:
	go run cmd/migration/*.go -dir=migrations $(DB_DRIVER) ${DB_STRING} up

migrate-reset:
	go run cmd/migration/*.go -dir=migrations ${DB_DRIVER} ${DB_STRING} reset

run:
	go run cmd/api/*.go

build:
	docker build -t ${IMAGE_NAME}:${IMAGE_TAG} .