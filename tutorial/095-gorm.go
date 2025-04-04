package main

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Author struct {
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
}

type Blog struct {
	ID     int
	Author Author `gorm:"embedded"`
	/* When embedding, fields are copied to `Blog`
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	*/
	Upvotes int32
}

// gorm.Model definition columns
type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Setting default Values
type User struct {
	ID   string `gorm:"default:gen_random_uuid()"` // Updated to use uuid_generate_v4()
	Name string `gorm:"default:galeone"`
	Age  int64  `gorm:"default:18"`
}

func connectToDatabase() {
	print := fmt.Println
	databaseString := "host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(databaseString), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	_ = db // Use the database connection
	err = db.AutoMigrate(&User{})
	if err != nil {
		print(err)
	}
	// db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})

	user := &User{Name: "bob", Age: 69}
	result := db.Create(user)
	print(user.ID /* inserted data's primary key */, result.Error, result.RowsAffected)

	users := []User{{Name: "bob"}, {Name: "sydney"}, {Name: "troy"}}
	db.Create(users)
	for _, user := range users {
		print(user.ID)
	}
}
