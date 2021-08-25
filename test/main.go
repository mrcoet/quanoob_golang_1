package main

import (
	"database/sql"
	"fmt"
	"math/big"
	"math/rand"
	"runtime"
	"strings"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mrcoet/quanbit_golang/work"
	"github.com/vsergeev/btckeygenie/btckey"
)

func main() {
	runtime.GOMAXPROCS(-1)
	stime := time.Now()
	work.DropRock(4, "")
	ftime := time.Now()
	fmt.Println("Finished in: ", ftime.Sub(stime))

	// sstime := time.Now()
	// fmt.Println(work.DropRockTestOne())
	// fftime := time.Now()
	// fmt.Println("Finished in: ", fftime.Sub(sstime))
	var waitGroup sync.WaitGroup

	ssstime := time.Now()
	i := 0
	waitGroup.Add(100)
	for i < 100 {
		go GetFullSync(&waitGroup)
		i++
	}
	waitGroup.Wait()
	ffftime := time.Now()
	fmt.Println("Finished in: ", ffftime.Sub(ssstime))

	sstime := time.Now()
	ii := 0
	for ii < 100 {
		GetFull()
		ii++
	}
	fftime := time.Now()
	fmt.Println("Finished in: ", fftime.Sub(sstime))

}

func GetFullSync(wg *sync.WaitGroup) {
	hexaValues := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(hexaValues), func(i, j int) { hexaValues[i], hexaValues[j] = hexaValues[j], hexaValues[i] })

	n := new(big.Int)
	n.SetString(strings.Join(hexaValues, ""), 16)
	pKey := btckey.NewPrivateKey(n)
	// fmt.Println(pKey.ToAddress())
	// fmt.Println(pKey.ToAddressUncompressed())

	db, err := sql.Open("sqlite3", "./addrs.db")
	checkErr(err)

	q := "SELECT * FROM blockchair WHERE address == \"" + pKey.ToAddress() + "\"" + "OR address == \"" + pKey.ToAddressUncompressed() + "\""
	rows, err := db.Query(q)
	checkErr(err)

	var addr string
	for rows.Next() {
		err = rows.Scan(&addr)
		checkErr(err)
		fmt.Println(addr)
		fmt.Println(n)
	}
	wg.Done()
}

func GetFull() {
	hexaValues := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(hexaValues), func(i, j int) { hexaValues[i], hexaValues[j] = hexaValues[j], hexaValues[i] })

	n := new(big.Int)
	n.SetString(strings.Join(hexaValues, ""), 16)
	pKey := btckey.NewPrivateKey(n)
	// fmt.Println(pKey.ToAddress())
	// fmt.Println(pKey.ToAddressUncompressed())

	db, err := sql.Open("sqlite3", "./addrs.db")
	checkErr(err)

	q := "SELECT * FROM blockchair WHERE address == \"" + pKey.ToAddress() + "\"" + "OR address == \"" + pKey.ToAddressUncompressed() + "\""
	rows, err := db.Query(q)
	checkErr(err)

	var addr string
	for rows.Next() {
		err = rows.Scan(&addr)
		checkErr(err)
		fmt.Println(addr)
		fmt.Println(n)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
