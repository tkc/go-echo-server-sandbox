# go-echo-server-sandbox
echo server sandbox using ORM mapper and html template.

## Install Vendor Files

```
glide install
```

# MySQL Database Config
please edit config.yaml

```
app: local
port: :8080

database:
  name: <database name>
  user_name: <user name>
  password: <pass word>
```

# Migrate
```
go run ./migrate/migrate.go
```

# Model
```
type User struct {
		Id        int64
		Name      string
		Age       int
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt *time.Time
}
```

# Unit testing
To run the  test suites, you can run the command.

```
./client_ci.sh
```
