# Makefile для создания миграций!

# Переменные
DB_DSN := "postgres://postgres:postgres@localhost:5435/taskdb?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

# Создание новой миграции
migrate-new:
	migrate create -ext sql -dir ./migrations $(NAME)

# Линтер анализатор кодовой базы
lint:
	golangci-lint run --out-format=colored-line-number

# Применение миграций
migrate:
	$(MIGRATE) up

# Откат миграций
migrate-down:
	$(MIGRATE) down

# Запуск приложения
run:
	go run cmd/main.go

# Генерация кода из OpenAPI-файла для tasks
gen-tasks:
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go


# Генерация кода из OpenAPI-файла для users
gen-users:
	oapi-codegen -config openapi/.openapi -include-tags users -package users openapi/openapi.yaml > ./internal/web/users/api.gen.go