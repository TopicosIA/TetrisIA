package main
import (
  "time"
  "github.com/nsf/termbox-go"
)

func NuevoTablero() {
	tablero = &Tablero{}
	tablero.Limpiar()
}

func (tablero *Tablero) Limpiar() {
	tablero.width = len(boards[tablero.boardsIndex].colors)
	tablero.height = len(boards[tablero.boardsIndex].colors[0])
	tablero.colors = make([][]termbox.Attribute, len(boards[tablero.boardsIndex].colors))
	for i := 0; i < len(boards[tablero.boardsIndex].colors); i++ {
		tablero.colors[i] = make([]termbox.Attribute, len(boards[tablero.boardsIndex].colors[i]))
		copy(tablero.colors[i], boards[tablero.boardsIndex].colors[i])
	}
	tablero.rotation = make([][]int, len(boards[tablero.boardsIndex].rotation))
	for i := 0; i < len(boards[tablero.boardsIndex].rotation); i++ {
		tablero.rotation[i] = make([]int, len(boards[tablero.boardsIndex].rotation[i]))
		copy(tablero.rotation[i], boards[tablero.boardsIndex].rotation[i])
	}
  //la figura que aparece en el cuadrito
	tablero.figuraPrevia = NuevaFigura()
	tablero.figuraActual = NuevaFigura()

}
func (tablero *Tablero) DibujarTablero() {
	for i := 0; i < tablero.width; i++ {
		for j := 0; j < tablero.height; j++ {
			if tablero.colors[i][j] != blankColor {
				interfaz.DibujaFigura(i, j, tablero.colors[i][j], tablero.rotation[i][j])
			}
		}
	}
}
func (tablero *Tablero) SetColor(x int, y int, color termbox.Attribute, rotation int) {
	tablero.colors[x][y] = color
	tablero.rotation[x][y] = rotation
}

//si se encuentra dentro del tablero
func PosicionValida(x int, y int) bool {
	return x >= 0 && x < tablero.width && y >= 0 && y < tablero.height
}

//si la posicion del bloque es valida en el tablero
func PosicionValidaFigura(x int, y int, enTablero bool) bool {
	if x < 0 || x >= tablero.width || y >= tablero.height {
		return false
	}
	if enTablero {
		if y < 0 {
			return false
		}
	} else {
		if y < -2 {
			return false
		}
	}
	if y > -1 {
		if tablero.colors[x][y] != blankColor {
			return false
		}
	}
	return true
}
//manda llamar a la funcion dibuja figura de figura
//pone la imagen en el cuadrito de la pantalla
func (tablero *Tablero) DibujarFiguraPrevia() {
	tablero.figuraPrevia.DibujaFigura(VistaPreviaFigura)
}

func (tablero *Tablero) DibujarFiguraActual(){
  tablero.figuraActual.DibujaFigura(FiguraActual)
}
func (tablero *Tablero) DibujarSombraFigura(){
  figura := tablero.figuraActual.MoverAbajo()
  if !figura.LocacionValida(false){
    return
  }
  //manda la sombra hasta abajo
  for figura.LocacionValida(false) {
		figura.MoverAbajo()
	}
  //sube una posicion la figura
	figura.MoverArriba()
	figura.DibujaFigura(SombraFigura)
}

func (tablero *Tablero) MoverIzquierda() {
	tablero.distAlPiso = 0
	figura := tablero.figuraActual.MoverIzq()
	if figura.LocacionValida(false) {//para que no se salga del tablero
		tablero.figuraActual = figura
		tablero.StartLockDelayIfBottom()
	}
}

func (tablero *Tablero) MoverFiguraAbajo() {
	figura := tablero.figuraActual.MoverAbajo()
	if figura.LocacionValida(false) {
		tablero.distAlPiso = 0
		tablero.figuraActual = figura
		if !tablero.StartLockDelayIfBottom() {
			motor.ResetTimer(0)
		}
		return
	}
	if !tablero.figuraActual.LocacionValida(true) {
		motor.GameOver()
		return
	}
	//tablero.sigFigura()
}

//falta ver que hace bien
func (tablero *Tablero) StartLockDelayIfBottom() bool {
	figura := tablero.figuraActual.MoverAbajo()
	if figura.LocacionValida(false) {
		return false
	}
	motor.ResetTimer(300 * time.Millisecond)
	return true
}

func (tablero *Tablero) MoverDerecha() {
  tablero.distAlPiso = 0
	figura := tablero.figuraActual.MoverDer()
	if figura.LocacionValida(false) {
		tablero.figuraActual = figura
		tablero.StartLockDelayIfBottom()
	}
}

func (tablero *Tablero) sigFigura() {
	//engine.AddScore(board.dropDistance)

	tablero.figuraActual.PonerEnTablero()

	//board.deleteCheck()

	if !tablero.figuraPrevia.LocacionValida(false) {
		tablero.figuraPrevia.MoverArriba()
		if !tablero.figuraPrevia.LocacionValida(false) {
			motor.GameOver()
			return
		}
	}

	tablero.figuraActual = tablero.figuraPrevia
	tablero.figuraPrevia = NuevaFigura()
	//engine.AiGetBestQueue()
	motor.ResetTimer(0)
}
