package utils

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/chandumlg/DBHelper1/utils/misc"
	"github.com/jinzhu/gorm"
)

var batch string

type InsertHandleT struct {
	db *gorm.DB
}

/*
TearDown releases all the resources
*/
func (insertHandler *InsertHandleT) TearDown() {
	insertHandler.db.Close()
}

func GetConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", 5433, "postgres", "postgres", "postgresDB")

}

func (insertHandler *InsertHandleT) SetUp() {
	var err error
	cs := GetConnectionString()
	insertHandler.db, err = gorm.Open("postgres", cs)
	if err != nil {
		fmt.Println("Failed to connect to DB")
		return
	}

	dat, err := ioutil.ReadFile("batch.json")
	check(err)
	batch = string(dat)
}

func (insertHandler *InsertHandleT) Insert() {

	//Creating new entries in split_jobs table
	sqlStatement := "INSERT INTO event_uploads (\"writeKey\", \"event\") VALUES "
	for i := 0; i <= 1000; i++ {
		sqlStatement = sqlStatement + fmt.Sprintf("('%s', '%s'),", strconv.Itoa(i), batch)
	}

	sqlStatement = sqlStatement[:len(sqlStatement)-1]
	errs := insertHandler.db.Exec(sqlStatement).GetErrors()
	if len(errs) != 0 {
		misc.AssertError(errs[0])
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
