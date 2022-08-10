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
pub struct SessionOrgTypeMaxLimit {
    /// What is the current value
    #[serde(rename = "current_value", skip_serializing_if = "Option::is_none")]
    pub current_value: Option<f64>,
    /// Description of the limit
    #[serde(rename = "description", skip_serializing_if = "Option::is_none")]
    pub description: Option<String>,
    /// Has the limit been met or exceeded
    #[serde(rename = "limit_exceeded", skip_serializing_if = "Option::is_none")]
    pub limit_exceeded: Option<bool>,
    /// What is the max limit value
    #[serde(rename = "limit_value", skip_serializing_if = "Option::is_none")]
    pub limit_value: Option<f64>,
    /// How many items can be added to the current value before it is exceeded
    #[serde(rename = "remaining_capacity", skip_serializing_if = "Option::is_none")]
    pub remaining_capacity: Option<f64>,
    /// The time window in seconds that the limit is calcuated on
    #[serde(rename = "time_window_in_seconds", skip_serializing_if = "Option::is_none")]
    pub time_window_in_seconds: Option<i32>,
}

impl SessionOrgTypeMaxLimit {
    pub fn new() -> SessionOrgTypeMaxLimit {
        SessionOrgTypeMaxLimit {
            current_value: None,
            description: None,
            limit_exceeded: None,
            limit_value: None,
            remaining_capacity: None,
            time_window_in_seconds: None,
        }
    }
}


