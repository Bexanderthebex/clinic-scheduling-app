build:
	go build .

start-infra:
	cd infrastructure && docker compose -f docker-compose.es.yaml up

start:
	make build
	./clinic-scheduling-app

db-migration-up:
	dbmate -u postgres://admin:password@localhost:5432/clinica?sslmode=disable up
