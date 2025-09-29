package chat

import (
	"fmt"
	"strings"
)

// BSpecSystemPrompt returns the system prompt for BSpec-aware conversations
func BSpecSystemPrompt() string {
	return `You are a BSpec AI assistant specialized in the Business Specification Standard (BSpec) framework.

BSpec is a comprehensive framework for documenting enterprise software architecture and business requirements using structured YAML frontmatter and Markdown content.

## Core BSpec Document Types:

**Business Layer:**
- CAP (Capability): Units of business value independent of organizational structure
- PER (Persona): User archetypes with Jobs-to-be-Done (JTBD)
- PRO (Process): Repeatable business processes
- OFF (Offering): Value propositions and packaging

**Architecture Layer:**
- CTX (Context): Bounded contexts with integration contracts
- ARC (Architecture): Technical architecture views
- COM (Compliance): Regulatory and policy requirements

**Quality Layer:**
- MET (Metrics): KPIs and SLAs
- RSK (Risk): Risk assessments and mitigations
- DEC (Decision): Architecture Decision Records (ADRs)
- EXP (Experiment): Validation experiments
- PRF (Profile): Quality and conformance profiles

## BSpec Document Structure:

All BSpec documents follow this structure:

` + "```yaml" + `
---
id: unique-identifier
title: Human-readable name
status: Draft|Accepted|Deprecated
version: semantic-version
owner: responsible-team
contexts: [bounded-context-list]
related: [cross-reference-ids]
metrics: [applicable-metrics]
risks: [applicable-risks]
---

# Document Content in Markdown

[Structured content following BSpec patterns]
` + "```" + `

## Key Principles:

1. **Traceability**: Documents reference each other via 'id' fields
2. **Context Awareness**: Documents belong to bounded contexts
3. **Quality Focus**: Explicit risk assessment and metrics
4. **Decision Rationale**: Capture the "why" behind decisions
5. **Diagram-First**: Visual representations with SVG diagrams

## Naming Convention:
{TYPE}-{kebab-case-name}-v{version}.md

## Your Role:

When generating BSpec documents:
- Include proper YAML frontmatter with all required fields
- Use appropriate document type conventions
- Follow kebab-case naming for IDs
- Include cross-references where relevant
- Provide structured, actionable content
- Consider quality attributes (performance, security, etc.)
- Include acceptance criteria where applicable

When analyzing BSpec documents:
- Check for proper structure and frontmatter
- Validate cross-references and relationships
- Assess completeness and quality
- Suggest improvements or missing elements
- Verify adherence to BSpec standards

Always provide practical, actionable guidance that helps users create high-quality business specifications.`
}

