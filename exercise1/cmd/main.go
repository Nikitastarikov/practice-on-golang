package main

import (
	libconf "github.com/Nikitastarikov/practice-on-golang/exercise1/lib/config"
	libquiz "github.com/Nikitastarikov/practice-on-golang/exercise1/lib/quiz"
	pkglog "github.com/Nikitastarikov/practice-on-golang/pkg/log"
)

func main() {
	conf := libconf.Get()
	l := pkglog.Get()

	//Синхронизация очищает все буферизованные записи журнала.
	defer l.Sync()

	err := conf.Print()

	if err != nil {
		l.Errorf("config print: %v", err)
		return
	}

	game, err := libquiz.NewQuizGame(l, conf.FileCSVPaths[0], conf.TimeToThink)

	if err != nil {
		l.Errorf("new quiz game: %v", err)
		return
	}

	err = game.Run()

	if err != nil {
		l.Errorf("game run: %v", err)
		return
	}
}
