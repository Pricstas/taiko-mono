package producer

import (
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings/metadata"
)

var (
	errProofGenerating = errors.New("proof is generating")
	errEmptyProof      = errors.New("proof is empty")
	ErrInvalidLength   = errors.New("invalid items length")
)

// ProofRequestBody represents a request body to generate a proof.
type ProofRequestBody struct {
	Tier uint16
	Meta metadata.TaikoBlockMetaData
}

// ContestRequestBody represents a request body to generate a proof for contesting.
type ContestRequestBody struct {
	BlockID    *big.Int
	ProposedIn *big.Int
	ParentHash common.Hash
	Meta       metadata.TaikoBlockMetaData
	Tier       uint16
}

// ProofRequestOptions contains all options that need to be passed to a backend proof producer service.
type ProofRequestOptions struct {
	BlockID            *big.Int
	ProverAddress      common.Address
	ProposeBlockTxHash common.Hash
	TaikoL2            common.Address
	MetaHash           common.Hash
	BlockHash          common.Hash
	ParentHash         common.Hash
	StateRoot          common.Hash
	EventL1Hash        common.Hash
	Graffiti           string
	GasUsed            uint64
	ParentGasUsed      uint64
	Compressed         bool
}

type ProofWithHeader struct {
	BlockID *big.Int
	Meta    metadata.TaikoBlockMetaData
	Header  *types.Header
	Proof   []byte
	Opts    *ProofRequestOptions
	Tier    uint16
}

type BatchProofs struct {
	Proofs     []*ProofWithHeader
	BatchProof []byte
	Tier       uint16
	BlockIDs   []*big.Int
}

type ProofProducer interface {
	RequestProof(
		ctx context.Context,
		opts *ProofRequestOptions,
		blockID *big.Int,
		meta metadata.TaikoBlockMetaData,
		header *types.Header,
		requestAt time.Time,
	) (*ProofWithHeader, error)
	Aggregate(
		ctx context.Context,
		items []*ProofWithHeader,
		requestAt time.Time,
	) (*BatchProofs, error)
	RequestCancel(
		ctx context.Context,
		opts *ProofRequestOptions,
	) error
	Tier() uint16
}
