.PHONY: compile_solidity_erc compile_go_erc compile_erc
.PHONY: compile_solidity_cerc compile_go_cerc compile_cerc
.PHONY: compile_solidity_sig_fake compile_go_sig_fake compile_sig_fake
.PHONY: compile_solidity_hash_fake compile_go_hash_fake compile_hash_fake
.PHONY: compile_solidity_swivel compile_go_swivel compile_swivel
.PHONY: compile

compile_solidity_erc:
	@echo "compiling ERC20 solidity source into abi and bin files"
	solc -o ./test/contracts/tokens --abi --bin --overwrite ./test/contracts/tokens/Erc20.sol

compile_go_erc:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/contracts/tokens/Erc20.abi --bin ./test/contracts/tokens/Erc20.bin -pkg tokens -type Erc20 -out ./test/contracts/tokens/erc20.go 

compile_erc: compile_solidity_erc compile_go_erc

compile_solidity_sig_fake:
	@echo "compiling Signature solidity source into abi and bin files"
	solc -o ./test/contracts/fakes --abi --bin --overwrite ./test/contracts/SigFake.sol

compile_go_sig_fake:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/contracts/fakes/SigFake.abi --bin ./test/contracts/fakes/SigFake.bin -pkg fakes -type SigFake -out ./test/contracts/fakes/sigfake.go 

compile_sig_fake: compile_solidity_sig_fake compile_go_sig_fake

compile_solidity_hash_fake:
	@echo "compiling Hash solidity source into abi and bin files"
	solc -o ./test/contracts/fakes --abi --bin --overwrite ./test/contracts/HashFake.sol

compile_go_hash_fake:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/contracts/fakes/HashFake.abi --bin ./test/contracts/fakes/HashFake.bin -pkg fakes -type HashFake -out ./test/contracts/fakes/hashfake.go 

compile_hash_fake: compile_solidity_hash_fake compile_go_hash_fake

compile_solidity_swivel:
	@echo "compiling Swivel solidity source into abi and bin files"
	solc -o ./test/contracts/swivel --abi --bin --overwrite ./test/contracts/Swivel.sol

compile_go_swivel:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/contracts/swivel/Swivel.abi --bin ./test/contracts/swivel/Swivel.bin -pkg swivel -type Swivel -out ./test/contracts/swivel/swivel.go 

compile_swivel: compile_solidity_swivel compile_go_swivel

compile: compile_sig_test compile_hash_test compile_swivel
