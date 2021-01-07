.PHONY: compile_solidity compile_go compile

compile_solidity:
	@echo "compiling solidity source into abi and bin files"
	solc -o ./test/contracts --abi --bin --overwrite ./test/contracts/Erc20.sol

compile_go:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/contracts/Erc20.abi -pkg token -type Token -out ./test/contracts/token/erc20.go -bin ./test/contracts/Erc20.bin

compile: compile_solidity compile_go
