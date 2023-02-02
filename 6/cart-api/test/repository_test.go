package test

import (
	"database/sql"
	"regexp"
	"richard-here/haioo-api/cart-api/database"
	"richard-here/haioo-api/cart-api/model"
	"richard-here/haioo-api/cart-api/repository"
	"testing"

	"github.com/go-test/deep"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Suite struct {
	suite.Suite
	DB   *database.DBInstance
	mock sqlmock.Sqlmock

	repository repository.Repo
	product    *model.Product
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)
	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	gormDb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}))
	s.DB = &database.DBInstance{Db: gormDb}
	require.NoError(s.T(), err)

	s.repository = repository.CreateRepository(*s.DB)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) Test_repository_GetProductsInCart() {
	var (
		name     = "test"
		quantity = 1
	)
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "products" WHERE "products"."deleted_at" IS NULL`,
	)).
		WillReturnRows(sqlmock.NewRows([]string{"code", "name", "quantity"}).
			AddRow(nil, name, quantity))

	res, err := s.repository.GetProductsInCart()
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal([]model.Product{{Code: uuid.Nil, Name: name, Quantity: quantity}}, res))
}

func (s *Suite) Test_repository_AddProductToCart() {
	var (
		id       = uuid.New()
		name     = "test"
		quantity = 2
	)
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		`INSERT INTO "products" ("code","name","quantity","deleted_at") VALUES ($1,$2,$3,$4)`,
	)).
		WithArgs(id, name, quantity, gorm.DeletedAt{}).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.repository.AddProductToCart(&model.Product{Code: id, Name: name, Quantity: quantity})
	require.NoError(s.T(), err)
}

func (s *Suite) Test_repository_UpdateProductInCart() {
	var (
		id       = uuid.New()
		name     = "test"
		quantity = 3
	)
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		`UPDATE "products" SET "name"=$1,"quantity"=$2,"deleted_at"=$3 WHERE "products"."deleted_at" IS NULL AND "code" = $4`,
	)).WithArgs(name, quantity, gorm.DeletedAt{}, id).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.repository.UpdateProductInCart(&model.Product{Code: id, Quantity: quantity, Name: name})
	require.NoError(s.T(), err)
}

func (s *Suite) Test_repository_VerifyProductExists() {
	var (
		name     = "test"
		quantity = 2
	)
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "products" WHERE name = $1 AND "products"."deleted_at" IS NULL`,
	)).
		WithArgs(name).
		WillReturnRows(sqlmock.NewRows([]string{"code", "name", "quantity"}).
			AddRow(nil, name, quantity))

	_, err := s.repository.VerifyProductExists(name)
	require.NoError(s.T(), err)
}
