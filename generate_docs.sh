#!/bin/bash

npx solidity-docgen --solc-module ~/node_modules/solc-0.8 -t ./docs/ -i test/marketplace -o docs/marketplace
npx solidity-docgen --solc-module ~/node_modules/solc-0.8 -t ./docs/ -i test/lender -o docs/lender
npx solidity-docgen --solc-module ~/node_modules/solc-0.8 -t ./docs/ -i test/redeemer -o docs/redeemer