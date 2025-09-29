//! Document type implementations
//!
//! This module contains all BSpec document type implementations.pub mod msn;
pub use msn::MsnDocument;pub mod vsn;
pub use vsn::VsnDocument;pub mod val;
pub use val::ValDocument;pub mod str;
pub use str::StrDocument;pub mod obj;
pub use obj::ObjDocument;pub mod mot;
pub use mot::MotDocument;pub mod pur;
pub use pur::PurDocument;pub mod thy;
pub use thy::ThyDocument;pub mod mkt;
pub use mkt::MktDocument;pub mod seg;
pub use seg::SegDocument;pub mod cmp;
pub use cmp::CmpDocument;pub mod pos;
pub use pos::PosDocument;pub mod trn;
pub use trn::TrnDocument;pub mod eco;
pub use eco::EcoDocument;pub mod opp;
pub use opp::OppDocument;pub mod thr;
pub use thr::ThrDocument;pub mod reg;
pub use reg::RegDocument;pub mod mac;
pub use mac::MacDocument;pub mod per;
pub use per::PerDocument;pub mod jtb;
pub use jtb::JtbDocument;pub mod cjm;
pub use cjm::CjmDocument;pub mod use;
pub use use::UseDocument;pub mod sto;
pub use sto::StoDocument;pub mod pai;
pub use pai::PaiDocument;pub mod gai;
pub use gai::GaiDocument;pub mod emp;
pub use emp::EmpDocument;pub mod fee;
pub use fee::FeeDocument;pub mod int;
pub use int::IntDocument;pub mod sur;
pub use sur::SurDocument;pub mod beh;
pub use beh::BehDocument;pub mod prd;
pub use prd::PrdDocument;pub mod svc;
pub use svc::SvcDocument;pub mod fea;
pub use fea::FeaDocument;pub mod rod;
pub use rod::RodDocument;pub mod req;
pub use req::ReqDocument;pub mod qua;
pub use qua::QuaDocument;pub mod uxd;
pub use uxd::UxdDocument;pub mod per;
pub use per::PerDocument;pub mod int;
pub use int::IntDocument;pub mod sup;
pub use sup::SupDocument;pub mod bmc;
pub use bmc::BmcDocument;pub mod rev;
pub use rev::RevDocument;pub mod prc;
pub use prc::PrcDocument;pub mod cst;
pub use cst::CstDocument;pub mod chn;
pub use chn::ChnDocument;pub mod rel;
pub use rel::RelDocument;pub mod res;
pub use res::ResDocument;pub mod act;
pub use act::ActDocument;pub mod prt;
pub use prt::PrtDocument;pub mod unt;
pub use unt::UntDocument;pub mod ltv;
pub use ltv::LtvDocument;pub mod cac;
pub use cac::CacDocument;pub mod prc;
pub use prc::PrcDocument;pub mod wfl;
pub use wfl::WflDocument;pub mod org;
pub use org::OrgDocument;pub mod rol;
pub use rol::RolDocument;pub mod tea;
pub use tea::TeaDocument;pub mod ski;
pub use ski::SkiDocument;pub mod pol;
pub use pol::PolDocument;pub mod sla;
pub use sla::SlaDocument;pub mod vnd;
pub use vnd::VndDocument;pub mod fac;
pub use fac::FacDocument;pub mod too;
pub use too::TooDocument;pub mod cap;
pub use cap::CapDocument;pub mod arc;
pub use arc::ArcDocument;pub mod sys;
pub use sys::SysDocument;pub mod dat;
pub use dat::DatDocument;pub mod api;
pub use api::ApiDocument;pub mod inf;
pub use inf::InfDocument;pub mod sec;
pub use sec::SecDocument;pub mod dev;
pub use dev::DevDocument;pub mod ana;
pub use ana::AnaDocument;pub mod fin;
pub use fin::FinDocument;pub mod bud;
pub use bud::BudDocument;pub mod for;
pub use for::ForDocument;pub mod fnd;
pub use fnd::FndDocument;pub mod inv;
pub use inv::InvDocument;pub mod val;
pub use val::ValDocument;pub mod met;
pub use met::MetDocument;pub mod rep;
pub use rep::RepDocument;pub mod aud;
pub use aud::AudDocument;pub mod tax;
pub use tax::TaxDocument;pub mod rsk;
pub use rsk::RskDocument;pub mod mit;
pub use mit::MitDocument;pub mod cmp;
pub use cmp::CmpDocument;pub mod gvn;
pub use gvn::GvnDocument;pub mod ctl;
pub use ctl::CtlDocument;pub mod cri;
pub use cri::CriDocument;pub mod eth;
pub use eth::EthDocument;pub mod sta;
pub use sta::StaDocument;pub mod gtm;
pub use gtm::GtmDocument;pub mod grw;
pub use grw::GrwDocument;pub mod scl;
pub use scl::SclDocument;pub mod exp;
pub use exp::ExpDocument;pub mod inn;
pub use inn::InnDocument;pub mod rnd;
pub use rnd::RndDocument;pub mod acq;
pub use acq::AcqDocument;pub mod exp;
pub use exp::ExpDocument;
// Re-export all document types
pub use crate::base::BSpecDocument;

