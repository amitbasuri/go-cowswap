package util

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

// SignTypedData Signs The TypedData
func SignTypedData(typedData apitypes.TypedData, privateKey *ecdsa.PrivateKey) (sig []byte, err error) {
	hash, err := EncodeForSigning(typedData)
	if err != nil {
		return
	}
	sig, err = crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return
	}
	sig[64] += 27
	return
}

// Encoding the data
func EncodeForSigning(typedData apitypes.TypedData) (hash common.Hash, err error) {
	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return
	}
	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return
	}
	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	hash = common.BytesToHash(crypto.Keccak256(rawData))
	return
}

func VerifySig(from, sigHex string, msg []byte) bool {
	sig := hexutil.MustDecode(sigHex)
	//msg = accounts.TextHash(msg)
	if sig[crypto.RecoveryIDOffset] == 27 || sig[crypto.RecoveryIDOffset] == 28 {
		sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}
	recovered, err := crypto.SigToPub(msg, sig)
	recoveredAddr1 := crypto.PubkeyToAddress(*recovered)
	fmt.Printf("the recovered address: %v \n", recoveredAddr1)
	if err != nil {
		return false
	}
	recoveredAddr := crypto.PubkeyToAddress(*recovered)
	return from == recoveredAddr.Hex()
}
