# Chirpy

Chirpy is a RESTful HTTP server written in Go that powers a simple social media backend — think a stripped-down Twitter/X clone. It handles user accounts, short posts ("chirps"), authentication, and third-party webhook integrations.

## Why Care?

Chirpy demonstrates core backend engineering concepts in a single, self-contained project:

- JWT-based authentication and refresh token flows
- Password hashing and secure credential storage
- Full CRUD operations backed by a PostgreSQL database
- Database migrations with Goose and type-safe queries with SQLC
- Webhook handling with API key authorization
- Clean routing and middleware patterns using Go's standard library

It's a solid reference for anyone learning how production HTTP servers are structured in Go.

## Prerequisites

- [Go](https://go.dev/dl/) 1.22+
- [PostgreSQL](https://www.postgresql.org/download/) 15+
- [Goose](https://github.com/pressly/goose) (for migrations)
- [SQLC](https://sqlc.dev/) (optional, for regenerating queries)

## Installation & Setup

### 1. Clone the repository

```bash
git clone https://github.com/<your-username>/chirpy.git
cd chirpy
```

### 2. Install Go dependencies

```bash
go mod download
```

### 3. Set up PostgreSQL

Create a database:

```sql
CREATE DATABASE chirpy;
```

### 4. Configure environment variables

Create a `.env` file in the project root:

```env
DB_URL="postgres://<user>:<password>@localhost:5432/chirpy?sslmode=disable"
JWT_SECRET="your-secret-key"
POLKA_KEY="your-webhook-api-key"
PLATFORM="dev"
```

### 5. Run database migrations

```bash
goose -dir sql/schema postgres "$DB_URL" up
```

### 6. Start the server

```bash
go build -o chirpy . && ./chirpy
```

The server will start on `http://localhost:8080`.

## API Overview

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/users` | Create a user |
| PUT | `/api/users` | Update user (auth required) |
| POST | `/api/login` | Log in, receive JWT |
| POST | `/api/refresh` | Refresh access token |
| POST | `/api/revoke` | Revoke refresh token |
| POST | `/api/chirps` | Create a chirp (auth required) |
| GET | `/api/chirps` | Get all chirps |
| GET | `/api/chirps/{chirpID}` | Get a single chirp by ID |
| DELETE | `/api/chirps/{chirpID}` | Delete a chirp (auth required) |
| POST | `/api/polka/webhooks` | Polka webhook handler (API key required) |

### Query Parameters — `GET /api/chirps`

| Parameter | Type | Description |
|-----------|------|-------------|
| `author_id` | `string` (UUID) | Filter chirps by a specific author |
| `sort` | `string` | Sort order for chirps. Accepts `asc` (default) or `desc` |

**Example:**
GET /api/chirps?author_id=123e4567-e89b-12d3-a456-426614174000&sort=desc
