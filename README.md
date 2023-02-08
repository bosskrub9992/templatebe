# templatebe

Generate migration file

```sh
    cd ./src/infrastructure/sqlcrepository/migration
    goose create <file_name> sql
```

Generating sqlc by command prompt

```sh
    cd ./src/infrastructure/sqlcrepository
    docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate
```

// TODO: make dockerfile work