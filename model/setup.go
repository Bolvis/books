package model

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func SetupModels() *gorm.DB {
	psqlInfo := fmt.Sprintf(psqlInfoFormat, host, port, user, password, dbName)

	db, err := gorm.Open(dbDriver, psqlInfo)

	if err != nil {
		panic(strings.Join([]string{"Can't connect to database", err.Error()}, " -> "))
	}

	return db
}
