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
  logger.Info("MotorJuego Run start")

	var event *termbox.Event

	//inicializamos el timer para poder pausar y todo el show
	motor.timer = time.NewTimer(motor.tickTime)
	motor.timer.Stop()

  sonido.timerCancion = time.NewTimer(sonido.tickTimeCancion)
  sonido.timerCancion.Stop()

  //preparamos el juego
  motor.ranking = NewRanking()
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
  			case <-motor.timer.C:
  				motor.tick() //desplazamos la pieza hacia abajo
        case <-sonido.timerCancion.C:
          sonido.Play()
      case <-motor.chanStop:
  				break loop
  			}
  		}
  	}
    logger.Info("MotorJuego Run end")
}

func (motor *MotorJuego) PreviewBoard() {
	motor.tableroInicio = true
}

func (motor *MotorJuego) Stop() {
	logger.Info("MotorJuego Stop start")

	if !motor.stopped {
		motor.stopped = true
		close(motor.chanStop)
	}
	motor.timer.Stop()
	//motor.aiTimer.Stop()

	logger.Info("MotorJuego Stop end")
}
func (motor *MotorJuego) NuevoJuego() {
	logger.Info("MotorJuego NewGame start")

	//iniciamos el nivel
	tablero.Limpiar()
	motor.tableroInicio = false
	motor.gameOver = false
	motor.tickTime = 480 * time.Millisecond
	motor.score = 0
	motor.nivel = 1
	motor.lineasBorradas = 0
  sonido.tickTimeCancion = 1* time.Second

loop:
	for {
		select {
		case <-motor.tecla.opciones_teclas:
		default:
			break loop
		}
	}
//hace que la pieza se desplace
	motor.QuitarPausa()

	logger.Info("MotorJuego NewGame end")
}

func (motor *MotorJuego) GameOver() {
	logger.Info("MotorJuego GameOver start")
	motor.Pause()
  //sonido.Stop()
	motor.gameOver = true
  interfaz.MostrarAnimacionGameOver()
loop:
	for {
		select {
		case <-motor.tecla.opciones_teclas:
		default:
			break loop
		}
	}

	motor.ranking.InsertScore(uint64(motor.score))
	motor.ranking.Save()

	logger.Info("MotorJuego GameOver end")
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
  /*if sonido.reproduciendo == true{
    sonido.ResetTimerS(0)
  }*/
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

func (motor *MotorJuego) AGregarLineasBorradas(lines int) {
	motor.lineasBorradas += lines
	if motor.lineasBorradas > 999999 {
		motor.lineasBorradas = 999999
	}

	switch lines {
	case 1:
		motor.SumaPts(40 * (motor.nivel + 1))
	case 2:
		motor.SumaPts(100 * (motor.nivel + 1))
	case 3:
		motor.SumaPts(300 * (motor.nivel + 1))
	case 4:
		motor.SumaPts(1200 * (motor.nivel + 1))
	}

	if motor.nivel < motor.lineasBorradas/10 {
		motor.LevelUp()
	}
}

func (motor *MotorJuego) SumaPts(pts int) {
	motor.score += pts
	if motor.score > 9999999 {
		motor.score = 9999999
	}
}

func (motor *MotorJuego) LevelUp() {
	if motor.nivel >= 30 {
		return
	}

	motor.nivel++
	switch {
	case motor.nivel > 29:
		motor.tickTime = 10 * time.Millisecond
	case motor.nivel > 25:
		motor.tickTime = 20 * time.Millisecond
	case motor.nivel > 19:
		// 50 to 30
		motor.tickTime = time.Duration(10*(15-motor.nivel/2)) * time.Millisecond
	case motor.nivel > 9:
		// 150 to 60
		motor.tickTime = time.Duration(10*(25-motor.nivel)) * time.Millisecond
	default:
		// 480 to 160
		motor.tickTime = time.Duration(10*(52-4*motor.nivel)) * time.Millisecond
	}
}
