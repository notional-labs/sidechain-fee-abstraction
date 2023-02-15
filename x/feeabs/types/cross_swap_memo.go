package types

import (
	"encoding/json"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
)

type OsmosisSpecialMemo struct {
	Wasm map[string]interface{} `json:"wasm"`
}

type OsmosisSwapMsg struct {
	OsmosisSwap Swap `json:"osmosis_swap"`
}
type Swap struct {
	InputCoin   sdk.Coin `json:"input_coin"`
	OutPutDenom string   `json:"output_denom"`
	Slippage    Twap     `json:"slippage"`
	Receiver    string   `json:"receiver"`
}

type Twap struct {
	Twap TwapRouter `json:"twap"`
}

type TwapRouter struct {
	SlippagePercentage string `json:"slippage_percentage"`
	WindowSeconds      uint64 `json:"window_seconds"`
}

func NewOsmosisSwapMsg(inputCoin sdk.Coin, outputDenom string, slippagePercentage string, windowSeconds uint64, receiver string) OsmosisSwapMsg {
	swap := Swap{
		InputCoin:   inputCoin,
		OutPutDenom: outputDenom,
		Slippage: Twap{
			Twap: TwapRouter{SlippagePercentage: slippagePercentage,
				WindowSeconds: windowSeconds,
			}},
		Receiver: receiver,
	}

	return OsmosisSwapMsg{
		OsmosisSwap: swap,
	}
}

// ParseMsgToMemo build a memo from msg, contractAddr, compatible with ValidateAndParseMemo in https://github.com/osmosis-labs/osmosis/blob/nicolas/crosschain-swaps-new/x/ibc-hooks/wasm_hook.go
func ParseMsgToMemo(msg OsmosisSwapMsg, contractAddr string, receiver string) (string, error) {
	// TODO: need to validate the msg && contract address
	memo := OsmosisSpecialMemo{
		Wasm: make(map[string]interface{}),
	}

	memo.Wasm["contract"] = contractAddr
	memo.Wasm["msg"] = msg
	memo.Wasm["receiver"] = receiver

	memo_marshalled, err := json.Marshal(&memo)
	if err != nil {
		return "", err
	}
	return string(memo_marshalled), nil
}

func BuildPacketMiddlewareMemo(inputToken sdk.Coin, outputDenom string, receiver string, hostChainConfig HostChainFeeAbsConfig) (string, error) {
	// TODO: this should be chain params.
	timeOut := time.Duration(1800000)
	retries := uint8(8)
	nextMemo, err := BuildNextMemo(inputToken, outputDenom, hostChainConfig.CrosschainSwapAddress, receiver)
	if err != nil {
		return "", nil
	}
	metadata := ForwardMetadata{
		Receiver: hostChainConfig.MiddlewareAddress,
		Port:     transfertypes.PortID,
		Channel:  hostChainConfig.HostZoneIbcTransferChannel,
		Timeout:  timeOut,
		Retries:  &retries,
		Next:     nextMemo,
	}

	// TODO: need to validate the msg && contract address.
	memo_marshalled, err := json.Marshal(&metadata)
	if err != nil {
		return "", err
	}
	return string(memo_marshalled), nil
}
