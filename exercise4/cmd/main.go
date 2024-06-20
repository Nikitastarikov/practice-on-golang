package main

import (
	"fmt"
	"regexp"
	"strings"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	libconf "github.com/Nikitastarikov/practice-on-golang/exercise4/lib/config"
	libentity "github.com/Nikitastarikov/practice-on-golang/exercise4/lib/entity"
	libpsql "github.com/Nikitastarikov/practice-on-golang/exercise4/lib/psql"
	librepository "github.com/Nikitastarikov/practice-on-golang/exercise4/lib/repository"
	pkgerr "github.com/Nikitastarikov/practice-on-golang/pkg/error"
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

	db, dbClear, err := libpsql.SqlDB(conf.Psql)

	if err != nil {
		l.Fatal(err)
	}

	defer dbClear()

	err = db.Ping()

	if err != nil {
		l.Fatal(err)
	}

	err = libpsql.UpMigration(conf.Psql.Dbname, db)

	if err != nil {
		l.Fatal(err)
	}

	repo, err := librepository.NewRepository(db)

	if err != nil {
		l.Fatal(err)
	}

	normalizationPhoneNumbers(l, repo)
}

func normalizationPhoneNumbers(log pkglog.Logger, repo *librepository.Repository) {
	l := log.Named("normalizationPhoneNumbers")

	phoneList, err := repo.GetPhoneList(nil, nil)

	if err != nil {
		l.Errorf("get phone list: %v", err)
		return
	}

	outputPhoneList(phoneList)

	for i := range phoneList {
		normNumber, err := normalizationNumber(phoneList[i].Number)

		if err != nil {
			l.Errorf("normalization number: %v", err)
			return
		}

		fmt.Println(*normNumber)

		if *normNumber != phoneList[i].Number {
			_, err = repo.UpdatePhoneById(phoneList[i].Id, *normNumber)

			if err != nil {
				errCode := pkgerr.ErrorCode(err)

				if errCode != pkgerr.UniqueViolation {
					l.Errorf("update phone by id:%v: %v", phoneList[i].Id)
					return
				}

				err = repo.DeletePhoneById(phoneList[i].Id)

				if err != nil {
					l.Errorf("delete phone by id:%v: %v", phoneList[i].Id, err)
					return
				}
			}
		}
	}

	phoneList, err = repo.GetPhoneList(nil, nil)

	if err != nil {
		l.Errorf("get phone list: %v", err)
		return
	}

	outputPhoneList(phoneList)
}

func outputPhoneList(phoneList []*libentity.Phone) {
	fmt.Println()
	fmt.Println("phoneList:")
	for i := range phoneList {
		fmt.Println(phoneList[i].Id, phoneList[i].Number)
	}
	fmt.Println()
}

func normalizationNumber(number string) (*string, error) {
	reg, err := regexp.Compile(`\d*`)

	if err != nil {
		return nil, err
	}

	s := reg.FindAllString(number, -1)

	normNumber := strings.Join(s, "")

	return &normNumber, nil
}
