package precompile

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

var (
	errInvalidInputLength = errors.New("invalid input length")
	errInputTooLarge      = errors.New("input numbers too large")
	errOverflow           = errors.New("multiplication would overflow")
)

const (
	MultiplyBaseGas    uint64 = 100
	MultiplyPerByteGas uint64 = 10
)

// Multiply is a precompiled contract that multiplies two 256-bit integers
type Multiply struct{}

// RequiredGas calculates the gas required to execute the pre-compiled contract
func (m *Multiply) RequiredGas(input []byte) uint64 {
	return MultiplyBaseGas + uint64(len(input))*MultiplyPerByteGas
}

// Run executes the precompiled contract
func (m *Multiply) Run(input []byte) ([]byte, error) {
	if len(input) != 64 {
		return nil, errInvalidInputLength
	}

	num1 := new(big.Int).SetBytes(input[:32])
	num2 := new(big.Int).SetBytes(input[32:64])

	if num1.BitLen() > 256 || num2.BitLen() > 256 {
		return nil, errInputTooLarge
	}

	if num1.BitLen()+num2.BitLen() > 256 {
		return nil, errOverflow
	}

	result := new(big.Int).Mul(num1, num2)

	return common.LeftPadBytes(result.Bytes(), 32), nil
}

// Register adds the multiply precompile to the given address in the provided PrecompiledContracts map
func Register(precompiles vm.PrecompiledContracts, address common.Address) {
	precompiles[address] = &Multiply{}
}
