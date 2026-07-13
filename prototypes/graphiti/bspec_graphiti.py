#!/usr/bin/env python
"""Prototype BSpec ingestion/search helpers for Graphiti.

The script is intentionally standalone: parsing and JSONL export work without
Graphiti, Neo4j, or an LLM key. Ingestion/search commands load Graphiti only
when those commands are used.
"""

from __future__ import annotations

import argparse
import asyncio
import json
import os
import re
import sys
import tarfile
import uuid
from collections import Counter, defaultdict
from dataclasses import dataclass
from datetime import datetime, timezone
from pathlib import Path
from typing import Any
from urllib import error, request

try:
    import yaml
except ModuleNotFoundError:  # pragma: no cover - only needed at runtime
    yaml = None


REPO_ROOT = Path(__file__).resolve().parents[2]
DEFAULT_SPEC_DIR = REPO_ROOT / "spec" / "v1"
DEFAULT_EXPORT_PATH = Path(__file__).resolve().parent / "bspec-episodes.jsonl"
DEFAULT_PROJECT_EXPORT_PATH = Path(__file__).resolve().parent / "bspec-project-episodes.jsonl"
DEFAULT_GROUP_ID = "bspec-v1"
DEFAULT_PROJECT_GROUP_ID = "bspec-project"
DEFAULT_REFERENCE_TIME = "2025-09-30T00:00:00+00:00"


RELATIONSHIP_FIELDS = ("depends_on", "enables", "conflicts_with", "related", "supersedes")
RELATIONSHIP_LABELS = {
    "depends_on": "DEPENDS_ON",
    "enables": "ENABLES",
    "conflicts_with": "CONFLICTS_WITH",
    "related": "RELATED_TO",
    "supersedes": "SUPERSEDES",
}


@dataclass(frozen=True)
class RelationshipRef:
    """A normalized BSpec document relationship target."""

    pattern: str
    code: str

    def as_dict(self) -> dict[str, str]:
        return {"pattern": self.pattern, "code": self.code}


@dataclass(frozen=True)
class BSpecDocumentType:
    """Parsed BSpec document type metadata."""

    code: str
    name: str
    domain_name: str
    domain_slug: str
    status: str
    version: str
    last_updated: str | None
    source_path: Path
    abstract: str
    purpose: str
    metadata_schema: dict[str, Any]
    relationships: dict[str, list[RelationshipRef]]
    quality_standards: str
    relationship_guidance: str


@dataclass(frozen=True)
class ProjectRelationshipRef:
    """A normalized relationship target from a user-authored BSpec document."""

    pattern: str
    code: str | None
    document_id: str | None

    def as_dict(self) -> dict[str, str]:
        result = {"pattern": self.pattern}
        if self.code:
            result["code"] = self.code
        if self.document_id:
            result["document_id"] = self.document_id
        return result


@dataclass(frozen=True)
class BSpecProjectDocument:
    """Parsed BSpec project document metadata and body."""

    document_id: str
    title: str
    doc_type: str
    status: str
    version: str
    owner: str
    created: str | None
    updated: str | None
    domain: str
    source_path: str
    content: str
    metadata: dict[str, Any]
    relationships: dict[str, list[ProjectRelationshipRef]]


@dataclass(frozen=True)
class BSpecProject:
    """A user-authored BSpec archive or extracted project directory."""

    input_label: str
    manifest: dict[str, Any]
    documents: list[BSpecProjectDocument]


def repo_relative(path: Path) -> str:
    return str(path.resolve().relative_to(REPO_ROOT))


def display_path(path: Path) -> str:
    try:
        return repo_relative(path)
    except ValueError:
        return str(path)


def stable_uuid(kind: str, key: str) -> str:
    return str(uuid.uuid5(uuid.NAMESPACE_URL, f"https://bspec.dev/graphiti/{kind}/{key}"))


def compact_text(value: str, limit: int = 1600) -> str:
    value = re.sub(r"\s+", " ", value).strip()
    if len(value) <= limit:
        return value
    return value[: limit - 1].rstrip() + "..."


def field_value(content: str, label: str) -> str | None:
    match = re.search(rf"^\*\*{re.escape(label)}:\*\*\s*(.+?)\s*$", content, re.MULTILINE)
    return match.group(1).strip() if match else None


def extract_section(content: str, heading: str) -> str:
    pattern = rf"^##\s+{re.escape(heading)}\s*$\n(?P<body>.*?)(?=^##\s+|\Z)"
    matches = [m.group("body").strip() for m in re.finditer(pattern, content, re.MULTILINE | re.DOTALL)]
    seen: set[str] = set()
    unique = []
    for match in matches:
        normalized = compact_text(match)
        if normalized and normalized not in seen:
            unique.append(normalized)
            seen.add(normalized)
    return "\n\n".join(unique)


