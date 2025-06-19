# ObservaQuiz Development Plan

## Project Overview

ObservaQuiz is a full-stack observability quiz application with:

- **Contracts**: API specifications and coordination (this repository)
- **Frontend**: React/TypeScript app with OpenTelemetry instrumentation (`observaquiz-ui`)
- **Backend**: Go AWS Lambda API with OpenTelemetry (`observaquiz-api`)
- **Purpose**: Educational quiz about observability while demonstrating observability practices

## Architecture Principles

- **Dependency Inversion**: Frontend and backend depend on contracts, not each other
- **Contract-First Development**: Define APIs before implementation
- **Independent Deployment**: Each service can be developed and deployed separately
- **Test-Driven Development**: Red → Green → Refactor cycle with contract validation
- **Parallel Development**: Frontend and backend teams can work independently using contracts

## Current State Assessment

### Frontend (observaquiz-ui)

- ✅ React/TypeScript with esbuild
- ✅ OpenTelemetry instrumentation implemented
- ✅ Jest testing framework configured
- ✅ Basic test structure exists
- ⚠️ Limited unit test coverage
- ⚠️ No integration tests
- ⚠️ Dependencies may need updates

### Backend (observaquiz-api)

- ✅ Go 1.21.5 with AWS Lambda
- ✅ OpenTelemetry instrumentation
- ✅ AWS SAM for local development
- ❌ No visible unit tests
- ❌ No integration tests
- ⚠️ Manual testing via HTTP files only

## Phase 1: Contract Implementation (Dependency Inversion)

**Goal**: Implement dependency inversion by making both repositories depend on contracts

### [ ] 1.1 Contract Validation
- [x] Define API contracts using OpenAPI specification
- [x] Create shared TypeScript/Go type definitions
- [x] Set up contract testing framework
- [x] Define telemetry data contracts
- [x] Create integration test specifications

### [ ] 1.2 Frontend Contract Integration (Frontend Agent)
**Assignable to: Frontend Agent**
- [ ] Add observaquiz-contracts as dependency to package.json
- [ ] Replace existing ad-hoc types with contract types
- [ ] Update QuestionSetRetrieval.tsx to use contract Question type
- [ ] Update respondToAnswer.ts to use contract AnswerResponse type
- [ ] Migrate fetchFromBackend calls to use contract API endpoints
- [ ] Update all component props to use contract interfaces
- [ ] Add contract validation in API client layer
- [ ] Create fake implementations of contract APIs for development
- [ ] Update tests to use contract types and fake APIs

### [ ] 1.3 Backend Contract Integration (Backend Agent)
**Assignable to: Backend Agent**
- [ ] Add observaquiz-contracts as Go module dependency
- [ ] Replace existing Question struct with contract types
- [ ] Update get_questions.go to implement contract API
- [ ] Migrate answer submission handlers to contract format
- [ ] Update all HTTP handlers to match OpenAPI specification
- [ ] Add contract validation middleware
- [ ] Ensure JSON serialization matches contract schemas
- [ ] Implement contract compliance tests
- [ ] Add API endpoint validation against OpenAPI spec

### [ ] 1.4 Contract Compliance Testing
- [ ] Run contract tests against frontend implementation
- [ ] Run contract tests against backend implementation
- [ ] Fix any contract violations found
- [ ] Validate API request/response schemas match exactly
- [ ] Test error response formats comply with contracts
- [ ] Verify telemetry events match contract specifications

## Phase 2: Environment Setup & Validation (Foundation)

**Goal**: Ensure both projects can run locally and are ready for TDD

### [ ] 2.1 Frontend Environment Setup (Frontend Agent)
**Assignable to: Frontend Agent**
- [ ] Install Node.js dependencies
- [ ] Verify build process works
- [ ] Run existing tests
- [ ] Start local development server
- [ ] Validate OpenTelemetry integration
- [ ] Set up contract fake server for independent development
- [ ] Configure development environment to use contract fakes

