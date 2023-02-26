# templatebe

### Prerequisite

    - golang 1.20 or newer
    - Docker
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

### Layout
    - cmd: (command) contains multiple main.go file
    - config: contains config file
    - lib: (library) contains packages used in package src but are rarely changed
    - src: contains packages that is usually changed
        - model: contains request and response struct of the api
        - domain: contains database struct
        - router: manage route name 
        - api: contains handler that is used for parsing request and return response
        - controller: contains business logics and repository interfaces
        - repository: contains structs that implement interfaces that is needed by controller 

    *api will only send/receive data to controller in model format.
    *controller will only send/receive data to repository in domain format.

### Design pattern
    - Clean architecture

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