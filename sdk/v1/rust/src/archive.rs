//! BSpec Archive (.bspec file) handling
//!
//! This module provides functionality for creating, reading, and manipulating
//! .bspec archive files that contain BSpec documents and assets.
//!
//! **GENERATED CODE - DO NOT EDIT MANUALLY**
//! Generated at: 2025-09-28T16:00:48.186517

use chrono::{DateTime, Utc};
use serde::{Deserialize, Serialize};
use std::collections::HashMap;
use std::fs::File;
use std::io::{Read, Write, BufReader, BufWriter};
use std::path::{Path, PathBuf};
use zip::{ZipArchive, ZipWriter, CompressionMethod};

use crate::base::{BSpecDocument, DocumentType, ConformanceLevel, IndustryProfile};
use crate::documents::*;
use crate::error::{BSpecError, Result};

/// Manifest for a BSpec archive
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ArchiveManifest {
    /// Archive format version
    pub format_version: String,
    /// BSpec version used
    pub bspec_version: String,
    /// Archive name/title
    pub name: String,
    /// Archive description
    pub description: String,
    /// Archive author
    pub author: String,
    /// Creation timestamp
    pub created_at: DateTime<Utc>,
    /// Last update timestamp
    pub updated_at: DateTime<Utc>,
    /// Total number of documents
    pub total_documents: u32,
    /// Document types included
    pub document_types: Vec<String>,
    /// Business domains covered
    pub domains: Vec<String>,
    /// Conformance level
    pub conformance_level: String,
    /// Industry profile
    pub industry_profile: String,
    /// Custom metadata
    #[serde(default)]
    pub metadata: HashMap<String, serde_json::Value>,
}

impl Default for ArchiveManifest {
    fn default() -> Self {
        let now = Utc::now();
        Self {
            format_version: "1.0.0".to_string(),
            bspec_version: "1.0.0".to_string(),
            name: String::new(),
            description: String::new(),
            author: String::new(),
            created_at: now,
            updated_at: now,
            total_documents: 0,
            document_types: Vec::new(),
            domains: Vec::new(),
            conformance_level: "bronze".to_string(),
            industry_profile: "software-saas".to_string(),
            metadata: HashMap::new(),
        }
    }
}

/// A BSpec archive containing documents and assets
#[derive(Debug)]
pub struct BSpecArchive {
    /// Archive manifest
    pub manifest: ArchiveManifest,
    /// Documents in the archive (path -> document)
    pub documents: HashMap<String, Box<dyn BSpecDocument>>,
    /// Assets in the archive (path -> data)
    pub assets: HashMap<String, Vec<u8>>,
    /// Computed data (path -> JSON)
    pub computed: HashMap<String, serde_json::Value>,
}

impl BSpecArchive {
    /// Create a new empty archive
    pub fn new(name: String) -> Self {
        let mut manifest = ArchiveManifest::default();
        manifest.name = name;

        Self {
            manifest,
            documents: HashMap::new(),
            assets: HashMap::new(),
            computed: HashMap::new(),
        }
    }

    /// Load archive from .bspec file
    pub fn load_from_file<P: AsRef<Path>>(path: P) -> Result<Self> {
        let file = File::open(path)?;
        let mut archive = ZipArchive::new(BufReader::new(file))?;

        // Read manifest
        let mut manifest_file = archive.by_name("manifest.json")?;
        let mut manifest_data = String::new();
        manifest_file.read_to_string(&mut manifest_data)?;
        let manifest: ArchiveManifest = serde_json::from_str(&manifest_data)?;

        let mut bspec_archive = Self {
            manifest,
            documents: HashMap::new(),
            assets: HashMap::new(),
            computed: HashMap::new(),
        };

        // Read documents
        for i in 0..archive.len() {
            let mut file = archive.by_index(i)?;
            let path = file.name().to_string();

            if path.starts_with("documents/") && path.ends_with(".json") {
                let mut content = String::new();
                file.read_to_string(&mut content)?;

                // Determine document type from path or content
                let document = bspec_archive.deserialize_document(&content)?;
                bspec_archive.documents.insert(path, document);
            } else if path.starts_with("assets/") {
                let mut data = Vec::new();
                file.read_to_end(&mut data)?;
                bspec_archive.assets.insert(path, data);
            } else if path.starts_with("computed/") && path.ends_with(".json") {
                let mut content = String::new();
                file.read_to_string(&mut content)?;
                let computed_data: serde_json::Value = serde_json::from_str(&content)?;
                bspec_archive.computed.insert(path, computed_data);
            }
        }

        Ok(bspec_archive)
    }

