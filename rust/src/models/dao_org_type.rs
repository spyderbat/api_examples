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
pub struct DaoOrgType {
    /// Type of organization
    #[serde(rename = "description", skip_serializing_if = "Option::is_none")]
    pub description: Option<String>,
    #[serde(rename = "org_uid", skip_serializing_if = "Option::is_none")]
    pub org_uid: Option<String>,
    #[serde(rename = "policy", skip_serializing_if = "Option::is_none")]
    pub policy: Option<Box<crate::models::DaoOrgTypePolicy>>,
    #[serde(rename = "uid", skip_serializing_if = "Option::is_none")]
    pub uid: Option<String>,
}

impl DaoOrgType {
    pub fn new() -> DaoOrgType {
        DaoOrgType {
            description: None,
            org_uid: None,
            policy: None,
            uid: None,
        }
    }
}


