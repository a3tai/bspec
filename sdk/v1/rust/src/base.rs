//! Base types and traits for BSpec documents
//!
//! **GENERATED CODE - DO NOT EDIT MANUALLY**
//! Generated at: 2025-09-28T16:00:48.121625

use chrono::{DateTime, Utc};
use serde::{Deserialize, Serialize};
use std::collections::HashMap;
use url::Url;
use uuid::Uuid;

use crate::error::{BSpecError, Result};

/// All possible document types in BSpec v1.0.0
#[derive(Debug, Clone, PartialEq, Eq, Hash, Serialize, Deserialize)]
pub enum DocumentType {    /// Mission: Why the organization exists and its core purpose
    Msn,    /// Vision: The future state the organization aims to create
    Vsn,    /// Values: Guiding principles that shape culture and decisions
    Val,    /// Strategy: How the organization will achieve its vision and compete
    Str,    /// Objectives: Specific, measurable goals with timeframes (OKRs)
    Obj,    /// Moats: Competitive advantages that protect market position
    Mot,    /// Purpose: Social impact and stakeholder value beyond profit
    Pur,    /// Theory of Change: Logic model connecting activities to outcomes
    Thy,    /// Market Definition: TAM/SAM/SOM, market boundaries and sizing
    Mkt,    /// Market Segments: Customer categories and targeting strategy
    Seg,    /// Competitive Analysis: Landscape mapping and competitor profiles
    Cmp,    /// Positioning: Unique value proposition and market position
    Pos,    /// Trends: Market forces and changes shaping the industry
    Trn,    /// Ecosystem: Partners, suppliers, distributors, and ecosystem players
    Eco,    /// Opportunities: Identified market gaps and growth potential
    Opp,    /// Threats: External risks to market position and business model
    Thr,    /// Regulatory Environment: Laws, regulations, and compliance requirements
    Reg,    /// Macro Environment: Economic, political, social, technological factors
    Mac,    /// Personas: Detailed user archetypes with characteristics and needs
    Per,    /// Jobs-to-be-Done: Specific outcomes customers hire products to achieve
    Jtb,    /// Customer Journey: End-to-end experience from awareness to advocacy
    Cjm,    /// Use Cases: Specific scenarios where customers apply the solution
    Use,    /// User Stories: Individual requirements from user perspective
    Sto,    /// Pain Points: Customer problems and frustrations being solved
    Pai,    /// Gains: Benefits and value customers receive
    Gai,    /// Empathy Maps: Deep understanding of customer thoughts and feelings
    Emp,    /// Feedback: Customer input, reviews, and satisfaction data
    Fee,    /// Interviews: Direct customer research and insights
    Int,    /// Surveys: Quantitative customer research and data
    Sur,    /// Behaviors: Customer usage patterns and behavioral insights
    Beh,    /// Products: Physical or digital products offered
    Prd,    /// Services: Service offerings and support capabilities
    Svc,    /// Features: Specific capabilities of products and services
    Fea,    /// Roadmap: Planned development and enhancement timeline
    Rod,    /// Requirements: Functional and non-functional specifications
    Req,    /// Quality Standards: Quality metrics, testing, and assurance practices
    Qua,    /// User Experience: Design principles and user interface standards
    Uxd,    /// Performance: Speed, reliability, and performance characteristics
    Per,    /// Integrations: Connections with other systems and platforms
    Int,    /// Support: Customer support and success processes
    Sup,    /// Business Model Canvas: Complete business model overview
    Bmc,    /// Revenue Streams: How money is generated from customers
    Rev,    /// Pricing: Pricing strategy, models, and tactics
    Prc,    /// Cost Structure: Major cost categories and cost drivers
    Cst,    /// Channels: Distribution and sales channel strategy
    Chn,    /// Customer Relationships: How relationships are built and maintained
    Rel,    /// Key Resources: Critical assets required for the business model
    Res,    /// Key Activities: Essential processes for value creation
    Act,    /// Key Partnerships: Strategic alliances and supplier relationships
    Prt,    /// Unit Economics: Per-customer or per-unit financial metrics
    Unt,    /// Lifetime Value: Customer lifetime value models and drivers
    Ltv,    /// Customer Acquisition: Customer acquisition cost and strategies
    Cac,    /// Processes: Repeatable workflows that drive business outcomes
    Prc,    /// Workflows: Detailed task sequences within processes
    Wfl,    /// Organization: Reporting relationships and organizational design
    Org,    /// Roles: Individual position definitions and responsibilities
    Rol,    /// Teams: Team structures, compositions, and dynamics
    Tea,    /// Skills: Required capabilities and competency frameworks
    Ski,    /// Policies: Rules and guidelines governing behavior
    Pol,    /// Service Levels: Performance commitments and standards
    Sla,    /// Vendors: Supplier and partner relationship management
    Vnd,    /// Facilities: Physical spaces, equipment, and infrastructure
    Fac,    /// Tools: Software, systems, and operational tools
    Too,    /// Capabilities: Organizational capabilities and competencies
    Cap,    /// Architecture: High-level system design and component relationships
    Arc,    /// Systems: Software systems, platforms, and applications
    Sys,    /// Data Models: Data structures, schemas, and information architecture
    Dat,    /// APIs: Interface specifications, protocols, and integrations
    Api,    /// Infrastructure: Deployment, hosting, and runtime environment
    Inf,    /// Security: Security architecture, controls, and compliance
    Sec,    /// Development: Software development lifecycle and practices
    Dev,    /// Analytics: Data analytics, business intelligence, and insights
    Ana,    /// Financial Model: Comprehensive P&L, balance sheet, cash flow projections
    Fin,    /// Budget: Resource allocation and spending plans
    Bud,    /// Forecasts: Forward-looking financial predictions and scenarios
    For,    /// Funding: Capital requirements and financing strategy
    Fnd,    /// Investment: Capital allocation and investment decisions
    Inv,    /// Valuation: Business valuation models and assumptions
    Val,    /// Metrics: Key performance indicators and measurement frameworks
    Met,    /// Reporting: Financial reporting and dashboard requirements
    Rep,    /// Audit: Financial controls, compliance, and audit processes
    Aud,    /// Tax Strategy: Tax planning, structure, and compliance
    Tax,    /// Risks: Identified threats to business success
    Rsk,    /// Mitigations: Risk response strategies and controls
    Mit,    /// Compliance: Regulatory obligations and adherence requirements
    Cmp,    /// Governance: Decision-making structure and oversight
    Gvn,    /// Controls: Internal controls and process safeguards
    Ctl,    /// Crisis Management: Crisis response and business continuity plans
    Cri,    /// Ethics: Ethical guidelines and moral standards
    Eth,    /// Stakeholders: Stakeholder mapping and management
    Sta,    /// Go-to-Market: Launch and customer acquisition strategy
    Gtm,    /// Growth Model: How the business scales customers and revenue
    Grw,    /// Scaling: Operational scaling approach and constraints
    Scl,    /// Experiments: Validation tests and learning initiatives
    Exp,    /// Innovation: Future product and service development pipeline
    Inn,    /// Research: Investigation and development activities
    Rnd,    /// Acquisitions: M&A strategy and integration planning
    Acq,    /// Expansion: Geographic, market, or product expansion plans
    Exp,}