def extract_metadata_schema_block(content: str) -> str:
    section_match = re.search(
        r"^##\s+Document Metadata Schema\s*$\n(?P<body>.*?)(?=^##\s+|\Z)",
        content,
        re.MULTILINE | re.DOTALL,
    )
    if not section_match:
        return ""

    section = section_match.group("body")
    match = re.search(r"```(?:yaml)?\s*\n(?P<body>.*?)\n```", section, re.DOTALL)
    if not match:
        return ""

    block = match.group("body").strip()
    if block.startswith("---"):
        block = block[3:].lstrip()
    if block.endswith("---"):
        block = block[:-3].rstrip()
    return block


def fallback_schema_fields(schema_block: str) -> dict[str, Any]:
    fields: dict[str, Any] = {}
    for line in schema_block.splitlines():
        match = re.match(r"^([a-zA-Z_][\w-]*):\s*(.*?)\s*(?:#.*)?$", line.strip())
        if match:
            fields[match.group(1)] = match.group(2).strip()
    return fields


def parse_metadata_schema(schema_block: str) -> dict[str, Any]:
    if not schema_block:
        return {}

    if yaml is None:
        return fallback_schema_fields(schema_block)

    try:
        parsed = yaml.safe_load(schema_block)
    except Exception:
        return fallback_schema_fields(schema_block)

    return parsed if isinstance(parsed, dict) else fallback_schema_fields(schema_block)


def parse_ref_token(token: str) -> RelationshipRef | None:
    cleaned = token.strip().strip("\"'").strip()
    if not cleaned or cleaned in {"[]", "None", "null"}:
        return None

    code_match = re.match(r"([A-Z]{3})", cleaned)
    if not code_match:
        return None

    return RelationshipRef(pattern=cleaned, code=code_match.group(1))


def parse_inline_refs(value: Any) -> list[RelationshipRef]:
    if value is None:
        return []

    if isinstance(value, list):
        refs = [parse_ref_token(str(item)) for item in value]
        return [ref for ref in refs if ref is not None]

    value_text = str(value).strip()
    if not value_text or value_text == "[]":
        return []

    if value_text.startswith("[") and value_text.endswith("]"):
        value_text = value_text[1:-1]

    refs = [parse_ref_token(part) for part in value_text.split(",")]
    return [ref for ref in refs if ref is not None]


def normalize_scalar(value: str) -> str:
    return value.strip().strip("\"'").strip()


def parse_inline_list(value: str) -> list[str]:
    inner = value.strip()
    if inner.startswith("[") and inner.endswith("]"):
        inner = inner[1:-1]
    return [normalize_scalar(part) for part in inner.split(",") if normalize_scalar(part)]


def fallback_frontmatter(frontmatter: str) -> dict[str, Any]:
    data: dict[str, Any] = {}
    current_key: str | None = None

    for raw_line in frontmatter.splitlines():
        if not raw_line.strip() or raw_line.lstrip().startswith("#"):
            continue

        if not raw_line.startswith((" ", "\t")) and ":" in raw_line:
            key, raw_value = raw_line.split(":", 1)
            key = key.strip()
            value = raw_value.strip()
            current_key = key

            if not value:
                data[key] = []
            elif value.startswith("[") and value.endswith("]"):
                data[key] = parse_inline_list(value)
            else:
                data[key] = normalize_scalar(value)
            continue

        if current_key and raw_line.strip().startswith("- "):
            value = normalize_scalar(raw_line.strip()[2:])
            existing = data.get(current_key)
            if not isinstance(existing, list):
                existing = [] if existing in (None, "") else [existing]
            existing.append(value)
            data[current_key] = existing

    return data


def parse_frontmatter(frontmatter: str) -> dict[str, Any]:
    if yaml is None:
        return fallback_frontmatter(frontmatter)

    try:
        parsed = yaml.safe_load(frontmatter) or {}
    except Exception:
        return fallback_frontmatter(frontmatter)

    return parsed if isinstance(parsed, dict) else fallback_frontmatter(frontmatter)


def split_frontmatter(content: str) -> tuple[dict[str, Any], str]:
    if not content.startswith("---\n"):
        return {}, content

    parts = content[4:].split("\n---\n", 1)
    if len(parts) != 2:
        return {}, content

    return parse_frontmatter(parts[0]), parts[1].strip()


