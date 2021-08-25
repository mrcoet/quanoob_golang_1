package work

import (
	"database/sql"
	"fmt"
	"math/big"

	_ "github.com/mattn/go-sqlite3"

	"github.com/vsergeev/btckeygenie/btckey"
)

func CheckHex(xBytes string) {
	n := new(big.Int)
	n.SetString(xBytes, 16)
	pKey := btckey.NewPrivateKey(n)

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
		fmt.Println(xBytes)
	}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func GetAddr() {
	xBytes := "f912f9da5b2810f4107b33c38c5b7ccba26a770e6e5e45de1da446608df38992"
	n := new(big.Int)
	n.SetString(xBytes, 16)
	pKey := btckey.NewPrivateKey(n)
	fmt.Println(pKey.ToAddress())
	fmt.Println(pKey.ToAddressUncompressed())
}
