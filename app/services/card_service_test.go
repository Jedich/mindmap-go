package services

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"mindmap-go/app/services/forms"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"mindmap-go/app/repository"
	"mindmap-go/internal/database"
)

func createMockDB(t *testing.T) (*database.Database, sqlmock.Sqlmock) {
	var db *sql.DB
	var err error

	db, mock, err := sqlmock.New() // mock sql.DB
	assert.NoError(t, err)

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	assert.NoError(t, err)

	mockDB := database.Database{
		Connection: gormDB,
	}

	return &mockDB, mock
}

func TestCreateCardByService(t *testing.T) {
	db, mock := createMockDB(t)
	card := &forms.CardForm{
		Name:      "a",
		Text:      "",
		Color:     "",
		ParentID:  nil,
		CreatorID: 0,
		MapID:     0,
	}
	//mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()
	const sqlInsert = "INSERT INTO `cards` (`created_at`,`updated_at`,`deleted_at`,`name`,`text_data`,`color`,`parent_id`,`creator_id`,`map_id`) VALUES (?,?,?,?,?,?,?,?,?)"
	mock.ExpectExec(regexp.QuoteMeta(sqlInsert)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, "a", "", "", nil, 0, 0).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	newCard, err := NewCardService(repository.NewCardRepository(db)).CreateCard(card)

	assert.NoError(t, err)
	assert.NotEmpty(t, newCard)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestGetCardByMapIDNotExistByService(t *testing.T) {
	db, mock := createMockDB(t)
	mock.ExpectQuery("SELECT(.*)").WithArgs(0).
		WillReturnRows(sqlmock.NewRows(nil))

	l, err := repository.NewCardRepository(db).GetCardsByMapID(0)

	assert.NoError(t, err)
	assert.Empty(t, l)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestGetCardByIDNotExistByService(t *testing.T) {
	db, mock := createMockDB(t)
	mock.ExpectQuery("SELECT(.*)").WithArgs(1).
		WillReturnRows(sqlmock.NewRows(nil))

	l, err := NewCardService(repository.NewCardRepository(db)).GetCardByID(1)

	assert.NoError(t, err)
	assert.Empty(t, l)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestGetCardByMapIDExistsByService(t *testing.T) {
	db, mock := createMockDB(t)
	rows := sqlmock.NewRows([]string{`created_at`, `updated_at`, `deleted_at`, `name`, `text_data`, `color`, `parent_id`, `creator_id`, `map_id`}).
		AddRow(time.Now(), time.Now(), nil, "a", "", "", nil, 0, 0)
	mock.ExpectQuery("SELECT(.*)").WithArgs(0).
		WillReturnRows(rows)

	l, err := NewCardService(repository.NewCardRepository(db)).GetCardsByMapID(0)
	assert.NoError(t, err)
	assert.NotEmpty(t, l)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestRandomFilename(t *testing.T) {
	filename := NewCardService(repository.NewCardRepository(nil)).GetRandomFilename(5)
	assert.True(t, len([]rune(filename)) == 5)
}
