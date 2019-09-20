package extensions

import (
	sdk "github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/types"
)

/*

You can declare types here or import them directly into codec
If you need SDK types, import the stripped-down version of the sdk's types

*/

// nameservice types
type MsgSetName struct {
	Name  string         `json:"name"`
	Value string         `json:"value"`
	Owner sdk.AccAddress `json:"owner"`
}
type MsgBuyName struct {
	Name  string         `json:"name"`
	Bid   sdk.Coins      `json:"bid"`
	Buyer sdk.AccAddress `json:"buyer"`
}
type MsgDeleteName struct {
	Name  string         `json:"name"`
	Owner sdk.AccAddress `json:"owner"`
}
type QueryResResolve struct {
	Value string `json:"value"`
}

type Whois struct {
	Value string         `json:"value"`
	Owner sdk.AccAddress `json:"owner"`
	Price sdk.Coins      `json:"price"`
}

type QueryResNames []string

// HelloChain types
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
