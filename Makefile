build:
	go build .

start-infra:
	cd infrastructure && docker-compose up

start:
	make build
	./clinic-scheduling-app

db-migration-up:
	export DATABASE_URL=postgres://admin-user:admin@localhost:5432/clinica?sslmode=disable
	dbmate -e DATABASE_URL up
