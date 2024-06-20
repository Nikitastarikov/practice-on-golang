package main

import (
	libconf "example.com/m/v2/exercise4/lib"
	pkglog "example.com/m/v2/pkg/log"
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
