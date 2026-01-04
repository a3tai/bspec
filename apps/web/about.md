# About BSpec

BSpec is an open-source standard for describing any business as a structured, machine-readable knowledge graph.

## The Problem

Business documentation is fragmented, inconsistent, and difficult for both humans and AI systems to work with effectively. Every company documents differently, making it hard to:

- Compare businesses (for investors, acquirers, partners)
- Ensure complete documentation (are we covering everything?)
- Enable AI systems to understand and reason about businesses
- Transfer knowledge when people leave
- Maintain consistency as companies grow

## The Solution

BSpec provides a universal language - 112 standardized document types organized across 12 comprehensive business domains. Similar to how SAFE agreements standardized early-stage investing, BSpec standardizes business documentation.

### Key Innovation: Atomic + Connected

1. **Atomic Documents**: Each document covers exactly one business concern
2. **Rich Relationships**: Documents declare dependencies, enablements, and conflicts
3. **Knowledge Graph**: Together, documents form an interconnected model of your business
4. **Machine-Readable**: YAML frontmatter enables tools and AI to parse and analyze

## Who Is It For?

### Founders & Business Leaders
Document your business in a structured way that enables clear thinking, better communication, and AI-powered analysis.

### Investors & Advisors
Evaluate businesses with consistent, standardized documentation. Automate due diligence and portfolio analysis.

### AI Developers
Build tools that work with any BSpec-compliant business. The standard format provides rich context for AI reasoning.

### Consultants & Coaches
Help clients document their business systematically. Use conformance levels to guide the documentation journey.

## Project Status

**Current Version**: 1.0.0

BSpec v1.0 is stable and ready for use. The specification includes:
- 112 document types across 12 domains
- Complete relationship model
- Conformance levels (Bronze, Silver, Gold)
- Industry-specific profiles
- SDKs for TypeScript, Python, and Go
- CLI tool
- MCP server for AI integration

## Open Source

BSpec is released under the MIT License. The project includes:

- **Specification**: Complete documentation of the standard
- **SDKs**: Type-safe libraries for TypeScript, Python, and Go
- **CLI Tool**: Command-line tool for validation and analysis
- **MCP Server**: Integration with AI assistants like Claude
- **Documentation**: This website and comprehensive guides

### Repository Structure

```
bspec/
├── spec/v1/              # Complete specification
├── sdk/                  # Generated SDKs
│   ├── typescript/
│   ├── python/
│   └── go/
├── sdk/cli/              # CLI tool (Go)
├── apps/
│   ├── web/              # This documentation site
│   └── mcp/              # MCP server
└── scripts/              # Code generation tools
```

## Contributing

We welcome contributions! Here's how you can help:

- **Documentation**: Improve guides, add examples, fix typos
- **Tools**: Build new tools that work with BSpec
- **SDKs**: Add support for new languages
- **Feedback**: Share your experience using BSpec

See [CONTRIBUTING.md](https://github.com/a3tai/bspec/blob/main/CONTRIBUTING.md) for guidelines.

## Community

- **GitHub**: [a3tai/bspec](https://github.com/a3tai/bspec)
- **Discussions**: [GitHub Discussions](https://github.com/a3tai/bspec/discussions)
- **Issues**: [GitHub Issues](https://github.com/a3tai/bspec/issues)

## License

MIT License - see [LICENSE](https://github.com/a3tai/bspec/blob/main/LICENSE) for details.

## Acknowledgments

BSpec was inspired by:
- SAFE agreements (standardization of early-stage investing)
- OpenAPI/Swagger (standardization of API documentation)
- Model Context Protocol (standardization of AI-system integration)

The project stands on the shoulders of giants in open-source standards and documentation.
