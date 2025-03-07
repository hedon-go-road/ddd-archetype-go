#!/bin/bash

set -e

# pre-commit
if ! command -v pre-commit &> /dev/null
then
    echo "pre-commit not found, installing..."
    brew install pre-commit
else
    echo "pre-commit is already installed"
fi

# typos
if ! command -v typos &> /dev/null
then
    echo "Installing typos..."
    brew install typos-cli
else
    echo "typos is already installed"
fi

# swag
if ! command -v swag &> /dev/null
then
    echo "Installing swag(v1.16.3)..."
    go install github.com/swaggo/swag/cmd/swag@v1.16.3
else
    echo "swag(v1.16.3) is already installed"
fi

# pre-commit hooks
echo "Installing pre-commit hooks..."
pre-commit install

echo "All done! pre-commit hooks are installed and configured."
