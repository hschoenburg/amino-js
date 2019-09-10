package extensions

import (
	amino "github.com/tendermint/go-amino"
)

func RegisterCodec(codec *amino.Codec) {
	codec.RegisterConcrete(Greeting{}, "hellochain/Greeting", nil)
	codec.RegisterConcrete(MsgGreet{}, "hellochain/MsgGreet", nil)
	codec.RegisterConcrete(GreetingsList{}, "hellochain/GreetingsList", nil)
	codec.RegisterConcrete(QueryResGreetings{}, "hellochain/QueryResGreetings", nil)
}