    /// Save archive to .bspec file
    pub fn save_to_file<P: AsRef<Path>>(&mut self, path: P) -> Result<()> {
        // Update manifest
        self.update_manifest();

        let file = File::create(path)?;
        let mut zip = ZipWriter::new(BufWriter::new(file));

        // Write manifest
        zip.start_file("manifest.json", Default::default())?;
        let manifest_json = serde_json::to_string_pretty(&self.manifest)?;
        zip.write_all(manifest_json.as_bytes())?;

        // Write documents
        for (path, document) in &self.documents {
            zip.start_file(path, Default::default())?;
            let doc_json = document.to_json()?;
            zip.write_all(doc_json.as_bytes())?;
        }

        // Write assets
        for (path, data) in &self.assets {
            zip.start_file(path, Default::default())?;
            zip.write_all(data)?;
        }

        // Write computed data
        for (path, data) in &self.computed {
            zip.start_file(path, Default::default())?;
            let computed_json = serde_json::to_string_pretty(data)?;
            zip.write_all(computed_json.as_bytes())?;
        }

        zip.finish()?;
        Ok(())
    }

    /// Add a document to the archive
    pub fn add_document(&mut self, document: Box<dyn BSpecDocument>) -> Result<()> {
        let doc_type = document.document_type();
        let id = document.metadata().id.clone();
        let path = format!("documents/{}.json", id);

        self.documents.insert(path, document);
        self.update_manifest();

        Ok(())
    }

    /// Remove a document from the archive
    pub fn remove_document(&mut self, document_id: &str) -> bool {
        let path = format!("documents/{}.json", document_id);
        if self.documents.remove(&path).is_some() {
            self.update_manifest();
            true
        } else {
            false
        }
    }

    /// Add an asset to the archive
    pub fn add_asset<P: AsRef<Path>>(&mut self, path: P, data: Vec<u8>) -> Result<()> {
        let asset_path = format!("assets/{}", path.as_ref().display());
        self.assets.insert(asset_path, data);
        Ok(())
    }

    /// Query documents by type
    pub fn query_by_type(&self, doc_type: DocumentType) -> Vec<&Box<dyn BSpecDocument>> {
        self.documents
            .values()
            .filter(|doc| doc.document_type() == doc_type)
            .collect()
    }

    /// Query documents by status
    pub fn query_by_status(&self, status: crate::base::DocumentStatus) -> Vec<&Box<dyn BSpecDocument>> {
        self.documents
            .values()
            .filter(|doc| doc.metadata().status == status)
            .collect()
    }

    /// Search documents by content
    pub fn search_content(&self, query: &str) -> Vec<&Box<dyn BSpecDocument>> {
        let query_lower = query.to_lowercase();
        self.documents
            .values()
            .filter(|doc| {
                doc.content().to_lowercase().contains(&query_lower)
                    || doc.metadata().title.to_lowercase().contains(&query_lower)
            })
            .collect()
    }

    /// Get document by ID
    pub fn get_document(&self, id: &str) -> Option<&Box<dyn BSpecDocument>> {
        let path = format!("documents/{}.json", id);
        self.documents.get(&path)
    }

    /// Get all documents
    pub fn get_all_documents(&self) -> Vec<&Box<dyn BSpecDocument>> {
        self.documents.values().collect()
    }

    /// Check conformance level
    pub fn check_conformance(&self, level: ConformanceLevel) -> bool {
        let doc_count = self.documents.len() as u32;
        doc_count >= level.min_documents()
    }

    /// Update manifest with current archive state
    fn update_manifest(&mut self) {
        self.manifest.updated_at = Utc::now();
        self.manifest.total_documents = self.documents.len() as u32;

        // Collect unique document types
        let mut doc_types: Vec<String> = self.documents
            .values()
            .map(|doc| doc.document_type().as_str().to_string())
            .collect();
        doc_types.sort();
        doc_types.dedup();
        self.manifest.document_types = doc_types;

        // Collect unique domains
        let mut domains: Vec<String> = self.documents
            .values()
            .map(|doc| doc.document_type().domain().as_str().to_string())
            .collect();
        domains.sort();
        domains.dedup();
        self.manifest.domains = domains;
    }