def project_ref_token(token: str) -> ProjectRelationshipRef | None:
    cleaned = normalize_scalar(token)
    if not cleaned or cleaned in {"[]", "None", "null"}:
        return None

    code: str | None = None
    code_match = re.match(r"([A-Z]{3})(?:[-_*]|$)", cleaned)
    if code_match:
        code = code_match.group(1)
    else:
        id_code_match = re.match(r"([a-z]{3})(?:[-_]|$)", cleaned)
        if id_code_match:
            code = id_code_match.group(1).upper()

    document_id = None if "*" in cleaned else cleaned
    return ProjectRelationshipRef(pattern=cleaned, code=code, document_id=document_id)


def parse_project_refs(value: Any) -> list[ProjectRelationshipRef]:
    if value is None:
        return []

    if isinstance(value, list):
        refs = [project_ref_token(str(item)) for item in value]
        return [ref for ref in refs if ref is not None]

    value_text = str(value).strip()
    if not value_text or value_text == "[]":
        return []

    return [ref for ref in (project_ref_token(part) for part in parse_inline_list(value_text)) if ref is not None]


def extract_project_relationships(metadata: dict[str, Any]) -> dict[str, list[ProjectRelationshipRef]]:
    return {field: parse_project_refs(metadata.get(field)) for field in RELATIONSHIP_FIELDS}


def extract_relationships(schema_block: str, metadata_schema: dict[str, Any]) -> dict[str, list[RelationshipRef]]:
    relationships: dict[str, list[RelationshipRef]] = {field: [] for field in RELATIONSHIP_FIELDS}

    for field in RELATIONSHIP_FIELDS:
        relationships[field] = parse_inline_refs(metadata_schema.get(field))

    for line in schema_block.splitlines():
        match = re.match(r"^\s*(depends_on|enables|conflicts_with|related|supersedes):\s*(.*?)\s*(?:#.*)?$", line)
        if match:
            field, value = match.groups()
            refs = parse_inline_refs(value)
            if refs:
                relationships[field] = refs

    return relationships


def reference_time_for(document: BSpecDocumentType) -> str:
    if document.last_updated and re.match(r"^\d{4}-\d{2}-\d{2}$", document.last_updated):
        return f"{document.last_updated}T00:00:00+00:00"
    return DEFAULT_REFERENCE_TIME


def parse_spec_file(path: Path) -> BSpecDocumentType:
    content = path.read_text(encoding="utf-8")
    schema_block = extract_metadata_schema_block(content)
    metadata_schema = parse_metadata_schema(schema_block)

    code = field_value(content, "Document Type Code") or path.name.split("-")[0]
    name = field_value(content, "Document Type Name") or code
    domain_name = field_value(content, "Domain") or path.parent.name.replace("-", " ").title()

    return BSpecDocumentType(
        code=code,
        name=name,
        domain_name=domain_name,
        domain_slug=path.parent.name,
        status=field_value(content, "Status") or "Draft",
        version=field_value(content, "Version") or "1.0.0",
        last_updated=field_value(content, "Last Updated"),
        source_path=path,
        abstract=compact_text(extract_section(content, "Abstract")),
        purpose=compact_text(extract_section(content, "Purpose and Scope")),
        metadata_schema=metadata_schema,
        relationships=extract_relationships(schema_block, metadata_schema),
        quality_standards=compact_text(extract_section(content, "Quality Standards")),
        relationship_guidance=compact_text(
            "\n\n".join(
                filter(
                    None,
                    [
                        extract_section(content, "Relationship Guidelines"),
                        extract_section(content, "Document Relationships"),
                    ],
                )
            )
        ),
    )


def load_documents(spec_dir: Path) -> list[BSpecDocumentType]:
    paths = sorted(spec_dir.glob("*/*-spec.md"))
    return [parse_spec_file(path) for path in paths]


def load_bspec_version(spec_dir: Path) -> str:
    version_file = spec_dir / "version.txt"
    if version_file.exists():
        return version_file.read_text(encoding="utf-8").strip()
    return "1.0.0"


def project_reference_time(document: BSpecProjectDocument) -> str:
    for value in (document.updated, document.created):
        if value and re.match(r"^\d{4}-\d{2}-\d{2}$", value):
            return f"{value}T00:00:00+00:00"
    return datetime.now(timezone.utc).replace(microsecond=0).isoformat()


def parse_project_document(content: str, source_path: str) -> BSpecProjectDocument:
    metadata, markdown = split_frontmatter(content)
    path_fallback = Path(source_path).stem
    doc_type = str(metadata.get("type") or path_fallback.split("-")[0] or "DOC").upper()
    document_id = str(metadata.get("id") or path_fallback)

    return BSpecProjectDocument(
        document_id=document_id,
        title=str(metadata.get("title") or document_id),
        doc_type=doc_type,
        status=str(metadata.get("status") or "Draft"),
        version=str(metadata.get("version") or "1.0.0"),
        owner=str(metadata.get("owner") or "Unknown"),
        created=str(metadata.get("created")) if metadata.get("created") else None,
        updated=str(metadata.get("updated")) if metadata.get("updated") else None,
        domain=str(metadata.get("domain") or "unspecified"),
        source_path=source_path,
        content=compact_text(markdown, 2400),
        metadata=metadata,
        relationships=extract_project_relationships(metadata),
    )


