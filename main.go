package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

// User represents a user in the database
type User struct {
	ID   int64
	Name string
	Age  int
}

func main() {
	// Create a new engine for MySQL
	engine, err := xorm.NewEngine("mysql", "username:password@tcp(localhost:3306)/database_name?charset=utf8")
	if err != nil {
		panic(err)
	}

	// Create the users table if it doesn't exist
	err = engine.Sync2(new(User))
	if err != nil {
		panic(err)
	}

	// Insert a new user into the database
	user := &User{Name: "John Doe", Age: 30}
	_, err = engine.Insert(user)
	if err != nil {
		panic(err)
	}

	// Query all users from the database
	users := make([]User, 0)
	err = engine.Find(&users)
	if err != nil {
		panic(err)
	}

	// Print the retrieved users
	for _, u := range users {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", u.ID, u.Name, u.Age)
	}
}