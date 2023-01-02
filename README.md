# templatebe

Generate migration file

```sh
    cd ./pkg/infrastructure/sqlcrepository/migration
    goose create <file_name> sql
```

Generating sqlc by command prompt

```sh
    cd ./pkg/infrastructure/sqlcrepository
    docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate
```

todo
- add validator
- add logger
- add wire