def load_project_from_directory(input_path: Path) -> BSpecProject:
    manifest_path = input_path / "manifest.json"
    manifest: dict[str, Any] = {}
    documents_root = input_path / "documents"

    if manifest_path.exists():
        manifest = json.loads(manifest_path.read_text(encoding="utf-8"))
    else:
        documents_root = input_path

    documents: list[BSpecProjectDocument] = []
    if documents_root.exists():
        for path in sorted(documents_root.rglob("*.md")):
            if path.name.startswith("._") or path.name == ".DS_Store":
                continue
            source_path = display_path(path)
            if manifest_path.exists():
                source_path = f"{display_path(input_path)}/documents/{path.relative_to(documents_root)}"
            documents.append(parse_project_document(path.read_text(encoding="utf-8"), source_path))

    return BSpecProject(input_label=display_path(input_path), manifest=manifest, documents=documents)


def load_project_from_bspec(input_path: Path) -> BSpecProject:
    manifest: dict[str, Any] = {}
    documents: list[BSpecProjectDocument] = []

    def archive_name(member_name: str) -> str:
        while member_name.startswith("./"):
            member_name = member_name[2:]
        return member_name

    with tarfile.open(input_path, "r:gz") as archive:
        for member in archive.getmembers():
            if archive_name(member.name) == "manifest.json":
                manifest_file = archive.extractfile(member)
                if manifest_file:
                    manifest = json.loads(manifest_file.read().decode("utf-8"))
                break

        for member in sorted(archive.getmembers(), key=lambda item: item.name):
            member_name = archive_name(member.name)
            member_path = Path(member_name)
            if member_path.name.startswith("._") or member_path.name == ".DS_Store":
                continue
            if not member.isfile() or not member_name.startswith("documents/") or not member_name.endswith(".md"):
                continue
            document_file = archive.extractfile(member)
            if not document_file:
                continue
            source_path = f"{display_path(input_path)}!/{member_name}"
            documents.append(parse_project_document(document_file.read().decode("utf-8"), source_path))

    return BSpecProject(input_label=display_path(input_path), manifest=manifest, documents=documents)


def load_bspec_project(input_path: Path) -> BSpecProject:
    if input_path.is_dir():
        return load_project_from_directory(input_path)
    return load_project_from_bspec(input_path)


def document_episode(document: BSpecDocumentType, bspec_version: str, group_id: str) -> dict[str, Any]:
    body = {
        "kind": "bspec.document_type",
        "standard": "BSpec",
        "standard_version": bspec_version,
        "code": document.code,
        "name": document.name,
        "domain": {"slug": document.domain_slug, "name": document.domain_name},
        "status": document.status,
        "version": document.version,
        "source_path": repo_relative(document.source_path),
        "abstract": document.abstract,
        "purpose": document.purpose,
        "metadata_fields": sorted(document.metadata_schema.keys()),
        "relationships": {
            field: [ref.as_dict() for ref in refs]
            for field, refs in document.relationships.items()
            if refs
        },
        "quality_standards": document.quality_standards,
        "relationship_guidance": document.relationship_guidance,
    }

    return {
        "uuid": stable_uuid("document-type", f"{document.domain_slug}/{document.code}"),
        "group_id": group_id,
        "name": f"BSpec {bspec_version} document type {document.code}: {document.name}",
        "episode_body": body,
        "source": "json",
        "source_description": f"BSpec document type specification: {repo_relative(document.source_path)}",
        "reference_time": reference_time_for(document),
    }


def domain_episode(
    domain_slug: str,
    domain_name: str,
    documents: list[BSpecDocumentType],
    bspec_version: str,
    group_id: str,
) -> dict[str, Any]:
    documents = sorted(documents, key=lambda doc: doc.code)
    relationship_counts = Counter(
        field for document in documents for field, refs in document.relationships.items() if refs
    )

    body = {
        "kind": "bspec.domain",
        "standard": "BSpec",
        "standard_version": bspec_version,
        "slug": domain_slug,
        "name": domain_name,
        "document_type_count": len(documents),
        "document_types": [{"code": doc.code, "name": doc.name} for doc in documents],
        "relationship_counts": dict(sorted(relationship_counts.items())),
    }

    return {
        "uuid": stable_uuid("domain", domain_slug),
        "group_id": group_id,
        "name": f"BSpec {bspec_version} domain {domain_name}",
        "episode_body": body,
        "source": "json",
        "source_description": f"BSpec domain specification: spec/v1/{domain_slug}",
        "reference_time": DEFAULT_REFERENCE_TIME,
    }


