# PKX API

A clean architecture Go API with JWT authentication.

## Project Structure

```
.
├── cmd/
│   └── api/                 # Application entry point
├── internal/
│   ├── domain/             # Enterprise business rules
│   │   ├── entity/         # Business entities
│   │   └── repository/     # Repository interfaces
│   ├── application/        # Application business rules
│   │   └── service/        # Use cases
│   ├── infrastructure/     # Frameworks and drivers
│   │   ├── persistence/    # Database implementations
│   │   └── web/           # Web framework implementations
│   └── presentation/       # Interface adapters
│       └── http/          # HTTP handlers
└── sql/                   # SQL queries for sqlc
```

## Setup

1. Install dependencies:

```bash
go mod download
```

2. Install sqlc:

```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

3. Generate database code:

```bash
sqlc generate
```

4. Run the application:

```bash
go run cmd/api/main.go
```

## Dependencies

- Echo (Web Framework)
- sqlc (SQL Compiler)
- JWT-Go (JWT Implementation)
- PostgreSQL (Database)
