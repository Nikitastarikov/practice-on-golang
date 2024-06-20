package libpsql

import (
	"database/sql"
	"fmt"

	"github.com/Nikitastarikov/practice-on-golang/exercise4/lib/config"
)

func SqlDB(conf *config.Psql) (*sql.DB, func(), error) {
	// sslmode - использовать ли ssl сертификат
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.Dbname,
	)

	fmt.Printf("psqlInfo = %v\n", psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		return nil, nil, err
	}

	dbClear := func() {
		db.Close()
	}

	return db, dbClear, nil
}
