start-server:
	@go run cmd/server/main.go

start-cron:
	@go run cmd/cron/main.go

start-queue:
	@go run cmd/queue/main.go

# Install swaggo: https://github.com/swaggo/echo-swagger
swag-install:
	@go install github.com/swaggo/swag/cmd/swag@latest

# Regenerate swaggo
swag-generate:
	@swag fmt && swag init --parseDependency -d cmd/server,internal/controller -o cmd/server/docs

# Usage: make migration-new name=add_balance_column_users
migration-new:
	@migrate create -ext sql -dir ./migrations ${name}

# Usage: make migration-up database=postgres://blehbleh...
migration-up:
	@migrate -database ${database} -source file://migrations up

# Usage: make migration-up database=postgres://blehbleh...
migration-down:
	@migrate -database ${database} -source file://migrations down

# Usage: make mock-repo src=user
mock-repo:
	@mockgen -source=internal/repository/${src}/repository.go -destination=internal/repository/mock/mock_${src}/repository.go

# Usage: make mock-service src=user
mock-service:
	@mockgen -source=internal/service/${src}/service.go -destination=internal/service/mock/mock_${src}/service.go

# Usage: make mock-utils
mock-utils:
	@mockgen -source=internal/lib/utils/utils.go -destination=internal/lib/utils/mock/utils.go
