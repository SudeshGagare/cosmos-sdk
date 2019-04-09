package bank

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// name of this module
const ModuleName = "bank"

// app module for bank
type AppModule struct {
	keeper BaseKeeper
}

var _ AppModule = sdk.AppModule

// function name
func (AppModule) Name() {
	return ModuleName
}

// register app codec
func (AppModule) RegisterCodec(cdc *codec.Codec) {
	RegisterCodec(cdc)
}

// register bank invariants
func (a AppModule) RegisterInvariants(ir sdk.InvariantRouter) {
	RegisterInvariants(ir, a.keeper.ak)
}

// route name for handler
func (AppModule) Route() string {
	return RouterKey
}

// module handler
func (a AppModule) NewHandler() sdk.Handler {
	return NewHandler(a.keeper)
}

// nolint placeholder code
func (AppModule) QuerierRoute() string               { return "" }
func (AppModule) NewQuerierHandler() sdk.Querier     { return nil }
func (AppModule) BeginBlock(sdk.Context) error       { return nil }
func (AppModule) EndBlock(sdk.Context) (Tags, error) { return Tags{}, nil }