impl DocumentType {
    /// Get the string representation of the document type
    pub fn as_str(&self) -> &'static str {
        match self {            Self::Msn => "MSN",            Self::Vsn => "VSN",            Self::Val => "VAL",            Self::Str => "STR",            Self::Obj => "OBJ",            Self::Mot => "MOT",            Self::Pur => "PUR",            Self::Thy => "THY",            Self::Mkt => "MKT",            Self::Seg => "SEG",            Self::Cmp => "CMP",            Self::Pos => "POS",            Self::Trn => "TRN",            Self::Eco => "ECO",            Self::Opp => "OPP",            Self::Thr => "THR",            Self::Reg => "REG",            Self::Mac => "MAC",            Self::Per => "PER",            Self::Jtb => "JTB",            Self::Cjm => "CJM",            Self::Use => "USE",            Self::Sto => "STO",            Self::Pai => "PAI",            Self::Gai => "GAI",            Self::Emp => "EMP",            Self::Fee => "FEE",            Self::Int => "INT",            Self::Sur => "SUR",            Self::Beh => "BEH",            Self::Prd => "PRD",            Self::Svc => "SVC",            Self::Fea => "FEA",            Self::Rod => "ROD",            Self::Req => "REQ",            Self::Qua => "QUA",            Self::Uxd => "UXD",            Self::Per => "PER",            Self::Int => "INT",            Self::Sup => "SUP",            Self::Bmc => "BMC",            Self::Rev => "REV",            Self::Prc => "PRC",            Self::Cst => "CST",            Self::Chn => "CHN",            Self::Rel => "REL",            Self::Res => "RES",            Self::Act => "ACT",            Self::Prt => "PRT",            Self::Unt => "UNT",            Self::Ltv => "LTV",            Self::Cac => "CAC",            Self::Prc => "PRC",            Self::Wfl => "WFL",            Self::Org => "ORG",            Self::Rol => "ROL",            Self::Tea => "TEA",            Self::Ski => "SKI",            Self::Pol => "POL",            Self::Sla => "SLA",            Self::Vnd => "VND",            Self::Fac => "FAC",            Self::Too => "TOO",            Self::Cap => "CAP",            Self::Arc => "ARC",            Self::Sys => "SYS",            Self::Dat => "DAT",            Self::Api => "API",            Self::Inf => "INF",            Self::Sec => "SEC",            Self::Dev => "DEV",            Self::Ana => "ANA",            Self::Fin => "FIN",            Self::Bud => "BUD",            Self::For => "FOR",            Self::Fnd => "FND",            Self::Inv => "INV",            Self::Val => "VAL",            Self::Met => "MET",            Self::Rep => "REP",            Self::Aud => "AUD",            Self::Tax => "TAX",            Self::Rsk => "RSK",            Self::Mit => "MIT",            Self::Cmp => "CMP",            Self::Gvn => "GVN",            Self::Ctl => "CTL",            Self::Cri => "CRI",            Self::Eth => "ETH",            Self::Sta => "STA",            Self::Gtm => "GTM",            Self::Grw => "GRW",            Self::Scl => "SCL",            Self::Exp => "EXP",            Self::Inn => "INN",            Self::Rnd => "RND",            Self::Acq => "ACQ",            Self::Exp => "EXP",        }
    }

    /// Get the display name of the document type
    pub fn display_name(&self) -> &'static str {
        match self {            Self::Msn => "Mission",            Self::Vsn => "Vision",            Self::Val => "Values",            Self::Str => "Strategy",            Self::Obj => "Objectives",            Self::Mot => "Moats",            Self::Pur => "Purpose",            Self::Thy => "Theory of Change",            Self::Mkt => "Market Definition",            Self::Seg => "Market Segments",            Self::Cmp => "Competitive Analysis",            Self::Pos => "Positioning",            Self::Trn => "Trends",            Self::Eco => "Ecosystem",            Self::Opp => "Opportunities",            Self::Thr => "Threats",            Self::Reg => "Regulatory Environment",            Self::Mac => "Macro Environment",            Self::Per => "Personas",            Self::Jtb => "Jobs-to-be-Done",            Self::Cjm => "Customer Journey",            Self::Use => "Use Cases",            Self::Sto => "User Stories",            Self::Pai => "Pain Points",            Self::Gai => "Gains",            Self::Emp => "Empathy Maps",            Self::Fee => "Feedback",            Self::Int => "Interviews",            Self::Sur => "Surveys",            Self::Beh => "Behaviors",            Self::Prd => "Products",            Self::Svc => "Services",            Self::Fea => "Features",            Self::Rod => "Roadmap",            Self::Req => "Requirements",            Self::Qua => "Quality Standards",            Self::Uxd => "User Experience",            Self::Per => "Performance",            Self::Int => "Integrations",            Self::Sup => "Support",            Self::Bmc => "Business Model Canvas",            Self::Rev => "Revenue Streams",            Self::Prc => "Pricing",            Self::Cst => "Cost Structure",            Self::Chn => "Channels",            Self::Rel => "Customer Relationships",            Self::Res => "Key Resources",            Self::Act => "Key Activities",            Self::Prt => "Key Partnerships",            Self::Unt => "Unit Economics",            Self::Ltv => "Lifetime Value",            Self::Cac => "Customer Acquisition",            Self::Prc => "Processes",            Self::Wfl => "Workflows",            Self::Org => "Organization",            Self::Rol => "Roles",            Self::Tea => "Teams",            Self::Ski => "Skills",            Self::Pol => "Policies",            Self::Sla => "Service Levels",            Self::Vnd => "Vendors",            Self::Fac => "Facilities",            Self::Too => "Tools",            Self::Cap => "Capabilities",            Self::Arc => "Architecture",            Self::Sys => "Systems",            Self::Dat => "Data Models",            Self::Api => "APIs",            Self::Inf => "Infrastructure",            Self::Sec => "Security",            Self::Dev => "Development",            Self::Ana => "Analytics",            Self::Fin => "Financial Model",            Self::Bud => "Budget",            Self::For => "Forecasts",            Self::Fnd => "Funding",            Self::Inv => "Investment",            Self::Val => "Valuation",            Self::Met => "Metrics",            Self::Rep => "Reporting",            Self::Aud => "Audit",            Self::Tax => "Tax Strategy",            Self::Rsk => "Risks",            Self::Mit => "Mitigations",            Self::Cmp => "Compliance",            Self::Gvn => "Governance",            Self::Ctl => "Controls",            Self::Cri => "Crisis Management",            Self::Eth => "Ethics",            Self::Sta => "Stakeholders",            Self::Gtm => "Go-to-Market",            Self::Grw => "Growth Model",            Self::Scl => "Scaling",            Self::Exp => "Experiments",            Self::Inn => "Innovation",            Self::Rnd => "Research",            Self::Acq => "Acquisitions",            Self::Exp => "Expansion",        }
    }

    /// Get the purpose/description of the document type
    pub fn purpose(&self) -> &'static str {
        match self {            Self::Msn => "Why the organization exists and its core purpose",            Self::Vsn => "The future state the organization aims to create",            Self::Val => "Guiding principles that shape culture and decisions",            Self::Str => "How the organization will achieve its vision and compete",            Self::Obj => "Specific, measurable goals with timeframes (OKRs)",            Self::Mot => "Competitive advantages that protect market position",            Self::Pur => "Social impact and stakeholder value beyond profit",            Self::Thy => "Logic model connecting activities to outcomes",            Self::Mkt => "TAM/SAM/SOM, market boundaries and sizing",            Self::Seg => "Customer categories and targeting strategy",            Self::Cmp => "Landscape mapping and competitor profiles",            Self::Pos => "Unique value proposition and market position",            Self::Trn => "Market forces and changes shaping the industry",            Self::Eco => "Partners, suppliers, distributors, and ecosystem players",            Self::Opp => "Identified market gaps and growth potential",            Self::Thr => "External risks to market position and business model",            Self::Reg => "Laws, regulations, and compliance requirements",            Self::Mac => "Economic, political, social, technological factors",            Self::Per => "Detailed user archetypes with characteristics and needs",            Self::Jtb => "Specific outcomes customers hire products to achieve",            Self::Cjm => "End-to-end experience from awareness to advocacy",            Self::Use => "Specific scenarios where customers apply the solution",            Self::Sto => "Individual requirements from user perspective",            Self::Pai => "Customer problems and frustrations being solved",            Self::Gai => "Benefits and value customers receive",            Self::Emp => "Deep understanding of customer thoughts and feelings",            Self::Fee => "Customer input, reviews, and satisfaction data",            Self::Int => "Direct customer research and insights",            Self::Sur => "Quantitative customer research and data",            Self::Beh => "Customer usage patterns and behavioral insights",            Self::Prd => "Physical or digital products offered",            Self::Svc => "Service offerings and support capabilities",            Self::Fea => "Specific capabilities of products and services",            Self::Rod => "Planned development and enhancement timeline",            Self::Req => "Functional and non-functional specifications",            Self::Qua => "Quality metrics, testing, and assurance practices",            Self::Uxd => "Design principles and user interface standards",            Self::Per => "Speed, reliability, and performance characteristics",            Self::Int => "Connections with other systems and platforms",            Self::Sup => "Customer support and success processes",            Self::Bmc => "Complete business model overview",            Self::Rev => "How money is generated from customers",            Self::Prc => "Pricing strategy, models, and tactics",            Self::Cst => "Major cost categories and cost drivers",            Self::Chn => "Distribution and sales channel strategy",            Self::Rel => "How relationships are built and maintained",            Self::Res => "Critical assets required for the business model",            Self::Act => "Essential processes for value creation",            Self::Prt => "Strategic alliances and supplier relationships",            Self::Unt => "Per-customer or per-unit financial metrics",            Self::Ltv => "Customer lifetime value models and drivers",            Self::Cac => "Customer acquisition cost and strategies",            Self::Prc => "Repeatable workflows that drive business outcomes",            Self::Wfl => "Detailed task sequences within processes",            Self::Org => "Reporting relationships and organizational design",            Self::Rol => "Individual position definitions and responsibilities",            Self::Tea => "Team structures, compositions, and dynamics",            Self::Ski => "Required capabilities and competency frameworks",            Self::Pol => "Rules and guidelines governing behavior",            Self::Sla => "Performance commitments and standards",            Self::Vnd => "Supplier and partner relationship management",            Self::Fac => "Physical spaces, equipment, and infrastructure",            Self::Too => "Software, systems, and operational tools",            Self::Cap => "Organizational capabilities and competencies",            Self::Arc => "High-level system design and component relationships",            Self::Sys => "Software systems, platforms, and applications",            Self::Dat => "Data structures, schemas, and information architecture",            Self::Api => "Interface specifications, protocols, and integrations",            Self::Inf => "Deployment, hosting, and runtime environment",            Self::Sec => "Security architecture, controls, and compliance",            Self::Dev => "Software development lifecycle and practices",            Self::Ana => "Data analytics, business intelligence, and insights",            Self::Fin => "Comprehensive P&L, balance sheet, cash flow projections",            Self::Bud => "Resource allocation and spending plans",            Self::For => "Forward-looking financial predictions and scenarios",            Self::Fnd => "Capital requirements and financing strategy",            Self::Inv => "Capital allocation and investment decisions",            Self::Val => "Business valuation models and assumptions",            Self::Met => "Key performance indicators and measurement frameworks",            Self::Rep => "Financial reporting and dashboard requirements",            Self::Aud => "Financial controls, compliance, and audit processes",            Self::Tax => "Tax planning, structure, and compliance",            Self::Rsk => "Identified threats to business success",            Self::Mit => "Risk response strategies and controls",            Self::Cmp => "Regulatory obligations and adherence requirements",            Self::Gvn => "Decision-making structure and oversight",            Self::Ctl => "Internal controls and process safeguards",            Self::Cri => "Crisis response and business continuity plans",            Self::Eth => "Ethical guidelines and moral standards",            Self::Sta => "Stakeholder mapping and management",            Self::Gtm => "Launch and customer acquisition strategy",            Self::Grw => "How the business scales customers and revenue",            Self::Scl => "Operational scaling approach and constraints",            Self::Exp => "Validation tests and learning initiatives",            Self::Inn => "Future product and service development pipeline",            Self::Rnd => "Investigation and development activities",            Self::Acq => "M&A strategy and integration planning",            Self::Exp => "Geographic, market, or product expansion plans",        }
    }

    /// Get the domain this document type belongs to
    pub fn domain(&self) -> Domain {
        match self {            Self::Msn => Domain::Strategic,            Self::Vsn => Domain::Strategic,            Self::Val => Domain::Strategic,            Self::Str => Domain::Strategic,            Self::Obj => Domain::Strategic,            Self::Mot => Domain::Strategic,            Self::Pur => Domain::Strategic,            Self::Thy => Domain::Strategic,            Self::Mkt => Domain::Market,            Self::Seg => Domain::Market,            Self::Cmp => Domain::Market,            Self::Pos => Domain::Market,            Self::Trn => Domain::Market,            Self::Eco => Domain::Market,            Self::Opp => Domain::Market,            Self::Thr => Domain::Market,            Self::Reg => Domain::Market,            Self::Mac => Domain::Market,            Self::Per => Domain::Customer,            Self::Jtb => Domain::Customer,            Self::Cjm => Domain::Customer,            Self::Use => Domain::Customer,            Self::Sto => Domain::Customer,            Self::Pai => Domain::Customer,            Self::Gai => Domain::Customer,            Self::Emp => Domain::Customer,            Self::Fee => Domain::Customer,            Self::Int => Domain::Customer,            Self::Sur => Domain::Customer,            Self::Beh => Domain::Customer,            Self::Prd => Domain::Product,            Self::Svc => Domain::Product,            Self::Fea => Domain::Product,            Self::Rod => Domain::Product,            Self::Req => Domain::Product,            Self::Qua => Domain::Product,            Self::Uxd => Domain::Product,            Self::Per => Domain::Product,            Self::Int => Domain::Product,            Self::Sup => Domain::Product,            Self::Bmc => Domain::BusinessModel,            Self::Rev => Domain::BusinessModel,            Self::Prc => Domain::BusinessModel,            Self::Cst => Domain::BusinessModel,            Self::Chn => Domain::BusinessModel,            Self::Rel => Domain::BusinessModel,            Self::Res => Domain::BusinessModel,            Self::Act => Domain::BusinessModel,            Self::Prt => Domain::BusinessModel,            Self::Unt => Domain::BusinessModel,            Self::Ltv => Domain::BusinessModel,            Self::Cac => Domain::BusinessModel,            Self::Prc => Domain::Operations,            Self::Wfl => Domain::Operations,            Self::Org => Domain::Operations,            Self::Rol => Domain::Operations,            Self::Tea => Domain::Operations,            Self::Ski => Domain::Operations,            Self::Pol => Domain::Operations,            Self::Sla => Domain::Operations,            Self::Vnd => Domain::Operations,            Self::Fac => Domain::Operations,            Self::Too => Domain::Operations,            Self::Cap => Domain::Operations,            Self::Arc => Domain::Technology,            Self::Sys => Domain::Technology,            Self::Dat => Domain::Technology,            Self::Api => Domain::Technology,            Self::Inf => Domain::Technology,            Self::Sec => Domain::Technology,            Self::Dev => Domain::Technology,            Self::Ana => Domain::Technology,            Self::Fin => Domain::Financial,            Self::Bud => Domain::Financial,            Self::For => Domain::Financial,            Self::Fnd => Domain::Financial,            Self::Inv => Domain::Financial,            Self::Val => Domain::Financial,            Self::Met => Domain::Financial,            Self::Rep => Domain::Financial,            Self::Aud => Domain::Financial,            Self::Tax => Domain::Financial,            Self::Rsk => Domain::Risk,            Self::Mit => Domain::Risk,            Self::Cmp => Domain::Risk,            Self::Gvn => Domain::Risk,            Self::Ctl => Domain::Risk,            Self::Cri => Domain::Risk,            Self::Eth => Domain::Risk,            Self::Sta => Domain::Risk,            Self::Gtm => Domain::Growth,            Self::Grw => Domain::Growth,            Self::Scl => Domain::Growth,            Self::Exp => Domain::Growth,            Self::Inn => Domain::Growth,            Self::Rnd => Domain::Growth,            Self::Acq => Domain::Growth,            Self::Exp => Domain::Growth,        }
    }
}

