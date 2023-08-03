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
pub struct DashboardSearch {
    /// UI supplied JSON object
    #[serde(rename = "data", skip_serializing_if = "Option::is_none")]
    pub data: Option<::std::collections::HashMap<String, serde_json::Value>>,
    /// Description of query (max 64 chars)
    #[serde(rename = "description", skip_serializing_if = "Option::is_none")]
    pub description: Option<String>,
    /// Are notifications generated from this search
    #[serde(rename = "notify", skip_serializing_if = "Option::is_none")]
    pub notify: Option<bool>,
    /// Frequency of notification in seconds. One of 300, 3600, or 86400.
    #[serde(rename = "notify_frequency", skip_serializing_if = "Option::is_none")]
    pub notify_frequency: Option<i32>,
    /// Org UID
    #[serde(rename = "org_uid", skip_serializing_if = "Option::is_none")]
    pub org_uid: Option<String>,
    /// Search query to run
    #[serde(rename = "search", skip_serializing_if = "Option::is_none")]
    pub search: Option<String>,
    /// User supplied tags
    #[serde(rename = "tags", skip_serializing_if = "Option::is_none")]
    pub tags: Option<Vec<String>>,
    /// UID for the DashboardSearch
    #[serde(rename = "uid", skip_serializing_if = "Option::is_none")]
    pub uid: Option<String>,
}

impl DashboardSearch {
    pub fn new() -> DashboardSearch {
        DashboardSearch {
            data: None,
            description: None,
            notify: None,
            notify_frequency: None,
            org_uid: None,
            search: None,
            tags: None,
            uid: None,
        }
    }
}


