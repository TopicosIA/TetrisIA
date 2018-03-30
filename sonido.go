package main

import (
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"os"
	"time"
)
func NuevoSonido() {
		sonido = &Sonido{
			tickTimeCancion: time.Second,
			reproduciendo: false,
		}
}
func (sonido *Sonido) Play() {
	//tetris audio

	f, _ := os.Open("tetris.mp3")
	s, format, _ := mp3.Decode(f)
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(s)
}

func (sonido *Sonido) Stop() {
	f, _ := os.Open("pum.mp3")
	s, format, _ := mp3.Decode(f)
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(s)

}

func (sonido *Sonido) ResetTimerS(duracion time.Duration) {
	if !sonido.timerCancion.Stop() {
		select {
		case <-sonido.timerCancion.C:
		default:
		}
	}
	if duracion == 0 {
		// duration 0 means tick time
		sonido.timerCancion.Reset(sonido.tickTimeCancion)//duracion del tickTime antes de parar
	} else {
		// duration is lock delay
		sonido.timerCancion.Reset(duracion)
	}
}
