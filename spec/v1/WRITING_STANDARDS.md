# BSpec Professional Writing Standards

## Purpose

This document establishes authoritative writing and structural standards for all BSpec document type specifications. These standards ensure professional quality comparable to IETF RFCs and MBA-level business specifications.

## Professional Tone Requirements

### 1. Formal Business Language
- **Use**: "Organizations establish systematic approaches to..."
- **Avoid**: "Companies figure out how to..."
- **Use**: "This specification defines requirements for..."
- **Avoid**: "This document tells you how to..."

### 2. Normative Language (RFC 2119 Style)
- **MUST/SHALL**: Absolute requirement
- **MUST NOT/SHALL NOT**: Absolute prohibition  
- **SHOULD/RECOMMENDED**: Strong recommendation
- **SHOULD NOT/NOT RECOMMENDED**: Strong discouragement
- **MAY/OPTIONAL**: Truly optional

### 3. Executive Voice
- Write for C-level executives and senior managers
- Assume business sophistication but explain technical concepts
- Focus on strategic business value and operational impact
- Use quantifiable outcomes where possible

## Document Structure Standards

### Required Sections (in order)

1. **Document Header**
   ```markdown
   # {CODE}: {Full Name} Specification
   
   **Document Type Code:** {CODE}
   **Document Type Name:** {Full Name}
   **Domain:** {Domain Name}
   **Version:** 1.0.0
   **Status:** Draft
   **Last Updated:** {Date}
   ```

2. **Abstract** (NEW - add to all specs)
   ```markdown
   ## Abstract
   
   This specification defines the {document type} document type within the BSpec 1.0 Universal Business Specification Standard. It establishes normative requirements, structured templates, and implementation guidance for organizations documenting {specific business function}. This specification enables systematic, machine-readable documentation of {specific domain} that supports strategic planning, operational execution, and organizational alignment.
   ```

3. **Purpose and Scope**
   - Must be 2-3 sentences
   - Must use normative language
   - Must specify what is included AND excluded

4. **Document Metadata Schema**
   - Must be valid YAML
   - Must include all standard fields
   - Must document domain-specific extensions

5. **Content Structure Template**
   - Must provide complete template
   - Must use consistent formatting
   - Must include normative guidance

6. **Quality Standards** (Bronze/Silver/Gold)
   - Must use checkboxes
   - Must be measurable criteria
   - Must build progressively

7. **Implementation Guidelines**
   - Must include creation process
   - Must include review process  
   - Must include measurement approaches

8. **Validation Requirements**
   - Must provide complete checklist
   - Must be actionable items
   - Must cover all quality dimensions

### Optional but Recommended Sections

9. **Industry Variations**
   - Software/SaaS, Physical Products, Services, Nonprofit
   - Domain-specific adaptations

10. **Common Implementation Patterns**
    - Success patterns with examples
    - Anti-patterns with corrections

11. **Relationship Guidelines** 
    - Typical dependencies and enablements
    - Common conflicts and resolutions

## Language and Style Guide

### Professional Vocabulary

| Instead Of | Use |
|------------|-----|
| "figure out" | "determine", "establish", "define" |
| "deal with" | "address", "manage", "handle" |
| "make sure" | "ensure", "verify", "confirm" |
| "help" | "enable", "facilitate", "support" |
| "get" | "obtain", "acquire", "secure" |
| "really important" | "critical", "essential", "vital" |

### Sentence Structure
- **Preferred**: Active voice with clear subject-verb-object
- **Avoid**: Passive voice unless absolutely necessary
- **Length**: 15-25 words per sentence average
- **Complexity**: One main idea per sentence

### Paragraph Structure  
- **Opening**: Topic sentence stating the main point
- **Body**: Supporting details and examples
- **Closing**: Transition to next concept or summary

## Content Quality Standards

### Completeness Requirements
- All sections must be fully developed
- No placeholder text (e.g., "TBD", "Coming soon")
- All examples must be realistic and industry-relevant
- All references must be accurate and complete

### Accuracy Requirements
- Technical information must be current and correct
- Business concepts must reflect modern practices
- Industry examples must be authentic and relevant
- Compliance information must be up-to-date

### Consistency Requirements
- Terminology must be consistent across all specifications
- Structure must follow the established template exactly
- Formatting must use standard markdown conventions
- Metadata schema must match the domain standards

## Review and Validation Process

### Technical Review
1. **Structure**: Follows template exactly
2. **Completeness**: All required sections present and complete
3. **Accuracy**: Technical and business information is correct
4. **Consistency**: Terminology and formatting match standards

### Editorial Review
1. **Grammar**: Correct grammar and punctuation
2. **Style**: Professional tone and appropriate vocabulary
3. **Clarity**: Clear communication of complex concepts
4. **Flow**: Logical organization and smooth transitions

### Business Review
1. **Relevance**: Content addresses real business needs
2. **Practicality**: Implementation guidance is actionable
3. **Value**: Specification provides clear business value
4. **Completeness**: Covers all necessary aspects of the domain

## Metadata Standards

### Standard Fields (Required)
```yaml
id: {TYPE}-{identifier}
title: "{Document Type} — {Specific Instance}"
type: {TYPE}
status: Draft|Review|Approved|Active|Deprecated
version: 1.0.0
owner: {responsible-role}
stakeholders: [stakeholder-list]
domain: {domain-name}
priority: Critical|High|Medium|Low
scope: {scope-definition}
horizon: Strategic|Tactical|Operational
visibility: Internal|Restricted|Public

depends_on: [document-type-codes]
enables: [document-type-codes]

success_criteria:
  - "Measurable success criterion 1"
  - "Measurable success criterion 2"

assumptions:
  - "Key assumption underlying this document type"

review_cycle: {frequency}
```

### Domain-Specific Extensions
Each domain may add specific metadata fields relevant to that business area, but core fields must remain consistent across all specifications.

## Implementation Guidelines

### Creation Process
1. Copy the standard template
2. Complete all required sections using professional language
3. Add domain-specific content and metadata
4. Validate against quality standards
5. Submit for technical and editorial review

### Maintenance Process
1. Regular review according to specified cycle
2. Update based on industry evolution and feedback
3. Maintain version history and change documentation
4. Ensure backwards compatibility when possible

This standard applies to all BSpec document type specifications and must be followed to ensure professional quality and consistency across the entire specification framework.