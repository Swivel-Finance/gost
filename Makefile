.PHONY: compile_solidity_mock_erc compile_go_mock_erc compile_mock_erc
.PHONY: compile_solidity_mock_compound_token compile_go_mock_compound_token compile_mock_compound_token
.PHONY: compile_solidity_mock_erc_4626 compile_go_mock_erc_4626 compile_mock_erc_4626
.PHONY: compile_solidity_mock_yearn_vault compile_go_mock_yearn_vault compile_mock_yearn_vault
.PHONY: compile_solidity_mock_aave_token compile_go_mock_aave_token compile_mock_aave_token
.PHONY: compile_solidity_mock_aave_pool compile_go_mock_aave_pool compile_mock_aave_pool
.PHONY: compile_solidity_mock_euler_token compile_go_mock_euler_token compile_mock_euler_token
.PHONY: compile_solidity_mock_creator compile_go_mock_creator
.PHONY: compile_solidity_mock_marketplace compile_go_mock_marketplace
.PHONY: compile_mocks

.PHONY: compile_solidity_sig_fake compile_go_sig_fake compile_sig_fake
.PHONY: compile_solidity_hash_fake compile_go_hash_fake compile_hash_fake
.PHONY: compile_fakes

.PHONY: compile_solidity_swivel_test compile_go_swivel_test compile_swivel_test
.PHONY: compile_solidity_marketplace_test compile_go_marketplace_test compile_marketplace_test
.PHONY: compile_solidity_vaulttracker_test compile_go_vaulttracker_test compile_vaulttracker_test
.PHONY: compile_solidity_creator_test compile_go_creator_test compile_creator_test
.PHONY: compile_test

.PHONY: clean_test_abi clean_test_bin clean_test_go
.PHONY: clean_test

.PHONY: clean_build_sol clean_build_abi clean_build_bin clean_build_go

.PHONY: clean_build

.PHONY: copy_zctoken_to_build copy_vaulttracker_to_build copy_creator_to_build copy_marketplace_to_build copy_swivel_to_build
.PHONY: copy_to_build
.PHONY: compile_creator_build compile_marketplace_build compile_swivel_build compile_build

.PHONY: all

# ------------------------------------------------------- MOCKS -------------------------------------------------------------

# TODO should prob add _20 here...
compile_solidity_mock_erc:
	@echo "compiling Mock ERC20 solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/Erc20.sol

compile_go_mock_erc:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/Erc20.abi --bin ./test/mocks/Erc20.bin -pkg mocks -type Erc20 -out ./test/mocks/erc20.go 

compile_mock_erc: compile_solidity_mock_erc compile_go_mock_erc

# compound token
compile_solidity_mock_compound_token:
	@echo "compiling Mock Compound Token source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/CompoundToken.sol

compile_go_mock_compound_token:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/CompoundToken.abi --bin ./test/mocks/CompoundToken.bin -pkg mocks -type CompoundToken -out ./test/mocks/compoundtoken.go 

compile_mock_compound_token: compile_solidity_mock_compound_token compile_go_mock_compound_token

# erc4626 tokenized vault
compile_solidity_mock_erc_4626:
	@echo "compiling Mock Tokenized Vault source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/Erc4626.sol

compile_go_mock_erc_4626:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/Erc4626.abi --bin ./test/mocks/Erc4626.bin -pkg mocks -type Erc4626 -out ./test/mocks/erc4626.go 

compile_mock_erc_4626: compile_solidity_mock_erc_4626 compile_go_mock_erc_4626

# yearn vault
compile_solidity_mock_yearn_vault:
	@echo "compiling Mock Yearn Vault source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/YearnVault.sol

compile_go_mock_yearn_vault:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/YearnVault.abi --bin ./test/mocks/YearnVault.bin -pkg mocks -type YearnVault -out ./test/mocks/yearnvault.go 

compile_mock_yearn_vault: compile_solidity_mock_yearn_vault compile_go_mock_yearn_vault

# aave token
compile_solidity_mock_aave_token:
	@echo "compiling Mock Aave Token source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/AaveToken.sol

compile_go_mock_aave_token:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/AaveToken.abi --bin ./test/mocks/AaveToken.bin -pkg mocks -type AaveToken -out ./test/mocks/aavetoken.go 

compile_mock_aave_token: compile_solidity_mock_aave_token compile_go_mock_aave_token

# aave pool
compile_solidity_mock_aave_pool:
	@echo "compiling Mock Aave Pool source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/AavePool.sol

compile_go_mock_aave_pool:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/AavePool.abi --bin ./test/mocks/AavePool.bin -pkg mocks -type AavePool -out ./test/mocks/aavepool.go 

