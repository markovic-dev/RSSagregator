<div align="center">

  <h1>⚡ RSS Feed Aggregator & Scraper Engine ⚡</h1>

  <p>
    <b>High-Performance, Concurrent Go Backend REST API</b>
  </p>

  <p>
    A robust backend service featuring a real-time concurrent background scraper, built using modern Go architecture and best practices.
  </p>

  <br />

  <!-- BADGES -->
  <a href="https://golang.org/">
    <img src="https://img.shields.io/badge/Go-1.20+-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go" />
  </a>
  <a href="https://www.postgresql.org/">
    <img src="https://img.shields.io/badge/PostgreSQL-15+-4169E1?style=for-the-badge&logo=postgresql&logoColor=white" alt="PostgreSQL" />
  </a>
  <a href="https://github.com/go-chi/chi">
    <img src="https://img.shields.io/badge/Router-Chi_v5-000000?style=for-the-badge&logo=go&logoColor=white" alt="Chi Router" />
  </a>
  <a href="https://sqlc.dev/">
    <img src="https://img.shields.io/badge/SQL-sqlc_Type_Safe-FF6C37?style=for-the-badge&logo=postgresql&logoColor=white" alt="sqlc" />
  </a>

</div>

---

## 🚀 Overview

This project is a full-featured backend system for RSS feed aggregation and data processing. It runs a concurrent background worker that periodically fetches, parses XML structures, and stores the latest articles in a PostgreSQL database in parallel.

### ✨ Key Features

- ⚡ **Ultra-Fast REST API:** Powered by the lightweight and fast `go-chi/chi` router.
- 🔐 **Custom Auth Middleware:** Secure user authentication using custom API keys (`Authorization: ApiKey <KEY>`).
- 🔄 **Concurrent Scraper Engine:** Background worker utilizing Go `goroutines` and `sync.WaitGroup` to safely fetch multiple RSS feeds in parallel.
- 🗄️ **Type-Safe SQL Queries:** Uses `sqlc` to generate 100% type-safe Go code directly from raw SQL queries.
- 🛠️ **Database Migrations:** Database schema management using `goose`.
- 🔗 **Many-to-Many Relationships:** Robust feed following/unfollowing logic per user.

---

## 🛠️ Tech Stack

* **Language:** [Go (Golang)](https://golang.org/)
* **Database:** [PostgreSQL](https://www.postgresql.org/)
* **Router:** [go-chi/chi](https://github.com/go-chi/chi)
* **SQL Generator:** [sqlc](https://sqlc.dev/)
* **Migrations:** [goose](https://github.com/pressly/goose)
* **UUID:** `github.com/google/uuid`

---

## 📡 API Endpoints Overview

| Method | Endpoint | Auth | Description |
| :---: | :--- | :---: | :--- |
| `POST` | `/v1/users` | ❌ | Create a new user and generate an API key |
| `GET` | `/v1/users` | 🔑 | Fetch current authenticated user's profile |
| `POST` | `/v1/feeds` | 🔑 | Create a new RSS feed |
| `GET` | `/v1/feeds` | ❌ | Get all registered RSS feeds |
| `POST` | `/v1/feed_follows` | 🔑 | Follow a specific RSS feed by `feed_id` |
| `GET` | `/v1/feed_follows` | 🔑 | List all RSS feeds followed by the user |
| `DELETE` | `/v1/feed_follows/{feedFollowID}` | 🔑 | Unfollow an RSS feed by follow record ID |
| `GET` | `/v1/posts` | 🔑 | Retrieve latest parsed articles from followed feeds |

---

## 💻 Quick Start

### 1. Prerequisites
Make sure you have the following installed on your system:
- [Go (v1.20+)](https://golang.org/doc/install)
- [PostgreSQL](https://www.postgresql.org/download/)
- [goose CLI](https://github.com/pressly/goose) (for database migrations)

### 2. Clone the Repository
```bash
git clone [https://github.com/YOUR_USERNAME/rssagg.git](https://github.com/YOUR_USERNAME/rssagg.git)
cd rssagg
```

### 3. Environment Configuration (`.env`)
Create a `.env` file in the root directory (you can copy `.env.example`):
```env
PORT=8080
DB_URL=postgres://postgres:postgres@localhost:5432/rssagg?sslmode=disable
```

### 4. Run Database Migrations
```bash
cd sql/schema
goose postgres "postgres://postgres:postgres@localhost:5432/rssagg?sslmode=disable" up
cd ../..
```

### 5. Run the Application
```bash
go run .
```

The server will start at `http://localhost:8080` along with the background scraper worker! 🚀

---

## 🧪 Authentication & Testing

When sending requests to protected routes, make sure to include the `Authorization` header:

```http
Authorization: ApiKey YOUR_API_KEY_HERE
```

---

<div align="center">
  <sub>Built with ❤️ and Go.</sub>
</div>
