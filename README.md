![Go](https://github.com/tkc/go-echo-server-sandbox/workflows/Go/badge.svg)

# go-echo-server-sandbox

echo server sandbox using ORM mapper and html template.

- github.com/labstack/echo
- github.com/jinzhu/gorm
- github.com/pilu/fresh


## Requirements

### Go

```
$ brew install goenv
$ goenv install 1.13.x
$ goenv global 1.13.x
$ goenv rehash
```

## Dependency

```sh
$ make dep
```

## Install golangci-lint

```
$ brew install golangci/tap/golangci-lint
$ brew upgrade golangci/tap/golangci-lint
```
[more...](https://github.com/golangci/golangci-lint#macos)

## Lint

```
$ golangci-lint run 
```

## Database (MySQL) Config

`note: cp config.yaml.example config.yaml`

update database config

``` yaml
app: local
port: :8080

database:
  name: <database name>
  name: <user name>
  password: <user passWord>
```

## Migration

``` sh
$ go run ./migrate/migrate.go
```

## Serve and Hot Reload

``` sh
$ fresh
```

## Create New User

``` sh
$ curl http://localhost:8080/user \
  -X POST \
  -H "Content-Type: application/json" \
  -d '{"Name": "userName","Age": 1}'
```

## Testing

``` sh
$ go test ./...
```
