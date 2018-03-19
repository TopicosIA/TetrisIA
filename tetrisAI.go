package main

import (
	"time"
	"math/rand"
	"os"
	"path/filepath"
	"gopkg.in/inconshreveable/log15.v2"
)

func main() {

	baseDir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	logger = log15.New()
	if baseDir != "" {
		logger.SetHandler(log15.Must.FileHandler(baseDir+"/tetrisAI.log", log15.LogfmtFormat()))
	}

	rand.Seed(time.Now().UnixNano())
	CrearFiguras()
	NuevoTablero()
	NuevaInterfaz()
	NuevoMotorJuego()

	motor.Run()
	interfaz.Stop()

}
