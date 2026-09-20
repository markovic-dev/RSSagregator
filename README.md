<div align="center">

  <h1>⚡ RSS Feed Aggregator & Scraper Engine ⚡</h1>

  <p>
    <b>High-Performance, Concurrent Go Backend REST API</b>
  </p>

  <p>
    Snažan backend servis sa automatskim pozadinskim scraper-om u realnom vremenu, izgrađen po modernim Go standardima i arhitekturi.
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

## 🚀 O Projektu

Ovaj projekt predstavlja kompletan backend sustav za agregaciju i obradu RSS izvora vijesti. Sustav pokreće konurentni pozadinski radnik (*background worker*) koji u definiranima vremenskim intervalima paralelno preuzima, parsira XML strukture i skladišti najnovije članke u bazu podataka.

### ✨ Ključne Mogućnosti

- ⚡ **Ultra-brzi REST API:** Rutiranje zasnovano na laganom i brzom `go-chi/chi` ruteru.
- 🔐 **Custom Auth Middleware:** Sigurna autentifikacija korisnika putem jedinstvenih API ključeva (`Authorization: ApiKey <KEY>`).
- 🔄 **Concurrent Scraper Engine:** Pozadinski radnik koji koristi Go gorutine (`goroutines`) i `sync.WaitGroup` za bezbedno i paralelno preuzimanje više RSS feed-ova odjednom.
- 🗄️ **Type-Safe SQL Queries:** Korištenje `sqlc` generatora koji pretvara čiste SQL upite u 100% type-safe Go kod.
- 🛠️ **Database Migrations:** Upravljanje šemom baze podataka pomoću `goose` alata.
- 🔗 **Many-to-Many Relacije:** Napredno praćenje i otpraćivanje feed-ova po korisniku.

---

## 🛠️ Tehnološki Stog

* **Jezik:** [Go (Golang)](https://golang.org/)
* **Baza Podataka:** [PostgreSQL](https://www.postgresql.org/)
* **Router:** [go-chi/chi](https://github.com/go-chi/chi)
* **SQL Generator:** [sqlc](https://sqlc.dev/)
* **Migracije:** [goose](https://github.com/pressly/goose)
* **UUID:** `github.com/google/uuid`

---

## 📡 API Endpoints Overview

| Metoda | Endpoint | Auth | Opis |
| :---: | :--- | :---: | :--- |
| `POST` | `/v1/users` | ❌ | Registracija novog korisnika i generiranje API ključa |
| `GET` | `/v1/users` | 🔑 | Dohvaćanje profila trenutno autentificiranog korisnika |
| `POST` | `/v1/feeds` | 🔑 | Kreiranje novog RSS feed-a u sustavu |
| `GET` | `/v1/feeds` | ❌ | Pregled svih registriranih RSS feed-ova |
| `POST` | `/v1/feed_follows` | 🔑 | Zapraćivanje određenog RSS feed-a po `feed_id` |
| `GET` | `/v1/feed_follows` | 🔑 | Izlistavanje svih feed-ova koje korisnik prati |
| `DELETE` | `/v1/feed_follows/{feedFollowID}` | 🔑 | Otpraćivanje feed-a na temelju ID-a zapraćivanja |
| `GET` | `/v1/posts` | 🔑 | Preuzimanje najnovijih obrađenih članaka za zapraćene feed-ove |

---

## 💻 Brzo Pokretanje (Quick Start)

### 1. Preduvjeti
Uverite se da imate instalirano sljedeće na vašem sustavu:
- [Go (v1.20+)](https://golang.org/doc/install)
- [PostgreSQL](https://www.postgresql.org/download/)
- [goose CLI](https://github.com/pressly/goose) (za pokretanje migracija)

### 2. Kloniranje Repozitorijuma
```bash
git clone [https://github.com/TVOJE_USERNAME/rssagg.git](https://github.com/TVOJE_USERNAME/rssagg.git)
cd rssagg
```

### 3. Konfiguracija Okruženja (`.env`)
Kreirajte `.env` datoteku u korijenu projekta (možete kopirati iz `.env.example`):
```env
PORT=8080
DB_URL=postgres://postgres:postgres@localhost:5432/rssagg?sslmode=disable
```

### 4. Pokretanje Baznih Migracija
```bash
cd sql/schema
goose postgres "postgres://postgres:postgres@localhost:5432/rssagg?sslmode=disable" up
cd ../..
```

### 5. Pokretanje Aplikacije
```bash
go run .
```

Server će se pokrenuti na `http://localhost:8080` zajedno sa pozadinskim scraper radnikom! 🚀

---

## 🧪 Testiranje putem HTTP Klijenta

Prilikom slanja zahtjeva na zaštićene rute, obavezno proslijedite `Authorization` zaglavlje:

```http
Authorization: ApiKey VAŠ_API_KLJUČ_HERE
```

---

<div align="center">
  <sub>Izgrađeno sa ❤️ i Go jezikom.</sub>
</div>
