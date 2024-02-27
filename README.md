# Nasdaq BE
Project is applied Clean Architecture and inspired from and based on [Golang Clean Architecture REST API example](https://github.com/AleksK1NG/Go-Clean-Architecture-REST-API)

#### Full list what has been used:
* [gin](https://github.com/gin-gonic/gin) - Web framework
* [gorm](https://gorm.io/) - Extensions to database/sql.
* [caarlos0](https://github.com/caarlos0/env) - Go configuration with fangs
* [zap](https://github.com/uber-go/zap) - Logger
* [validator](https://github.com/go-playground/validator) - Go Struct and Field validation
* [jwt-go](https://github.com/golang-jwt/jwt) - JSON Web Tokens (JWT)
* [uuid](https://github.com/google/uuid) - UUID
* [testify](https://github.com/stretchr/testify) - Testing toolkit
* [mockery](https://github.com/vektra/mockery) - Mocking framework

This project has 4 layer :

- Models Layer
- Repository Layer
- Usecase Layer
- Delivery Layer

### Requirement
- Golang 1.20 or highest
- MySQL 8.0 or highest
- Redis 6 or highest 

### Configurations
- Create `.env` file in the root directory following `.env.example`
  - Update `MYSQL_URI`,  following your configuration in docker-compose.yml
- Create `.env.testing` file in the root directory following `.env.example`
  - Update `MYSQL_URI`,  following your configuration in docker-compose.testing.yml

### Run API

```bash
make setup
make run
```

### Run Test

```bash
make test
```

### Database
- Redo migrations

```bash
make mysql-redo
```

- Create a new migration file

```bash
make migrate-create name=<filename>
```

### Tear down everything

```bash
make down
```