def relationship_rule_episodes(
    documents: list[BSpecDocumentType],
    bspec_version: str,
    group_id: str,
) -> list[dict[str, Any]]:
    episodes: list[dict[str, Any]] = []
    for document in documents:
        for field, refs in document.relationships.items():
            for ref in refs:
                rule_key = f"{document.domain_slug}/{document.code}/{field}/{ref.pattern}"
                body = {
                    "kind": "bspec.relationship_rule",
                    "standard": "BSpec",
                    "standard_version": bspec_version,
                    "source_document_type": {
                        "code": document.code,
                        "name": document.name,
                        "domain": document.domain_slug,
                    },
                    "relationship": {
                        "field": field,
                        "label": RELATIONSHIP_LABELS[field],
                    },
                    "target_document_type": {
                        "code": ref.code,
                        "pattern": ref.pattern,
                    },
                    "source_path": repo_relative(document.source_path),
                }
                episodes.append(
                    {
                        "uuid": stable_uuid("relationship-rule", rule_key),
                        "group_id": group_id,
                        "name": (
                            f"BSpec relationship rule {document.code} "
                            f"{RELATIONSHIP_LABELS[field]} {ref.pattern}"
                        ),
                        "episode_body": body,
                        "source": "json",
                        "source_description": f"BSpec relationship rule from {repo_relative(document.source_path)}",
                        "reference_time": reference_time_for(document),
                    }
                )
    return episodes


def project_archive_episode(project: BSpecProject, group_id: str) -> dict[str, Any]:
    document_types = sorted({document.doc_type for document in project.documents if document.doc_type})
    domains = sorted({document.domain for document in project.documents if document.domain})
    relationship_counts = Counter(
        field for document in project.documents for field, refs in document.relationships.items() for _ in refs
    )
    manifest_name = project.manifest.get("name") or Path(project.input_label).stem or "BSpec Project"

    body = {
        "kind": "bspec.archive",
        "standard": "BSpec",
        "input": project.input_label,
        "manifest": project.manifest,
        "document_count": len(project.documents),
        "document_types": document_types,
        "domains": domains,
        "relationship_counts": dict(sorted(relationship_counts.items())),
    }

    return {
        "uuid": stable_uuid("project", f"{project.input_label}/{manifest_name}"),
        "group_id": group_id,
        "name": f"BSpec project {manifest_name}",
        "episode_body": body,
        "source": "json",
        "source_description": f"BSpec project archive: {project.input_label}",
        "reference_time": DEFAULT_REFERENCE_TIME,
    }


def project_document_episode(document: BSpecProjectDocument, group_id: str) -> dict[str, Any]:
    body = {
        "kind": "bspec.document",
        "standard": "BSpec",
        "id": document.document_id,
        "title": document.title,
        "type": document.doc_type,
        "status": document.status,
        "version": document.version,
        "owner": document.owner,
        "created": document.created,
        "updated": document.updated,
        "domain": document.domain,
        "source_path": document.source_path,
        "metadata": document.metadata,
        "relationships": {
            field: [ref.as_dict() for ref in refs]
            for field, refs in document.relationships.items()
            if refs
        },
        "content": document.content,
    }

    return {
        "uuid": stable_uuid("project-document", f"{group_id}/{document.document_id}/{document.source_path}"),
        "group_id": group_id,
        "name": f"BSpec document {document.document_id}: {document.title}",
        "episode_body": body,
        "source": "json",
        "source_description": f"BSpec document: {document.source_path}",
        "reference_time": project_reference_time(document),
    }


def project_relationship_fact_episodes(project: BSpecProject, group_id: str) -> list[dict[str, Any]]:
    documents_by_id = {document.document_id: document for document in project.documents}
    episodes: list[dict[str, Any]] = []

    for document in project.documents:
        for field, refs in document.relationships.items():
            for ref in refs:
                target_document = documents_by_id.get(ref.document_id or "")
                relation_key = f"{document.document_id}/{field}/{ref.pattern}"
                body = {
                    "kind": "bspec.document_relationship",
                    "standard": "BSpec",
                    "source_document": {
                        "id": document.document_id,
                        "title": document.title,
                        "type": document.doc_type,
                    },
                    "relationship": {
                        "field": field,
                        "label": RELATIONSHIP_LABELS[field],
                    },
                    "target_document": {
                        "id": ref.document_id,
                        "type": ref.code,
                        "pattern": ref.pattern,
                        "title": target_document.title if target_document else None,
                    },
                    "source_path": document.source_path,
                }
                episodes.append(
                    {
                        "uuid": stable_uuid("project-relationship", f"{group_id}/{relation_key}"),
                        "group_id": group_id,
                        "name": (
                            f"BSpec document relationship {document.document_id} "
                            f"{RELATIONSHIP_LABELS[field]} {ref.pattern}"
                        ),
                        "episode_body": body,
                        "source": "json",
                        "source_description": f"BSpec relationship from {document.source_path}",
                        "reference_time": project_reference_time(document),
                    }
                )

    return episodes


