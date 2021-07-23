package cmd

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/enesusta/tercuman/pkg/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Should_Retrieve_All_Translation_That_Sqlite_Has(t *testing.T) {
	// Given
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()
	sqliteClient := NewSqliteClient(db)

	mockAudio := []byte{1, 2, 3}
	mockTranslation := model.Translation{Word: "any", Translations: "her", Audio: mockAudio}
	rows := sqlmock.NewRows([]string{"word", "translations", "audio"}).
		AddRow("any", "her", mockAudio)

	mock.ExpectQuery(GET_QUERY).WillReturnRows(rows)

	// When
	translations := sqliteClient.RetrieveTranslations()

	// Then
	assert.NotNil(t, translations)
	assert.NoError(t, err)
	assert.Equal(t, mockTranslation, translations[0])
	assert.Equal(t, len(translations), 1)
}

func Test_Should_Retrieve_Translation_With_Given_Word(t *testing.T) {
	// Given
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()
	sqliteClient := NewSqliteClient(db)

	mockAudio := []byte{1, 2, 3}
	mockTranslation := model.Translation{Word: "any", Translations: "her", Audio: mockAudio}
	rows := sqlmock.NewRows([]string{"word", "translations", "audio"}).
		AddRow("any", "her", mockAudio)

	mock.ExpectQuery(GET_QUERY).WithArgs("any").WillReturnRows(rows)

	// When
	translation := sqliteClient.RetrieveTranslation("any")

	// Then
	assert.NotNil(t, translation)
	assert.NoError(t, err)
	assert.Equal(t, mockTranslation, translation)
}
