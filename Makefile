.PHONY: compile_solidity_sig compile_go_sig compile_sig

compile_solidity_sig:
	@echo "compiling Signature solidity source into abi and bin files"
	solc -o ./test/contracts --abi --bin --overwrite ./test/contracts/Sig.sol

compile_go_sig:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/contracts/Sig.abi --bin ./test/contracts/Sig.bin -pkg swivel -type Sig -out ./test/contracts/swivel/sig.go 

compile_sig: compile_solidity_sig compile_go_sig
