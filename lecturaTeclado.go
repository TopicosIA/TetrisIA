package main

import (
	"runtime"

	"github.com/nsf/termbox-go"
)
func NuevaTecla() *TeclaPresionada {
	return &TeclaPresionada{
		tecla_salir:     make(chan struct{}, 1),
		opciones_teclas: make(chan *termbox.Event, 8),
	}
}
//Lectura de tecla
func (tecla *TeclaPresionada) Run() {
	logger.Info("KeyInput Run start")

loop:
	for {
		select {
		case <-tecla.tecla_salir:
			break loop
		default:
		}
		event := termbox.PollEvent()
		if event.Type == termbox.EventKey && len(tecla.opciones_teclas) < 8 {
			select {
			case <-tecla.tecla_salir:
				break loop
			case tecla.opciones_teclas <- &event:
			}
		}
	}

	logger.Info("KeyInput Run end")
}

func (tecla *TeclaPresionada) ProcesarEvento(evento *termbox.Event) {
	if evento.Key == termbox.KeyCtrlI {
		// ctrl i to log stack trace
		buffer := make([]byte, 1<<16)
		length := runtime.Stack(buffer, true)
		logger.Debug("Stack trace", "buffer", string(buffer[:length]))
		return
	}

	if evento.Ch == 'q' || evento.Key == termbox.KeyCtrlC {
		if !tecla.stopped {
			tecla.stopped = true
			close(tecla.tecla_salir)
		}
		motor.Stop()
		return
	}
	//si ya perdimos o va iniciando al presionar space inicia el juego
	if motor.tableroInicio || motor.gameOver {
		if evento.Ch == 0{
			switch evento.Key{
			case termbox.KeySpace:
				motor.NuevoJuego()
			}
		}
		return
	}
  if motor.pausado{
    if evento.Ch == 'p'{
      motor.QuitarPausa()
    }
    return
  }

	//opciones de teclas
	if evento.Ch == 0 {
		switch evento.Key {
		case termbox.KeySpace:
			tablero.DescensoFigura()
		case termbox.KeyArrowUp:
			tablero.RotarFiguraDerecha()
		case termbox.KeyArrowDown:
			///board.MinoMoveDown()
		case termbox.KeyArrowLeft:
			tablero.MoverIzquierda()
		case termbox.KeyArrowRight:
			tablero.MoverDerecha()
		}
	} else {
		switch evento.Ch {
		case 's':
			tablero.RotarFiguraIzquierda()
		case 'd':
			tablero.RotarFiguraDerecha()
		case 'p':
			motor.Pause()
		case 'i':
			//engine.EnabledAi()
		}
	}
}
