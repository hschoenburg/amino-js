package extensions

import (
	g "github.com/cosmos/hellochain/x/greeter/types"
	amino "github.com/tendermint/go-amino"
)

func RegisterCodec(codec *amino.Codec) {
	codec.RegisterConcrete(g.Greeting{}, "hellochain/Greeting", nil)
	codec.RegisterConcrete(g.MsgGreet{}, "hellochain/Greeting", nil)
	codec.RegisterConcrete(g.GreetingsList{}, "hellochain/GreetingsList", nil)
	codec.RegisterConcrete(g.QueryResGreetings{}, "hellochain/QueryResGreetings", nil)
}
