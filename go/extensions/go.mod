module github.com/cosmos/amino-js/go/extensions

go 1.12

require (
	github.com/cosmos/amino-js/go/lib v0.0.0
	github.com/cosmos/hellochain v0.0.0-20190823001912-b2edd8984ccd
	github.com/tendermint/go-amino v0.15.0
)

replace github.com/cosmos/amino-js/go/lib => ../lib