compile_mock_aave_pool: compile_solidity_mock_aave_pool compile_go_mock_aave_pool

# euler token
compile_solidity_mock_euler_token:
	@echo "compiling Mock Euler Token source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/EulerToken.sol

compile_go_mock_euler_token:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/EulerToken.abi --bin ./test/mocks/EulerToken.bin -pkg mocks -type EulerToken -out ./test/mocks/eulertoken.go 

compile_mock_euler_token: compile_solidity_mock_euler_token compile_go_mock_euler_token

compile_solidity_mock_creator:
	@echo "compiling Mock Creator solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/Creator.sol

# include go bindings for mock ZcToken and VaultTracker here because they are needed in testing
compile_go_mock_creator:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/Creator.abi --bin ./test/mocks/Creator.bin -pkg mocks -type Creator -out ./test/mocks/creator.go 
	abigen --abi ./test/creator/ZcToken.abi --bin ./test/creator/ZcToken.bin -pkg mocks -type ZcToken -out ./test/mocks/zctoken.go 
	abigen --abi ./test/creator/VaultTracker.abi --bin ./test/creator/VaultTracker.bin -pkg mocks -type VaultTracker -out ./test/mocks/vaulttracker.go 

compile_mock_creator: compile_solidity_mock_creator compile_go_mock_creator

compile_solidity_mock_marketplace:
	@echo "compiling Mock MarketPlace solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/MarketPlace.sol

compile_go_mock_marketplace:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/MarketPlace.abi --bin ./test/mocks/MarketPlace.bin -pkg mocks -type MarketPlace -out ./test/mocks/marketplace.go 

compile_mock_marketplace: compile_solidity_mock_marketplace compile_go_mock_marketplace

compile_mocks: compile_mock_erc compile_mock_compound_token compile_mock_erc_4626 compile_mock_yearn_vault compile_mock_aave_token compile_mock_aave_pool compile_mock_euler_token compile_mock_creator compile_mock_marketplace

# -------------------------------------------------- FAKES -----------------------------------------------------------------------------------------------

compile_solidity_sig_fake:
	@echo "compiling Sig and Fake source into abi and bin files"
	solc -o ./test/fakes --abi --bin --overwrite ./test/fakes/SigFake.sol

compile_go_sig_fake:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/fakes/SigFake.abi --bin ./test/fakes/SigFake.bin -pkg fakes -type SigFake -out ./test/fakes/sigfake.go 

compile_sig_fake: compile_solidity_sig_fake compile_go_sig_fake

compile_solidity_hash_fake:
	@echo "compiling Hash and Fake solidity source into abi and bin files"
	solc -o ./test/fakes --abi --bin --overwrite ./test/fakes/HashFake.sol

compile_go_hash_fake:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/fakes/HashFake.abi --bin ./test/fakes/HashFake.bin -pkg fakes -type HashFake -out ./test/fakes/hashfake.go 

compile_hash_fake: compile_solidity_hash_fake compile_go_hash_fake

compile_fakes: compile_sig_fake compile_hash_fake

# ----------------------------------------------- CONTRACTS (TEST) ------------------------------------------------------------------------------------------

compile_solidity_swivel_test:
	@echo "compiling Swivel solidity source into abi and bin files"
	solc -o ./test/swivel --optimize --optimize-runs=10000 --abi --bin --overwrite ./test/swivel/Swivel.sol

compile_go_swivel_test:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/swivel/Swivel.abi --bin ./test/swivel/Swivel.bin -pkg swivel -type Swivel -out ./test/swivel/swivel.go 

compile_swivel_test: compile_solidity_swivel_test compile_go_swivel_test

compile_solidity_marketplace_test:
	@echo "compiling MarketPlace solidity source into abi and bin files"
	solc -o ./test/marketplace --optimize --optimize-runs=10000 --abi --bin --overwrite ./test/marketplace/MarketPlace.sol

compile_go_marketplace_test:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/marketplace/MarketPlace.abi --bin ./test/marketplace/MarketPlace.bin -pkg marketplace -type MarketPlace -out ./test/marketplace/marketplace.go 

compile_marketplace_test: compile_solidity_marketplace_test compile_go_marketplace_test

compile_solidity_creator_test:
	@echo "compiling Creator solidity source into abi and bin files"
	solc -o ./test/creator --optimize --optimize-runs=10000 --abi --bin --overwrite ./test/creator/Creator.sol

compile_go_creator_test:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/creator/Creator.abi --bin ./test/creator/Creator.bin -pkg creator -type Creator -out ./test/creator/creator.go 

