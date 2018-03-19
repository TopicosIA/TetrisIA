package main

import (
	"math/rand"

	"github.com/nsf/termbox-go"
)

func NuevaFigura() *Figura {
  rotacionFiguras := figuras.bolsaFiguras[figuras.bolsaRandom[figuras.indiceBolsa]]
  figuras.indiceBolsa++
  if figuras.indiceBolsa > 6{
    figuras.indiceBolsa=0
    figuras.bolsaRandom = rand.Perm(7)
  }
  figura := &Figura {
    rotacionFiguras: rotacionFiguras,
    len: len(rotacionFiguras[0]),
  }
  figura.x = tablero.width/2 - (figura.len+1)/2 //centramos la PIEZA
  figura.y = -1
  return figura
}

func (figura *Figura) MoverIzq() *Figura{
  nuevaFigura := *figura
	nuevaFigura.MI()
  return &nuevaFigura
}
func (figura *Figura) MI() {
	figura.x--
}

func (figura *Figura) MoverDer() *Figura{
  nuevaFigura := *figura
	nuevaFigura.MD()
  return &nuevaFigura
}
func (figura *Figura) MD() {
	figura.x++
}

func (figura *Figura) CopiaRotadaDerecha() *Figura{
  nuevaFigura := *figura
	nuevaFigura.RotarDerecha()
  return &nuevaFigura
}
func (figura *Figura) RotarDerecha() {
	figura.rotacion++
	//regresamos a la figura inicial
  if figura.rotacion > 3{
    figura.rotacion=0
  }
}

func (figura *Figura) CopiaRotadaIzquierda() *Figura{
  nuevaFigura := *figura
	nuevaFigura.RotarIzquierda()
  return &nuevaFigura
}

func (figura *Figura) RotarIzquierda() {
	if figura.rotacion < 1{
    figura.rotacion = 3
		return
  }
  figura.rotacion--
}

func (figura *Figura) MoverAbajo() *Figura {
	newFigura := *figura
	newFigura.MAbajo()
	return &newFigura
}

func (figura *Figura) MAbajo() {
	figura.y++
}
func (figura *Figura) MoverArriba(){
	figura.y--
}


func (figura *Figura) LocacionValida(enTablero bool) bool {
	//obtenemos todas las figuras las 4 rotaciones
	figuras := figura.rotacionFiguras[figura.rotacion]
	for i := 0; i < figura.len; i++ {
		for j := 0; j < figura.len; j++ {
			if figuras[i][j] != blankColor {
				//esta funcion esta en tablero
				if !PosicionValidaFigura(figura.x+i, figura.y+j, enTablero) {
					return false
				}
			}
		}
	}
	return true
}
//pone la pieza en el tablero
func (figura *Figura) PonerEnTablero() {
	figuras := figura.rotacionFiguras[figura.rotacion]
	for i := 0; i < figura.len; i++ {
		for j := 0; j < figura.len; j++ {
			if figuras[i][j] != blankColor {
				//esta en tablero la funcion
				tablero.SetColor(figura.x+i, figura.y+j, figuras[i][j], figura.rotacion)
			}
		}
	}
}

func (figura *Figura) DibujaFigura(tipoFigura TipoFigura) {
	figuras := figura.rotacionFiguras[figura.rotacion]
	for i := 0; i < figura.len; i++ {
		for j := 0; j < figura.len; j++ {
			if figuras[i][j] != blankColor {
				switch tipoFigura {
				case VistaPreviaFigura:
					interfaz.DibujaVistaPreviaFigura(i, j, figuras[i][j], figura.rotacion, figura.len)
				case SombraFigura:
					interfaz.DibujaFigura(figura.x+i, figura.y+j, blankColor, figura.rotacion)
				case FiguraActual:
					if PosicionValida(figura.x+i, figura.y+j) {
						interfaz.DibujaFigura(figura.x+i, figura.y+j, figuras[i][j], figura.rotacion)
					}
				}
			}
		}
	}
}

func (figura *Figura) isMinoAtLocation(x int, y int) bool {
	xIndex := x - figura.x
	yIndex := y - figura.y
	if xIndex < 0 || xIndex >= figura.len || yIndex < 0 || yIndex >= figura.len {
		return false
	}

	figuras := figura.rotacionFiguras[figura.rotacion]
	if figuras[xIndex][yIndex] != blankColor {
		return true
	}

	return false
}

func (figura *Figura) minoOverlap(figura1 *Figura) bool {
	figuras := figura.rotacionFiguras[figura.rotacion]
	for i := 0; i < figura.len; i++ {
		for j := 0; j < figura.len; j++ {
			if figuras[i][j] != blankColor {
				if figura1.isMinoAtLocation(figura.x+i, figura.y+j) {
					return true
				}
			}
		}
	}
	return false
}

func (figura *Figura) getMinoColorAtLocation(x int, y int) termbox.Attribute {
	xIndex := x - figura.x
	yIndex := y - figura.y
	if xIndex < 0 || xIndex >= figura.len || yIndex < 0 || yIndex >= figura.len {
		return blankColor
	}

	figuras := figura.rotacionFiguras[figura.rotacion]
	return figuras[xIndex][yIndex]
}
