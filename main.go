//test3
//test2
//test
package main

import (
	"fmt"
	"time"

	"github.com/chandumlg/DBHelper1/utils"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var insertHandler utils.InsertHandleT

func main() {

	insertHandler.SetUp()

	t1 := time.Now()

	loopCount := 100

	for i := 0; i < loopCount; i++ {
		insertHandler.Insert()
		time.Sleep(time.Duration(100))
	}

	t2 := time.Now()
	diff := t2.Sub(t1)

	fmt.Println(diff)
}
