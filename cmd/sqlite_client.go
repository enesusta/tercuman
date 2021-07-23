package cmd

import (
	"database/sql"
	"github.com/enesusta/tercuman/pkg/model"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

//go:generate mockery --name=SqliteClient --output=../mocks/sqliteclientmock

type sqliteClient struct {
	conn *sql.DB
}

type SqliteClient interface {
	RetrieveTranslations() []model.Translation
	RetrieveTranslation(word string) model.Translation
}

const (
	GET_QUERY           string = "select word, translations, audio from tercuman"
	GET_WITH_WORD_QUERY        = "select word, translations, audio from tercuman where word = ?"
)

func NewSqliteClient(conn *sql.DB) SqliteClient {
	return &sqliteClient{conn: conn}
}

func (s sqliteClient) RetrieveTranslations() []model.Translation {
	var translations []model.Translation
	rows, err := s.conn.Query(GET_QUERY)
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()
	for rows.Next() {
		var word string
		var translationsRaw string
		var audio []byte

		err = rows.Scan(&word, &translationsRaw, &audio)
		translation := model.Translation{Word: word, Translations: translationsRaw, Audio: audio}
		translations = append(translations, translation)
	}

	return translations
}

func (s sqliteClient) RetrieveTranslation(word string) model.Translation {
	rows, err := s.conn.Query(GET_WITH_WORD_QUERY, word)
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()
	if rows.Next() {
		var wordRaw string
		var translationsRaw string
		var audio []byte

		err = rows.Scan(&wordRaw, &translationsRaw, &audio)
		translation := model.Translation{Word: word, Translations: translationsRaw, Audio: audio}
		return translation
	}

	return model.Translation{}
}
