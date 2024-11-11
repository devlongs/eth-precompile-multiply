# Ethereum Multiply Precompile

This repository implements a custom precompiled contract for Ethereum that performs multiplication of two 256-bit integers.

## Overview

The precompile is located at address `0x14` and implements a simple multiplication operation with proper overflow checks and gas calculations.

## Features

- Multiplies two 256-bit integers
- Includes overflow protection
- Gas cost calculation based on input size
- Comprehensive test suite
- Solidity interface for easy integration

## Installation

1. Clone the repository:
```bash
git clone https://github.com/devlongs/eth-precompile-multiply.git
cd eth-precompile-multiply
```

2. Install dependencies:
```bash
go mod tidy
```

3. Run the test suit
```bash
go test ./...
```

## Usage
### In Go
```go
import "github.com/devlongs/eth-precompile-multiply/precompile"

// Register the precompile
address := common.BytesToAddress([]byte{0x14})
precompiles := vm.PrecompiledContracts{}
precompile.Register(precompiles, address)
```

### In Solidity
```solidity
import "./Multiply.sol";

contract MyContract {
    IMultiply constant multiply = IMultiply(0x14);

    function multiply(uint256 a, uint256 b) external view returns (uint256) {
        return multiply.multiply(a, b);
    }
}
```


