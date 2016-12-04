# go-echo-server-sandbox
echo server sandbox using ORM mapper and html template.


```
github.com/labstack/echo
github.com/jinzhu/gorm
```

# MySQL Database Config
please edit config.go

```
const DbUser = "<user>"
const DbPass = "<pass>"
const DbName = "<dbName>"
const DbHost = "127.0.0.1"
const DbPort = "3306"
const DbLocate = "Asia%2FTokyo"
```

# model
user table
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
To run the  test suites, you can run the command:

```
./client_ci.sh
```
