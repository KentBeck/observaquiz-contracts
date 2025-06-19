#!/bin/bash

# ObservaQuiz Development Setup Script
# Sets up the complete development environment for all repositories

set -e

echo "🚀 Setting up ObservaQuiz development environment..."

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${GREEN}✓${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}⚠${NC} $1"
}

print_error() {
    echo -e "${RED}✗${NC} $1"
}

# Check if we're in the contracts repo
if [ ! -f "package.json" ] || ! grep -q "observaquiz-contracts" package.json; then
    print_error "This script must be run from the observaquiz-contracts repository root"
    exit 1
fi

# Install contracts dependencies
print_status "Installing contracts dependencies..."
npm install

# Validate OpenAPI specification
print_status "Validating API contracts..."
npm run validate:openapi

# Check if frontend repo exists
if [ -d "../observaquiz-ui" ]; then
    print_status "Setting up frontend (observaquiz-ui)..."
    cd ../observaquiz-ui
    
    if [ -f "package.json" ]; then
        npm install
        print_status "Frontend dependencies installed"
    else
        print_warning "Frontend package.json not found"
    fi
    
    cd ../observaquiz-contracts
else
    print_warning "Frontend repository (observaquiz-ui) not found in parent directory"
    echo "  Clone it with: git clone https://github.com/honeycombio/observaquiz-ui.git ../observaquiz-ui"
fi

# Check if backend repo exists
if [ -d "../observaquiz-api" ]; then
    print_status "Setting up backend (observaquiz-api)..."
    cd ../observaquiz-api
    
    if [ -f "go.mod" ]; then
        go mod download
        print_status "Backend dependencies downloaded"
    else
        print_warning "Backend go.mod not found"
    fi
    
    cd ../observaquiz-contracts
else
    print_warning "Backend repository (observaquiz-api) not found in parent directory"
    echo "  Clone it with: git clone https://github.com/honeycombio/observaquiz-api.git ../observaquiz-api"
fi

# Generate types from contracts
print_status "Generating types from contracts..."
if command -v openapi-typescript &> /dev/null; then
    npm run generate:ts
    print_status "TypeScript types generated"
else
    print_warning "openapi-typescript not found, skipping TypeScript generation"
    echo "  Install with: npm install -g openapi-typescript"
fi

# Check for required tools
echo ""
echo "🔧 Checking development tools..."

# Node.js
if command -v node &> /dev/null; then
    NODE_VERSION=$(node --version)
    print_status "Node.js $NODE_VERSION"
else
    print_error "Node.js not found - please install Node.js 18+"
fi

# Go
if command -v go &> /dev/null; then
    GO_VERSION=$(go version | cut -d' ' -f3)
    print_status "Go $GO_VERSION"
else
    print_error "Go not found - please install Go 1.21+"
fi

# AWS SAM CLI
if command -v sam &> /dev/null; then
    SAM_VERSION=$(sam --version | cut -d' ' -f4)
    print_status "AWS SAM CLI $SAM_VERSION"
else
    print_warning "AWS SAM CLI not found - needed for backend development"
    echo "  Install from: https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/install-sam-cli.html"
fi

echo ""
echo "🎉 Development environment setup complete!"
echo ""
echo "Next steps:"
echo "1. Review plan.md for development priorities"
echo "2. Set up environment variables (see README files in each repo)"
echo "3. Run contract tests: npm test"
echo "4. Start development servers in each repository"
echo ""
echo "Happy coding! 🚀"
