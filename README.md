# ObservaQuiz Contracts

This repository contains the shared contracts, interfaces, and coordination documentation for the ObservaQuiz project.

## Repository Structure

```
observaquiz-contracts/     # This repo - coordination & contracts
├── plan.md               # Master development plan
├── api-contracts/        # API specifications
├── types/               # Shared type definitions
├── integration-tests/   # End-to-end contract validation
├── docs/               # Architecture & coordination docs
└── scripts/            # Cross-repo development tools
```

## Related Repositories

- **Frontend**: [observaquiz-ui](https://github.com/honeycombio/observaquiz-ui) - React/TypeScript implementation
- **Backend**: [observaquiz-api](https://github.com/honeycombio/observaquiz-api) - Go AWS Lambda implementation

## Dependency Inversion Principle

This repository implements dependency inversion for the ObservaQuiz project:

- **Frontend** depends on contracts defined here (not on backend implementation)
- **Backend** implements contracts defined here (not coupled to frontend)
- **Integration tests** validate that implementations satisfy contracts

## Development Workflow

1. **Contract-First Development**: Define APIs and types here first
2. **Independent Implementation**: Frontend and backend implement against stable contracts
3. **Contract Testing**: Validate implementations satisfy contracts
4. **Coordination**: Use `plan.md` to coordinate work across repositories

## Getting Started

1. Clone this repository
2. Review `plan.md` for current development priorities
3. Check `api-contracts/` for latest API specifications
4. Run integration tests to validate current state

## Contract-Driven TDD

1. **Red**: Write failing contract test
2. **Green**: Implement in frontend/backend to satisfy contract
3. **Refactor**: Improve implementation while maintaining contract compliance

This ensures both sides develop against stable, well-defined interfaces.
