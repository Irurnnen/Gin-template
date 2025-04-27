package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestRepository_Ping(t *testing.T) {
	// Mock database
	db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	assert.NoError(t, err)
	defer db.Close()

	// Mock ping
	mock.ExpectPing()

	// Initialize repository
	logger := zap.NewNop()
	repo := NewRepositoryDB(sqlx.NewDb(db, "sqlmock"), logger)

	// Call method
	err = repo.Ping()

	// Assertions
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
