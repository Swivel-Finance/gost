##### Smart contract testing with Geth via the Golang ABIGEN
```
                _______                                 
               /______/\                                
               \__::::\/                                
      _______      ______       ______       _________  
     /______/\    /_____/\     /_____/\     /________/\ 
     \::::__\/__  \:::_ \ \    \::::_\/_    \__.::.__\/     ,--.
      \:\ /____/\  \:\ \ \ \    \:\/___/\      \::\ \      |  oo|
       \:\\_  _\/   \:\ \ \ \    \_::._\:\      \::\ \     |  ~~|  o  o  o  o  o  o  o  o  o  o  
        \:\_\ \ \    \:\_\ \ \     /____\:\      \::\ \    |/\/\|
         \_____\/     \_____\/     \_____\/       \__\/ 
```
## Getting Started
This project contains the Swivel Smart Contracts and Libraries which have been compiled to their `abi` and `bin` components, those then transformed into
golang bindings via the Geth `abigen` tool. You can see the commands used to perform these tasks in the `Makefile`.

Tests are located in `/pkg/testing/`. 

Existing tests and any newly created are always run via `go test [-v] ./...` from the project root.

Gost only depends on Geth itself. We _do_ list `solc` as being necessary, but that isn't _exactly_ true.
If you are in possesion of the `abi` and `bin` files of any smart contract (regardless of language used) the
Makefile's `compile_go` type steps can be made for them. 

### Solidity Compiler
First, assure that `solc` is in your $PATH. If not, [make it so](https://docs.soliditylang.org/en/v0.8.0/installing-solidity.html). As stated above,
this is assuming your smart contracts are `.sol` and you do not already have the `.abi` and `.bin` files. Add new Makefile rules for other languages
(Vyper for example) and compilers.

### Geth ABIGEN
Second, you will need `abigen` in your $PATH. You can, for example, follow the instructions [here](https://github.com/ethereum/go-ethereum).
Note that you only need to use the `devtools` rule from the go-ethereum Makefile. i.e. `make devtools` will suffice.

### Project Structure
As always we use the [project layout](https://github.com/golang-standards/project-layout) guidelines. This interpretation places the
Solidity files into `test/contracts` with their compiled artifacts going into aptly named subdirectories for ease of use with
golang's package conventions (see Makefile). Some notes on the residents here:

* Swivel.sol. This is the whole point of this repo. Develop and test this contract (and it's libraries). Everything else here supports this file.
* Sig.sol. Imbeddable library used by Swivel.sol
* Hash.sol. Imbeddable library used by Swivel.sol
* Abstracts.sol. The external interfaces declared by Swivel.sol for our two token dependencies.
* SigFake.sol and HashFake.sol. Faux contracts which exist simply to compile and test the Sig and Hash libraries in isolation.
* Subdirectories:
    * /fakes. Compiled artifacts for the Fake.sol files
    * /tokens. Mock contracts (and their artifacts) used strictly to stub our token dependencies
    * /swivel. Compiled artifacts for the Swivel contracts.

#### The Makefile
This is the way.

### Compiling and Testing Your Contracts
* Compiling:
    * `make clean`. Removes any compiled artifacts from all directories
    * `make compile`. Compiles all contracts and libraries.
    * `make <foo>`. There are sub-commands for all possible steps so if you are adding new ones, **follow the pattern**
* Testing:
    * `go test ./...` from root (as stated). Add the -v flag if you expect to see any logging you may be doing.
    * You can of course test at the package level, but the test run is fast so... whateves...
    * This author runs individual tests via the excellent vim-go tool while developing. Those of you with cuter tools are on your own...

