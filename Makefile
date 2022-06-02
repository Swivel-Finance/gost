.PHONY: compile_solidity_mock_erc compile_go_mock_erc compile_mock_erc
.PHONY: compile_solidity_mock_yield compile_go_mock_yield compile_mock_yield
.PHONY: compile_solidity_mock_yield_token compile_go_mock_yield_token compile_mock_yield_token
.PHONY: compile_solidity_mock_swivel compile_go_mock_swivel compile_mock_swivel
.PHONY: compile_solidity_mock_pendle compile_go_mock_pendle compile_mock_pendle
.PHONY: compile_solidity_mock_pendle_token compile_go_mock_pendle_token compile_mock_pendle_token
.PHONY: compile_solidity_mock_tempus compile_go_mock_tempus compile_mock_tempus
.PHONY: compile_solidity_mock_tempus_token compile_go_mock_tempus_token compile_mock_tempus_token
.PHONY: compile_solidity_mock_sense compile_go_mock_sense compile_mock_sense
.PHONY: compile_solidity_mock_apwine_token compile_go_mock_apwine_token compile_mock_apwine_token
.PHONY: compile_solidity_mock_notional_token compile_go_mock_notional_token compile_mock_notional_token
.PHONY: compile_solidity_mock_apwine compile_go_mock_apwine compile_mock_apwine
.PHONY: compile_solidity_mock_sense_token compile_go_mock_sense_token compile_mock_sense_token
.PHONY: compile_solidity_mock_element_token compile_go_mock_element_token compile_mock_element_token
.PHONY: compile_solidity_mock_marketplace compile_go_mock_marketplace compile_mock_marketplace
.PHONY: compile_mock_deps
.PHONY: compile_mocks

.PHONY: compile_solidity_zct compile_go_zct compile_zct
.PHONY: compile_tokens

.PHONY: compile_solidity_marketplace_test compile_go_marketplace_test compile_marketplace_test
.PHONY: compile_solidity_lender_test compile_go_lender_test compile_lender_test
.PHONY: compile_solidity_redeemer_test compile_go_redeemer_test compile_redeemer_test
.PHONY: compile_test

.PHONY: clean_test_abi clean_test_bin clean_test_go
.PHONY: clean_test

.PHONY: clean_build_sol clean_build_abi clean_build_bin clean_build_go
.PHONY: clean_build

.PHONY: copy_to_build

.PHONY: compile_marketplace_build compile_lender_build compile_redeemer_build
.PHONY: compile_build

.PHONY: all

# Mocks
compile_solidity_mock_erc:
	@echo "compiling Mock ERC20 solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/Erc20.sol

compile_go_mock_erc:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/Erc20.abi --bin ./test/mocks/Erc20.bin -pkg mocks -type Erc20 -out ./test/mocks/erc20.go 

compile_mock_erc: compile_solidity_mock_erc compile_go_mock_erc

compile_solidity_mock_yield:
	@echo "compiling Mock Yield solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/Yield.sol

compile_go_mock_yield:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/Yield.abi --bin ./test/mocks/Yield.bin -pkg mocks -type Yield -out ./test/mocks/yield.go 

compile_mock_yield: compile_solidity_mock_yield compile_go_mock_yield

compile_solidity_mock_yield_token:
	@echo "compiling Mock YieldToken solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/YieldToken.sol

compile_go_mock_yield_token:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/YieldToken.abi --bin ./test/mocks/YieldToken.bin -pkg mocks -type YieldToken -out ./test/mocks/yieldtoken.go 

compile_mock_yield_token: compile_solidity_mock_yield_token compile_go_mock_yield_token

compile_solidity_mock_pendle_token:
	@echo "compiling Mock Pendle Token solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/PendleToken.sol

compile_go_mock_pendle_token:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/PendleToken.abi --bin ./test/mocks/PendleToken.bin -pkg mocks -type PendleToken -out ./test/mocks/pendletoken.go 

compile_mock_pendle_token: compile_solidity_mock_pendle_token compile_go_mock_pendle_token

compile_solidity_mock_pendle:
	@echo "compiling Mock Pendle solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/Pendle.sol

compile_go_mock_pendle:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/Pendle.abi --bin ./test/mocks/Pendle.bin -pkg mocks -type Pendle -out ./test/mocks/pendle.go 

compile_mock_pendle: compile_mock_pendle_token compile_solidity_mock_pendle compile_go_mock_pendle

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

compile_mock_element: compile_mock_element_token compile_solidity_mock_element compile_go_mock_element

compile_solidity_mock_swivel:
	@echo "compiling Mock Swivel solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/Swivel.sol

compile_go_mock_swivel:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/Swivel.abi --bin ./test/mocks/Swivel.bin -pkg mocks -type Swivel -out ./test/mocks/swivel.go 