### [ ] 2.2 Backend Environment Setup (Backend Agent)
**Assignable to: Backend Agent**
- [ ] Verify Go environment (Go 1.21.5+)
- [ ] Install AWS SAM CLI
- [ ] Set up environment.json with required keys
- [ ] Build and run locally
- [ ] Test API endpoints manually
- [ ] Set up contract validation in local development
- [ ] Configure OpenAPI spec validation middleware

### [ ] 2.3 Integration Validation

- [ ] Run both frontend and backend together
- [ ] Verify end-to-end communication follows contracts
- [ ] Confirm telemetry data flows correctly
- [ ] Document any setup issues found

## Phase 3: Testing Infrastructure (TDD Foundation)

**Goal**: Establish robust testing capabilities for TDD workflow

### [ ] 3.1 Frontend Testing Enhancement (Frontend Agent)
**Assignable to: Frontend Agent**
- [ ] Audit existing Jest configuration
- [ ] Add React Testing Library if missing
- [ ] Create test utilities for OpenTelemetry faking
- [ ] Set up component testing patterns
- [ ] Add test coverage reporting
- [ ] Create contract fake utilities for testing
- [ ] Add contract compliance tests for API client
- [ ] Set up automated testing against contract fakes

### [ ] 3.2 Backend Testing Infrastructure (Backend Agent)
**Assignable to: Backend Agent**
- [ ] Set up Go testing framework
- [ ] Create test utilities for Lambda handlers
- [ ] Add HTTP testing helpers
- [ ] Set up OpenTelemetry test fakes
- [ ] Configure test database/storage if needed
- [ ] Add contract compliance testing framework
- [ ] Create API endpoint validation tests
- [ ] Set up automated contract validation in CI

### [ ] 3.3 Integration Testing Setup (Coordination)
**Requires: Both agents to complete their contract integration**
- [ ] Set up end-to-end testing framework
- [ ] Create test data fixtures
- [ ] Add API contract testing
- [ ] Set up test environment automation
- [ ] Validate frontend and backend work together via contracts

## Phase 4: Code Quality & Standards (Clean Foundation)

**Goal**: Establish code quality standards and tooling

### [ ] 4.1 Frontend Code Quality (Frontend Agent)
**Assignable to: Frontend Agent**
- [ ] Add/configure ESLint with TypeScript rules
- [ ] Add Prettier for code formatting
- [ ] Set up pre-commit hooks
- [ ] Add TypeScript strict mode if not enabled
- [ ] Audit and update dependencies
- [ ] Add contract type validation in linting
- [ ] Set up automated code quality checks

### [ ] 4.2 Backend Code Quality (Backend Agent)
**Assignable to: Backend Agent**
- [ ] Add golangci-lint configuration
- [ ] Set up gofmt/goimports automation
- [ ] Add pre-commit hooks for Go
- [ ] Review and update Go dependencies
- [ ] Add security scanning (gosec)
- [ ] Add contract validation in code quality checks
- [ ] Set up automated Go code quality pipeline

### [ ] 4.3 Documentation Standards (Coordination)
**Requires: Input from both agents**
- [ ] Create/update API documentation
- [ ] Document development workflow
- [ ] Add code commenting standards
- [ ] Create troubleshooting guides
- [ ] Document contract-first development process

## Phase 5: CI/CD Pipeline (Automation)

**Goal**: Automate testing and deployment processes

### [ ] 5.1 Frontend CI/CD (Frontend Agent)
**Assignable to: Frontend Agent**
- [ ] Set up GitHub Actions workflows for frontend
- [ ] Add automated testing on PR
- [ ] Add code quality checks
- [ ] Set up dependency vulnerability scanning
- [ ] Add build artifact generation
- [ ] Add contract validation in CI pipeline
- [ ] Set up automated deployment to staging/prod

### [ ] 5.2 Backend CI/CD (Backend Agent)
**Assignable to: Backend Agent**
- [ ] Set up GitHub Actions workflows for backend
- [ ] Add automated testing on PR
- [ ] Add code quality checks
- [ ] Set up dependency vulnerability scanning
- [ ] Add build artifact generation
- [ ] Add contract compliance validation in CI
- [ ] Validate existing Pulumi deployment
- [ ] Set up automated deployments with rollback capabilities

