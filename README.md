# Auth Service with Google OAuth and JWT

A full-featured authorization service implemented in Go using Gin, JWT, and Google OAuth 2.0. The project is containerized with Docker and uses PostgreSQL as a database. Live reload is supported via CompileDaemon.

---

## Project Structure

```
.
├── Dockerfile
├── docker-compose.yml
├── go.mod / go.sum
├── .env 
├── /api 
├── /infrastructure 
├── /adapters
├── /interfaces 
├── /handlers 
├── /models
└── /utils
```

---

### 1. User registration and login
- Manual registration 
- Login with password
- Password hashing with `bcrypt`
- Email confirmation 

### 2. Google OAuth 2.0
- User can log in via Google
- Receive `email`, `name`, `picture`
- Save new user in DB if it doesn't exist yet (`FindOrCreateUser`)
- Generate JWT token and automatically log user in
- JWT is sent as HTTP-only cookie

### 3. JWT Authorization
- Generate JWT token with role and email
- Middleware verifies the token and extracts data from it

###  4. PostgreSQL
- PostgreSQL 15 is used
- `auth_postgres` container
- `users` table
- All data (including `picture`, `role`, `verified`) is saved

---

##  Docker and Dev Workflow

### Dockerfile
- `golang:1.24-alpine` is used for lightweight image
- `CompileDaemon` is installed for automatic rebuilding
- `CMD` starts `CompileDaemon`, which monitors code changes and rebuilds the service

###  docker-compose.yml
- Two services: `app` and `db`
- `app`:
- is built from Dockerfile
- mounts local code (live update)
- uses `.env` file
- `db`:
- PostgreSQL 15
- port 5433 out (to avoid conflict with local Postgres)
- saves data to volume


## .env example
```
GOOGLE_CLIENT_ID=
GOOGLE_CLIENT_SECRET=
JWT_SECRET=
DB_USER=
DB_PASSWORD=
DB_NAME=
DB_PORT=
DB_HOST=
```

---

## To run the project

```bash
docker-compose up --build
```
