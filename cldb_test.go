package clightningdb

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"testing"
)

const dbpath = "./lightningd.sqlite3"

func TestUtxoset(t *testing.T) {
	u := &utxoset{}
	runtest(t, u, "utxoset", false)
}

func TestTx(t *testing.T) {
	tr := &tx{}
	rows, db := runtest(t, tr, "transactions", true)

	for rows.Next() {
		err := scanToStruct(tr, rows, db)
		log.Printf("txid: %s\n", tr.IdStr())
		log.Printf("tx: %s\n", tr.Raw())
		checkErr(err, t)
	}

}

func TestOutputs(t *testing.T) {
	o := &outputs{}
	runtest(t, o, "outputs", false)
}

func TestVars(t *testing.T) {
	log.Println("TODO: *********** VARS TEST NEEDS FIXED ************")
}

func TestShachains(t *testing.T) {
	s := &shachains{}
	runtest(t, s, "shachains", false)
}

func TestShachainsKnown(t *testing.T) {
	s := &shachain_known{}
	runtest(t, s, "shachain_known", false)
}

func TestPeers(t *testing.T) {
	p := &peers{}
	runtest(t, p, "peers", false)
}

func TestChannelConfigs(t *testing.T) {
	p := &channel_configs{}
	runtest(t, p, "channel_configs", false)
}

func TestChannels(t *testing.T) {
	c := &channels{}
	runtest(t, c, "channels", false)
}

func TestInvoices(t *testing.T) {
	i := &invoices{}
	runtest(t, i, "invoices", false)
}

func TestPayments(t *testing.T) {
	p := &payments{}
	runtest(t, p, "payments", false)
}

func checkErr(err error, t *testing.T) {
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func runtest(t *testing.T, entity interface{}, table string, more bool) (*sql.Rows, *sql.DB) {
	db, err := sql.Open("sqlite3", dbpath)
	checkErr(err, t)

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s", table))
	checkErr(err, t)
	cols, _ := rows.Columns()
	log.Printf("query columns: %s\n", cols)

	for rows.Next() {
		s := entity

		err = scanToStruct(s, rows, db)
		log.Printf("row: %v\n", s)
		checkErr(err, t)
	}

	if more {
		rows, err = db.Query(fmt.Sprintf("SELECT * FROM %s", table))
		checkErr(err, t)
	}

	return rows, db
}
