/*
 * Spyderbat API UI & Public APIs
 *
 * Restful APIs for use by UI & customers.
 *
 * The version of the OpenAPI document: 1.0.0
 * Contact: apisupport@spyderbat.com
 * Generated by: https://openapi-generator.tech
 */

/// OrcApiRuntimeDetails : Runtime details



#[derive(Clone, Debug, PartialEq, Default, Serialize, Deserialize)]
pub struct OrcApiRuntimeDetails {
    #[serde(rename = "agent_arch", skip_serializing_if = "Option::is_none")]
    pub agent_arch: Option<String>,
    #[serde(rename = "agent_version", skip_serializing_if = "Option::is_none")]
    pub agent_version: Option<String>,
    #[serde(rename = "boot_time", skip_serializing_if = "Option::is_none")]
    pub boot_time: Option<f64>,
    #[serde(rename = "cloud_image_id", skip_serializing_if = "Option::is_none")]
    pub cloud_image_id: Option<String>,
    #[serde(rename = "cloud_instance_id", skip_serializing_if = "Option::is_none")]
    pub cloud_instance_id: Option<String>,
    #[serde(rename = "cloud_region", skip_serializing_if = "Option::is_none")]
    pub cloud_region: Option<String>,
    #[serde(rename = "cloud_tags", skip_serializing_if = "Option::is_none")]
    pub cloud_tags: Option<Vec<String>>,
    #[serde(rename = "cloud_type", skip_serializing_if = "Option::is_none")]
    pub cloud_type: Option<String>,
    #[serde(rename = "cpu_cores", skip_serializing_if = "Option::is_none")]
    pub cpu_cores: Option<i32>,
    #[serde(rename = "cpu_make", skip_serializing_if = "Option::is_none")]
    pub cpu_make: Option<String>,
    #[serde(rename = "cpu_model", skip_serializing_if = "Option::is_none")]
    pub cpu_model: Option<String>,
    #[serde(rename = "hostname", skip_serializing_if = "Option::is_none")]
    pub hostname: Option<String>,
    #[serde(rename = "ip_addresses", skip_serializing_if = "Option::is_none")]
    pub ip_addresses: Option<Vec<String>>,
    #[serde(rename = "mac_addresses", skip_serializing_if = "Option::is_none")]
    pub mac_addresses: Option<Vec<String>>,
    #[serde(rename = "os_name", skip_serializing_if = "Option::is_none")]
    pub os_name: Option<String>,
    #[serde(rename = "os_pretty_name", skip_serializing_if = "Option::is_none")]
    pub os_pretty_name: Option<String>,
    #[serde(rename = "request_ip", skip_serializing_if = "Option::is_none")]
    pub request_ip: Option<String>,
    #[serde(rename = "src_uid", skip_serializing_if = "Option::is_none")]
    pub src_uid: Option<String>,
    #[serde(rename = "uname", skip_serializing_if = "Option::is_none")]
    pub uname: Option<String>,
}

impl OrcApiRuntimeDetails {
    /// Runtime details
    pub fn new() -> OrcApiRuntimeDetails {
        OrcApiRuntimeDetails {
            agent_arch: None,
            agent_version: None,
            boot_time: None,
            cloud_image_id: None,
            cloud_instance_id: None,
            cloud_region: None,
            cloud_tags: None,
            cloud_type: None,
            cpu_cores: None,
            cpu_make: None,
            cpu_model: None,
            hostname: None,
            ip_addresses: None,
            mac_addresses: None,
            os_name: None,
            os_pretty_name: None,
            request_ip: None,
            src_uid: None,
            uname: None,
        }
    }
}


