package main

import (
	"database/sql"
	"fmt"
	"github.com/enesusta/tercuman/cmd"
	"github.com/enesusta/tercuman/internal"
	"os"
)

const (
	AUDIO_OPTION string = "--audio"
)

func main() {
	conn, err := sql.Open("sqlite3", "./tercuman.sqlite")
	if err != nil {
		panic(err)
	}

	sqliteClient := cmd.NewSqliteClient(conn)
	audioPlayer := internal.NewAudioPlayer()

	args := os.Args[1:]
	word := args[0]
	audioCondition := args[1] == AUDIO_OPTION

	translation := sqliteClient.RetrieveTranslation(word)

	if audioCondition {
		fmt.Printf("%s = %s\n", translation.Word, translation.Translations)
		audioPlayer.PlayAudio(translation.Audio)
	} else {
		fmt.Printf("%s = %s\n", translation.Word, translation.Translations)
	}

}
