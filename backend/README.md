# LTI ATS Backend

## Overview

Go-based backend service for the LTI Applicant Tracking System, implementing Domain-Driven Design (DDD) principles for robust candidate profile management and recruitment workflows.

## Project Structure

```
backend/
├── application/         # Application layer
│   └── usecases/       # Business logic use cases
│       └── create_candidate_profile.go
├── domain/             # Domain layer
│   ├── models/         # Domain entities
│   │   └── user.go
│   ├── repositories/   # Repository interfaces
│   └── services/       # Domain services
├── infrastructure/     # Infrastructure layer
│   ├── http/          # HTTP server setup
│   │   └── server.go
│   └── persistence/   # Database implementations
│       └── user_repository_impl.go
├── interfaces/         # Interface adapters
│   ├── controllers/   # HTTP handlers
│   │   └── candidate_controller.go
│   ├── presenters/    # Response formatters
│   └── repositories/  # Repository implementations
└── main.go            # Application entry point
```

## Technologies

- Go 1.22.4
- PostgreSQL
- CORS middleware
- Docker
- Testify for testing

## Getting Started

1. **Environment Setup**
```bash
cp .env.example .env
```

2. **Install Dependencies**
```bash
go mod download
```

3. **Database Setup**
```bash
docker compose up -d
```

4. **Run Migrations**
```bash
docker compose run migrate
```

5. **Start Server**
```bash
go run main.go
```

## API Endpoints

### POST /api/candidates
Creates a new candidate profile.

**Request Body:**
```json
{
  "name": "string",
  "contact_info": {
    "email": "string",
    "phone": "string"
  },
  "resume_link": "string",
  "current_position": "string",
  "location": "string"
}
```

## Testing

Run all tests:
```bash
go test ./...
```

Run specific package tests:
```bash
go test ./domain/...
go test ./application/...
```

## Project Documentation

- API Documentation
- Database Schema
- Architecture Decisions

## Contributing

See CONTRIBUTING.md in the root directory.

## License

This project is licensed under the MIT License - see the LICENSE file for details.