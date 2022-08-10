/*
 * Spyderbat API UI & Public APIs
 *
 * Restful APIs for use by UI & customers.
 *
 * The version of the OpenAPI document: 1.0.0
 * Contact: apisupport@spyderbat.com
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Default, Serialize, Deserialize)]
pub struct SrcUpdateInput {
    /// User supplied description of the source
    #[serde(rename = "description", skip_serializing_if = "Option::is_none")]
    pub description: Option<String>,
    #[serde(rename = "last_data", skip_serializing_if = "Option::is_none")]
    pub last_data: Option<String>,
    /// The end of the last chunk ingested from the agent
    #[serde(rename = "last_ingest_chunk_end_time", skip_serializing_if = "Option::is_none")]
    pub last_ingest_chunk_end_time: Option<String>,
    /// The end of the last chunk stored from the agent
    #[serde(rename = "last_stored_chunk_end_time", skip_serializing_if = "Option::is_none")]
    pub last_stored_chunk_end_time: Option<String>,
    /// User supplied name of the source
    #[serde(rename = "name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    /// Resource name used for RBAC
    #[serde(rename = "resource_name", skip_serializing_if = "Option::is_none")]
    pub resource_name: Option<String>,
    #[serde(rename = "resource_policy", skip_serializing_if = "Option::is_none")]
    pub resource_policy: Option<Box<crate::models::ResourcePolicy>>,
    /// Description of the runtime of the source
    #[serde(rename = "runtime_description", skip_serializing_if = "Option::is_none")]
    pub runtime_description: Option<String>,
    #[serde(rename = "runtime_details", skip_serializing_if = "Option::is_none")]
    pub runtime_details: Option<Box<crate::models::OrcApiRuntimeDetails>>,
    /// User supplied tags
    #[serde(rename = "tags", skip_serializing_if = "Option::is_none")]
    pub tags: Option<Vec<String>>,
    /// Type of source
    #[serde(rename = "type", skip_serializing_if = "Option::is_none")]
    pub _type: Option<String>,
    /// Valid from date, the first date this object was valid
    #[serde(rename = "valid_from", skip_serializing_if = "Option::is_none")]
    pub valid_from: Option<String>,
    /// Valid to date, the date this object is valid to
    #[serde(rename = "valid_to", skip_serializing_if = "Option::is_none")]
    pub valid_to: Option<String>,
}

impl SrcUpdateInput {
    pub fn new() -> SrcUpdateInput {
        SrcUpdateInput {
            description: None,
            last_data: None,
            last_ingest_chunk_end_time: None,
            last_stored_chunk_end_time: None,
            name: None,
            resource_name: None,
            resource_policy: None,
            runtime_description: None,
            runtime_details: None,
            tags: None,
            _type: None,
            valid_from: None,
            valid_to: None,
        }
    }
}


