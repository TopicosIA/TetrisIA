package main

import (
  "time"
  "fmt"
  "math/rand"
  "github.com/nsf/termbox-go"
)


func NuevaInterfaz(){
  err := termbox.Init()
	if err != nil {
		panic(err)
	}

	termbox.SetInputMode(termbox.InputEsc)
	termbox.Flush()
	interfaz = &Interfaz{}
}

func (interfaz *Interfaz) Stop() {
	logger.Info("View Stop start")

	termbox.Close()

	logger.Info("View Stop end")
}

func (interfaz *Interfaz) dibujaFondo() {

  posX := posX //comenzamos en la posicion X
	posY := posY //comenzamos en la posicion Y
	xFinal := posX + tablero.width*2 + 4
	yFinal := posY + tablero.height + 2
  // dibujando marco del tablero
	for x := posX; x < xFinal; x++ {
		for y := posY; y < yFinal; y++ {
			if x == posX || x == posX+1 || x == xFinal-1 || x == xFinal-2 ||
				y == posY || y == yFinal-1 {
				termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorWhite)
			} else {
				termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorBlack)
			}
		}
	}
  // vista previa de la siguiente figura

  posX = posX + tablero.width*2 + 8
	posY = posY
  //dibuja una linea azul para el fondo de las letras
  for x := posX; x < posX+15; x++ {
    termbox.SetCell(x, posY, ' ', termbox.ColorDefault, termbox.ColorBlue)
  }
  interfaz.dibujaTexto(posX+3, posY, "SIG. PIEZA", termbox.ColorWhite, termbox.ColorBlue)
  posY+=2
  xFinal = posX + 14
	yFinal = posY + 6
  for x := posX; x < xFinal; x++ {
		for y := posY; y < yFinal; y++ {
			if x == posX || x == posX+1 || x == xFinal-1 || x == xFinal-2 ||
				y == posY || y == yFinal-1 {
				termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorWhite)
			} else {
				termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorBlack)
			}
		}
	}
  //Fondo de Score
  posX=4
  posY=5
  for x := posX; x < posX+12; x++ {
    termbox.SetCell(x, posY, ' ', termbox.ColorDefault, termbox.ColorBlue)
  }
  for x := posX; x < posX+12; x++ {
    termbox.SetCell(x, posY+1, ' ', termbox.ColorDefault, termbox.ColorWhite)
  }
  //fondo de NIVEL
  posX=4
  posY=10
  for x := posX; x < posX+12; x++ {
    termbox.SetCell(x, posY, ' ', termbox.ColorDefault, termbox.ColorBlue)
  }
  for x := posX; x < posX+12; x++ {
    termbox.SetCell(x, posY+1, ' ', termbox.ColorDefault, termbox.ColorWhite)
  }
  //fondo de LINEAS BORRADAS
  posX=4
  posY=15
  for x := posX; x < posX+12; x++ {
    termbox.SetCell(x, posY, ' ', termbox.ColorDefault, termbox.ColorBlue)
  }
  for x := posX; x < posX+12; x++ {
    termbox.SetCell(x, posY+1, ' ', termbox.ColorDefault, termbox.ColorWhite)
  }
  //fondo control de teclas
  for x := 9; x < 62; x++ {
    termbox.SetCell(x,26, ' ', termbox.ColorDefault, termbox.ColorWhite)
  }
}

func (interfaz *Interfaz) dibujaTexto(x int, y int, texto string, fg termbox.Attribute, bg termbox.Attribute) {
	for index, char := range texto {
		termbox.SetCell(x+index, y, rune(char), fg, bg)
	}
}

func (interfaz *Interfaz) dibujarTextoCentroTablero(y int, text string, fg termbox.Attribute, bg termbox.Attribute) {
	x := tablero.width - (len(text)+1)/2 + posX + 2
	for index, char := range text {
		termbox.SetCell(index+x, y, rune(char), fg, bg)
	}
}

func (interfaz *Interfaz) dibujaTextoTablero() {
	posX := posX + tablero.width*2 + 8
	posY := posY + 18

	interfaz.dibujaTexto(7, 5, "SCORE", termbox.ColorWhite, termbox.ColorBlue)
	interfaz.dibujaTexto(5, 6, fmt.Sprintf("%5d",motor.score), termbox.ColorBlack, termbox.ColorWhite)

  interfaz.dibujaTexto(7, 10, "NIVEL", termbox.ColorWhite, termbox.ColorBlue)
	interfaz.dibujaTexto(6, 11, fmt.Sprintf("%4d", motor.nivel), termbox.ColorBlack, termbox.ColorWhite)

  interfaz.dibujaTexto(6, 15, "LINEAS", termbox.ColorWhite, termbox.ColorBlue)
	interfaz.dibujaTexto(4, 16, fmt.Sprintf("%6d", motor.lineasBorradas), termbox.ColorBlack, termbox.ColorWhite)
  //
  interfaz.dibujaTexto(12, 25, "     ←       s      <SPC>     d,↑      →", termbox.ColorWhite, termbox.ColorBlack)
  interfaz.dibujaTexto(12, 26, " izquierda    ↺     bajar      ↻    derecha", termbox.ColorBlack, termbox.ColorWhite)

  interfaz.dibujaTexto(posX, posY, "j    - jugar IA", termbox.ColorWhite, termbox.ColorBlack)
	posY++
	interfaz.dibujaTexto(posX, posY, "e    - entrenar IA", termbox.ColorWhite, termbox.ColorBlack)
  posY++
	interfaz.dibujaTexto(posX, posY, "p    - pausa", termbox.ColorWhite, termbox.ColorBlack)
	posY++
	interfaz.dibujaTexto(posX, posY, "q    - salir", termbox.ColorWhite, termbox.ColorBlack)
}


