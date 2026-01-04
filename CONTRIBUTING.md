# Contributing to BSpec

Thank you for your interest in contributing to BSpec! This document provides guidelines and information for contributors.

## Ways to Contribute

### Documentation
- Improve existing guides and explanations
- Add examples and use cases
- Fix typos and clarify language
- Translate documentation

### Code
- Fix bugs in SDKs or CLI
- Add new features
- Improve test coverage
- Optimize performance

### Specification
- Propose new document types
- Suggest improvements to existing types
- Report inconsistencies or gaps

## Getting Started

1. **Fork the repository** on GitHub
2. **Clone your fork** locally
3. **Create a branch** for your changes
4. **Make your changes** following our guidelines
5. **Test your changes** thoroughly
6. **Submit a pull request**

## Development Setup

```bash
# Clone the repository
git clone https://github.com/a3tai/bspec.git
cd bspec

# Install dependencies for the web documentation
cd apps/web
bun install

# Install dependencies for the MCP server
cd ../mcp
bun install

# Build the CLI
cd ../../sdk/cli
make build
```

## Code Style

- **TypeScript/JavaScript**: Follow the existing ESLint configuration
- **Python**: Follow PEP 8 guidelines
- **Go**: Use `gofmt` for formatting
- **Markdown**: Use consistent heading levels and formatting

## Commit Messages

Use clear, descriptive commit messages:

```
type(scope): brief description

- More details if needed
- Reference issues with #123
```

Types: `feat`, `fix`, `docs`, `style`, `refactor`, `test`, `chore`

## Pull Request Process

1. Update documentation for any changed functionality
2. Add tests for new features
3. Ensure all tests pass
4. Request review from maintainers
5. Address feedback promptly

## Reporting Issues

When reporting issues, please include:

- Clear description of the problem
- Steps to reproduce
- Expected vs actual behavior
- Environment details (OS, versions)
- Relevant logs or screenshots

## Questions?

- Open a [GitHub Discussion](https://github.com/a3tai/bspec/discussions) for questions
- Check existing issues before creating new ones
- Join the community conversation

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

