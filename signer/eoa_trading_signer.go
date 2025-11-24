package signer

import "github.com/ivanzzeth/ethsig"

type EOATradingSigner interface {
	ethsig.TypedDataSigner
	TransactionSignerAndAddrGetter
}
