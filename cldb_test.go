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
	runcltest(t, u, "utxoset", []string{})
}

func TestTx(t *testing.T) {
	tr := &tx{}
	rows := runtest(t, tr, "transactions", true)

	for rows.Next() { // TODO: remove need for separate test
		err := scanToStruct(tr, rows)
		log.Printf("txid: %s\n", tr.IdStr())
		log.Printf("tx: %s\n", tr.Raw())
		checkErr(err, t)
	}

}

func TestOutputs(t *testing.T) {
	o := &outputs{}
	runcltest(t, o, "outputs", []string{})
}

func TestVars(t *testing.T) {
	log.Println("TODO: *********** VARS TEST NEEDS FIXED ************")
}

func TestShachains(t *testing.T) {
	s := &shachains{}
	runcltest(t, s, "shachains", []string{})
}

func TestShachainsKnown(t *testing.T) {
	s := &shachain_known{}
	runcltest(t, s, "shachain_known", []string{})
}

func TestPeers(t *testing.T) {
	p := &peers{}
	runcltest(t, p, "peers", []string{})
}

func TestChannelConfigs(t *testing.T) {
	p := &channel_configs{}
	runcltest(t, p, "channel_configs", []string{})
}

func TestChannels(t *testing.T) {
	c := &channels{}
	runcltest(t, c, "channels", []string{})
}

func TestInvoices(t *testing.T) {
	i := &invoices{}
	runcltest(t, i, "invoices", []string{})
}

func TestPayments(t *testing.T) {
	p := &payments{}
	runcltest(t, p, "payments", []string{})
}

func TestBlocks(t *testing.T) {
	b := &blocks{}
	runcltest(t, b, "blocks", []string{})
}

func TestChannelTxs(t *testing.T) {
	b := &channeltxs{}
	runcltest(t, b, "channeltxs", []string{})
}

func TestForwards(t *testing.T) {
	b := &forwarded_payments{}
	runcltest(t, b, "forwarded_payments", []string{})
}

func TestPartialFields(t *testing.T) {
	p := &payments{}
	runcltest(t, p, "payments", []string{"id", "timestamp", "status",
		"payment_hash", "destination", "msatoshi"})
}

func TestStructAccess(t *testing.T) {
	p := &payments{}

	sdb, err := sql.Open("sqlite3", dbpath)
	checkErr(err, t)
	db := &cldb{sdb}
	rows := db.queryFields("payments", []string{"id", "timestamp", "status",
		"payment_hash", "destination", "msatoshi"}, p)
	for _, r := range rows {
		pay := r.(payments)
		log.Printf("%x %d\n", pay.Payment_hash, pay.Msatoshi)
	}

}

func TestListPeers(t *testing.T) {
	sdb, err := sql.Open("sqlite3", dbpath)
	if err != nil {

	}
	db := &cldb{sdb}
	db.listPeers()

}

func checkErr(err error, t *testing.T) {
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func runcltest(t *testing.T, entity cl, table string, fields []string) {
	sdb, err := sql.Open("sqlite3", dbpath)
	checkErr(err, t)
	db := &cldb{sdb}
	rows := db.queryFields(table, fields, entity)
	for _, r := range rows {
		log.Printf("%v\n", r)
	}
}

func runtest(t *testing.T, entity interface{}, table string, more bool) *sql.Rows {
	db, err := sql.Open("sqlite3", dbpath)
	checkErr(err, t)

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s", table))
	checkErr(err, t)
	cols, _ := rows.Columns()
	log.Printf("query columns: %s\n", cols)

	for rows.Next() {
		s := entity

		err = scanToStruct(s, rows)
		log.Printf("row: %v\n", s)
		checkErr(err, t)
	}

	if more {
		rows, err = db.Query(fmt.Sprintf("SELECT * FROM %s", table))
		checkErr(err, t)
	}

	return rows
}
