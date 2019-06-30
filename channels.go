package clightningdb

type channels struct {
	Id                               int
	Peer_id                          int
	Short_channel_id                 []byte
	Channel_config_local             int
	Channel_config_remote            int
	State                            int
	Funder                           int
	Channel_flags                    int
	Minimum_depth                    int
	Next_index_local                 int
	Next_index_remote                int
	Next_htlc_id                     int
	Funding_tx_id                    []byte
	Funding_tx_outnum                int
	Funding_satoshi                  int
	Funding_locked_remote            int
	Push_msatoshi                    int
	Msatoshi_local                   int
	Fundingkey_remote                []byte
	Revocation_basepoint_remote      []byte
	Payment_basepoint_remote         []byte
	Htlc_basepoint_remote            []byte
	Delayed_payment_basepoint_remote []byte
	Per_commit_remote                []byte
	Old_per_commit_remote            []byte
	Local_feerate_per_kw             int
	Remote_feerate_per_kw            int
	Shachain_remote_id               int
	Shutdown_scriptpubkey_remote     []byte
	Shutdown_keyidx_local            int
	Last_sent_commit_state           int
	Last_sent_commit_id              int
	Last_tx                          []byte
	Last_sig                         []byte
	Closing_fee_received             int
	Closing_sig_received             []byte
	First_blocknum                   int
	Last_was_revoke                  int
	In_payments_offered              int
	In_payments_fulfilled            int
	In_msatoshi_offered              int
	In_msatoshi_fulfilled            int
	Out_payments_offered             int
	Out_payments_fulfilled           int
	Out_msatoshi_offered             int
	Out_msatoshi_fulfilled           int
	Min_possible_feerate             int
	Max_possible_feerate             int
	Msatoshi_to_us_min               int
	Msatoshi_to_us_max               int
	Future_per_commitment_point      []byte
	Last_sent_commit                 []byte
	Feerate_base                     int
	Feerate_ppm                      int
	Remote_upfront_shutdown_script   []byte
	Remote_ann_node_sig              []byte
	Remote_ann_bitcoin_sig           []byte
}

func (c *channels) String() string {
	return structString(c)
}
