package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User model
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Email     string    `gorm:"size:100;uniqueIndex;not null" json:"email"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func main() {
	// Database connection
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// Connect to database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}

	// Auto migrate - creates the users table
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Failed to migrate:", err)
	}

	fmt.Println("Database and users table created successfully!")

	// Optional: Insert a sample user
	sampleUser := User{
		Name:  "John Doe",
		Email: "john@example.com",
		Age:   30,
	}

	result := db.Create(&sampleUser)
	if result.Error == nil {
		fmt.Printf("Sample user created with ID: %d\n", sampleUser.ID)
	} else {
		fmt.Println("Sample user creation failed (may already exist):", result.Error)
	}

	// Show existing users
	var users []User
	db.Find(&users)
	fmt.Printf("\n📊 Total users in database: %d\n", len(users))
	for _, u := range users {
		fmt.Printf("  - ID: %d, Name: %s, Email: %s, Age: %d\n", u.ID, u.Name, u.Email, u.Age)
	}
}
