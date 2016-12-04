# go-echo-server-sandbox
echo server sandbox using ORM mapper and html template.


github.com/labstack/echo

github.com/jinzhu/gorm


# model

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