    /// Deserialize a document from JSON based on its type
    fn deserialize_document(&self, json: &str) -> Result<Box<dyn BSpecDocument>> {
        // Parse JSON to determine document type
        let value: serde_json::Value = serde_json::from_str(json)?;

        // Try to get document type from the JSON
        let doc_type_str = value
            .get("type")
            .or_else(|| value.get("document_type"))
            .and_then(|v| v.as_str())
            .ok_or_else(|| BSpecError::InvalidArchive("Missing document type".to_string()))?;

        let doc_type: DocumentType = doc_type_str.parse()?;

        // Deserialize based on type
        match doc_type {            DocumentType::Msn => {
                let doc: MsnDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Vsn => {
                let doc: VsnDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Val => {
                let doc: ValDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Str => {
                let doc: StrDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Obj => {
                let doc: ObjDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Mot => {
                let doc: MotDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Pur => {
                let doc: PurDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Thy => {
                let doc: ThyDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Mkt => {
                let doc: MktDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Seg => {
                let doc: SegDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Cmp => {
                let doc: CmpDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Pos => {
                let doc: PosDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Trn => {
                let doc: TrnDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Eco => {
                let doc: EcoDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Opp => {
                let doc: OppDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Thr => {
                let doc: ThrDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Reg => {
                let doc: RegDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Mac => {
                let doc: MacDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Per => {
                let doc: PerDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Jtb => {
                let doc: JtbDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Cjm => {
                let doc: CjmDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Use => {
                let doc: UseDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Sto => {
                let doc: StoDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Pai => {
                let doc: PaiDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Gai => {
                let doc: GaiDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Emp => {
                let doc: EmpDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Fee => {
                let doc: FeeDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Int => {
                let doc: IntDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Sur => {
                let doc: SurDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Beh => {
                let doc: BehDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Prd => {
                let doc: PrdDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Svc => {
                let doc: SvcDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Fea => {
                let doc: FeaDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Rod => {
                let doc: RodDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Req => {
                let doc: ReqDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Qua => {
                let doc: QuaDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Uxd => {
                let doc: UxdDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Per => {
                let doc: PerDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Int => {
                let doc: IntDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Sup => {
                let doc: SupDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Bmc => {
                let doc: BmcDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Rev => {
                let doc: RevDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Prc => {
                let doc: PrcDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Cst => {
                let doc: CstDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Chn => {
                let doc: ChnDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Rel => {
                let doc: RelDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Res => {
                let doc: ResDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Act => {
                let doc: ActDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Prt => {
                let doc: PrtDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Unt => {
                let doc: UntDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Ltv => {
                let doc: LtvDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Cac => {
                let doc: CacDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Prc => {
                let doc: PrcDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Wfl => {
                let doc: WflDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Org => {
                let doc: OrgDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Rol => {
                let doc: RolDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Tea => {
                let doc: TeaDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Ski => {
                let doc: SkiDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Pol => {
                let doc: PolDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Sla => {
                let doc: SlaDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Vnd => {
                let doc: VndDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Fac => {
                let doc: FacDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Too => {
                let doc: TooDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Cap => {
                let doc: CapDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Arc => {
                let doc: ArcDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Sys => {
                let doc: SysDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Dat => {
                let doc: DatDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Api => {
                let doc: ApiDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Inf => {
                let doc: InfDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Sec => {
                let doc: SecDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Dev => {
                let doc: DevDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Ana => {
                let doc: AnaDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Fin => {
                let doc: FinDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Bud => {
                let doc: BudDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::For => {
                let doc: ForDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Fnd => {
                let doc: FndDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Inv => {
                let doc: InvDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Val => {
                let doc: ValDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Met => {
                let doc: MetDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Rep => {
                let doc: RepDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Aud => {
                let doc: AudDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Tax => {
                let doc: TaxDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Rsk => {
                let doc: RskDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Mit => {
                let doc: MitDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Cmp => {
                let doc: CmpDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Gvn => {
                let doc: GvnDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Ctl => {
                let doc: CtlDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Cri => {
                let doc: CriDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Eth => {
                let doc: EthDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Sta => {
                let doc: StaDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Gtm => {
                let doc: GtmDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Grw => {
                let doc: GrwDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Scl => {
                let doc: SclDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Exp => {
                let doc: ExpDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Inn => {
                let doc: InnDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Rnd => {
                let doc: RndDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Acq => {
                let doc: AcqDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }            DocumentType::Exp => {
                let doc: ExpDocument = serde_json::from_str(json)?;
                Ok(Box::new(doc))
            }        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use tempfile::NamedTempFile;

    #[test]
    fn test_new_archive() {
        let archive = BSpecArchive::new("Test Archive".to_string());
        assert_eq!(archive.manifest.name, "Test Archive");
        assert_eq!(archive.documents.len(), 0);
    }

    #[test]
    fn test_add_document() {
        let mut archive = BSpecArchive::new("Test".to_string());

        let mission = MsnDocument::new(
            "MSN-test-001".to_string(),
            "Test Mission".to_string(),
            "Test Owner".to_string(),
        );

        archive.add_document(Box::new(mission)).unwrap();
        assert_eq!(archive.documents.len(), 1);
        assert_eq!(archive.manifest.total_documents, 1);
    }

    #[test]
    fn test_save_and_load() {
        let mut archive = BSpecArchive::new("Test Archive".to_string());

        let mission = MsnDocument::new(
            "MSN-test-001".to_string(),
            "Test Mission".to_string(),
            "Test Owner".to_string(),
        );
        archive.add_document(Box::new(mission)).unwrap();

        // Save to temporary file
        let temp_file = NamedTempFile::new().unwrap();
        archive.save_to_file(temp_file.path()).unwrap();

        // Load from file
        let loaded_archive = BSpecArchive::load_from_file(temp_file.path()).unwrap();
        assert_eq!(loaded_archive.manifest.name, "Test Archive");
        assert_eq!(loaded_archive.documents.len(), 1);
    }
}