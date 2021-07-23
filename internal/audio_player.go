package internal

import (
	"bytes"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
	"io"
	"log"
)

//go:generate mockery --name=AudioPlayer --output=../mocks/audioplayermock

type audioPlayer struct { }

type AudioPlayer interface {
	PlayAudio(audio []byte)
}

func NewAudioPlayer() AudioPlayer {
	return &audioPlayer{}
}

func (a audioPlayer) PlayAudio(audio []byte) {
	reader := bytes.NewReader(audio)

	d, err := mp3.NewDecoder(reader)
	if err != nil {
		log.Println(err)
	}

	c, err := oto.NewContext(d.SampleRate(), 2, 2, 8192)
	if err != nil {
		log.Println(err)
	}

	defer c.Close()

	p := c.NewPlayer()
	defer p.Close()

	if _, err := io.Copy(p, d); err != nil {
		log.Println(err)
	}
}




