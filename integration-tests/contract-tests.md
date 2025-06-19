# Contract Tests

This directory contains integration tests that validate both frontend and backend implementations comply with the defined contracts.

## Test Categories

### API Contract Tests
- Validate API endpoints match OpenAPI specification
- Test request/response schemas
- Verify error handling contracts
- Check authentication requirements

### Type Contract Tests  
- Validate TypeScript types match Go types
- Test serialization/deserialization
- Verify enum values are consistent
- Check optional field handling

### Telemetry Contract Tests
- Validate telemetry events are emitted correctly
- Test event schema compliance
- Verify OpenTelemetry integration
- Check data flow end-to-end

## Running Contract Tests

```bash
# Install dependencies
npm install

# Run all contract tests
npm test

# Run specific test suite
npm test -- --grep "API Contract"
npm test -- --grep "Type Contract" 
npm test -- --grep "Telemetry Contract"

# Run tests against local services
npm run test:local

# Run tests against staging
npm run test:staging
```

## Test Structure

```
integration-tests/
├── api/                 # API contract tests
│   ├── health.test.js
│   ├── quiz.test.js
│   └── auth.test.js
├── types/               # Type contract tests
│   ├── serialization.test.js
│   └── validation.test.js
├── telemetry/           # Telemetry contract tests
│   ├── events.test.js
│   └── tracing.test.js
├── fixtures/            # Test data
│   ├── questions.json
│   └── sessions.json
└── utils/               # Test utilities
    ├── api-client.js
    └── mock-server.js
```

## Contract Validation Rules

1. **API Contracts**: All endpoints must match OpenAPI spec exactly
2. **Type Contracts**: TypeScript and Go types must be equivalent
3. **Telemetry Contracts**: All events must follow defined schema
4. **Error Contracts**: Error responses must be consistent across services

## Adding New Contract Tests

1. Define the contract in the appropriate specification file
2. Add test cases for both success and failure scenarios
3. Include edge cases and boundary conditions
4. Verify both frontend and backend implementations
5. Update documentation with new contract requirements
