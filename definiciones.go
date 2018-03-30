package main

import (
	"time"
	"github.com/nsf/termbox-go"
	"gopkg.in/inconshreveable/log15.v2"
)

const (
	blankColor      = termbox.ColorBlack
	posX    = 20
	posY    = 2
	rankingFileName = "/tetris2.db"

	VistaPreviaFigura TipoFigura = iota
	FiguraActual         				 = iota
	SombraFigura             	   = iota
)
type (
  TipoFigura       int
  BloquesFigura    [][]termbox.Attribute //para crear la figura
	RotacionFiguras  [4]BloquesFigura // 4 posibilidades

  Figuras struct {
		bolsaFiguras  [7]RotacionFiguras
		bolsaRandom   []int
		indiceBolsa   int
	}

  Figura struct {
		x            int
		y            int
		len          int
		rotacion     int
		rotacionFiguras RotacionFiguras
	}

  Tablero struct {
		boardsIndex  int
		width        int
		height       int
		colors       [][]termbox.Attribute
		rotation     [][]int
	  figuraPrevia *Figura
		figuraActual *Figura
		distAlPiso   int
  }

  Interfaz struct {
  }

	Sonido struct {
		tickTimeCancion time.Duration
		timerCancion 	  *time.Timer
		reproduciendo 	bool
  }

  Boards struct {
		colors   [][]termbox.Attribute
		rotation [][]int
  }
  TeclaPresionada struct {
		stopped      		bool
		tecla_salir     chan struct{}
		opciones_teclas chan *termbox.Event
	}
Ranking struct{
	pts []uint64 //puntos realizados en lapartida
	//podemos meter una variable string para el nombre
}
  MotorJuego struct {
		stopped      		 bool
		chanStop     		 chan struct{}
		tecla     			 *TeclaPresionada
		ranking          *Ranking
		timer          	 *time.Timer
		tickTime      	 time.Duration
		pausado       	 bool
		gameOver         bool
		tableroInicio    bool
		score            int
		nivel            int
		lineasBorradas   int
		aiEnabled    		 bool
}
)
var (
	boards 		[]Boards
	baseDir 	string
	logger  	log15.Logger
	figuras   *Figuras
	tablero   *Tablero
	interfaz  *Interfaz
	sonido    *Sonido
	motor  		*MotorJuego
)