impl std::fmt::Display for DocumentType {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "{}", self.as_str())
    }
}

impl std::str::FromStr for DocumentType {
    type Err = BSpecError;

    fn from_str(s: &str) -> Result<Self> {
        match s.to_uppercase().as_str() {            "MSN" => Ok(Self::Msn),            "VSN" => Ok(Self::Vsn),            "VAL" => Ok(Self::Val),            "STR" => Ok(Self::Str),            "OBJ" => Ok(Self::Obj),            "MOT" => Ok(Self::Mot),            "PUR" => Ok(Self::Pur),            "THY" => Ok(Self::Thy),            "MKT" => Ok(Self::Mkt),            "SEG" => Ok(Self::Seg),            "CMP" => Ok(Self::Cmp),            "POS" => Ok(Self::Pos),            "TRN" => Ok(Self::Trn),            "ECO" => Ok(Self::Eco),            "OPP" => Ok(Self::Opp),            "THR" => Ok(Self::Thr),            "REG" => Ok(Self::Reg),            "MAC" => Ok(Self::Mac),            "PER" => Ok(Self::Per),            "JTB" => Ok(Self::Jtb),            "CJM" => Ok(Self::Cjm),            "USE" => Ok(Self::Use),            "STO" => Ok(Self::Sto),            "PAI" => Ok(Self::Pai),            "GAI" => Ok(Self::Gai),            "EMP" => Ok(Self::Emp),            "FEE" => Ok(Self::Fee),            "INT" => Ok(Self::Int),            "SUR" => Ok(Self::Sur),            "BEH" => Ok(Self::Beh),            "PRD" => Ok(Self::Prd),            "SVC" => Ok(Self::Svc),            "FEA" => Ok(Self::Fea),            "ROD" => Ok(Self::Rod),            "REQ" => Ok(Self::Req),            "QUA" => Ok(Self::Qua),            "UXD" => Ok(Self::Uxd),            "PER" => Ok(Self::Per),            "INT" => Ok(Self::Int),            "SUP" => Ok(Self::Sup),            "BMC" => Ok(Self::Bmc),            "REV" => Ok(Self::Rev),            "PRC" => Ok(Self::Prc),            "CST" => Ok(Self::Cst),            "CHN" => Ok(Self::Chn),            "REL" => Ok(Self::Rel),            "RES" => Ok(Self::Res),            "ACT" => Ok(Self::Act),            "PRT" => Ok(Self::Prt),            "UNT" => Ok(Self::Unt),            "LTV" => Ok(Self::Ltv),            "CAC" => Ok(Self::Cac),            "PRC" => Ok(Self::Prc),            "WFL" => Ok(Self::Wfl),            "ORG" => Ok(Self::Org),            "ROL" => Ok(Self::Rol),            "TEA" => Ok(Self::Tea),            "SKI" => Ok(Self::Ski),            "POL" => Ok(Self::Pol),            "SLA" => Ok(Self::Sla),            "VND" => Ok(Self::Vnd),            "FAC" => Ok(Self::Fac),            "TOO" => Ok(Self::Too),            "CAP" => Ok(Self::Cap),            "ARC" => Ok(Self::Arc),            "SYS" => Ok(Self::Sys),            "DAT" => Ok(Self::Dat),            "API" => Ok(Self::Api),            "INF" => Ok(Self::Inf),            "SEC" => Ok(Self::Sec),            "DEV" => Ok(Self::Dev),            "ANA" => Ok(Self::Ana),            "FIN" => Ok(Self::Fin),            "BUD" => Ok(Self::Bud),            "FOR" => Ok(Self::For),            "FND" => Ok(Self::Fnd),            "INV" => Ok(Self::Inv),            "VAL" => Ok(Self::Val),            "MET" => Ok(Self::Met),            "REP" => Ok(Self::Rep),            "AUD" => Ok(Self::Aud),            "TAX" => Ok(Self::Tax),            "RSK" => Ok(Self::Rsk),            "MIT" => Ok(Self::Mit),            "CMP" => Ok(Self::Cmp),            "GVN" => Ok(Self::Gvn),            "CTL" => Ok(Self::Ctl),            "CRI" => Ok(Self::Cri),            "ETH" => Ok(Self::Eth),            "STA" => Ok(Self::Sta),            "GTM" => Ok(Self::Gtm),            "GRW" => Ok(Self::Grw),            "SCL" => Ok(Self::Scl),            "EXP" => Ok(Self::Exp),            "INN" => Ok(Self::Inn),            "RND" => Ok(Self::Rnd),            "ACQ" => Ok(Self::Acq),            "EXP" => Ok(Self::Exp),            _ => Err(BSpecError::InvalidDocumentType(s.to_string())),
        }
    }
}

