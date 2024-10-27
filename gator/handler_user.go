package main

import (
    "database/sql"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "time"
)

// User represents the user model in the database
type User struct {
    ID        uuid.UUID
    CreatedAt time.Time
    UpdatedAt time.Time
    Name      string
}

// Global database connection
var db *sql.DB // Ensure to initialize this with your database connection

// userExists checks if a user already exists in the database
func userExists(name string) (bool, error) {
    var exists bool
    query := "SELECT exists(SELECT 1 FROM users WHERE name = $1)"
    err := db.QueryRow(query, name).Scan(&exists)
    if err != nil {
        return false, fmt.Errorf("error checking user existence: %v", err)
    }
    return exists, nil
}

// registerUser creates a new user in the database
func registerUser(name string) error {
    exists, err := userExists(name)
    if err != nil {
        return err
    }
    if exists {
        return fmt.Errorf("user '%s' already exists", name)
    }

    // Generate a new UUID for the user
    newUserID := uuid.New()
    now := time.Now()

    // Insert the new user into the database
    _, err = db.Exec("INSERT INTO users (id, created_at, updated_at, name) VALUES ($1, $2, $3, $4)",
        newUserID, now, now, name)
    if err != nil {
        return fmt.Errorf("error registering user: %v", err)
    }
    return nil
}

// loginUser authenticates a user by their name
func loginUser(name string) error {
    exists, err := userExists(name)
    if err != nil {
        return err
    }
    if !exists {
        return fmt.Errorf("couldn't find user: %s", name)
    }

    // Handle login logic here (e.g., setting current user, etc.)
    fmt.Printf("User '%s' logged in successfully.\n", name)
    return nil
}

// Example usage
func main() {
    // Initialize your database connection (db = ...)
    // Ensure db is properly connected before calling any functions

    // Example user interactions
    username := "kahya"
    
    // Register the user
    err := registerUser(username)
    if err != nil {
        fmt.Println("Register Error:", err)
    } else {
        fmt.Println("User registered successfully.")
    }

    // Attempt to log in the user
    err = loginUser(username)
    if err != nil {
        fmt.Println("Login Error:", err)
    }

    // Attempt to log in an unknown user
    err = loginUser("unknown")
    if err != nil {
        fmt.Println("Login Error:", err)
    }
}