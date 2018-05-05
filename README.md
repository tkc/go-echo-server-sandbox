# go-echo-server-sandbox

echo server sandbox using ORM mapper and html template.

- github.com/labstack/echo
- github.com/jinzhu/gorm
- github.com/pilu/fresh


## Install Vendor Files

```sh
glide install
```

## MySQL Database Config

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

## User Model

``` go
type User struct {
		Id        int64
		Name      string
		Age       int
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt *time.Time
}
```

## Run Dev Server 

``` sh
$ fresh
```

## Create User

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
