package repository

import (
	"database/sql"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
)

// NewSQLMock returns mocked database connection
func NewSQLMock() (*sql.DB, sqlmock.Sqlmock, error) {
	return sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
}

// CreateOriginalItemMock ...
func CreateOriginalItemMock(mock sqlmock.Sqlmock, arg CreateOriginalItemParams, expectedErr error) {
	if expectedErr != nil {
		mock.ExpectExec(createOriginalItem).
			WithArgs(arg.ID, arg.Name, arg.Path, arg.URL).
			WillReturnError(expectedErr)
	} else {
		t, _ := time.Parse("Mon Jan 2 15:04:05 -0700 MST 2006", "Mon Jan 2 15:04:05 -0700 MST 2006")
		mock.ExpectQuery(createOriginalItem).
			WithArgs(arg.ID, arg.Name, arg.Path, arg.URL).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "name", "path", "url", "created_at"}).
					AddRow(arg.ID, arg.Name, arg.Path, arg.URL, t),
			)
	}
}

// GetOriginalItemByIDMock ...
func GetOriginalItemByIDMock(mock sqlmock.Sqlmock, arg OriginalItem, expectedErr error) {
	if expectedErr != nil {
		mock.ExpectQuery(getOriginalItemByID).
			WithArgs(arg.ID).
			WillReturnError(expectedErr)
	} else {
		mock.ExpectQuery(getOriginalItemByID).
			WithArgs(arg.ID).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "name", "path", "url", "created_at"}).
					AddRow(arg.ID, arg.Name, arg.Path, arg.URL, arg.CreatedAt),
			)
	}
}

// GetOriginalItemsListMock ...
func GetOriginalItemsListMock(mock sqlmock.Sqlmock, arg GetOriginalItemsListParams, items []OriginalItem, expectedErr error) {
	if expectedErr != nil {
		mock.ExpectQuery(getOriginalItemsList).
			WithArgs(arg.Limit, arg.Offset).
			WillReturnError(expectedErr)
	} else {
		rows := sqlmock.NewRows([]string{"id", "name", "path", "url", "created_at"})
		for _, item := range items {
			rows.AddRow(item.ID, item.Name, item.Path, item.URL, item.CreatedAt)
		}
		mock.ExpectQuery(getOriginalItemsList).
			WithArgs(arg.Limit, arg.Offset).
			WillReturnRows(rows)
	}
}

// CreateResizedItemMock ...
func CreateResizedItemMock(mock sqlmock.Sqlmock, arg CreateResizedItemParams, expectedErr error) {
	if expectedErr != nil {
		mock.ExpectExec(createResizedItem).
			WithArgs(arg.ID, arg.OID, arg.Name, arg.Path, arg.URL, arg.Width, arg.Height).
			WillReturnError(expectedErr)
	} else {
		t, _ := time.Parse("Mon Jan 2 15:04:05 -0700 MST 2006", "Mon Jan 2 15:04:05 -0700 MST 2006")
		mock.ExpectQuery(createResizedItem).
			WithArgs(arg.ID, arg.OID, arg.Name, arg.Path, arg.URL, arg.Width, arg.Height).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "oid", "name", "path", "url", "width", "height", "created_at"}).
					AddRow(arg.ID, arg.OID, arg.Name, arg.Path, arg.URL, arg.Width, arg.Height, t),
			)
	}
}

// GetResizedItemByIDMock ...
func GetResizedItemByIDMock(mock sqlmock.Sqlmock, arg ResizedItem, expectedErr error) {
	if expectedErr != nil {
		mock.ExpectQuery(getResizedItemByID).
			WithArgs(arg.ID).
			WillReturnError(expectedErr)
	} else {
		mock.ExpectQuery(getResizedItemByID).
			WithArgs(arg.ID).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "oid", "name", "path", "url", "width", "height", "created_at"}).
					AddRow(arg.ID, arg.OID, arg.Name, arg.Path, arg.URL, arg.Width, arg.Height, arg.CreatedAt),
			)
	}
}

// GetResizedItemsListMock ...
func GetResizedItemsListMock(mock sqlmock.Sqlmock, oid uuid.UUID, items []ResizedItem, expectedErr error) {
	if expectedErr != nil {
		mock.ExpectQuery(getResizedItemsList).
			WithArgs(oid).
			WillReturnError(expectedErr)
	} else {
		rows := sqlmock.NewRows([]string{"id", "oid", "name", "path", "url", "width", "height", "created_at"})
		for _, item := range items {
			rows.AddRow(item.ID, item.OID, item.Name, item.Path, item.URL, item.Width, item.Height, item.CreatedAt)
		}
		mock.ExpectQuery(getResizedItemsList).
			WithArgs(oid).
			WillReturnRows(rows)
	}
}
