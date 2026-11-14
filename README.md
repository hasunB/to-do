# Go To-Do API

A robust and scalable RESTful API for a To-Do application built with Go, the Gin web framework, and PostgreSQL.

---

## 🚀 Features

- **Fast & Lightweight:** Built using [Gin](https://gin-gonic.com/), a high-performance HTTP web framework.
- **PostgreSQL Connection Pool:** Utilizes [`pgx/v5`](https://github.com/jackc/pgx) for efficient, production-grade database connection pooling.
- **Clean Architecture:** Modular structure separating domain logic, infrastructure, and delivery interfaces.
- **Live Reloading:** Pre-configured with [Air](https://github.com/air-verse/air) for fast local development.
- **Environment Configuration:** Managed via [`godotenv`](https://github.com/joho/godotenv) for secure, twelve-factor app configuration.

---

## 🛠️ Tech Stack

- **Language:** [Go](https://go.dev/) (1.27+)
- **Router / Web Framework:** [Gin Gonic](https://github.com/gin-gonic/gin)
- **Database:** [PostgreSQL](https://www.postgresql.org/)
- **Database Driver / Pool:** [pgx v5](https://github.com/jackc/pgx)
- **Dev Tool:** [Air](https://github.com/air-verse/air)

---

## 📁 Project Structure

```text
.
├── cmd/
│   └── api/
│       └── main.go                  # Application entry point
├── internals/
│   ├── domain/                      # Business entities & interfaces
│   ├── infrastructure/
│   │   ├── config/                  # Environment & app configuration
│   │   │   └── config.go
│   │   └── database/                # Database connection & pooling
│   │       └── postgres.go
│   ├── interface/                   # Delivery layer (HTTP handlers, routing)
│   │   ├── http/
│   │   │   ├── controllers/
│   │   │   └── middleware/
│   │   └── requests/
│   └── Utils/                       # Helper utilities
├── migrations/                      # SQL schema migration files
├── .air.toml                        # Hot-reload configuration for Air
├── .env.example                     # Example environment variables
├── go.mod                           # Go module definition
└── README.md
```

---

## ⚙️ Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) installed (version 1.22+ recommended)
- [PostgreSQL](https://www.postgresql.org/download/) installed and running
- [Air](https://github.com/air-verse/air) (optional, for live reloading)
  ```powershell
  go install github.com/air-verse/air@latest
  ```

---

### Installation & Setup

1. **Clone the repository:**
   ```bash
   git clone https://github.com/HasunBandara/to-do.git
   cd to-do
   ```

2. **Configure environment variables:**
   Copy the example environment file and fill in your PostgreSQL credentials:
   ```powershell
   copy .env.example .env
   ```

   Edit `.env`:
   ```env
   DATABASE_URL=postgres://username:password@localhost:5432/go-todo?sslmode=disable
   PORT=8080
   ```
   > **Note:** If your password contains special characters like `@`, ensure they are URL-encoded (e.g. `@` becomes `%40`).

3. **Install dependencies:**
   ```bash
   go mod download
   ```

---

### Running the Application

#### Development Mode (with Live Reloading)
```bash
air
```

#### Production / Standard Mode
```bash
go run ./cmd/api
```

The server will start listening at `http://localhost:8080`.

---

## 📡 API Endpoints

| Method | Endpoint | Description | Status |
| :--- | :--- | :--- | :--- |
| `GET` | `/` | Health check & database connection status | Public |

### Sample Response (`GET /`)
```json
{
  "database": "connected",
  "message": "pong is pong",
  "status": "success"
}
```

---

## 📄 License

This project is licensed under the [MIT License](LICENSE).
