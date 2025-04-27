package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Irurnnen/gin-template/internal/database"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestHelloRepository_GetHelloMessage(t *testing.T) {
	// Mock database
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Mock query
	rows := sqlmock.NewRows([]string{"message"}).AddRow("Hello World")
	mock.ExpectQuery("SELECT 'Hello World' AS message").WillReturnRows(rows)

	// Initialize repository
	provider := database.NewProviderDB(sqlx.NewDb(db, "sqlmock"))
	logger := zap.NewNop()
	repo := NewHelloRepository(provider, logger)

	// Call method
	message, err := repo.GetHelloMessage()

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, "Hello World", message)
	assert.NoError(t, mock.ExpectationsWereMet())
}
