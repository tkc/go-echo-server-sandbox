# go-echo-server-sandbox

echo server sandbox using ORM mapper and html template.

- github.com/labstack/echo
- github.com/jinzhu/gorm
- github.com/pilu/fresh


## Requirements

### Go

```
brew install goenv
goenv install 1.13.x
goenv global 1.13.x
goenv rehash
```

## Install Vendor Files

```sh
glide install
>>>>>>> master
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

``` sh
$ go run ./migrate/migrate.go
```

## Serve and Hot Reload

``` sh
$ fresh
```

## Create Mew User

``` sh
$ curl http://localhost:8080/user \
  -X POST \
  -H "Content-Type: application/json" \
  -d '{"Name": "userName","Age": 1}'
```

## Testing

``` sh
$ go test ./models/user/ -v
```