/// Business domains in BSpec
#[derive(Debug, Clone, PartialEq, Eq, Hash, Serialize, Deserialize)]
pub enum Domain {    /// Strategic Foundation: The fundamental "why" and "what" of the organization
    Strategic,    /// Market & Environment: Understanding the external landscape and competitive dynamics
    Market,    /// Customer & Value: Understanding who you serve and what value you deliver
    Customer,    /// Product & Service: What the organization delivers to create value
    Product,    /// Business Model: How value is created, delivered, and captured
    BusinessModel,    /// Operations & Execution: How work gets done and the organization operates
    Operations,    /// Technology & Data: Technical architecture and information systems
    Technology,    /// Financial & Investment: Financial planning, performance, and funding
    Financial,    /// Risk & Governance: Risk management and organizational governance
    Risk,    /// Growth & Innovation: How the organization grows and evolves
    Growth,}

impl Domain {
    /// Get the string representation of the domain
    pub fn as_str(&self) -> &'static str {
        match self {            Self::Strategic => "strategic",            Self::Market => "market",            Self::Customer => "customer",            Self::Product => "product",            Self::BusinessModel => "business-model",            Self::Operations => "operations",            Self::Technology => "technology",            Self::Financial => "financial",            Self::Risk => "risk",            Self::Growth => "growth",        }
    }

    /// Get the display name of the domain
    pub fn display_name(&self) -> &'static str {
        match self {            Self::Strategic => "Strategic Foundation",            Self::Market => "Market & Environment",            Self::Customer => "Customer & Value",            Self::Product => "Product & Service",            Self::BusinessModel => "Business Model",            Self::Operations => "Operations & Execution",            Self::Technology => "Technology & Data",            Self::Financial => "Financial & Investment",            Self::Risk => "Risk & Governance",            Self::Growth => "Growth & Innovation",        }
    }

    /// Get the emoji for the domain
    pub fn emoji(&self) -> &'static str {
        match self {            Self::Strategic => "🎯",            Self::Market => "🌍",            Self::Customer => "👥",            Self::Product => "📦",            Self::BusinessModel => "💰",            Self::Operations => "⚙️",            Self::Technology => "🔧",            Self::Financial => "📊",            Self::Risk => "⚠️",            Self::Growth => "📈",        }
    }
}