compile_mock_swivel: compile_solidity_mock_swivel compile_go_mock_swivel

compile_solidity_mock_marketplace:
	@echo "compiling Mock MarketPlace solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/MarketPlace.sol

compile_go_mock_marketplace:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/MarketPlace.abi --bin ./test/mocks/MarketPlace.bin -pkg mocks -type MarketPlace -out ./test/mocks/marketplace.go 

compile_mock_marketplace: compile_solidity_mock_marketplace compile_go_mock_marketplace

compile_solidity_mock_zc_token:
	@echo "compiling Mock ZcToken solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/ZcToken.sol

compile_go_mock_zc_token:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/ZcToken.abi --bin ./test/mocks/ZcToken.bin -pkg mocks -type ZcToken -out ./test/mocks/zctoken.go 

compile_mock_zc_token: compile_solidity_mock_zc_token compile_go_mock_zc_token

compile_solidity_mock_tempus:
	@echo "compiling Mock Tempus solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/Tempus.sol

compile_go_mock_tempus:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/Tempus.abi --bin ./test/mocks/Tempus.bin -pkg mocks -type Tempus -out ./test/mocks/tempus.go 

compile_mock_tempus: compile_solidity_mock_tempus compile_go_mock_tempus

compile_solidity_mock_tempus_token:
	@echo "compiling Mock Tempus Token solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/TempusToken.sol

compile_go_mock_tempus_token:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/TempusToken.abi --bin ./test/mocks/TempusToken.bin -pkg mocks -type TempusToken -out ./test/mocks/tempustoken.go 

compile_mock_tempus_token: compile_solidity_mock_tempus_token compile_go_mock_tempus_token

compile_solidity_mock_sense:
	@echo "compiling Mock Sense solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/Sense.sol

compile_go_mock_sense:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/Sense.abi --bin ./test/mocks/Sense.bin -pkg mocks -type Sense -out ./test/mocks/sense.go 

compile_mock_sense: compile_solidity_mock_sense compile_go_mock_sense

compile_solidity_mock_sense_token:
	@echo "compiling Mock SenseToken solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/SenseToken.sol

compile_go_mock_sense_token:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/SenseToken.abi --bin ./test/mocks/SenseToken.bin -pkg mocks -type SenseToken -out ./test/mocks/sensetoken.go 

compile_mock_sense_token: compile_solidity_mock_sense_token compile_go_mock_sense_token

compile_solidity_mock_apwine_token:
	@echo "compiling Mock APWine Token solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/APWineToken.sol

compile_go_mock_apwine_token:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/APWineToken.abi --bin ./test/mocks/APWineToken.bin -pkg mocks -type APWineToken -out ./test/mocks/apwinetoken.go 

compile_mock_apwine_token: compile_solidity_mock_apwine_token compile_go_mock_apwine_token

compile_solidity_mock_apwine:
	@echo "compiling Mock APWine solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/APWine.sol

compile_go_mock_apwine:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/APWine.abi --bin ./test/mocks/APWine.bin -pkg mocks -type APWine -out ./test/mocks/apwine.go 

# TODO: Add the notional mock implementation here

compile_solidity_mock_notional_token:
	@echo "compiling Mock APWine Token solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/NotionalToken.sol

compile_go_mock_notional_token:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/NotionalToken.abi --bin ./test/mocks/NotionalToken.bin -pkg mocks -type NotionalToken -out ./test/mocks/notionaltoken.go 

compile_mock_notional_token: compile_solidity_mock_notional_token compile_go_mock_notional_token

compile_mock_apwine: compile_mock_apwine_token compile_solidity_mock_apwine compile_go_mock_apwine

compile_mock_deps: compile_mock_erc compile_mock_zc_token

compile_mocks: compile_mock_deps compile_mock_marketplace compile_mock_swivel compile_mock_yield compile_mock_yield_token compile_mock_element compile_mock_pendle compile_mock_tempus compile_mock_sense compile_mock_sense_token compile_mock_apwine compile_mock_tempus_token

# Real Tokens
compile_solidity_zct:
	@echo "compiling ZCT solidity source into abi and bin files"
	solc -o ./test/tokens --abi --bin --overwrite ./test/tokens/ZcToken.sol

compile_go_zct:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/tokens/ZcToken.abi --bin ./test/tokens/ZcToken.bin -pkg tokens -type ZcToken -out ./test/tokens/zctoken.go 

compile_zct: compile_solidity_zct compile_go_zct

compile_tokens: compile_zct

# Fakes

# Contracts for unit testing (with mocks)
compile_solidity_marketplace_test:
	@echo "compiling MarketPlace solidity source into abi and bin files"
	solc -o ./test/marketplace --optimize --optimize-runs=15000 --abi --bin --overwrite ./test/marketplace/MarketPlace.sol

