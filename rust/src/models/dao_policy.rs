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
pub struct DaoPolicy {
    /// Allowed providers (use direct for direct sign-in)
    #[serde(rename = "allowed_providers")]
    pub allowed_providers: Vec<String>,
    /// Authentication policy
    #[serde(rename = "auth_policy", skip_serializing_if = "Option::is_none")]
    pub auth_policy: Option<serde_json::Value>,
    /// comment
    #[serde(rename = "comment", skip_serializing_if = "Option::is_none")]
    pub comment: Option<String>,
    /// Regular expression for email addresses
    #[serde(rename = "email_regex")]
    pub email_regex: String,
    #[serde(rename = "signup_policy")]
    pub signup_policy: Box<crate::models::DaoSignupPolicy>,
}

impl DaoPolicy {
    pub fn new(allowed_providers: Vec<String>, email_regex: String, signup_policy: crate::models::DaoSignupPolicy) -> DaoPolicy {
        DaoPolicy {
            allowed_providers,
            auth_policy: None,
            comment: None,
            email_regex,
            signup_policy: Box::new(signup_policy),
        }
    }
}