def build_episodes(
    spec_dir: Path,
    group_id: str,
    include_relationship_rules: bool = False,
) -> list[dict[str, Any]]:
    documents = load_documents(spec_dir)
    bspec_version = load_bspec_version(spec_dir)

    by_domain: dict[str, list[BSpecDocumentType]] = defaultdict(list)
    for document in documents:
        by_domain[document.domain_slug].append(document)

    episodes: list[dict[str, Any]] = []
    for domain_slug, domain_documents in sorted(by_domain.items()):
        domain_name = domain_documents[0].domain_name
        episodes.append(domain_episode(domain_slug, domain_name, domain_documents, bspec_version, group_id))

    episodes.extend(document_episode(document, bspec_version, group_id) for document in documents)

    if include_relationship_rules:
        episodes.extend(relationship_rule_episodes(documents, bspec_version, group_id))

    return episodes


def build_project_episodes(
    input_path: Path,
    group_id: str,
    include_relationship_facts: bool = False,
) -> list[dict[str, Any]]:
    project = load_bspec_project(input_path)
    episodes = [project_archive_episode(project, group_id)]
    episodes.extend(project_document_episode(document, group_id) for document in project.documents)
    if include_relationship_facts:
        episodes.extend(project_relationship_fact_episodes(project, group_id))
    return episodes


def write_jsonl(path: Path, episodes: list[dict[str, Any]]) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    with path.open("w", encoding="utf-8") as output:
        for episode in episodes:
            output.write(json.dumps(episode, sort_keys=True) + "\n")


def read_jsonl(path: Path, limit: int | None = None) -> list[dict[str, Any]]:
    episodes: list[dict[str, Any]] = []
    with path.open("r", encoding="utf-8") as input_file:
        for line in input_file:
            if line.strip():
                episodes.append(json.loads(line))
                if limit and len(episodes) >= limit:
                    break
    return episodes


def print_summary(spec_dir: Path) -> None:
    documents = load_documents(spec_dir)
    bspec_version = load_bspec_version(spec_dir)
    domain_counts = Counter(document.domain_slug for document in documents)
    relationship_counts = Counter(
        field for document in documents for field, refs in document.relationships.items() for _ in refs
    )

    print(f"BSpec version: {bspec_version}")
    print(f"Document types: {len(documents)}")
    print(f"Domains: {len(domain_counts)}")
    print(f"Relationship patterns: {sum(relationship_counts.values())}")
    print()
    print("Domains")
    for domain_slug, count in sorted(domain_counts.items()):
        domain_name = next(doc.domain_name for doc in documents if doc.domain_slug == domain_slug)
        print(f"  - {domain_slug}: {domain_name} ({count})")
    print()
    print("Relationships")
    for field in RELATIONSHIP_FIELDS:
        print(f"  - {field}: {relationship_counts[field]}")


def print_project_summary(input_path: Path) -> None:
    project = load_bspec_project(input_path)
    manifest_name = project.manifest.get("name") or Path(project.input_label).stem or "BSpec Project"
    domain_counts = Counter(document.domain for document in project.documents)
    type_counts = Counter(document.doc_type for document in project.documents)
    relationship_counts = Counter(
        field for document in project.documents for field, refs in document.relationships.items() for _ in refs
    )

    print(f"BSpec project: {manifest_name}")
    print(f"Input: {project.input_label}")
    print(f"Documents: {len(project.documents)}")
    print(f"Document types: {len(type_counts)}")
    print(f"Domains: {len(domain_counts)}")
    print(f"Relationship facts: {sum(relationship_counts.values())}")
    print()
    print("Document types")
    for doc_type, count in sorted(type_counts.items()):
        print(f"  - {doc_type}: {count}")
    print()
    print("Domains")
    for domain, count in sorted(domain_counts.items()):
        print(f"  - {domain}: {count}")
    print()
    print("Relationships")
    for field in RELATIONSHIP_FIELDS:
        print(f"  - {field}: {relationship_counts[field]}")


