package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/ibc/02-client/exported"
	"github.com/cosmos/cosmos-sdk/x/ibc/02-client/types/errors"
	host "github.com/cosmos/cosmos-sdk/x/ibc/24-host"
	ibctypes "github.com/cosmos/cosmos-sdk/x/ibc/types"
)

var _ sdk.Msg = MsgCreateClient{}

// MsgCreateClient defines a message to create an IBC client
type MsgCreateClient struct {
	ClientID       string                  `json:"client_id" yaml:"client_id"`
	ClientType     string                  `json:"client_type" yaml:"client_type"`
	ConsensusState exported.ConsensusState `json:"consensus_state" yaml:"consensus_address"`
	Signer         sdk.AccAddress          `json:"address" yaml:"address"`
}

// NewMsgCreateClient creates a new MsgCreateClient instance
func NewMsgCreateClient(id, clientType string, consensusState exported.ConsensusState, signer sdk.AccAddress) MsgCreateClient {
	return MsgCreateClient{
		ClientID:       id,
		ClientType:     clientType,
		ConsensusState: consensusState,
		Signer:         signer,
	}
}

// Route implements sdk.Msg
func (msg MsgCreateClient) Route() string {
	return ibctypes.RouterKey
}

// Type implements sdk.Msg
func (msg MsgCreateClient) Type() string {
	return "create_client"
}

// ValidateBasic implements sdk.Msg
func (msg MsgCreateClient) ValidateBasic() sdk.Error {
	if err := host.DefaultClientIdentifierValidator(msg.ClientID); err != nil {
		return sdk.ConvertError(err)
	}
	if _, err := exported.ClientTypeFromString(msg.ClientType); err != nil {
		return sdk.ConvertError(errors.ErrInvalidClientType(errors.DefaultCodespace, err.Error()))
	}
	if msg.ConsensusState == nil {
		return sdk.ConvertError(errors.ErrInvalidConsensus(errors.DefaultCodespace))
	}
	if msg.Signer.Empty() {
		return sdk.ErrInvalidAddress("empty address")
	}
	return nil
}

// GetSignBytes implements sdk.Msg
func (msg MsgCreateClient) GetSignBytes() []byte {
	return sdk.MustSortJSON(SubModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements sdk.Msg
func (msg MsgCreateClient) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}

var _ sdk.Msg = MsgUpdateClient{}

// MsgUpdateClient defines a message to update an IBC client
type MsgUpdateClient struct {
	ClientID string          `json:"client_id" yaml:"client_id"`
	Header   exported.Header `json:"header" yaml:"header"`
	Signer   sdk.AccAddress  `json:"address" yaml:"address"`
}

// NewMsgUpdateClient creates a new MsgUpdateClient instance
func NewMsgUpdateClient(id string, header exported.Header, signer sdk.AccAddress) MsgUpdateClient {
	return MsgUpdateClient{
		ClientID: id,
		Header:   header,
		Signer:   signer,
	}
}

// Route implements sdk.Msg
func (msg MsgUpdateClient) Route() string {
	return ibctypes.RouterKey
}

// Type implements sdk.Msg
func (msg MsgUpdateClient) Type() string {
	return "update_client"
}

// ValidateBasic implements sdk.Msg
func (msg MsgUpdateClient) ValidateBasic() sdk.Error {
	if err := host.DefaultClientIdentifierValidator(msg.ClientID); err != nil {
		return sdk.ConvertError(err)
	}
	if msg.Header == nil {
		return sdk.ConvertError(errors.ErrInvalidHeader(errors.DefaultCodespace))
	}
	if msg.Signer.Empty() {
		return sdk.ErrInvalidAddress("empty address")
	}
	return nil
}

// GetSignBytes implements sdk.Msg
func (msg MsgUpdateClient) GetSignBytes() []byte {
	return sdk.MustSortJSON(SubModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements sdk.Msg
func (msg MsgUpdateClient) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}
