package tests

import (
	"backend/initializers"
	"backend/models"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("DB_URL", "postgres://test_user:test_pass@localhost:5432/todo_test?sslmode=disable")
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.DB.AutoMigrate(&models.Category{}, &models.Todo{})
	initializers.DB.Exec("TRUNCATE TABLE todos RESTART IDENTITY CASCADE")
	initializers.DB.Exec("TRUNCATE TABLE categories RESTART IDENTITY CASCADE")

	code := m.Run()
	os.Exit(code)
}