func (interfaz *Interfaz) dibujarPantallaInicio() {
	y := posY + 3
	interfaz.dibujarTextoCentroTablero(y, "SPACE para iniciar", termbox.ColorWhite, termbox.ColorBlack)
	if motor.tableroInicio {
		return
	}
}

func (interfaz *Interfaz) refrescarPantalla() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	interfaz.dibujaFondo()
	interfaz.dibujaTextoTablero()
  //diferentes estados del juego
  if motor.tableroInicio {
		tablero.DibujarTablero()
		interfaz.dibujarPantallaInicio()
	} else if motor.gameOver {
		interfaz.dibujarGameOver()
		interfaz.dibujaRankingScores()
	} else if motor.pausado {
		interfaz.dibujarPausa()
	} else {//empieza el juego
    tablero.DibujarTablero()
		tablero.DibujarFiguraPrevia()
    tablero.DibujarFiguraActual()
    tablero.DibujarSombraFigura()
	}

	termbox.Flush()
}
func (interfaz *Interfaz) DibujaVistaPreviaFigura(x int, y int, color termbox.Attribute, rotacion int, len int) {
	var char1 rune
	var char2 rune
	if rotacion < 2 {
		char1 = '▓'
		char2 = ' '
	} else {
		char1 = ' '
		char2 = '▓'
	}
	pos_X := 2*x + 2*tablero.width + posX + 11 + (4 - len)
  //acomodamos la figura en el cuadrito
	termbox.SetCell(pos_X, y+posY+4, char1, color, color^termbox.AttrBold)
	termbox.SetCell(pos_X+1, y+posY+4, char2, color, color^termbox.AttrBold)
}

func (interfaz *Interfaz) dibujarPausa() {
	y := (tablero.height+1)/2 + posY
	interfaz.dibujarTextoCentroTablero(y, "Pausa", termbox.ColorWhite, termbox.ColorBlack)
}

func (interfaz *Interfaz) DibujaFigura(x int, y int, color termbox.Attribute, rotacion int) {
	var char1 rune
	var char2 rune
	if rotacion < 2 {
		char1 = '▓'
		char2 = ' '
	} else {
		char1 = ' '
		char2 = '▓'
	}
	if color == blankColor {
		// para la sombra de la figura
		termbox.SetCell(2*x+posX+2, y+posY+1, char1, termbox.ColorBlack|termbox.AttrBold, termbox.ColorWhite)
		termbox.SetCell(2*x+posX+3, y+posY+1, char2, termbox.ColorBlack|termbox.AttrBold, termbox.ColorWhite)
	} else {
		termbox.SetCell(2*x+posX+2, y+posY+1, char1, color, color^termbox.AttrBold)
		termbox.SetCell(2*x+posX+3, y+posY+1, char2, color, color^termbox.AttrBold)
	}
}

func (interfaz *Interfaz) dibujarGameOver() {
	y := posY + 2
	interfaz.dibujarTextoCentroTablero(y, " GAME OVER", termbox.ColorWhite, termbox.ColorBlack)
	y += 2
	interfaz.dibujarTextoCentroTablero(y, "SPACE para iniciar", termbox.ColorWhite, termbox.ColorBlack)
}

//colorear linea del color
func (interfaz *Interfaz) colorizeLine(y int, color termbox.Attribute) {
	for x := 0; x < tablero.width; x++ {
		termbox.SetCell(x*2+posX+2, y+posY+1, ' ', termbox.ColorDefault, color)
		termbox.SetCell(x*2+posX+3, y+posY+1, ' ', termbox.ColorDefault, color)
	}
}


func (interfaz *Interfaz) ShowDeleteAnimation(lines []int) {
	interfaz.refrescarPantalla()

	for times := 0; times < 3; times++ {
		for _, y := range lines {
			interfaz.colorizeLine(y, termbox.ColorGreen)
		}
		termbox.Flush()
		time.Sleep(140 * time.Millisecond)

    interfaz.refrescarPantalla()
		time.Sleep(140 * time.Millisecond)
	}
}
//Al perder muestra dos formas de borrar la pantalla
func (interfaz *Interfaz) MostrarAnimacionGameOver() {
	logger.Info("View ShowGameOverAnimation start")

	switch rand.Intn(3) {
	case 0:
		for y := tablero.height - 1; y >= 0; y-- {
			interfaz.colorizeLine(y, termbox.ColorBlack)
			termbox.Flush()
			time.Sleep(60 * time.Millisecond)
		}

	case 1:
		for y := 0; y < tablero.height; y++ {
			interfaz.colorizeLine(y, termbox.ColorBlack)
			termbox.Flush()
			time.Sleep(60 * time.Millisecond)
		}
  }
	logger.Info("View ShowGameOverAnimation end")
}

func (interfaz *Interfaz) dibujaRankingScores() {
	y := posY + 7
	for index, line := range motor.ranking.pts {
		interfaz.dibujarTextoCentroTablero(y+index, fmt.Sprintf("%1d: %6d", index+1, line), termbox.ColorWhite, termbox.ColorBlack)
	}
}
