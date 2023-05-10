package services

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"mindmap-go/app/repository"
	"mindmap-go/app/services/forms"
	"regexp"
	"testing"
	"time"
)

func TestCreateMap(t *testing.T) {
	db, mock := createMockDB(t)
	card := &forms.MapForm{
		CreatorID:   0,
		Name:        "",
		Description: "",
	}
	//mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()
	sqlInsert := "INSERT INTO `maps` (`created_at`,`updated_at`,`deleted_at`,`name`,`desc`,`creator_id`) VALUES (?,?,?,?,?,?)"
	mock.ExpectExec(regexp.QuoteMeta(sqlInsert)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, "Unnamed map", "", 0).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlInsert = "INSERT INTO `cards` (`created_at`,`updated_at`,`deleted_at`,`name`,`text_data`,`color`,`parent_id`,`creator_id`,`map_id`) VALUES (?,?,?,?,?,?,?,?,?)"
	mock.ExpectExec(regexp.QuoteMeta(sqlInsert)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, "Mind Map", "Example long description for your new mind map.", "", nil, 0, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	newCard, err := NewMapService(repository.NewMapRepository(db)).CreateMap(card)

	assert.NoError(t, err)
	assert.NotEmpty(t, newCard)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestGetMapByIDExists(t *testing.T) {
	db, mock := createMockDB(t)
	rows := sqlmock.NewRows([]string{`created_at`, `updated_at`, `deleted_at`, `name`, `desc`, `creator_id`}).
		AddRow(time.Now(), time.Now(), nil, "Unnamed map", "", 0)
	mock.ExpectQuery("SELECT(.*)").WithArgs(1, 0).
		WillReturnRows(rows)

	l, err := NewMapService(repository.NewMapRepository(db)).GetMapByID(1, 0)
	assert.NoError(t, err)
	assert.NotEmpty(t, l)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
