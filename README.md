# templatebe

### Prerequisite

    - golang 1.20
    - wire

### Run program

```sh
    go run ./cmd/serve
```

### Generate migration file

```sh
    cd ./src/infrastructure/sqlcrepository/migration
    goose create <file_name> sql
```

### Generating sqlc by command prompt

```sh
    cd ./src/infrastructure/sqlcrepository
    docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate
```

### Inject Dependency

1. install wire

```sh
    go install github.com/google/wire/cmd/wire@latest
```

2. run wire
```sh
    cd ./cmd/serve
    wire
```

### Docker

```sh
    docker build -t serve .
    docker run serve
```