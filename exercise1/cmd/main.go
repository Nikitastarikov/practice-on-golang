package main

import (
	libconf "example.com/m/v2/exercise1/lib/config"
	libquiz "example.com/m/v2/exercise1/lib/quiz"
	pkglog "example.com/m/v2/pkg/log"
)

func main() {
	conf := libconf.Get()
	l := pkglog.Get()

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
