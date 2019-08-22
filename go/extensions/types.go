package extensions

import (
	sdk "github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = (*TxCreateMarket)(nil)

type TxCreateMarket struct {
	Account sdk.AccAddress
	Market  string
}

// Greeeting is the paylod of MsgGreet
type Greeting struct {
	Sender    sdk.AccAddress `json:"sender" yaml:"sender"`
	Recipient sdk.AccAddress `json:"recipient" yaml:"recipient"`
	Body      string         `json:"body" yaml:"body"`
}