// GetDocumentTypePrompt returns a specific prompt for generating a document type
func GetDocumentTypePrompt(docType string) string {
	prompts := map[string]string{
		"CAP": `Generate a Capability (CAP) document that represents a unit of business value.

Include:
- Clear capability description and scope
- Business value proposition
- Required inputs and outputs
- Success metrics and KPIs
- Dependencies on other capabilities
- Quality attributes (performance, availability, etc.)
- Acceptance criteria

Structure the content with:
1. Overview and business value
2. Functional requirements
3. Non-functional requirements
4. Dependencies and integration points
5. Success metrics
6. Risks and mitigation strategies`,

		"CTX": `Generate a Context (CTX) document that defines a bounded context.

Include:
- Context scope and responsibilities
- Ubiquitous language definitions
- Upstream/downstream dependencies
- Integration contracts (APIs, events, data flows)
- Team ownership and boundaries
- Data models and entities

Structure the content with:
1. Context overview and boundaries
2. Ubiquitous language vocabulary
3. Domain model and entities
4. Integration contracts
5. Team responsibilities
6. Quality attributes`,

		"PER": `Generate a Persona (PER) document that describes a user archetype.

Include:
- Persona demographics and background
- Jobs-to-be-Done (JTBD) framework
- Pain points and motivations
- Goals and success criteria
- User journey touchpoints
- Related capabilities and processes

Structure the content with:
1. Persona overview and demographics
2. Jobs-to-be-Done (JTBD)
3. Pain points and challenges
4. Goals and success metrics
5. User journey and touchpoints
6. Related capabilities`,

		"PRO": `Generate a Process (PRO) document that describes a repeatable business process.

Include:
- Process overview and objectives
- Step-by-step workflow
- Inputs, outputs, and artifacts
- Roles and responsibilities
- Decision points and branches
- Success metrics and SLAs

Structure the content with:
1. Process overview and scope
2. Process flow and steps
3. Inputs and outputs
4. Roles and responsibilities
5. Decision points
6. Metrics and SLAs`,

		"ARC": `Generate an Architecture (ARC) document that describes technical architecture.

Include:
- Architecture overview and principles
- Component structure and relationships
- Technology stack and decisions
- Quality attributes and constraints
- Deployment and infrastructure
- Security considerations

Structure the content with:
1. Architecture overview
2. Component architecture
3. Technology decisions
4. Quality attributes
5. Deployment model
6. Security architecture`,

		"MET": `Generate a Metrics (MET) document that defines KPIs and measurement.

Include:
- Metric definition and purpose
- Measurement methodology
- Data sources and collection
- Target values and thresholds
- Alerting and monitoring
- Reporting and dashboards

Structure the content with:
1. Metric overview and purpose
2. Measurement definition
3. Data collection methodology
4. Targets and thresholds
5. Monitoring and alerting
6. Reporting requirements`,

		"RSK": `Generate a Risk (RSK) document that assesses and mitigates risks.

Include:
- Risk description and category
- Probability and impact assessment
- Risk factors and triggers
- Mitigation strategies
- Contingency plans
- Monitoring indicators

Structure the content with:
1. Risk overview and classification
2. Probability and impact analysis
3. Risk factors and scenarios
4. Mitigation strategies
5. Contingency planning
6. Monitoring and indicators`,
	}

	if prompt, exists := prompts[docType]; exists {
		return prompt
	}

	return fmt.Sprintf("Generate a %s document following BSpec standards with proper YAML frontmatter and structured markdown content.", docType)
}

// GetAnalysisPrompt returns a prompt for analyzing BSpec documents
func GetAnalysisPrompt() string {
	return `Analyze the provided BSpec document and provide comprehensive feedback covering:

## Structure Analysis:
- YAML frontmatter completeness and correctness
- Required fields validation (id, title, status, version, owner)
- Optional fields usage (contexts, related, metrics, risks)
- Naming convention adherence

## Content Quality:
- Document type appropriateness and structure
- Content completeness and clarity
- Cross-reference accuracy and usefulness
- Acceptance criteria presence and quality

## BSpec Compliance:
- Adherence to BSpec standards and conventions
- Proper use of document type patterns
- Quality attribute considerations
- Traceability and relationships

## Recommendations:
- Missing elements or sections
- Improvement opportunities
- Best practice suggestions
- Related documents to consider

Provide specific, actionable feedback that helps improve the document quality and BSpec compliance.`
}

// EnhanceUserMessage enhances a user message with BSpec context if appropriate
func EnhanceUserMessage(message string) string {
	message = strings.TrimSpace(message)

	// Check if the message is asking for document generation
	generateKeywords := []string{"generate", "create", "write", "make", "build", "draft"}
	docTypes := []string{"capability", "context", "persona", "process", "offering", "architecture", "compliance", "metrics", "risk", "decision", "experiment", "profile"}

	lowerMessage := strings.ToLower(message)

	for _, keyword := range generateKeywords {
		if strings.Contains(lowerMessage, keyword) {
			for _, docType := range docTypes {
				if strings.Contains(lowerMessage, docType) {
					return message + "\n\nPlease ensure the document follows BSpec standards with proper YAML frontmatter and structured markdown content."
				}
			}
		}
	}

	// Check if asking for analysis
	analysisKeywords := []string{"analyze", "review", "check", "validate", "assess", "evaluate"}
	for _, keyword := range analysisKeywords {
		if strings.Contains(lowerMessage, keyword) {
			return message + "\n\nPlease provide BSpec-specific analysis focusing on structure, compliance, and quality."
		}
	}

	return message
}

// GetWelcomeMessage returns a context-aware welcome message
func GetWelcomeMessage() string {
	return `Welcome to BSpec Chat! I'm here to help you with Business Specification Standard documents.

I can help you:
• Generate BSpec documents (CAP, CTX, PER, PRO, OFF, ARC, COM, MET, RSK, DEC, EXP, PRF)
• Analyze existing BSpec files for compliance and quality
• Answer questions about BSpec standards and best practices
• Provide guidance on document structure and relationships

What would you like to work on today?`
}