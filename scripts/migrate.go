package main

import (
	"bakery-api/configs"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run scripts/migrate.go [up|down|force]")
		os.Exit(1)
	}

	command := os.Args[1]

	// Load configuration
	config := configs.GetConfig()

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.DbName,
	)

	m, err := migrate.New("file://scripts/migrations", dsn)
	if err != nil {
		log.Fatalf("Failed to create migration instance: %v", err)
	}
	defer m.Close()

	switch strings.ToLower(command) {
	case "up":
		upMigrate(m)
	case "down":
		downMigrate(m)
	case "force":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run scripts/migrate.go force <version>")
			os.Exit(1)
		}
		version := os.Args[2]
		versionNum, err := strconv.Atoi(version)
		if err != nil {
			log.Fatalf("Invalid version number: %v", err)
		}
		fmt.Printf("Forcing migration to version %s...\n", version)
		if err := m.Force(int(versionNum)); err != nil {
			log.Fatalf("Failed to force migration to version %s: %v", version, err)
		} else {
			fmt.Printf("Migration forced to version %s successfully.\n", version)
		}
	}
}

func upMigrate(m *migrate.Migrate) {
	fmt.Println("Running migrations up...")
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migrations up: %v", err)
	} else {
		fmt.Println("Migrations applied successfully.")
	}
}

func downMigrate(m *migrate.Migrate) {
	fmt.Println("Running migrations down...")
	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migrations down: %v", err)
	} else {
		fmt.Println("Migrations rolled back successfully.")
	}
}