/// Document status enumeration
#[derive(Debug, Clone, PartialEq, Eq, Hash, Serialize, Deserialize)]
pub enum DocumentStatus {
    /// Document is in draft state
    Draft,
    /// Document is under review
    Review,
    /// Document has been accepted
    Accepted,
    /// Document is deprecated
    Deprecated,
}

impl std::fmt::Display for DocumentStatus {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            Self::Draft => write!(f, "Draft"),
            Self::Review => write!(f, "Review"),
            Self::Accepted => write!(f, "Accepted"),
            Self::Deprecated => write!(f, "Deprecated"),
        }
    }
}

impl std::str::FromStr for DocumentStatus {
    type Err = BSpecError;

    fn from_str(s: &str) -> Result<Self> {
        match s.to_lowercase().as_str() {
            "draft" => Ok(Self::Draft),
            "review" => Ok(Self::Review),
            "accepted" => Ok(Self::Accepted),
            "deprecated" => Ok(Self::Deprecated),
            _ => Err(BSpecError::InvalidDocumentStatus(s.to_string())),
        }
    }
}

/// Conformance levels for BSpec projects
#[derive(Debug, Clone, PartialEq, Eq, Hash, Serialize, Deserialize)]
pub enum ConformanceLevel {    /// Bronze: Minimum Viable Business Spec
    Bronze,    /// Silver: Investment Ready
    Silver,    /// Gold: Operational Excellence
    Gold,}

