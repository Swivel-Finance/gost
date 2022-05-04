.PHONY: compile_solidity_mock_erc compile_go_mock_erc compile_mock_erc
.PHONY: compile_solidity_mock_yield_token compile_go_mock_yield_token compile_mock_yield_token
.PHONY: compile_solidity_mock_element_token compile_go_mock_element_token compile_mock_element_token
.PHONY: compile_solidity_mock_pendle_data compile_go_mock_pendle_data compile_mock_pendle_data
.PHONY: compile_solidity_mock_pendle_router compile_go_mock_pendle_router compile_mock_pendle_router
.PHONY: compile_solidity_mock_pendle_yield_token compile_go_mock_pendle_yield_token compile_mock_pendle_yield_token
.PHONY: compile_solidity_mock_market_place compile_go_mock_market_place compile_mock_market_place
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

compile_solidity_mock_yield_token:
	@echo "compiling Mock YToken solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/YieldToken.sol

compile_go_mock_yield_token:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/YieldToken.abi --bin ./test/mocks/YieldToken.bin -pkg mocks -type YieldToken -out ./test/mocks/yieldtoken.go 

compile_mock_yield_token: compile_solidity_mock_yield_token compile_go_mock_yield_token

compile_solidity_mock_pendle_data:
	@echo "compiling Mock Pendle Data solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/PendleData.sol

compile_go_mock_pendle_data:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/PendleData.abi --bin ./test/mocks/PendleData.bin -pkg mocks -type PendleData -out ./test/mocks/pendledata.go 

compile_mock_pendle_data: compile_solidity_mock_pendle_data compile_go_mock_pendle_data

compile_solidity_mock_pendle_router:
	@echo "compiling Mock Pendle Router solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/PendleRouter.sol

compile_go_mock_pendle_router:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/PendleRouter.abi --bin ./test/mocks/PendleRouter.bin -pkg mocks -type PendleRouter -out ./test/mocks/pendlerouter.go 

compile_mock_pendle_router: compile_solidity_mock_pendle_router compile_go_mock_pendle_router

compile_solidity_mock_pendle_yield_token:
	@echo "compiling Mock Pendle Yield Token solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/PendleYieldToken.sol

compile_go_mock_pendle_yield_token:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/PendleYieldToken.abi --bin ./test/mocks/PendleYieldToken.bin -pkg mocks -type PendleYieldToken -out ./test/mocks/pendleyieldtoken.go 

compile_mock_pendle_yield_token: compile_solidity_mock_pendle_yield_token compile_go_mock_pendle_yield_token

compile_solidity_mock_element_token:
	@echo "compiling Mock ElementToken solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/ElementToken.sol

compile_go_mock_element_token:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/ElementToken.abi --bin ./test/mocks/ElementToken.bin -pkg mocks -type ElementToken -out ./test/mocks/elementtoken.go 

compile_mock_element_token: compile_solidity_mock_element_token compile_go_mock_element_token

compile_solidity_mock_element:
	@echo "compiling Mock Element solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/Element.sol

compile_go_mock_element:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/Element.abi --bin ./test/mocks/Element.bin -pkg mocks -type Element -out ./test/mocks/element.go 

compile_mock_element: compile_solidity_mock_element compile_go_mock_element

compile_solidity_mock_swivel:
	@echo "compiling Mock Swivel solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/Swivel.sol

compile_go_mock_swivel:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/Swivel.abi --bin ./test/mocks/Swivel.bin -pkg mocks -type Swivel -out ./test/mocks/swivel.go 

compile_mock_swivel: compile_solidity_mock_swivel compile_go_mock_swivel

compile_solidity_mock_market_place:
	@echo "compiling Mock MarketPlace solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/MarketPlace.sol

compile_go_mock_market_place:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/MarketPlace.abi --bin ./test/mocks/MarketPlace.bin -pkg mocks -type MarketPlace -out ./test/mocks/marketplace.go 

compile_mock_market_place: compile_solidity_mock_market_place compile_go_mock_market_place

compile_solidity_mock_zc_token:
	@echo "compiling Mock ZcToken solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/ZcToken.sol

compile_go_mock_zc_token:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/ZcToken.abi --bin ./test/mocks/ZcToken.bin -pkg mocks -type ZcToken -out ./test/mocks/zctoken.go 

compile_mock_zc_token: compile_solidity_mock_zc_token compile_go_mock_zc_token

compile_mocks: compile_mock_erc compile_mock_yield_token compile_mock_element_token compile_mock_element compile_mock_market_place compile_mock_zc_token compile_mock_swivel compile_mock_pendle_data compile_mock_pendle_router compile_mock_pendle_yield_token

# Real Tokens
# compile_solidity_zct:
# 	@echo "compiling ZCT solidity source into abi and bin files"
# 	solc -o ./test/tokens --abi --bin --overwrite ./test/tokens/ZcToken.sol

# compile_go_zct:
# 	@echo "compiling abi and bin files to golang"
# 	abigen --abi ./test/tokens/ZcToken.abi --bin ./test/tokens/ZcToken.bin -pkg tokens -type ZcToken -out ./test/tokens/zctoken.go 

# compile_zct: compile_solidity_zct compile_go_zct

# compile_tokens: compile_zct

# Fakes

# Contracts
compile_solidity_lender_test:
	@echo "compiling Lender solidity source into abi and bin files"
	solc -o ./test/lender --optimize --optimize-runs=15000 --abi --bin --overwrite ./test/lender/Lender.sol

compile_go_lender_test:
	@echo "compiling Lender abi and bin files to golang"
	abigen --abi ./test/lender/Lender.abi --bin ./test/lender/Lender.bin -pkg lender -type Lender -out ./test/lender/lender.go

compile_lender_test: compile_solidity_lender_test compile_go_lender_test

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
