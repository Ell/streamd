package main

import (
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/wav"
	"log"
	"os"
	"time"
)

const SampleRate beep.SampleRate = 44100

type AudioPlayer struct {
	assetsPath string
}

func NewAudioPlayer(assetsPath string) (*AudioPlayer, error) {
	err := speaker.Init(SampleRate, SampleRate.N(time.Second/10))
	if err != nil {
		return nil, err
	}

	return &AudioPlayer{
		assetsPath: assetsPath,
	}, nil
}

func (p *AudioPlayer) Play(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}

	streamer, _, err := wav.Decode(f)
	if err != nil {
		return err
	}

	defer func(streamer beep.StreamSeekCloser) {
		err := streamer.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(streamer)

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done

	return nil
}