compile_go_marketplace_test:
	@echo "compiling MarketPlace abi and bin files to golang"
	abigen --abi ./test/marketplace/MarketPlace.abi --bin ./test/marketplace/MarketPlace.bin -pkg MarketPlace -type MarketPlace -out ./test/marketplace/marketplace.go

compile_marketplace_test: compile_solidity_marketplace_test compile_go_marketplace_test

compile_solidity_lender_test:
	@echo "compiling Lender solidity source into abi and bin files"
	solc -o ./test/lender --optimize --optimize-runs=15000 --abi --bin --overwrite ./test/lender/Lender.sol

compile_go_lender_test:
	@echo "compiling Lender abi and bin files to golang"
	abigen --abi ./test/lender/Lender.abi --bin ./test/lender/Lender.bin -pkg lender -type Lender -out ./test/lender/lender.go

compile_lender_test: compile_solidity_lender_test compile_go_lender_test

compile_solidity_redeemer_test:
	@echo "compiling Redeemer solidity source into abi and bin files"
	solc -o ./test/redeemer --optimize --optimize-runs=15000 --abi --bin --overwrite ./test/redeemer/Redeemer.sol

compile_go_redeemer_test:
	@echo "compiling Redeemer abi and bin files to golang"
	abigen --abi ./test/redeemer/Redeemer.abi --bin ./test/redeemer/Redeemer.bin -pkg redeemer -type Redeemer -out ./test/redeemer/redeemer.go

compile_redeemer_test: compile_solidity_redeemer_test compile_go_redeemer_test
compile_test: compile_marketplace_test compile_lender_test compile_redeemer_test

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

clean_build_sol:
	@echo "removing sol files from build/ dirs"
	rm build/**/*.sol

clean_build_abi:
	@echo "removing abi files from build/ dirs"
	rm build/**/*.abi

clean_build_bin:
	@echo "removing bin files from build/ dirs"
	rm build/**/*.bin

clean_build_go:
	@echo "removing go files from build/ dirs"
	rm build/**/*.go

clean_build: clean_build_sol clean_build_abi clean_build_bin clean_build_go

# Copying to build, prepare for compiling build assets
copy_marketplace_to_build:
	@echo "copying MarketPlace files to build"
	cp test/marketplace/MarketPlace.sol build/marketplace
	cp test/marketplace/Interfaces.sol build/marketplace
	cp test/marketplace/Safe.sol build/marketplace
	cp test/tokens/Hash.sol build/marketplace
	cp test/tokens/PErc20.sol build/marketplace
	cp test/tokens/IPErc20.sol build/marketplace
	cp test/tokens/Erc2612.sol build/marketplace
	cp test/tokens/IErc2612.sol build/marketplace
	cp test/tokens/ZcToken.sol build/marketplace
	cp test/tokens/IZcToken.sol build/marketplace

copy_lender_to_build:
	@echo "copying Lender files to build"
	cp test/lender/Lender.sol build/lender
	cp test/lender/Interfaces.sol build/lender
	cp test/lender/MarketPlace.sol build/lender
	cp test/lender/Swivel.sol build/lender
	cp test/lender/Element.sol build/lender
	cp test/lender/Safe.sol build/lender
	cp test/lender/Cast.sol build/lender

copy_redeemer_to_build:
	@echo "copying Redeemer files to build"
	cp test/redeemer/Redeemer.sol build/redeemer
	cp test/redeemer/Interfaces.sol build/redeemer
	cp test/redeemer/MarketPlace.sol build/redeemer
	cp test/redeemer/Safe.sol build/redeemer

copy_to_build: copy_marketplace_to_build copy_lender_to_build copy_redeemer_to_build

# Contracts for production (without mocks)
compile_marketplace_build:
	solc -o ./build/marketplace --optimize --optimize-runs=15000 --abi --bin --overwrite ./build/marketplace/MarketPlace.sol
	abigen --abi ./build/marketplace/MarketPlace.abi --bin ./build/marketplace/MarketPlace.bin -pkg marketplace -type MarketPlace -out ./build/marketplace/marketplace.go

compile_lender_build:
	solc -o ./build/lender --optimize --optimize-runs=15000 --abi --bin --overwrite ./build/lender/Lender.sol
	abigen --abi ./build/lender/Lender.abi --bin ./build/lender/Lender.bin -pkg lender -type Lender -out ./build/lender/lender.go

compile_redeemer_build:
	solc -o ./build/redeemer --optimize --optimize-runs=15000 --abi --bin --overwrite ./build/redeemer/Redeemer.sol
	abigen --abi ./build/redeemer/Redeemer.abi --bin ./build/redeemer/Redeemer.bin -pkg redeemer -type Redeemer -out ./build/redeemer/redeemer.go

compile_build: compile_marketplace_build compile_lender_build compile_redeemer_build

all: clean_test compile_mocks compile_test clean_build copy_to_build compile_build
