package extensions

import (
	amino "github.com/tendermint/go-amino"
)

func RegisterCodec(codec *amino.Codec) {
	codec.RegisterConcrete(Greeting{}, "hellochain/Greeting", nil)
	codec.RegisterConcrete(MsgGreet{}, "hellochain/MsgGreet", nil)
	codec.RegisterConcrete(GreetingsList{}, "hellochain/GreetingsList", nil)
	codec.RegisterConcrete(QueryResGreetings{}, "hellochain/QueryResGreetings", nil)
	codec.RegisterConcrete(QueryResNames{}, "nameservice/QueryResNames", nil)
	codec.RegisterConcrete(MsgDeleteName{}, "nameservice/MsgDeleteName", nil)
	codec.RegisterConcrete(MsgBuyName{}, "nameservice/MsgBuyName", nil)
	codec.RegisterConcrete(MsgSetName{}, "nameservice/MsgSetName", nil)
	codec.RegisterConcrete(QueryResResolve{}, "nameservice/QueryResResolve", nil)
	codec.RegisterConcrete(Whois{}, "nameservice/Whois", nil)

}
