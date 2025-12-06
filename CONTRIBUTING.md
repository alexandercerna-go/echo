# Contributing to Echo

Thank you for your interest in contributing to the Echo project! This is a learning project designed to demonstrate various Go programming patterns.

## Getting Started

1. Fork the repository
2. Clone your fork: `git clone https://github.com/your-username/echo.git`
3. Create a feature branch: `git checkout -b feature/your-feature`
4. Make your changes and commit them

## Development Setup

```bash
# Install dependencies
go mod download

# Build the project
make build

# Run tests
make test

# Format code
make fmt

# Run linter
make lint
```

## Code Standards

- Follow Go's idiomatic style guidelines
- Run `make fmt` before committing
- Add comments to exported functions
- Include tests for new functionality
- Keep commits atomic and well-documented

## Making Changes

1. **Code Quality**: Ensure your code passes `make lint` and `make fmt`
2. **Testing**: Add tests if you add new features
3. **Documentation**: Update comments and README if needed
4. **Commit Messages**: Use clear, descriptive commit messages

### Example Commit Message
```
feat: add new echo method with performance analysis
docs: update README with new features
refactor: simplify timing logic
```

## Pull Request Process

1. Update the README.md with any new features or changes
2. Ensure all tests pass: `make test`
3. Ensure code is formatted: `make fmt`
4. Ensure linter passes: `make lint`
5. Submit your PR with a clear description of changes

## Questions or Suggestions?

Feel free to open an issue for questions, bug reports, or feature suggestions.

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