def post_json(url: str, payload: dict[str, Any]) -> dict[str, Any]:
    body = json.dumps(payload).encode("utf-8")
    req = request.Request(
        url,
        data=body,
        headers={"Content-Type": "application/json"},
        method="POST",
    )
    try:
        with request.urlopen(req, timeout=60) as response:
            raw = response.read().decode("utf-8")
            return json.loads(raw) if raw else {}
    except error.HTTPError as exc:
        details = exc.read().decode("utf-8")
        raise RuntimeError(f"POST {url} failed with HTTP {exc.code}: {details}") from exc


async def ingest_core(args: argparse.Namespace) -> None:
    try:
        from graphiti_core import Graphiti
        from graphiti_core.nodes import EpisodeType
    except ModuleNotFoundError as exc:
        raise SystemExit(
            "graphiti-core is not installed. Run: "
            "python -m pip install -r prototypes/graphiti/requirements.txt"
        ) from exc

    input_path = Path(args.input)
    episodes = read_jsonl(input_path, args.limit)
    graphiti = Graphiti(args.neo4j_uri, args.neo4j_user, args.neo4j_password)

    try:
        if hasattr(graphiti, "build_indices_and_constraints"):
            await graphiti.build_indices_and_constraints()

        for index, episode in enumerate(episodes, start=1):
            await graphiti.add_episode(
                uuid=episode["uuid"],
                group_id=args.group_id,
                name=episode["name"],
                episode_body=json.dumps(episode["episode_body"], sort_keys=True),
                source=EpisodeType.json,
                source_description=episode["source_description"],
                reference_time=datetime.fromisoformat(episode["reference_time"]),
            )
            print(f"[{index}/{len(episodes)}] queued {episode['name']}")
    finally:
        await graphiti.close()


async def search_core(args: argparse.Namespace) -> None:
    try:
        from graphiti_core import Graphiti
    except ModuleNotFoundError as exc:
        raise SystemExit(
            "graphiti-core is not installed. Run: "
            "python -m pip install -r prototypes/graphiti/requirements.txt"
        ) from exc

    graphiti = Graphiti(args.neo4j_uri, args.neo4j_user, args.neo4j_password)
    try:
        try:
            results = await graphiti.search(
                query=args.query,
                group_ids=[args.group_id],
                num_results=args.limit,
            )
        except TypeError:
            results = await graphiti.search(args.query)

        for result in results[: args.limit]:
            print(f"- {getattr(result, 'fact', result)}")
    finally:
        await graphiti.close()


def service_messages(episodes: list[dict[str, Any]], group_id: str) -> dict[str, Any]:
    return {
        "group_id": group_id,
        "messages": [
            {
                "uuid": episode["uuid"],
                "name": episode["name"],
                "role_type": "system",
                "role": "bspec",
                "content": json.dumps(episode["episode_body"], sort_keys=True),
                "timestamp": episode["reference_time"],
                "source_description": episode["source_description"],
            }
            for episode in episodes
        ],
    }


def command_summary(args: argparse.Namespace) -> int:
    print_summary(Path(args.spec_dir))
    return 0


def command_export(args: argparse.Namespace) -> int:
    episodes = build_episodes(
        Path(args.spec_dir),
        args.group_id,
        include_relationship_rules=args.include_relationship_rules,
    )
    write_jsonl(Path(args.output), episodes)
    print(f"Wrote {len(episodes)} Graphiti episodes to {args.output}")
    return 0


def command_summary_bspec(args: argparse.Namespace) -> int:
    print_project_summary(Path(args.input))
    return 0


def command_export_bspec(args: argparse.Namespace) -> int:
    episodes = build_project_episodes(
        Path(args.input),
        args.group_id,
        include_relationship_facts=args.include_relationship_facts,
    )
    write_jsonl(Path(args.output), episodes)
    print(f"Wrote {len(episodes)} Graphiti episodes to {args.output}")
    return 0


def command_ingest_service(args: argparse.Namespace) -> int:
    episodes = read_jsonl(Path(args.input), args.limit)
    payload = service_messages(episodes, args.group_id)
    url = args.service_url.rstrip("/") + "/messages"
    response = post_json(url, payload)
    print(json.dumps(response, indent=2, sort_keys=True))
    return 0


def command_search_service(args: argparse.Namespace) -> int:
    payload = {
        "group_ids": [args.group_id],
        "query": args.query,
        "max_facts": args.limit,
    }
    response = post_json(args.service_url.rstrip("/") + "/search", payload)
    for fact in response.get("facts", []):
        print(f"- {fact.get('fact')}")
    return 0