/// Create a document from JSON based on its type
pub fn from_json(json: &str) -> crate::Result<Box<dyn BSpecDocument>> {
    let value: serde_json::Value = serde_json::from_str(json)?;

    let doc_type_str = value
        .get("type")
        .or_else(|| value.get("document_type"))
        .and_then(|v| v.as_str())
        .ok_or_else(|| crate::BSpecError::InvalidArchive("Missing document type".to_string()))?;

    let doc_type: crate::base::DocumentType = doc_type_str.parse()?;

    match doc_type {        crate::base::DocumentType::Msn => {
            let doc: MsnDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Vsn => {
            let doc: VsnDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Val => {
            let doc: ValDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Str => {
            let doc: StrDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Obj => {
            let doc: ObjDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Mot => {
            let doc: MotDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Pur => {
            let doc: PurDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Thy => {
            let doc: ThyDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Mkt => {
            let doc: MktDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Seg => {
            let doc: SegDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Cmp => {
            let doc: CmpDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Pos => {
            let doc: PosDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Trn => {
            let doc: TrnDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Eco => {
            let doc: EcoDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Opp => {
            let doc: OppDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Thr => {
            let doc: ThrDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Reg => {
            let doc: RegDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Mac => {
            let doc: MacDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Per => {
            let doc: PerDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Jtb => {
            let doc: JtbDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Cjm => {
            let doc: CjmDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Use => {
            let doc: UseDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Sto => {
            let doc: StoDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Pai => {
            let doc: PaiDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Gai => {
            let doc: GaiDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Emp => {
            let doc: EmpDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Fee => {
            let doc: FeeDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Int => {
            let doc: IntDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Sur => {
            let doc: SurDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Beh => {
            let doc: BehDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Prd => {
            let doc: PrdDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Svc => {
            let doc: SvcDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Fea => {
            let doc: FeaDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Rod => {
            let doc: RodDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Req => {
            let doc: ReqDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Qua => {
            let doc: QuaDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Uxd => {
            let doc: UxdDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Per => {
            let doc: PerDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Int => {
            let doc: IntDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Sup => {
            let doc: SupDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Bmc => {
            let doc: BmcDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Rev => {
            let doc: RevDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Prc => {
            let doc: PrcDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Cst => {
            let doc: CstDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Chn => {
            let doc: ChnDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Rel => {
            let doc: RelDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Res => {
            let doc: ResDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Act => {
            let doc: ActDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Prt => {
            let doc: PrtDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Unt => {
            let doc: UntDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Ltv => {
            let doc: LtvDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Cac => {
            let doc: CacDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Prc => {
            let doc: PrcDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Wfl => {
            let doc: WflDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Org => {
            let doc: OrgDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Rol => {
            let doc: RolDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Tea => {
            let doc: TeaDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Ski => {
            let doc: SkiDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Pol => {
            let doc: PolDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Sla => {
            let doc: SlaDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Vnd => {
            let doc: VndDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Fac => {
            let doc: FacDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Too => {
            let doc: TooDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Cap => {
            let doc: CapDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Arc => {
            let doc: ArcDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Sys => {
            let doc: SysDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Dat => {
            let doc: DatDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Api => {
            let doc: ApiDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Inf => {
            let doc: InfDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Sec => {
            let doc: SecDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Dev => {
            let doc: DevDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Ana => {
            let doc: AnaDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Fin => {
            let doc: FinDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Bud => {
            let doc: BudDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::For => {
            let doc: ForDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Fnd => {
            let doc: FndDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Inv => {
            let doc: InvDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Val => {
            let doc: ValDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Met => {
            let doc: MetDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Rep => {
            let doc: RepDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Aud => {
            let doc: AudDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Tax => {
            let doc: TaxDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Rsk => {
            let doc: RskDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Mit => {
            let doc: MitDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Cmp => {
            let doc: CmpDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Gvn => {
            let doc: GvnDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Ctl => {
            let doc: CtlDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Cri => {
            let doc: CriDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Eth => {
            let doc: EthDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Sta => {
            let doc: StaDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Gtm => {
            let doc: GtmDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Grw => {
            let doc: GrwDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Scl => {
            let doc: SclDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Exp => {
            let doc: ExpDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Inn => {
            let doc: InnDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Rnd => {
            let doc: RndDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Acq => {
            let doc: AcqDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }        crate::base::DocumentType::Exp => {
            let doc: ExpDocument = serde_json::from_str(json)?;
            Ok(Box::new(doc))
        }    }
}