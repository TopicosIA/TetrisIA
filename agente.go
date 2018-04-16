package main

func NuevoAgente() *Agente {
	agente := Agente{}
	listaAcciones := make([]rune, 1)
	listaAcciones[0] = 'x'
	agente.listaAcciones = &listaAcciones
	return &agente
}

func (ag *Agente) tomarAccion() {
	if ag.nuevaLista != nil {
		ag.listaAcciones = ag.nuevaLista
		ag.nuevaLista = nil
		ag.index = 0
	}
	listaAcciones := *ag.listaAcciones
	// wasd + qe keyboard keys
	switch listaAcciones[ag.index] {
	case 'w':
		tablero.DescensoFigura()
	case 'a':
		tablero.MoverIzquierda()
	case 'd':
		tablero.MoverDerecha()
	case 'q':
		tablero.RotarFiguraIzquierda()
	case 'e':
		tablero.RotarFiguraDerecha()
	case 'x':
		return
	}
	ag.index++
	interfaz.refrescarPantalla()
}

func (ag *Agente) maxAccion() {
	maxR := -9999999
	maxAccion := make([]rune, 0, 0)
	currentMino := *tablero.figuraActual

	for slide1 := 0; slide1 < 5; slide1++ {
		for move1 := tablero.width; move1 >= 0; move1-- {
			for rotate1 := 0; rotate1 < 5; rotate1++ {

				listaAcciones, mino1 := tablero.getAcciones(rotate1, move1, slide1, &currentMino, nil)
				if mino1 == nil {
					continue
				}

				for slide2 := 0; slide2 < 5; slide2++ {
					for move2 := tablero.width; move2 >= 0; move2-- {
						for rotate2 := 0; rotate2 < 5; rotate2++ {

							_, mino2 := tablero.getAcciones(rotate2, move2, slide2, tablero.figuraPrevia, mino1)
							if mino2 == nil {
								continue
							}

							fullLines, holes, bumpy := tablero.getEstado(mino1, mino2)
							score := ag.r(fullLines, holes, bumpy)

							if score > maxR {
								maxR = score
								maxAccion = listaAcciones
							}

						}
					}
				}

			}
		}
	}
//si no hay una buena accion
	if len(maxAccion) < 1 {
		maxAccion = append(maxAccion, 'x')
	}

	ag.nuevaLista = &maxAccion
}

func (tablero *Tablero) getAcciones(rotate int, move int, slide int, mino1 *Figura, mino2 *Figura) ([]rune, *Figura) {
	listaAcciones := make([]rune, 0, 4)
	mino := *mino1

	if rotate%2 == 0 {
		rotate /= 2
		for i := 0; i < rotate; i++ {
			mino.RotarDerecha()
			listaAcciones = append(listaAcciones, 'e')
			if !mino.LocacionValida(false) || (mino2 != nil && mino2.minoOverlap(&mino)) {
				return listaAcciones, nil
			}
		}
	} else {
		rotate = rotate/2 + 1
		for i := 0; i < rotate; i++ {
			mino.RotarIzquierda()
			listaAcciones = append(listaAcciones, 'q')
			if !mino.LocacionValida(false) || (mino2 != nil && mino2.minoOverlap(&mino)) {
				return listaAcciones, nil
			}
		}
	}

	if move%2 == 0 {
		move /= 2
		for i := 0; i < move; i++ {
			mino.MI()
			listaAcciones = append(listaAcciones, 'a')
			if !mino.LocacionValida(false) || (mino2 != nil && mino2.minoOverlap(&mino)) {
				return listaAcciones, nil
			}
		}
	} else {
		move = move/2 + 1
		for i := 0; i < move; i++ {
			mino.MD()
			listaAcciones = append(listaAcciones, 'd')
			if !mino.LocacionValida(false) || (mino2 != nil && mino2.minoOverlap(&mino)) {
				return listaAcciones, nil
			}
		}
	}
	for mino.LocacionValida(false) && (mino2 == nil || !mino2.minoOverlap(&mino)) {
		mino.MAbajo()
	}
	mino.MoverArriba()
	listaAcciones = append(listaAcciones, 'w')

	if slide%2 == 0 {
		slide /= 2
		for i := 0; i < slide; i++ {
			mino.MI()
			listaAcciones = append(listaAcciones, 'a')
			if !mino.LocacionValida(false) || (mino2 != nil && mino2.minoOverlap(&mino)) {
				return listaAcciones, nil
			}
		}
	} else {
		slide = slide/2 + 1
		for i := 0; i < slide; i++ {
			mino.MD()
			listaAcciones = append(listaAcciones, 'd')
			if !mino.LocacionValida(false) || (mino2 != nil && mino2.minoOverlap(&mino)) {
				return listaAcciones, nil
			}
		}
	}

	if !mino.LocacionValida(true) {
		return listaAcciones, nil
	}
	listaAcciones = append(listaAcciones, 'x')
	return listaAcciones, &mino
}

func (tablero *Tablero) getEstado(mino1 *Figura, mino2 *Figura) (fullLines int, holes int, bumpy int) {
	// fullLines
	fullLinesY := make(map[int]bool, 2)
	for j := 0; j < tablero.height; j++ {
		isFullLine := true
		for i := 0; i < tablero.width; i++ {
			if tablero.colors[i][j] == blankColor && !mino1.isMinoAtLocation(i, j) && !mino2.isMinoAtLocation(i, j) {
				isFullLine = false
				break
			}
		}
		if isFullLine {
			fullLinesY[j] = true
			fullLines++
		}
	}

	// holes and bumpy
	indexLast := 0
	for i := 0; i < tablero.width; i++ {
		index := tablero.height
		indexOffset := 0
		for j := 0; j < tablero.height; j++ {
			if _, found := fullLinesY[j]; found {
				indexOffset++
			} else {
				if tablero.colors[i][j] != blankColor || mino1.isMinoAtLocation(i, j) || mino2.isMinoAtLocation(i, j) {
					index = j
					break
				}
			}
		}

		if i != 0 {
			diffrence := (index + fullLines - indexOffset) - indexLast
			if diffrence < 0 {
				diffrence = -diffrence
			}
			bumpy += diffrence

		}
		indexLast = index + fullLines - indexOffset

		index++
		for j := index; j < tablero.height; j++ {
			if tablero.colors[i][j] == blankColor && !mino1.isMinoAtLocation(i, j) && !mino2.isMinoAtLocation(i, j) {
				holes++
			}
		}
	}
	return
}

func (ag *Agente) r(fullLines int, holes int, bumpy int) (recompensa int) {
	if fullLines == 4 {
		recompensa += 1200
	}
	recompensa -= 75 * holes
	//recompensa -= 25 * bumpy
	return recompensa
}
