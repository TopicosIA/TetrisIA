package main

import (
  "time"
	"github.com/nsf/termbox-go"
)

func NuevoMotorJuego() {
	motor = &MotorJuego{
		chanStop: make(chan struct{}, 1),
		gameOver: false,
		tableroInicio: true,//para que se vea los mensajes de inicio
		tickTime: time.Hour,
	}
}
func (motor *MotorJuego) Run() {
  logger.Info("Engine Run start")

	var event *termbox.Event

	//inicializamos el timer para poder pausar y todo el show
	motor.timer = time.NewTimer(motor.tickTime)
	motor.timer.Stop()

	//preparamos el juego
  tablero.Limpiar()
	interfaz.refrescarPantalla()

	//creamos una go rutina para siempre estar escuchando al teclado
  motor.tecla = NuevaTecla()
	go motor.tecla.Run()

  loop:
  	for {
  		select {
  		case <-motor.chanStop:
  			break loop
  		default:
  			select {
  			case event = <-motor.tecla.opciones_teclas:
  				motor.tecla.ProcesarEvento(event)
  				interfaz.refrescarPantalla()
  			//case <-motor.timer.C:
  			//	motor.tick()
      case <-motor.chanStop:
  				break loop
  			}
  		}
  	}
    logger.Info("Engine Run end")
}

func (motor *MotorJuego) PreviewBoard() {
	motor.tableroInicio = true
}

func (motor *MotorJuego) Stop() {
	logger.Info("Engine Stop start")

	if !motor.stopped {
		motor.stopped = true
		close(motor.chanStop)
	}
	//motor.timer.Stop()
	//motor.aiTimer.Stop()

	logger.Info("Engine Stop end")
}
func (motor *MotorJuego) NuevoJuego() {
	logger.Info("Engine NewGame start")

	//iniciamos el nivel
	tablero.Limpiar()
	motor.tableroInicio = false
	motor.gameOver = false
	/*motor.tickTime = 480 * time.Millisecond
	motor.score = 0
	motor.nivel = 1
	motor.lineasBorradas = 0

loop:
	for {
		select {
		case <-motor.tecla.opciones_teclas:
		default:
			break loop
		}
	}

	motor.QuitarPausa()
*/
	logger.Info("Engine NewGame end")
}

func (motor *MotorJuego) GameOver() {
	logger.Info("Engine GameOver start")

	motor.Pause()
	motor.gameOver = true

	//view.ShowGameOverAnimation()

loop:
	for {
		select {
		case <-motor.tecla.opciones_teclas:
		default:
			break loop
		}
	}

	//engine.ranking.InsertScore(uint64(engine.score))
	//engine.ranking.Save()

	logger.Info("Engine GameOver end")
}

func (motor *MotorJuego) Pause() {
	if !motor.timer.Stop() {
		select {
		case <-motor.timer.C:
		default:
		}
	}
	motor.pausado = true
}

func (motor *MotorJuego) QuitarPausa(){
  motor.timer.Reset(motor.tickTime)
	motor.pausado = false
}
//hace bajar la figura cada cierto tiempo
func (motor *MotorJuego) tick() {
	tablero.MoverFiguraAbajo()
	interfaz.refrescarPantalla()
}

func (motor *MotorJuego) ResetTimer(duracion time.Duration) {
	if !motor.timer.Stop() {
		select {
		case <-motor.timer.C:
		default:
		}
	}
	if duracion == 0 {
		// duration 0 means tick time
		motor.timer.Reset(motor.tickTime)//duracion del tickTime antes de parar
	} else {
		// duration is lock delay
		motor.timer.Reset(duracion)
	}
}
