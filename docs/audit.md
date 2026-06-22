# BSpec 1.0.0 Audit — Findings and Recommendations

**Spec under review:** BSpec 1.0.0, Universal Business Specification Standard — 112 document types across 12 domains.
**Audit scope:** Overlap/redundancy, gaps, framework misalignment, naming, plus full legal lens (IP/trademark, framework copying, regulatory accuracy, open-spec posture).
**Method:** Full content review of representative documents in every domain via the BSpec MCP; cross-checked against established frameworks and licenses via Tavily and Exa.

The TL;DR is that BSpec is structurally well-organized but currently has three serious problems: (a) its own internal references are broken in dozens of places, (b) it derives heavily from named, copyrighted, or trademarked frameworks without attribution, and (c) several risk/governance/finance docs misstate the scope of real-world laws. None of these are fatal — all are fixable — but they need to be cleared before the spec is published, sold, or marketed as a standard.

---

## 1. Severity-ranked summary

**P0 — must fix before any public/commercial release**

- NPS / Net Promoter Score / Net Promoter System / Net Promoter — registered trademarks of Bain & Company, NICE Systems, and Fred Reichheld — are used by name across CUS, VST, SUP, MET, FEE, SUR, REP, SOC, IFL with no attribution. Bain's licensing page is explicit: the methodology is free to use, the marks are not, and commercial mark use requires Bain's written consent in the form of an NPS License.
- Internal type-code collisions in YAML frontmatter: at least five document types declare the wrong `type` and `id` prefix in their own metadata (BPO declares POS; MCH declares CHN; CIN declares INT; IFL declares INF; PSP declares PER, which collides with PER = Personas). This is a hard correctness bug that breaks any tooling that relies on `type`.
- 25+ phantom code targets in `depends_on`, `enables`, and Relationship Guidelines — references to codes that do not exist in the 112-code list (BMC, GTM, PRC, REL, GRW, STA, IND, NEG, SUC, TSK, KWD, WEB, LED, CON, BRA, SAL, EDU, TRA, CUL, CRI, BCR, MON, AUT, COL, INK, MIT, PLN, PRJ). The relationship graph is the central value claim of BSpec and is currently lying.
- Value Proposition Canvas (Strategyzer) is adapted as PAI/GAI/JTB with no attribution. Unlike the Business Model Canvas, VPC is **not** under Creative Commons — Strategyzer reserves IP for adaptations and inclusion in software applications, and requires express written permission for either.

**P1 — fix before broad distribution**

- Business Model Canvas: 8 of the 9 BMC building blocks appear as BSpec document types (KPT, KAC, KRS, CST, REV, CUS, CHN, SEG) with no attribution to Strategyzer / Osterwalder / Pigneur. BMC is under CC BY-SA 3.0, which permits derivatives but requires attribution + ShareAlike licensing of the derivative. BSpec is currently violating both.
- "BSpec" name has not been trademark-cleared. The term is in use as a product/spec name in other domains (notably Intel uses "bspec" internally for graphics specifications and there are commercial products by similar names). Run proper USPTO TESS and EUIPO clearance before announcing.
- BSpec itself has no declared license, no contribution terms, no copyright posture, no IP policy. For something positioning itself as a "Universal Standard," this is the table-stakes legal documentation and it is currently missing.
- Empathy Map's six-quadrant structure (THINKS / FEELS / SEES / SAYS / DOES / PAINS-GAINS) in EMP is Dave Gray / XPLANE / Gamestorming — typically released CC BY-SA 4.0. No attribution.
- Geoffrey Moore's positioning statement template ("For [target] who [need], [product] is the [category] that [benefit] unlike [competitor] because [reason]…") is reproduced verbatim in POS with no attribution to Moore or *Crossing the Chasm*.
- Tuckman's forming/storming/norming/performing stages appear verbatim as `maturity_level` enum values in TEA. No attribution.
- Kaplan & Norton's Balanced Scorecard (four perspectives) appears as a structured section in MET. No attribution.

**P2 — quality / accuracy issues to clean up**