### [ ] 5.3 Integration CI/CD (Coordination)
**Requires: Both agents to complete their CI/CD setup**
- [ ] Set up cross-repository integration testing
- [ ] Add contract validation across services
- [ ] Set up monitoring/alerting
- [ ] Create deployment coordination workflows

## Phase 6: Feature Development Readiness (TDD Ready)

**Goal**: Prepare for efficient feature development using TDD

### [ ] 6.1 Frontend Development Workflow (Frontend Agent)
**Assignable to: Frontend Agent**
- [ ] Create frontend feature branch strategy
- [ ] Set up TDD workflow documentation for React components
- [ ] Create frontend-specific issue/PR templates
- [ ] Set up frontend development environment scripts
- [ ] Add frontend debugging guides
- [ ] Document contract-driven frontend development patterns

### [ ] 6.2 Backend Development Workflow (Backend Agent)
**Assignable to: Backend Agent**
- [ ] Create backend feature branch strategy
- [ ] Set up TDD workflow documentation for Go/Lambda
- [ ] Create backend-specific issue/PR templates
- [ ] Set up backend development environment scripts
- [ ] Add backend debugging guides
- [ ] Document contract-driven backend development patterns

### [ ] 6.3 Architecture Documentation (Coordination)
**Requires: Input from both agents**
- [ ] Document current architecture
- [ ] Identify extension points for new features
- [ ] Create development patterns guide
- [ ] Document OpenTelemetry integration patterns
- [ ] Add performance considerations
- [ ] Document contract-first development workflow
- [ ] Create cross-team collaboration guidelines

## Success Criteria

- [ ] Both frontend and backend depend on and comply with contracts
- [ ] Frontend and backend can be developed independently by separate agents
- [ ] Both services run locally without issues
- [ ] Comprehensive test suites with >80% coverage
- [ ] All code quality tools passing
- [ ] CI/CD pipeline fully functional for both services
- [ ] Clear documentation for new developers
- [ ] TDD workflow established and documented
- [ ] Contract validation passes in all environments
- [ ] Cross-team collaboration workflows documented

## Next Steps After Plan Completion

Once this plan is complete, the project will be ready for:

- Feature development using TDD methodology
- Adding new quiz questions and formats
- Enhancing observability features
- Performance optimizations
- User experience improvements

## Agent Coordination Guidelines

### Frontend Agent Responsibilities
- All tasks marked "Frontend Agent" or "Assignable to: Frontend Agent"
- Focus on React/TypeScript implementation
- Develop against contract fakes for independence
- Ensure contract compliance in all frontend code
- Coordinate with Backend Agent through contract validation

### Backend Agent Responsibilities
- All tasks marked "Backend Agent" or "Assignable to: Backend Agent"
- Focus on Go/AWS Lambda implementation
- Implement contract specifications exactly
- Ensure API endpoints match OpenAPI specification
- Coordinate with Frontend Agent through contract validation

### Coordination Points
- **Contract Changes**: Must be agreed upon by both agents before implementation
- **Integration Testing**: Requires both agents to complete contract integration
- **Documentation**: Shared responsibility for architecture and workflow docs
- **CI/CD Integration**: Both agents must coordinate deployment pipelines

### Communication Protocol
1. **Contract First**: All API changes start with contract updates
2. **Independent Development**: Agents work in parallel using contracts and fakes
3. **Validation Gates**: Contract tests must pass before integration
4. **Coordination Reviews**: Regular sync on contract compliance and integration

## Notes

- Follow Kent Beck's TDD principles: Red → Green → Refactor
- Separate structural changes from behavioral changes
- Commit frequently with clear messages
- Prioritize working software over comprehensive documentation
- Embrace simplicity and avoid over-engineering
- **Contract-Driven Development**: Always implement contracts before features
- **Agent Independence**: Each agent should be able to work without blocking the other
