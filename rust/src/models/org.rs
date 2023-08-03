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
pub struct Org {
    /// Total number of active sources within the last 5 minutes
    #[serde(rename = "active_sources", skip_serializing_if = "Option::is_none")]
    pub active_sources: Option<i32>,
    /// Total number of active users within the last 7 days (which might be active on a different org)
    #[serde(rename = "active_users", skip_serializing_if = "Option::is_none")]
    pub active_users: Option<i32>,
    /// Name of the organization
    #[serde(rename = "name")]
    pub name: String,
    /// Organization Type
    #[serde(rename = "org_type_uid", skip_serializing_if = "Option::is_none")]
    pub org_type_uid: Option<String>,
    /// The email address of the user who owns this org
    #[serde(rename = "owner_email")]
    pub owner_email: String,
    /// The user UID who owns this organization
    #[serde(rename = "owner_uid", skip_serializing_if = "Option::is_none")]
    pub owner_uid: Option<String>,
    /// Resource name utilized by RBAC
    #[serde(rename = "resource_name", skip_serializing_if = "Option::is_none")]
    pub resource_name: Option<String>,
    #[serde(rename = "resource_policy", skip_serializing_if = "Option::is_none")]
    pub resource_policy: Option<Box<crate::models::ResourcePolicy>>,
    /// User supplied tags
    #[serde(rename = "tags", skip_serializing_if = "Option::is_none")]
    pub tags: Option<Vec<String>>,
    /// Total number of sources
    #[serde(rename = "total_sources", skip_serializing_if = "Option::is_none")]
    pub total_sources: Option<i32>,
    /// Total number of users
    #[serde(rename = "total_users", skip_serializing_if = "Option::is_none")]
    pub total_users: Option<i32>,
    /// Org UID
    #[serde(rename = "uid", skip_serializing_if = "Option::is_none")]
    pub uid: Option<String>,
    /// Valid from date, the first date this object was valid
    #[serde(rename = "valid_from", skip_serializing_if = "Option::is_none")]
    pub valid_from: Option<String>,
    /// Valid to date, the date this object is valid to
    #[serde(rename = "valid_to", skip_serializing_if = "Option::is_none")]
    pub valid_to: Option<String>,
}

impl Org {
    pub fn new(name: String, owner_email: String) -> Org {
        Org {
            active_sources: None,
            active_users: None,
            name,
            org_type_uid: None,
            owner_email,
            owner_uid: None,
            resource_name: None,
            resource_policy: None,
            tags: None,
            total_sources: None,
            total_users: None,
            uid: None,
            valid_from: None,
            valid_to: None,
        }
    }
}


