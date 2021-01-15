##### Smart contract testing with Geth via the Golang ABIGEN
```
                _______                                 
               /______/\                                
               \__::::\/                                
      _______      ______       ______       _________  
     /______/\    /_____/\     /_____/\     /________/\ 
     \::::__\/__  \:::_ \ \    \::::_\/_    \__.::.__\/     ,--.
      \:\ /____/\  \:\ \ \ \    \:\/___/\      \::\ \      |  oo| 
       \:\\_  _\/   \:\ \ \ \    \_::._\:\      \::\ \     |  ~~|
        \:\_\ \ \    \:\_\ \ \     /____\:\      \::\ \    |/\/\|
         \_____\/     \_____\/     \_____\/       \__\/ 
```
## Getting Started
This project contains a sample Erc20.sol contract file which has been compiled to its `abi` and `bin` components with those then transformed into
golang bindings via the Geth `abigen` tool. You can see the commands used to perform those tasks in the `Makefile`.

Tests for that sample contract are located in `/pkg/testing/`. 

Existing tests and any newly created are always run via `go test [-v] ./...` from the project root.

To compile and test your smart contracts, simply follow the example provided. Gost only depends on Geth itself. We _do_ list `solc` as being
necessary, but that isn't _exactly_ true. If you are in possesion of the `abi` and `bin` files of any smart contract (regardless of language used) the
Makefile's `compile_go` type steps can be made for them. 

### Solidity Compiler
First, assure that `solc` is in your $PATH. If not, [make it so](https://docs.soliditylang.org/en/v0.8.0/installing-solidity.html). As stated above,
this is assuming your smart contracts are `.sol` and you do not already have the `.abi` and `.bin` files. Add new Makefile rules for other languages
(Vyper for example) and compilers.

### Geth ABIGEN
Second, you will need `abigen` in your $PATH. You can, for example, follow the instructions [here](https://github.com/ethereum/go-ethereum).
Note that you only need to use the `devtools` rule from the go-ethereum Makefile. i.e. `make devtools` will suffice.

### Project Structure
TODO 

#### The Makefile
TODO

### Compiling and Testing Your Contracts
TODO