def command_ingest_core(args: argparse.Namespace) -> int:
    asyncio.run(ingest_core(args))
    return 0


def command_search_core(args: argparse.Namespace) -> int:
    asyncio.run(search_core(args))
    return 0


def add_common_graph_args(parser: argparse.ArgumentParser) -> None:
    parser.add_argument("--group-id", default=os.environ.get("BSPEC_GRAPHITI_GROUP_ID", DEFAULT_GROUP_ID))


def add_core_args(parser: argparse.ArgumentParser) -> None:
    parser.add_argument("--neo4j-uri", default=os.environ.get("NEO4J_URI", "bolt://localhost:7687"))
    parser.add_argument("--neo4j-user", default=os.environ.get("NEO4J_USER", "neo4j"))
    parser.add_argument("--neo4j-password", default=os.environ.get("NEO4J_PASSWORD", "password"))


def build_parser() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(description=__doc__)
    subparsers = parser.add_subparsers(dest="command", required=True)

    summary = subparsers.add_parser("summary", help="Summarize the BSpec graph model without Graphiti")
    summary.add_argument("--spec-dir", default=str(DEFAULT_SPEC_DIR))
    summary.set_defaults(func=command_summary)

    export = subparsers.add_parser("export", help="Export BSpec specs as Graphiti-ready JSONL episodes")
    export.add_argument("--spec-dir", default=str(DEFAULT_SPEC_DIR))
    export.add_argument("--output", default=str(DEFAULT_EXPORT_PATH))
    export.add_argument(
        "--include-relationship-rules",
        action="store_true",
        help="Also emit one episode per explicit relationship pattern",
    )
    add_common_graph_args(export)
    export.set_defaults(func=command_export)

    summary_bspec = subparsers.add_parser(
        "summary-bspec",
        help="Summarize an existing BSpec project directory or .bspec archive",
    )
    summary_bspec.add_argument("input", help="Path to an extracted BSpec project, documents directory, or .bspec file")
    summary_bspec.set_defaults(func=command_summary_bspec)

    export_bspec = subparsers.add_parser(
        "export-bspec",
        help="Export an existing BSpec project or .bspec archive as Graphiti-ready JSONL episodes",
    )
    export_bspec.add_argument("input", help="Path to an extracted BSpec project, documents directory, or .bspec file")
    export_bspec.add_argument("--output", default=str(DEFAULT_PROJECT_EXPORT_PATH))
    export_bspec.add_argument(
        "--include-relationship-facts",
        action="store_true",
        help="Also emit one episode per document relationship instance",
    )
    export_bspec.add_argument(
        "--group-id",
        default=os.environ.get("BSPEC_GRAPHITI_PROJECT_GROUP_ID", DEFAULT_PROJECT_GROUP_ID),
    )
    export_bspec.set_defaults(func=command_export_bspec)

    ingest_service = subparsers.add_parser("ingest-service", help="Ingest JSONL through Graphiti's FastAPI service")
    ingest_service.add_argument("--input", default=str(DEFAULT_EXPORT_PATH))
    ingest_service.add_argument("--service-url", default=os.environ.get("GRAPHITI_SERVICE_URL", "http://localhost:8000"))
    ingest_service.add_argument("--limit", type=int)
    add_common_graph_args(ingest_service)
    ingest_service.set_defaults(func=command_ingest_service)

    search_service = subparsers.add_parser("search-service", help="Search through Graphiti's FastAPI service")
    search_service.add_argument("query")
    search_service.add_argument("--service-url", default=os.environ.get("GRAPHITI_SERVICE_URL", "http://localhost:8000"))
    search_service.add_argument("--limit", type=int, default=10)
    add_common_graph_args(search_service)
    search_service.set_defaults(func=command_search_service)

    ingest_core_parser = subparsers.add_parser("ingest-core", help="Ingest JSONL through graphiti-core")
    ingest_core_parser.add_argument("--input", default=str(DEFAULT_EXPORT_PATH))
    ingest_core_parser.add_argument("--limit", type=int)
    add_common_graph_args(ingest_core_parser)
    add_core_args(ingest_core_parser)
    ingest_core_parser.set_defaults(func=command_ingest_core)

    search_core_parser = subparsers.add_parser("search-core", help="Search through graphiti-core")
    search_core_parser.add_argument("query")
    search_core_parser.add_argument("--limit", type=int, default=10)
    add_common_graph_args(search_core_parser)
    add_core_args(search_core_parser)
    search_core_parser.set_defaults(func=command_search_core)

    return parser


def main(argv: list[str] | None = None) -> int:
    parser = build_parser()
    args = parser.parse_args(argv)
    return args.func(args)


if __name__ == "__main__":
    sys.exit(main())
