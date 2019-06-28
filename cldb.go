package clightningdb

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"reflect"
)

type tx struct {
	Id          []byte
	Blockheight int
	Txindex     int
	Rawtx       []byte
}

func (t *tx) String() string {
	return fmt.Sprintf("{%x %d %d%x}", t.Id, t.Blockheight, t.Txindex, t.Rawtx)
}

func (t *tx) IdStr() string {
	return fmt.Sprintf("%x", t.Id)
}

func (t *tx) Raw() string {
	return fmt.Sprintf("%x", t.Rawtx)
}

type utxoset struct {
	Txid         []byte
	Outnum       int
	Blockheight  int
	Spendheight  int
	Txindex      int
	Scriptpubkey []byte
	Satoshis     int64
}

func (u *utxoset) String() string {
	return fmt.Sprintf("{%x %d %d %d %d %x %d}", u.Txid, u.Outnum, u.Blockheight, u.Spendheight, u.Txindex, u.Scriptpubkey, u.Satoshis)
}

type outputs struct {
	Prev_out_tx         []byte
	Prev_out_index      int
	Value               int
	Type                int
	Status              int
	Keyindex            int
	Channel_id          int
	Peer_id             []byte
	Commitment_point    []byte
	Confirmation_height int
	Spend_height        int
	Scriptpubkey        []byte
}

func (o *outputs) String() string {
	return fmt.Sprintf("{%x %d %d %d %d %d %d %x %x %d %d %x}", o.Prev_out_tx,
		o.Prev_out_index, o.Value, o.Type, o.Status, o.Keyindex, o.Channel_id, o.Peer_id,
		o.Commitment_point, o.Confirmation_height, o.Spend_height, o.Scriptpubkey)
}

type vars struct {
	Name string
	Val  string // this could be string or bytes???
}

func (v *vars) String() string {
	return fmt.Sprintf("{%s %s}", v.Name, v.Val)
}

type shachains struct {
	Id        int
	Min_index int
	Num_valid int
}

type shachain_known struct {
	Shachain_id int
	Pos         int
	Idx         int
	Hash        []byte
}

func (k *shachain_known) String() string {
	return fmt.Sprintf("{%d %d %d %x}", k.Shachain_id, k.Pos, k.Idx, k.Hash)
}

type peers struct {
	Id      int
	Node_id []byte
	Address string
}

func (p *peers) String() string {
	return fmt.Sprintf("{%d %x %s}", p.Id, p.Node_id, p.Address)
}

type channel_configs struct {
	Id                            int
	Dust_limit_satoshis           int
	Max_htlc_value_in_flight_msat int
	Channel_reserve_satoshis      int
	Htlc_minimum_msat             int
	To_self_delay                 int
	Max_accepted_htlcs            int
}

func scanToStruct(obj interface{}, rows *sql.Rows, db *sql.DB) error {
	s := reflect.ValueOf(obj).Elem()
	fields := make([]interface{}, 0)
	for i := 0; i < s.NumField(); i++ {
		var f interface{}
		fields = append(fields, &f)
	}

	err := rows.Scan(fields...)

	for i := 0; i < s.NumField(); i++ {
		var raw_value = *fields[i].(*interface{})
		setFieldValue(s.Field(i), raw_value)
	}

	return err

}

func setFieldValue(field reflect.Value, val interface{}) {
	if val == nil {
		return
	}
	switch field.Kind() {
	case reflect.String:
		field.SetString(val.(string))
	case reflect.Int, reflect.Int64:
		field.SetInt(val.(int64))
	case reflect.Slice:
		field.SetBytes(val.([]byte)) // BLOB
	}

}
