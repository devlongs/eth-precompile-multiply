package precompile

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestMultiply_RequiredGas(t *testing.T) {
	m := &Multiply{}

	tests := []struct {
		input    []byte
		expected uint64
	}{
		{make([]byte, 64), MultiplyBaseGas + 64*MultiplyPerByteGas},
		{make([]byte, 32), MultiplyBaseGas + 32*MultiplyPerByteGas},
	}

	for _, test := range tests {
		gas := m.RequiredGas(test.input)
		assert.Equal(t, test.expected, gas)
	}
}

func TestMultiply_Run(t *testing.T) {
	m := &Multiply{}

	tests := []struct {
		name     string
		num1     *big.Int
		num2     *big.Int
		expected *big.Int
		err      error
	}{
		{
			name:     "simple multiplication",
			num1:     big.NewInt(2),
			num2:     big.NewInt(3),
			expected: big.NewInt(6),
			err:      nil,
		},
		{
			name:     "zero multiplication",
			num1:     big.NewInt(0),
			num2:     big.NewInt(5),
			expected: big.NewInt(0),
			err:      nil,
		},
		{
			name:     "large numbers",
			num1:     new(big.Int).SetUint64(1 << 63),
			num2:     big.NewInt(2),
			expected: new(big.Int).Lsh(big.NewInt(1), 64),
			err:      nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			input := make([]byte, 64)
			copy(input[32-len(test.num1.Bytes()):32], test.num1.Bytes())
			copy(input[64-len(test.num2.Bytes()):], test.num2.Bytes())

			result, err := m.Run(input)

			if test.err != nil {
				assert.Error(t, err)
				assert.Equal(t, test.err, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, common.LeftPadBytes(test.expected.Bytes(), 32), result)
			}
		})
	}
}