impl ConformanceLevel {
    /// Get the minimum number of documents required for this level
    pub fn min_documents(&self) -> u32 {
        match self {            Self::Bronze => 12,            Self::Silver => 25,            Self::Gold => 45,        }
    }

    /// Get the emoji for this conformance level
    pub fn emoji(&self) -> &'static str {
        match self {            Self::Bronze => "🥉",            Self::Silver => "🥈",            Self::Gold => "🥇",        }
    }
}

/// Industry profiles for BSpec projects
#[derive(Debug, Clone, PartialEq, Eq, Hash, Serialize, Deserialize)]
pub enum IndustryProfile {    /// Software/SaaS: Software as a Service and cloud platforms
    SoftwareSaas,    /// Physical Product: Physical products and manufacturing businesses
    PhysicalProduct,    /// Service Business: Professional services and consulting businesses
    ServiceBusiness,    /// Nonprofit: Nonprofit organizations and social enterprises
    Nonprofit,}

/// Common metadata for all BSpec documents
#[derive(Debug, Clone, PartialEq, Serialize, Deserialize)]
pub struct DocumentMetadata {
    /// Unique identifier for the document
    pub id: String,
    /// Human-readable title
    pub title: String,
    /// Document status
    pub status: DocumentStatus,
    /// Semantic version
    pub version: String,
    /// Document owner/author
    pub owner: String,
    /// Creation timestamp
    pub created: DateTime<Utc>,
    /// Last update timestamp
    pub updated: DateTime<Utc>,
    /// Business domain
    pub domain: Domain,
    /// Related document IDs
    #[serde(default)]
    pub related: Vec<String>,
    /// Document tags
    #[serde(default)]
    pub tags: Vec<String>,
    /// Custom metadata fields
    #[serde(default)]
    pub custom: HashMap<String, serde_json::Value>,
}