- COM and LEG duplicate the same Federal / State / International compliance structure end-to-end — boundaries between the two are unclear.
- Multiple regulatory factual errors (see section 4) including a Section 302 vs Section 906 SOX swap in AUD, HIPAA presented as a generic security standard, GDPR "Right to Erasure" misnamed as "Right to be forgotten," PCI DSS and SOC 2 mis-classified as regulatory standards, and use of the superseded "FIN 48" name for what is now ASC 740-10.
- Several document types reference frameworks without attribution that aren't necessarily legally risky but are sloppy: Three Horizons (McKinsey), Stage-Gate (registered TM of Stage-Gate International), Lean Startup (Ries), Design Thinking (IDEO), CRISP-DM, ITIL P1–P4 incident tiers, ISO 25010 quality characteristics, 12 Brand Archetypes (Mark & Pearson / Jung), Three Lines of Defense (IIA), Activity-Based Costing (Kaplan/Cooper), Antifragility (Taleb), Stakeholder Analysis Matrix (Mendelow).
- ITIL, COBIT, CMMI are named without disclosing they are commercial trademarks of AXELOS / ISACA / SEI-CMU.
- Conspicuously **missing** named frameworks despite covering their territory: SWOT (OPP+THR are SWOT halves but the term never appears), Porter's Five Forces (CMP territory), Porter's Value Chain (VST territory — and they're CC BY-SA 4.0, the safest to borrow), Wardley Maps (strategy/ECO territory — and they're CC BY-SA 4.0, the safest to borrow), COSO ERM (RSK), COSO Internal Control (CTL), AARRR (PRF/MET), HEART (MET/UXD), Service Blueprint (SVC/CJM), Kano model (PRD/FEA), QFD (PRD), North Star Metric (MET/OBJ), Lean Canvas (BMC adjacency).

**P3 — naming / structural cleanup**

- Code "PER" collides between Personas (customer-value) and PSP's declared `type: PER` in metadata.
- STO ("User Stories") is an unintuitive code choice — STO suggests "Store"; USR or STR (already taken) would be more natural; common in industry is "US" but three letters are required, so USS or STY.
- JTB is used as a code but the universally recognized acronym is JTBD — keeping the three-letter constraint costs recognition; consider JBD or accept four-letter codes for this one case.
- VAL collides conceptually: BSpec uses VAL for "Organizational Values," but POS lists "VAL (Values)" in its dependencies in a context that obviously means Value Propositions. The Value Proposition concept is completely missing as a first-class document type — yet it's referenced as if it exists.
- THR vs THY (Threats vs Theory of Change) are easily confused at a glance — three-letter prefix collision risk.
- MOT for Moats — readable for English speakers, but MOA would mirror "moat" more naturally; minor.
- "LEA" is ambiguous in usage: WIS treats LEA as "Leadership" and INC treats LEA as "Learning" — but the code is defined as "Learning Organization." Pick one meaning and enforce.

---

## 2. Overlap and redundancy map

Document types that materially overlap and whose boundaries need a designated split:

- **LEG vs COM vs REG** — LEG includes a "Regulatory Compliance and Government Relations" section that duplicates COM's federal/state/international structure, and both reference REG's content. Recommendation: REG = external regulatory environment scan (what laws exist); COM = our operational compliance program; LEG = contracts, IP, litigation, corporate legal affairs. Strip regulatory compliance content out of LEG and reference COM.
- **POS vs BPO** — POS is "Market Positioning" (Market & Environment); BPO is "Brand Positioning" (Brand Marketing). Both define positioning statements, target customer framing, differentiation, and message hierarchy. In practice nearly identical. Recommendation: merge or define BPO as strictly brand-personality/voice positioning while POS handles category/competitive positioning.
- **CHN vs MCH** — CHN is "Channel Strategy" (Business Model = distribution channels); MCH is "Marketing Channel Strategy" (Brand Marketing). The same companies use the same channels for both distribution and marketing. Recommendation: keep CHN for sales/distribution channels and reframe MCH as "Media Mix" or "Marketing Mix" so the boundary is obvious.
- **CUS vs CJM vs FEE vs SUP** — Customer Relationships, Customer Journey Map, Feedback, Support Specification. All four overlap on the customer-experience-after-purchase space. Recommendation: CUS = the relationship operating model, CJM = the visualization, FEE = the input pipeline, SUP = the reactive ticket/support function. Add an explicit "Customer Success" type (currently a gap) for proactive lifecycle management.
- **PRD vs FEA vs REQ vs USE vs STO** — Product Requirements Document, Feature Specification, Requirements Specification, Use Cases, User Stories. PRD and REQ overlap almost completely; FEA is a finer-grained PRD; STO and USE both express user-facing scenarios. Recommendation: PRD = product-level; FEA = feature-level (children of PRD); REQ = the system-level functional/nonfunctional spec used by engineering; collapse USE into either CJM or STO.
- **ARC vs SYS vs INF vs DEV** — Architecture, Systems, Infrastructure, Development. Heavy overlap. Recommendation: ARC = decisions and principles, SYS = catalog of actual systems, INF = hosting/network/platform, DEV = engineering practices. Tighten descriptions accordingly.
- **KNO vs LRN vs LEA vs WIS** — Knowledge Management, Learning Records, Learning Organization, Wisdom Synthesis. Four documents for "stuff we learn." Recommendation: collapse to two. KNO = stored knowledge (codified, retrievable). LRN/LEA/WIS → one "Learning Practice" doc that covers how the org learns.
- **DEC vs HYP** — Decision Records and Hypothesis Management have very similar structure (context, alternatives, rationale, outcome, validation). Keep separate but tighten: DEC = decision is made; HYP = decision is deferred pending evidence.
- **INN vs RND vs EXP vs IGN** — Innovation Strategy, R&D, Experimentation, Insight Generation. Recommendation: INN = the portfolio strategy, RND = the function/team, EXP = the testing protocol, IGN = the insight-to-action process. Right now they read interchangeably.
- **POL vs CTL vs COM vs GOV** — Policies, Controls, Compliance, Governance. The standard separation (which BSpec gestures at but doesn't enforce) is GOV = oversight bodies and authorities, POL = written rules, CTL = mechanisms that enforce/detect, COM = the program that runs it all. Tighten the boundary statements.
- **STR vs OBJ vs MET** — Strategy, Strategic Objectives, Metrics. OBJ already contains OKR Key Results which are themselves metrics; MET re-covers the same ground. Recommendation: OBJ owns the objective hierarchy and its leading/lagging indicators; MET owns the measurement system (definitions, instrumentation, dashboards) used to feed OBJ.
- **VLU vs FIN vs FOR** — Valuation, Financial Model, Forecasts. Overlap on projections and methodology. VLU is the only one that explicitly does enterprise-valuation methods (DCF, comparables, etc.) — make that the sole purpose. FIN = the integrated three-statement model. FOR = scenario/forecast outputs.

---

## 3. Gaps in coverage

Things a real business needs that BSpec does not have:

- **Sales Strategy / Sales Operations** — there is no sales-process document type. SAL is referenced but doesn't exist.
- **Customer Success / Onboarding** — no proactive post-sale doc; SUP is reactive.
- **People / HR Strategy** — ROL, SKI, TEA exist but there is no umbrella People doc covering talent acquisition, compensation, benefits, performance management, succession.
- **Diversity, Equity, Inclusion** — increasingly required for board reporting and public companies; absent.
- **Sustainability / ESG / Climate disclosure** — increasingly required (SEC climate rule, CSRD, ISSB); absent.
- **Data Privacy** — distinct from SEC (security). GDPR/CCPA/PIPEDA require separate privacy programs; currently bundled inside SEC and DAT.
- **M&A / Corporate Development** — referenced in LEG but no first-class doc.
- **Pricing** — PRI exists, good. But discount strategy, deal desk, contract negotiation policies are missing.
- **Product Lifecycle / Sunsetting / End-of-Life** — ROD covers the future; nothing covers retirement.
- **Crisis Management / Business Continuity** — CRI and BCR are referenced (phantom codes); not implemented.
- **Value Proposition** — astonishingly missing as a first-class type, despite VPC (PAI/GAI/JTB) being the customer half. Add VPR or similar.
- **Stakeholder Map** — STA is referenced; absent.
- **SWOT** — absent as a synthesizing doc, despite OPP/THR being half of it.

---

## 4. Legal / regulatory accuracy issues

Specific factual errors found in the risk/governance/finance domain. These matter because the spec presents itself as authoritative.

- **SOX section mislabeling (AUD)**. AUD says Section 906 is "CEO/CFO certification requirements." That description fits Section 302. Section 906 is the *criminal* certification of financial reports (18 U.S.C. § 1350). 302 is the disclosure-controls certification; 404 is the ICFR. Swap fixed.
- **SOX scope (CTL, AUD, REP)**. SOX is presented as a generic compliance regime. SOX applies to U.S. public companies and certain foreign filers; private companies use SOX-like frameworks only voluntarily. Call this out.
- **HIPAA scope (SEC)**. HIPAA is listed as a generic security framework. HIPAA applies only to covered entities and business associates. Misleading for non-healthcare BSpec adopters.
- **GDPR "Right to be forgotten" (SEC)**. The GDPR right is "Right to Erasure" (Article 17). "Right to be forgotten" is the pre-GDPR Google Spain CJEU concept and isn't the official name in the regulation.
- **GDPR/CCPA equivalence (SEC, DAT, REP)**. Repeatedly bracketed as if equivalent. CCPA/CPRA is materially narrower than GDPR (no lawful-basis requirement, no DPO mandate, different scope, different consent model). Stop conflating.
- **PCI DSS classification (SEC)**. PCI DSS is a contractual obligation imposed by card brands, not a regulatory or governmental standard. Bundling it with ISO 27001 / SOC 2 / NIST is technically wrong.
- **SOC 2 (SEC)**. SOC 2 is an AICPA attestation report, not a certification or a standard. The doc refers to "SOC 2 certification" — there is no such thing.
- **ISO 27001 vs SOC 2 (SEC)**. Listed together as `certification_audits`. ISO 27001 is a certifiable standard; SOC 2 is an attestation. They are not the same legal artifact.
- **"FIN 48" (TAX)**. Superseded by ASC 740-10 in the FASB Accounting Standards Codification (2009). Still informally used but is dated in a 2025-stamped spec.
- **ERISA fiduciary insurance (INS)**. Listed as a federal "requirement." ERISA requires bonding under §412 but does not require fiduciary liability insurance — that's an optional product. Misframed.
- **K&R coverage (INS)**. Listed under "political risk." K&R is a separate specialty line. Mis-categorized.
- **FCPA / UK Bribery Act not named (LEG)**. LEG references "anti-corruption" without naming the two operative regimes most businesses face.
- **Whistleblower statutes not named (ETH, LEG)**. ETH mentions whistleblower programs without referencing SOX §806, Dodd-Frank §922, or EU Directive 2019/1937 — which actually create the protections.
- **10-Q filing deadline (REP, AUD)**. The "40-day filing deadline" applies to accelerated and large accelerated filers; non-accelerated filers have 45 days. Oversimplified.
- **"Right to be forgotten" + "PIPEDA" treatment (REG)**. PIPEDA referenced without noting Quebec Law 25, Alberta PIPA, BC PIPA which displace/supplement PIPEDA for certain entities.

---

## 5. Recommended remediation plan

In rough priority order:

**Immediate (P0)**
1. Fix the YAML metadata collisions: BPO, MCH, CIN, IFL, PSP must declare their own type codes in `type:` and `id:` prefixes. This is mechanical and breaks nothing.
2. Remove or replace every phantom code reference. Either add the missing type (e.g., GTM, REL, SAL, STA, MIT, CRI, BCR, CUL, MON) or rewrite the affected `depends_on` / `enables` lists to point to real codes.
3. Add a top-level LICENSE, CONTRIBUTING, and ATTRIBUTION file to the spec. Pick an explicit license — CC BY-SA 4.0 is the conservative choice if BSpec wants to borrow from CC BY-SA frameworks (BMC, Wardley, Empathy Map) since it's the ShareAlike obligation that flows down.
4. Remove all "NPS" / "Net Promoter Score" / "Net Promoter System" references unless attributed in the form Bain mandates: "Net Promoter®, NPS®, Net Promoter System®, Net Promoter Score℠ are registered trademarks/service marks of Bain & Company, Inc., NICE Systems, Inc., and Fred Reichheld." For commercial product/service marketing use, obtain an NPS License from Bain.
5. Rework PAI/GAI/JTB to either (a) attribute Strategyzer's Value Proposition Canvas explicitly and clear adaptation rights with Strategyzer AG (their VPC terms specifically require this for inclusion in software/tools), or (b) restructure these documents around a non-VPC ontology (e.g., problems / outcomes / contexts).

**Near-term (P1)**
6. Clear "BSpec" as a name via USPTO TESS and EUIPO; check that intellectual property domain (intellectual property OS, business OS, etc.) doesn't collide.
7. Add attribution to every named framework borrowed: BMC (Strategyzer, CC BY-SA 3.0), Empathy Map (Gray/XPLANE), Geoffrey Moore positioning template, Tuckman, Balanced Scorecard (Kaplan & Norton), Three Horizons (Baghai/Coley/White / McKinsey), Stage-Gate® (Cooper / Stage-Gate International), Lean Startup (Ries), Design Thinking (IDEO/d.school), CRISP-DM, ISO 25010, 12 Brand Archetypes (Mark & Pearson), Three Lines of Defense (IIA), Activity-Based Costing (Kaplan/Cooper), Customer Effort Score (CEB/Gartner).
8. Carry the same trademark notice format for ITIL® (AXELOS), COBIT® (ISACA), CMMI® (CMMI Institute / ISACA), Stage-Gate® on first mention.
9. Resolve the overlap pairs called out in section 2. The fastest win is collapsing KNO/LRN/LEA/WIS to two docs and merging POS/BPO.
10. Fill the major gaps in section 3, especially Value Proposition, Sales Strategy, Customer Success, People/HR Strategy, Data Privacy, ESG/Sustainability.

**Then (P2)**
11. Sweep the risk/governance/finance domains for the factual errors in section 4. Add scope notes ("HIPAA applies to covered entities…"; "SOX applies to U.S. public companies…"; "PCI DSS is a contractual standard…").
12. Replace "FIN 48" with "ASC 740-10."
13. Decide once and for all whether the spec is jurisdiction-agnostic or U.S.-centric; right now it claims universality but its specifics are U.S.-flavored.
14. Backfill the obvious missing frameworks where they'd be cited anyway: SWOT, Porter's Five Forces, Porter's Value Chain, COSO ERM, COSO IC, Wardley Maps (CC BY-SA 4.0 — the easiest to borrow legally).
15. Add an "Attribution Required" field to the document metadata schema so derivative frameworks always carry their lineage.

---

## 6. What's working well

So the report isn't all negative: the atomic-document principle, the depends_on / enables relationship graph (when references resolve), the Bronze / Silver / Gold maturity tiers, the YAML-frontmatter-plus-markdown design choice, the three-letter code scheme, and the breadth of coverage across 12 domains are all good design decisions. The spec is conceptually ambitious and structurally cleaner than most business taxonomies I've seen. The problems above are fixable polish-and-clearance work, not foundation work.

---

## Appendix — sources

- BSpec content pulled via the registered BSpec MCP (all 112 document types referenced).
- Strategyzer "Usage of our tools" page (strategyzer.com/legal/usage-of-our-tools) — confirms BMC is CC BY-SA 3.0 with attribution requirement; VPC is more restrictive copyright requiring express permission for adaptation and software inclusion.
- Bain & Company Net Promoter trademark and licensing page (netpromotersystem.com/resources/trademarks-and-licensing) — confirms NPS®, Net Promoter®, Net Promoter System®, NPS Prism® are registered trademarks; Net Promoter Score℠ is a service mark; commercial use of the marks requires Bain's written consent.
- USPTO TTABVUE record for JOBS-TO-BE-DONE (Serial #86737621, Strategyn Holdings LLC) — application abandoned after inter-partes decision, so "Jobs-to-be-Done" itself is not a registered trademark. "Outcome-Driven Innovation®" (ODI) remains a Strategyn trademark.
- XPLANE / Gamestorming Empathy Map Canvas pages — confirms the six-quadrant Empathy Map is Dave Gray's design, typically released CC BY-SA 4.0.
- Wardley Maps FAQ (wardleymaps.com) — confirms CC BY-SA 4.0.
