package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func NewRanking() *Ranking {
	ranking := &Ranking{
		pts: make([]uint64, 9),
	}

	if _, err := os.Stat(baseDir + rankingFileName); os.IsNotExist(err) {
		for i := 0; i < 9; i++ {
			ranking.pts[i] = 0
		}
		return ranking
	}

	scoreBytes, err := ioutil.ReadFile(baseDir + rankingFileName)
	if err != nil {
		logger.Error("NewRanking ReadFile", "error", err.Error())
	}

	scoreStrings := strings.Split(string(scoreBytes), ",")
	for index, scoreString := range scoreStrings {
		if index > 8 {
			break
		}
		score, err := strconv.ParseUint(scoreString, 10, 64)
		if err != nil {
			logger.Error("NewRanking ParseUint", "error", err.Error())
			score = 0
		}
		ranking.pts[index] = score
	}

	return ranking
}

func (ranking *Ranking) Save() {
	var buffer bytes.Buffer

	for i := 0; i < 9; i++ {
		if i != 0 {
			buffer.WriteRune(',')
		}
		buffer.WriteString(strconv.FormatUint(ranking.pts[i], 10))
	}

	ioutil.WriteFile(baseDir+rankingFileName, buffer.Bytes(), 0644)
}

func (ranking *Ranking) InsertScore(newScore uint64) {
	for index, score := range ranking.pts {
		//ordenamos el score de mayor a menor
		if newScore > score {
			ranking.slideScores(index)
			ranking.pts[index] = newScore
			return
		}
	}
}

func (ranking *Ranking) slideScores(index int) {
	for i := len(ranking.pts) - 1; i > index; i-- {
		ranking.pts[i] = ranking.pts[i-1]
	}
}
