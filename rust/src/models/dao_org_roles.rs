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
pub struct DaoOrgRoles {
    /// Default roles for the user
    #[serde(rename = "default_roles", skip_serializing_if = "Option::is_none")]
    pub default_roles: Option<Vec<String>>,
    /// Org UID
    #[serde(rename = "org_uid", skip_serializing_if = "Option::is_none")]
    pub org_uid: Option<String>,
}

impl DaoOrgRoles {
    pub fn new() -> DaoOrgRoles {
        DaoOrgRoles {
            default_roles: None,
            org_uid: None,
        }
    }
}