compile_creator_test: compile_solidity_creator_test compile_go_creator_test

compile_solidity_vaulttracker_test:
	@echo "compiling VaultTracker solidity source into abi and bin files"
	solc -o ./test/vaulttracker --abi --bin --overwrite ./test/vaulttracker/VaultTracker.sol

compile_go_vaulttracker_test:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/vaulttracker/VaultTracker.abi --bin ./test/vaulttracker/VaultTracker.bin -pkg vaulttracker -type VaultTracker -out ./test/vaulttracker/vaulttracker.go 

compile_vaulttracker_test: compile_solidity_vaulttracker_test compile_go_vaulttracker_test

compile_test: compile_fakes compile_swivel_test compile_marketplace_test compile_creator_test compile_vaulttracker_test compile_mocks

# --------------------------------------------- CLEANING ----------------------------------------------------------------------------------------

# test
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

# build
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

# ---------------------------------- COPY TO BUILD ----------------------------------------------------------------------------------------------

copy_creator_to_build:
	@echo "copying Creator files to creator build"
	cp test/creator/Creator.sol build/creator

copy_zctoken_to_build:
	@echo "copying ZcToken files to marketplace build"
	cp test/tokens/Erc20.sol build/creator
	cp test/tokens/IERC5095.sol build/creator
	cp test/tokens/IRedeemer.sol build/creator
	cp test/tokens/ZcToken.sol build/creator

copy_vaulttracker_to_build:
	@echo "copying vaulttracker files to marketplace build"
	cp test/libraries/ICERC20.sol build/creator
	cp test/libraries/FixedPointMathLib.sol build/creator
	cp test/libraries/LibCompound.sol build/creator
	cp test/libraries/LibFuse.sol build/creator
	cp test/libraries/Compounding.sol build/creator
	cp test/vaulttracker/Protocols.sol build/creator
	cp test/vaulttracker/VaultTracker.sol build/creator

copy_marketplace_to_build:
	@echo "copying marketplace files to marketplace build"
	cp test/libraries/ICERC20.sol build/marketplace
	cp test/libraries/FixedPointMathLib.sol build/marketplace
	cp test/libraries/LibCompound.sol build/marketplace
	cp test/libraries/LibFuse.sol build/marketplace
	cp test/libraries/Compounding.sol build/marketplace
	cp test/marketplace/Interfaces.sol build/marketplace
	cp test/marketplace/Protocols.sol build/marketplace
	cp test/marketplace/MarketPlace.sol build/marketplace

copy_swivel_to_build:
	@echo "copying swivel files to marketplace build"
	cp test/swivel/Interfaces.sol build/swivel
	cp test/swivel/Protocols.sol build/swivel
	cp test/swivel/Hash.sol build/swivel
	cp test/swivel/Sig.sol build/swivel
	cp test/swivel/Safe.sol build/swivel
	cp test/swivel/Swivel.sol build/swivel

copy_to_build: copy_zctoken_to_build copy_vaulttracker_to_build copy_creator_to_build copy_marketplace_to_build copy_swivel_to_build

# ---------------------------------- COMPILE BUILD ------------------------------------------------------------------------------------------

compile_creator_build:
	@echo "compiling Creator solidity build source into deploy ready files"
	solc -o ./build/creator --optimize --optimize-runs=2000 --abi --bin --overwrite ./build/creator/Creator.sol
	abigen --abi ./build/creator/Creator.abi --bin ./build/creator/Creator.bin -pkg creator -type Creator -out ./build/creator/creator.go 

compile_marketplace_build:
	@echo "compiling MarketPlace solidity build source into deploy ready files"
	solc -o ./build/marketplace --optimize --optimize-runs=10000 --abi --bin --overwrite ./build/marketplace/MarketPlace.sol
	abigen --abi ./build/marketplace/MarketPlace.abi --bin ./build/marketplace/MarketPlace.bin -pkg marketplace -type MarketPlace -out ./build/marketplace/marketplace.go 

compile_swivel_build:
	@echo "compiling Swivel solidity build source into deploy ready files"
	solc -o ./build/swivel --optimize --optimize-runs=10000 --abi --bin --overwrite ./build/swivel/Swivel.sol
	abigen --abi ./build/swivel/Swivel.abi --bin ./build/swivel/Swivel.bin -pkg swivel -type Swivel -out ./build/swivel/swivel.go 

compile_build: compile_creator_build compile_marketplace_build compile_swivel_build

# ------------------------------ ALL -------------------------------------------------------------------------------------------------

all: clean_test compile_test clean_build copy_to_build compile_build
