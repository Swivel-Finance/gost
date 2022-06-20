#!/bin/bash

# Generate documentation for the project.
npx solidity-docgen --solc-module ~/node_modules/solc-0.8 -t ./docs/ -i test/marketplace -o docs/marketplace
npx solidity-docgen --solc-module ~/node_modules/solc-0.8 -t ./docs/ -i test/lender -o docs/lender
npx solidity-docgen --solc-module ~/node_modules/solc-0.8 -t ./docs/ -i test/redeemer -o docs/redeemer

echo "Lender loc"
cloc ./test/lender/*.sol --by-file > lender_loc.txt

echo "Redeemer loc"
cloc ./test/redeemer/*.sol --by-file > redeemer_loc.txt

echo "Marketplace loc"
cloc ./test/marketplace/*.sol --by-file > marketplace_loc.txt