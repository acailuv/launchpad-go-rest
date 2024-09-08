start-server:
	@go run cmd/server/main.go

start-cron:
	@go run cmd/cron/main.go

# Usage: make migration-new name=add_balance_column_users
migration-new:
	@migrate create -ext sql -dir ./migrations ${name}

# Usage: make migration-up database=postgres://blehbleh...
migration-up:
	@migrate -database ${database} -source file://migrations up

# Usage: make migration-up database=postgres://blehbleh...
migration-down:
	@migrate -database ${database} -source file://migrations down