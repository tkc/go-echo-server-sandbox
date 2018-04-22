# go-echo-server-sandbox

echo server sandbox using ORM mapper and html template.

## Install Vendor Files

```
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

``` bash
$ go run ./migrate/migrate.go
```

## User Model

``` bash

type User struct {
		Id        int64
		Name      string
		Age       int
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt *time.Time
}
```

## Run Developing Server 

``` bash
$ fresh
```

## Create User

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
