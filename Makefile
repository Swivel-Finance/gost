.PHONY: compile_solidity_mock_erc compile_go_mock_erc compile_mock_erc
.PHONY: compile_solidity_mock_cerc compile_go_mock_cerc compile_mock_cerc
.PHONY: compile_solidity_mock_ferc compile_go_mock_ferc compile_mock_ferc
.PHONY: compile_mocks

# TODO under? api-specific sol?

.PHONY: compile_solidity_zct compile_go_zct compile_zct
.PHONY: compile_tokens

.PHONY: clean_test_abi clean_test_bin clean_test_go
.PHONY: clean_test

# .PHONY: clean_build_sol clean_build_abi clean_build_bin clean_build_go
# .PHONY: clean_build

# .PHONY: copy_to_build

.PHONY: all

# Mocks
compile_solidity_mock_erc:
	@echo "compiling Mock ERC20 solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/Erc20.sol

compile_go_mock_erc:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/Erc20.abi --bin ./test/mocks/Erc20.bin -pkg mocks -type Erc20 -out ./test/mocks/erc20.go 

compile_mock_erc: compile_solidity_mock_erc compile_go_mock_erc

compile_solidity_mock_cerc:
	@echo "compiling Mock CERC20 solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/CErc20.sol

compile_go_mock_cerc:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/CErc20.abi --bin ./test/mocks/CErc20.bin -pkg mocks -type CErc20 -out ./test/mocks/cerc20.go 

compile_mock_cerc: compile_solidity_mock_cerc compile_go_mock_cerc

compile_solidity_mock_ferc:
	@echo "compiling Mock FERC20 solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/FErc20.sol

compile_go_mock_ferc:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/FErc20.abi --bin ./test/mocks/FErc20.bin -pkg mocks -type FErc20 -out ./test/mocks/ferc20.go 

compile_mock_ferc: compile_solidity_mock_ferc compile_go_mock_ferc

compile_mocks: compile_mock_erc compile_mock_cerc compile_mock_ferc

# Real Tokens
compile_solidity_zct:
	@echo "compiling ZCT solidity source into abi and bin files"
	solc -o ./test/tokens --abi --bin --overwrite ./test/tokens/ZcToken.sol

compile_go_zct:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/tokens/ZcToken.abi --bin ./test/tokens/ZcToken.bin -pkg tokens -type ZcToken -out ./test/tokens/zctoken.go 

compile_zct: compile_solidity_zct compile_go_zct

compile_solidity_underlying:
	@echo "compiling Underlying solidity source into abi and bin files"
	solc -o ./test/tokens --abi --bin --overwrite ./test/tokens/Underlying.sol

compile_go_underlying:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/tokens/Underlying.abi --bin ./test/tokens/Underlying.bin -pkg tokens -type Underlying -out ./test/tokens/underlying.go

compile_underlying: compile_solidity_underlying compile_go_underlying

compile_tokens: compile_zct compile_underlying

# Fakes

# Contracts

# Cleaning
clean_test_abi:
	@echo "removing abi files from test/ dirs"
	rm test/**/*.abi

clean_test_bin:
	@echo "removing bin files from test/ dirs"
	rm test/**/*.bin

clean_test_go:
	@echo "removing generated go bindings from test/ dirs"
	rm test/**/*.go

clean_test: clean_test_abi clean_test_bin clean_test_go

# clean_build_sol:
# 	@echo "removing sol files from build/ dirs"
# 	rm build/**/*.sol
# 
# clean_build_abi:
# 	@echo "removing abi files from build/ dirs"
# 	rm build/**/*.abi
# 
# clean_build_bin:
# 	@echo "removing bin files from build/ dirs"
# 	rm build/**/*.bin
# 
# clean_build_go:
# 	@echo "removing go files from build/ dirs"
# 	rm build/**/*.go
# 
# clean_build: clean_build_sol clean_build_abi clean_build_bin clean_build_go

all: clean_test compile_test
