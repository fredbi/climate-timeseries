include hack/termcolor.mk

.PHONY: create-db
## Creates a database and apply migrations
create-db: export GOOSE_DRIVER=postgres
create-db: export GOOSE_DBSTRING=postgresql://postgres@localhost:5432/climate_timeseries?sslmode=disable
create-db:
	goose -dir ./db/migrations up
