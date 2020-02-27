# go-echo-server-sandbox

echo server sandbox using ORM mapper and html template.

## Requirements

### Go

```
brew install goenv
goenv install 1.13.x
goenv global 1.13.x
goenv rehash
```

## Database (MySQL) Config

``` yaml
app: local
port: :8080

database:
  name: <DataBaseName>
  name: <UserName>
  password: <PassWord>
```

## Migration

``` bash
$ go run ./migrate/migrate.go
```

## Serve and Hot Reload

``` bash
$ fresh
```

## Create Mew User

``` bash
$ curl http://localhost:8080/user \
  -X POST \
  -H "Content-Type: application/json" \
  -d '{"Name": "userName","Age": 1}'
```

## Testing

``` bash
$ go test ./models/user/ -v
```
