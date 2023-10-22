package repository_test

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/google/uuid"
	"github.com/ronnachate/inventory-api-go/domain"
	"github.com/ronnachate/inventory-api-go/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository domain.UserRepository
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)
	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)
	dialector := postgres.New(postgres.Config{
		Conn:       db,
		DriverName: "postgres",
	})

	s.DB, _ = gorm.Open(dialector, &gorm.Config{})
	s.repository = repository.NewUserRepository(s.DB)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) Test_repository_GetByID() {
	var (
		id   = uuid.UUID{}
		name = "test-name"
	)
	s.T().Run("success", func(t *testing.T) {
		s.mock.ExpectQuery(regexp.QuoteMeta(
			`SELECT * FROM "users" WHERE id = $1 ORDER BY "users"."id" LIMIT 1`)).
			WithArgs(id.String()).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
				AddRow(id.String(), name))

		res, err := s.repository.GetByID(context.Background(), id.String())

		require.NoError(t, err)
		assert.Equal(t, id, res.ID)
		assert.Equal(t, name, res.Name)
	})

	s.T().Run("error", func(t *testing.T) {
		s.mock.ExpectQuery(regexp.QuoteMeta(
			`SELECT * FROM "users" WHERE id = $1 ORDER BY "users"."id" LIMIT 1`)).
			WithArgs(id.String()).
			WillReturnError(sql.ErrNoRows)

		_, err := s.repository.GetByID(context.Background(), id.String())

		assert.Error(t, err)
	})
}

func (s *Suite) Test_repository_GetUsers() {
	var (
		id   = uuid.UUID{}
		name = "test-name"
		page = 1
		rows = 10
	)
	s.T().Run("success", func(t *testing.T) {
		s.mock.ExpectQuery(regexp.QuoteMeta(
			`SELECT * FROM "users" LIMIT 10`)).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
				AddRow(id.String(), name))

		users, err := s.repository.GetUsers(context.Background(), page, rows)

		require.NoError(t, err)
		assert.Equal(t, 1, len(users))
		assert.Equal(t, id, users[0].ID)
		assert.Equal(t, name, users[0].Name)
	})

	s.T().Run("error", func(t *testing.T) {
		s.mock.ExpectQuery(regexp.QuoteMeta(
			`SELECT * FROM "users" LIMIT 10`)).
			WillReturnError(sql.ErrNoRows)

		_, err := s.repository.GetUsers(context.Background(), page, rows)

		assert.Error(t, err)
	})
}
