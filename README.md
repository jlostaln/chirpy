# Chirpy

**Languages:**  
[🇬🇧 English](#english) • [🇫🇮 Suomi](#finnish)

<a name="english"></a>
## English Version

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
```
GET /api/chirps?author_id=123e4567-e89b-12d3-a456-426614174000&sort=desc
```

--------------------

# Chirpy

<a name="finnish"></a>
## Suomenkielinen versio

Chirpy on Go-kielellä kirjoitettu RESTful HTTP-palvelin, joka toimii yksinkertaisen sosiaalisen median backendina — ajattele riisuttua Twitter/X-kloonia. Se hallinnoi käyttäjätilejä, lyhyitä viestejä ("chirpejä"), autentikaatiota ja kolmannen osapuolen webhook-integraatioita.

## Miksi tämä on kiinnostava?

Chirpy demonstroi keskeisiä backend-ohjelmoinnin konsepteja yhdessä itsenäisessä projektissa:

- JWT-pohjainen autentikaatio ja refresh token -käytännöt
- Salasanojen hajautus ja turvallinen tunnistetietojen tallennus
- Täydet CRUD-operaatiot PostgreSQL-tietokantaan pohjautuen
- Tietokantamigraatiot Goosella ja tyyppiturvalliset kyselyt SQLC:llä
- Webhook-käsittely API-avainpohjaisella autorisoinnilla
- Selkeät routing- ja middleware-rakenteet Go:n standardikirjastolla

Hyvä referenssiprojekti kaikille, jotka haluavat oppia miten tuotantotason HTTP-palvelimet rakentuvat Go:lla.

## Vaatimukset

- [Go](https://go.dev/dl/) 1.22+
- [PostgreSQL](https://www.postgresql.org/download/) 15+
- [Goose](https://github.com/pressly/goose) (migraatioita varten)
- [SQLC](https://sqlc.dev/) (valinnainen, kyselyiden uudelleengenerointia varten)

## Asennus ja käyttöönotto

### 1. Kloonaa repositorio

```bash
git clone https://github.com/<käyttäjänimesi>/chirpy.git
cd chirpy
```

### 2. Asenna Go-riippuvuudet

```bash
go mod download
```

### 3. Luo PostgreSQL-tietokanta

```sql
CREATE DATABASE chirpy;
```

### 4. Määritä ympäristömuuttujat

Luo `.env`-tiedosto projektin juureen:

```env
DB_URL="postgres://<käyttäjä>:<salasana>@localhost:5432/chirpy?sslmode=disable"
JWT_SECRET="salainen-avaimesi"
POLKA_KEY="webhook-api-avaimesi"
PLATFORM="dev"
```

### 5. Aja tietokantamigraatiot

```bash
goose -dir sql/schema postgres "$DB_URL" up
```

### 6. Käynnistä palvelin

```bash
go build -o chirpy . && ./chirpy
```

Palvelin käynnistyy osoitteessa `http://localhost:8080`.

## API-kuvaus

| Metodi | Endpoint | Kuvaus |
|--------|----------|--------|
| POST | `/api/users` | Luo käyttäjä |
| PUT | `/api/users` | Päivitä käyttäjä (autentikaatio vaaditaan) |
| POST | `/api/login` | Kirjaudu sisään, vastaanota JWT |
| POST | `/api/refresh` | Päivitä access token |
| POST | `/api/revoke` | Peruuta refresh token |
| POST | `/api/chirps` | Luo chirp (autentikaatio vaaditaan) |
| GET | `/api/chirps` | Hae kaikki chirpit |
| GET | `/api/chirps/{chirpID}` | Hae yksittäinen chirp ID:llä |
| DELETE | `/api/chirps/{chirpID}` | Poista chirp (autentikaatio vaaditaan) |
| POST | `/api/polka/webhooks` | Polka webhook -käsittelijä (API-avain vaaditaan) |

### Query-parametrit — `GET /api/chirps`

| Parametri | Tyyppi | Kuvaus |
|-----------|--------|--------|
| `author_id` | `string` (UUID) | Suodata chirpit tietyn kirjoittajan mukaan |
| `sort` | `string` | Chirppien järjestys. Hyväksyy `asc` (oletus) tai `desc` |

**Esimerkki:**

```
GET /api/chirps?author_id=123e4567-e89b-12d3-a456-426614174000&sort=desc
```


