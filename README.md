# Igo Refrigeration API

API developed in Go for managing users, employees, clients, and other resources for a refrigeration company.  
The goal is to create a solid and scalable base with good organizational practices, using Docker and PostgreSQL.

---

## Technologies Used

- [Go](https://golang.org/) `v1.23`
- [Chi Router](https://github.com/go-chi/chi)
- [pgx](https://github.com/jackc/pgx) (PostgreSQL Driver)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- PostgreSQL

---

## How to run the project

### Prerequisites

- Docker

### 1. Clone the repository

```
git clone https://github.com/igosantana/igo-refrigeration-api.git
cd igo-refrigeration-api
```

### 2. Configure the env file

```
DATABASE_URL=postgres://user:password@postgres:5432/igorefrigeration
```

### 3. Bring up the containers

```
docker compose up -d
```

The application will be running at:

http://localhost:8090


## Author

Developed by Igo Santana


## License

This project is licensed under the MIT License.
