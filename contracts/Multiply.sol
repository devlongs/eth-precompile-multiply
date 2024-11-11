// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IMultiply {
    function multiply(uint256 a, uint256 b) external view returns (uint256);
}

contract MultiplyUser {
    IMultiply constant MULTIPLY = IMultiply(0x0000000000000000000000000000000000000014);

    function multiplyNumbers(uint256 a, uint256 b) external view returns (uint256) {
        bytes memory input = abi.encodePacked(a, b);
        (bool success, bytes memory result) = address(MULTIPLY).staticcall(input);
        require(success, "Precompile call failed");

        return abi.decode(result, (uint256));
    }
}