impl Default for DocumentMetadata {
    fn default() -> Self {
        let now = Utc::now();
        Self {
            id: Uuid::new_v4().to_string(),
            title: String::new(),
            status: DocumentStatus::Draft,
            version: "1.0.0".to_string(),
            owner: String::new(),
            created: now,
            updated: now,
            domain: Domain::Strategic,
            related: Vec::new(),
            tags: Vec::new(),
            custom: HashMap::new(),
        }
    }
}

/// Base trait for all BSpec documents
pub trait BSpecDocument: Send + Sync {
    /// Get the document type
    fn document_type(&self) -> DocumentType;

    /// Get the document metadata
    fn metadata(&self) -> &DocumentMetadata;

    /// Get mutable document metadata
    fn metadata_mut(&mut self) -> &mut DocumentMetadata;

    /// Get the document content
    fn content(&self) -> &str;

    /// Set the document content
    fn set_content(&mut self, content: String);

    /// Validate the document structure and content
    #[cfg(feature = "validation")]
    fn validate(&self) -> Result<()>;

    /// Convert document to JSON
    fn to_json(&self) -> Result<String> {
        serde_json::to_string_pretty(self).map_err(BSpecError::Serialization)
    }

    /// Convert document to YAML
    fn to_yaml(&self) -> Result<String> {
        serde_yaml::to_string(self).map_err(BSpecError::Serialization)
    }

    /// Update the modification timestamp
    fn touch(&mut self) {
        self.metadata_mut().updated = Utc::now();
    }

    /// Check if document is valid for the given conformance level
    fn meets_conformance(&self, level: ConformanceLevel) -> bool {
        match level {
            ConformanceLevel::Bronze => !self.metadata().title.is_empty() && !self.content().is_empty(),
            ConformanceLevel::Silver => {
                self.meets_conformance(ConformanceLevel::Bronze)
                    && self.metadata().status != DocumentStatus::Draft
                    && !self.metadata().owner.is_empty()
            }
            ConformanceLevel::Gold => {
                self.meets_conformance(ConformanceLevel::Silver)
                    && self.metadata().status == DocumentStatus::Accepted
                    && !self.metadata().related.is_empty()
            }
        }
    }
}