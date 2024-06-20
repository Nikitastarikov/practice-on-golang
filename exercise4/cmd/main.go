package main

import (
	libconf "github.com/Nikitastarikov/practice-on-golang/exercise4/lib"
	pkglog "github.com/Nikitastarikov/practice-on-golang/pkg/log"
)

func main() {
	conf := libconf.Get()
	l := pkglog.Get()

	//Синхронизация очищает все буферизованные записи журнала.
	defer l.Sync()

	err := conf.Print()

	if err != nil {
		l.Fatal(err)
	}
}
