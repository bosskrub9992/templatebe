# github.com/bosskrub9992/templatebe

a monorepo template for Golang
<!-- TODO: finish docker compose -->

### Prerequisite

    - Golang 1.20 or newer
    - Docker
    - wire

### Getting Started (Deployment approach)

1. start docker compose
```sh
    cd ./docker
    docker-compose up --build
```

2. drop and create database, run
```sh
    docker exec -i postgres psql -U postgres -c "drop database if exists postgres_template_be" && \
    docker exec -i postgres psql -U postgres -c "create database postgres_template_be"
```

3. migrate database, execute sql file in all service

### Getting Started (Manual approach)

1. init go module working space, run (run only 1st time)
```sh
    go work init
    go work use ./corelib
    go work use -r ./service/*
```

2. sync dependency between module, run
```sh
    go work sync
```

3. start postgres, run
```sh
    docker run --rm -p 5432:5432 -e POSTGRES_PASSWORD=postgres --name postgres postgres:15.2-alpine
```

4. drop and create database, run
```sh
    docker exec -i postgres psql -U postgres -c "drop database if exists postgres_template_be" && \
    docker exec -i postgres psql -U postgres -c "create database postgres_template_be"
```

5. migrate database, execute sql file in all service


6. start service
```sh
    cd ./service/bff
    go run ./cmd/serve
```

### Generate migration file

```sh
    cd ./service/bff/src/infrastructure/sqlcrepository/migration
    goose create <file_name> sql
```

### Generating sqlc by command prompt

```sh
    cd ./service/bff/src/infrastructure/sqlcrepository
    docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate
```

### Inject Dependency

1. install wire

```sh
    go install github.com/google/wire/cmd/wire@latest
```

2. run wire
```sh
    cd ./service/bff/cmd/serve
    wire
```

### Docker

```sh
    cd ./service/bff
    docker build -t serve .
    docker run serve
```

### How to update dependency of corelib

run
```sh
    go mod tidy
```

### Layout
    - corelib: (core library) contains packages that are rarely changed and can be used across repos in service folder
    - docker: contains docker compose file
    - service: contains all repos of microservices
        - bff: (backend for frontend) 
            - cmd: (command) contains multiple main.go file
            - config: contains config.yaml file
            - src: contains packages that is usually changed
                - api: contains handler that is used for parsing request and return response
                - config: contains config struct that read from config.yaml
                - controller: contains business logics and repository interfaces
                - domain: contains database struct
                - model: contains request and response struct of the api
                - repository: contains structs that implement interfaces that is needed by controller 
                - router: manage route names and middlewares

        *api will only send/receive data to controller in struct in the model package to eliminate dependency.
        *controller will only send/receive data to repository in struct in the domain package to eliminate dependency.

### Design pattern
    - Clean architecture
    - Hexagonal architecture (port & adapter)

### API convention
    - RESTful API
    - route name: hyphen case
    - json body: camelCase (smaller size than snake case)
    - query param: camelCase (smaller size than snake case)
    - path param: camelCase (smaller size than snake case)
    - time: ISO8601 in UTC

### Database convention
    - table name: lower + snake case, in plural
    - enum: lower + snake case
    - time: ISO8601 in UTC

### Application convention
    - folder name: don't use camel and snake case. It should be short and clearly and it should not duplicate with the variable name that is tend to be used in application (Go convention)
    - unique name such as "productId" should be "productID"
    - https://github.com/uber-go/guide/blob/master/style.md

### Testing

when interface in controller has changed, run
```sh
    mockgen -source=./src/controller/customer.go -destination=./src/repository/mockrepository/customer.go -package=mockrepository
```