# Robot Mining Tool

> fork from [PoWERC20Miner](https://github.com/PowERC-20/PoWERC20Miner)

## Usage

0. **Compile the Binary**
   - go mod download
   - go build .

1. **Configuration**:
   - Customize your Ethereum private key and the contract address in the `init()` function.
   - Adjust the number of mining workers as needed.

2. **Running the Tool**:
   - Compile the tool by executing `go build .` in your terminal.
   - Launch the tool by executing `./miner-go -privateKey REPLACE_WITH_YOUR_PRIVATE_KEY` in your terminal.
   - You can use optional flags for specific configurations, for example: `./miner-go -privateKey YOUR_PRIVATE_KEY -referralAddress REFERRAL_ADDRESS -contractAddress CONTRACT_ADDRESS -workerCount NUMBER_OF_WORKERS`.
  
## Declare

The project code is completely open source. The released version is compiled using Github Actions. If you have any questions about the code, please feel free to raise them or submit the code for security auditing anywhere.

## Contributing

Your contributions are welcome. Please adhere to the project's coding standards and include tests for any new features or fixes.
