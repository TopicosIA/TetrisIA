package main

import (
  "math/rand"

  "github.com/nsf/termbox-go"
)

func CrearFiguras(){
  //creamos cada figura
  figuraJ := BloquesFigura{
    []termbox.Attribute{termbox.ColorBlue, termbox.ColorBlue, blankColor},
		[]termbox.Attribute{blankColor, termbox.ColorBlue, blankColor},
		[]termbox.Attribute{blankColor, termbox.ColorBlue, blankColor},
  }
  figuraO := BloquesFigura{
    []termbox.Attribute{termbox.ColorYellow, termbox.ColorYellow},
		[]termbox.Attribute{termbox.ColorYellow, termbox.ColorYellow},
  }
  figuraL := BloquesFigura{
    []termbox.Attribute{blankColor, termbox.ColorWhite, blankColor},
		[]termbox.Attribute{blankColor, termbox.ColorWhite, blankColor},
		[]termbox.Attribute{termbox.ColorWhite, termbox.ColorWhite, blankColor},
  }
  figuraI := BloquesFigura{
    []termbox.Attribute{blankColor, termbox.ColorCyan, blankColor, blankColor},
		[]termbox.Attribute{blankColor, termbox.ColorCyan, blankColor, blankColor},
		[]termbox.Attribute{blankColor, termbox.ColorCyan, blankColor, blankColor},
		[]termbox.Attribute{blankColor, termbox.ColorCyan, blankColor, blankColor},
  }
  figuraS := BloquesFigura{
    []termbox.Attribute{blankColor, termbox.ColorGreen, blankColor},
		[]termbox.Attribute{termbox.ColorGreen, termbox.ColorGreen, blankColor},
		[]termbox.Attribute{termbox.ColorGreen, blankColor, blankColor},
  }
  figuraT := BloquesFigura{
    []termbox.Attribute{blankColor, termbox.ColorMagenta, blankColor},
		[]termbox.Attribute{termbox.ColorMagenta, termbox.ColorMagenta, blankColor},
		[]termbox.Attribute{blankColor, termbox.ColorMagenta, blankColor},
  }
  figuraZ := BloquesFigura{
    []termbox.Attribute{termbox.ColorRed, blankColor, blankColor},
		[]termbox.Attribute{termbox.ColorRed, termbox.ColorRed, blankColor},
		[]termbox.Attribute{blankColor, termbox.ColorRed, blankColor},
  }

  var rotacionFJ RotacionFiguras
  rotacionFJ[0] = figuraJ
  for i := 1; i < 4; i++ {
		rotacionFJ[i] = figuraRotada(rotacionFJ[i-1])
	}
  var rotacionFL RotacionFiguras
  rotacionFL[0] = figuraL
  for i := 1; i < 4; i++ {
		rotacionFL[i] = figuraRotada(rotacionFL[i-1])
	}
  var rotacionFI RotacionFiguras
  rotacionFI[0] = figuraI
  for i := 1; i < 4; i++ {
		rotacionFI[i] = figuraRotada(rotacionFI[i-1])
	}
  var rotacionFO RotacionFiguras
  rotacionFO[0] = figuraO
  rotacionFO[1] = figuraO
  rotacionFO[2] = figuraO
  rotacionFO[3] = figuraO

  var rotacionFT RotacionFiguras
  rotacionFT[0] = figuraT
  for i := 1; i < 4; i++ {
		rotacionFT[i] = figuraRotada(rotacionFT[i-1])
	}
  var rotacionFS RotacionFiguras
  rotacionFS[0] = figuraS
  for i := 1; i < 4; i++ {
		rotacionFS[i] = figuraRotada(rotacionFS[i-1])
	}

  var rotacionFZ RotacionFiguras
  rotacionFZ[0] = figuraZ
  for i := 1; i < 4; i++ {
		rotacionFZ[i] = figuraRotada(rotacionFZ[i-1])
	}

  figuras = &Figuras{
    bolsaFiguras: [7]RotacionFiguras{rotacionFI,rotacionFZ,rotacionFJ,rotacionFT,rotacionFO,rotacionFL,rotacionFS},
    bolsaRandom: rand.Perm(7),
  }
}
func figuraRotada(figura BloquesFigura) BloquesFigura {
	length := len(figura)
	nuevaFiguraBloques := make(BloquesFigura, length, length)
  //rellenamos la matriz de arreglos tipo atributo
	for i := 0; i < length; i++ {
		nuevaFiguraBloques[i] = make([]termbox.Attribute, length, length)
	}
  //rotamos los indices
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			nuevaFiguraBloques[length-j-1][i] = figura[i][j]
		}
	}

	return nuevaFiguraBloques
}
