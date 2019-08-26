package extensions

import (
	sdk "github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/types"
)

/*

You can declare types here or import them directly into codec
If you need SDK types, import the stripped-down version of the sdk's types

*/

// MsgGreet defines the MsgGreet Message
type MsgGreet struct {
	Body      string
	Sender    sdk.AccAddress
	Recipient sdk.AccAddress
}

type Greeting struct {
	Sender    sdk.AccAddress
	Recipient sdk.AccAddress
	Body      string
}

type QueryResGreetings map[string][]Greeting

type GreetingsList []Greeting
