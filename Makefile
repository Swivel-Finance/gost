.PHONY: compile_solidity_sig_test compile_go_sig_test compile_sig_test
.PHONY: compile_solidity_hash_test compile_go_hash_test compile_hash_test
.PHONY: compile_solidity_swivel compile_go_swivel compile_swivel
.PHONY: compile

compile_solidity_sig_test:
	@echo "compiling Signature solidity source into abi and bin files"
	solc -o ./test/contracts --abi --bin --overwrite ./test/contracts/SigTest.sol

compile_go_sig_test:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/contracts/SigTest.abi --bin ./test/contracts/SigTest.bin -pkg swivel -type SigTest -out ./test/contracts/swivel/sigtest.go 

compile_sig_test: compile_solidity_sig_test compile_go_sig_test

compile_solidity_hash_test:
	@echo "compiling Hash solidity source into abi and bin files"
	solc -o ./test/contracts --abi --bin --overwrite ./test/contracts/HashTest.sol

compile_go_hash_test:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/contracts/HashTest.abi --bin ./test/contracts/HashTest.bin -pkg swivel -type HashTest -out ./test/contracts/swivel/hashtest.go 

compile_hash_test: compile_solidity_hash_test compile_go_hash_test

compile_solidity_swivel:
	@echo "compiling Swivel solidity source into abi and bin files"
	solc -o ./test/contracts --abi --bin --overwrite ./test/contracts/Swivel.sol

compile_go_swivel:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/contracts/Swivel.abi --bin ./test/contracts/Swivel.bin -pkg swivel -type Swivel -out ./test/contracts/swivel/swivel.go 

compile_swivel: compile_solidity_swivel compile_go_swivel

compile: compile_sig_test compile_hash_test compile_swivel
